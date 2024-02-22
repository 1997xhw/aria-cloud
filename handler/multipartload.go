package handler

import (
	"aria-cloud/cache/redis"
	"aria-cloud/db"
	"aria-cloud/util"
	"fmt"
	redi "github.com/garyburd/redigo/redis"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// MUltipartUploadInfo 初始化信息
type MUltipartUploadInfo struct {
	Filehash   string
	FileSize   int
	UploadID   string
	ChunkSize  int
	ChunkCount int
}

func InitialMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析用户请求参数
	r.ParseForm()
	username := r.Form.Get("usename")
	filehash := r.Form.Get("filehash")
	filesize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "params invaild", nil).JSONBytes())
		return
	}
	// 2 获得redis的一个链接
	redisConn := redis.RedisPool().Get()
	defer redisConn.Close()
	// 3 生成分块上传的初始信息
	upInfo := MUltipartUploadInfo{
		Filehash:   filehash,
		FileSize:   filesize,
		UploadID:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  5 * 1024 * 1024,
		ChunkCount: int(math.Ceil(float64(filesize) / (5 * 1024 * 1024))),
	}
	// 4 将初始化信息写入到redis缓存
	redisConn.Do("HSET", "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount)
	redisConn.Do("HSET", "MP_"+upInfo.UploadID, "filehash", upInfo.Filehash)
	redisConn.Do("HSET", "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)
	// 将响应初始化数据返回到客户端
	w.Write(util.NewRespMsg(0, "OK", upInfo).JSONBytes())
}

func UploadPartHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析用户请求参数
	r.ParseForm()
	//username := r.Form.Get("usename")
	uploadID := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")
	// 2 获得redis链接池的链接
	redisConn := redis.RedisPool().Get()
	defer redisConn.Close()
	// 3 获得文件句柄，用于存储内容
	fpath := "/data/" + uploadID + "/" + chunkIndex
	os.MkdirAll(path.Dir(fpath), 0755)
	create, err := os.Create(fpath)
	if err != nil {
		w.Write(util.NewRespMsg(-1, "Upload part failed", nil).JSONBytes())
		return
	}
	defer create.Close()
	buf := make([]byte, 1024*1024)
	for {
		n, err := r.Body.Read(buf)
		create.Write(buf[:n])
		if err != nil {
			break
		}
	}
	// 4 更新redis缓存状态
	redisConn.Do("HSET", "MP_"+uploadID, "chkidx_"+chunkIndex, 1)
	// 5 返回处理结果到客户端
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}

// 通知上传合并
func CompleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析请求参数
	r.ParseForm()
	upid := r.Form.Get("uploadid")
	username := r.Form.Get("username")
	//filehash := r.Form.Get("filehash")
	//filesize := r.Form.Get("filesize")
	filename := r.Form.Get("filename")

	// 2 获得redis链接池中的一个链接
	redisConn := redis.RedisPool().Get()
	defer redisConn.Close()

	// 3 通过upload查询redis并判断是否所有分块上传完成
	values, err := redi.Values(redisConn.Do("HGETALL", "MP_"+upid))
	if err != nil {
		fmt.Println(err)
		w.Write(util.NewRespMsg(-1, err.Error(), nil).JSONBytes())
		return
	}
	chunkCnt := 0
	totalCnt := 0
	for i := 0; i < len(values); i += 2 {
		filed, _ := redi.String(values[i], nil)
		value, _ := redi.String(values[i+1], nil)
		if filed == "chunkcount" {
			totalCnt, _ = strconv.Atoi(value)
		} else if strings.HasPrefix(filed, "chkidx_") && value == "1" {
			chunkCnt++
		}
	}
	if totalCnt != chunkCnt {
		fmt.Println("文件未完全上传，请继续等待！")
		w.Write(util.NewRespMsg(-2, "invaild request", nil).JSONBytes())
		return
	}
	// 4 TODO 合并分块
	// 4.1 创建或打开目标文件
	//outputFile, err := os.Create("/data/" + upid + "/complete")
	filedir := "/data/" + filename
	outputFile, err := os.Create(filedir)
	if err != nil {
		w.Write(util.NewRespMsg(-1, "合并文件时发生异常："+err.Error(), nil).JSONBytes())
		return
	}
	defer outputFile.Close()

	// 4.2 按顺序读取每一个分块
	for i := 0; i < totalCnt; i++ {
		chunkFile, err := os.Open(fmt.Sprintf("/data/%s/%d", upid, i))
		if err != nil {
			w.Write(util.NewRespMsg(-1, fmt.Sprintf("读取文件块 %d 时发生异常：%v", i, err.Error()), nil).JSONBytes())
			return
		}
		// 4.3 将分块内容写入目标文件
		_, err = io.Copy(outputFile, chunkFile)
		chunkFile.Close()
		if err != nil {
			w.Write(util.NewRespMsg(-1, fmt.Sprintf("写入文件块 %d 时发生异常：%v", i, err.Error()), nil).JSONBytes())
			return
		}
	}
	// 5 更新唯一文件表及用户文件表
	info, err := outputFile.Stat()
	if err != nil {
		// 处理错误
		log.Fatal(err)
		return
	}
	db.OnFileUploadFinished(util.FileSha1(outputFile), filename, info.Size(), filedir)
	db.OnUserFileUploadFinished(username, util.FileSha1(outputFile), filename, info.Size())
	// 6 响应处理结果
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}