package main

import (
	ini "aria-cloud/lib"
	"aria-cloud/util"
	"bufio"
	"bytes"
	"fmt"
	jsonit "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type file struct {
	username  string
	token     string
	from      string
	hash      string
	size      string
	to        string
	base      string
	uploadID  string
	chunkSize int
	chunkCnt  int
}

func main() {
	conf := ini.LoadServerConfig()
	fmt.Println(conf)

	var tf = file{}
	tf.username = "xhw"
	tf.token = "c090f2bbe7a12b9784ac61e24f705dce65d8997a"
	//tf.from = "/Volumes/Akiiita/考研相关/李林最后4套卷【数二】【微信公众号：考研核心资料】免费分享.pdf"
	tf.from = "/Volumes/Akiiita/考研相关/2021张剑黄皮书英语一提高版2009-2016试卷版01【W】【微信公众号：考研核心资料】免费分享.pdf"
	tf.base = filepath.Base(tf.from)
	//读取上传文件
	upfile, err := os.Open(tf.from)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer upfile.Close()
	tf.hash = util.FileSha1(upfile)
	fmt.Println("上传文件的hash：" + tf.hash)
	tf.size = strconv.Itoa(int(util.GetFileSize(tf.from)))
	fmt.Println("上传文件的大小：" + strconv.Itoa(int(util.GetFileSize(tf.from)/1024/1024)) + " M")
	tf.to = conf.UploadLocation
	fmt.Println("上传文件要放在：" + tf.to)

	// 1. 请求初始化分块上传接口
	resp, err := http.PostForm(
		"http://localhost:8080/file/mpupload/init",
		url.Values{
			"username": {tf.username},
			"token":    {tf.token},
			"filehash": {tf.hash},
			"filesize": {tf.size},
			"fileto":   {tf.to},
		})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// 2. 验证【/file/mpupload/init】接口
	//得到uploadID以及服务端指定的分块大小chunkSize
	tf.uploadID = jsonit.Get(body, "data").Get("UploadID").ToString()
	tf.chunkSize = jsonit.Get(body, "data").Get("ChunkSize").ToInt()
	tf.chunkCnt = jsonit.Get(body, "data").Get("ChunkCount").ToInt()
	fmt.Printf("uploadid: %s  chunksize: %d chunkCount: %d\n", tf.uploadID, tf.chunkSize, tf.chunkCnt)

	// 3. 验证【/file/mpupload/uppart】请求分块上传接口
	tURL := "http://localhost:8080/file/mpupload/uppart?" +
		"username=" + tf.username + "&token=" + tf.token + "&uploadid=" + tf.uploadID + "&fileto=" + tf.to
	err = multipartUpload(tf.from, tURL, tf.chunkSize)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 4. 验证【/file/mpupload/complete】请求分块完成接口
	resp, err = http.PostForm(
		"http://localhost:8080/file/mpupload/complete",
		url.Values{
			"username": {tf.username},
			"token":    {tf.token},
			"filehash": {tf.hash},
			"filesize": {tf.size},
			"filename": {tf.to + tf.base},
			"uploadid": {tf.uploadID},
		})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Printf("complete result: %s\n", string(body))

}

func multipartUpload(filename string, targetURL string, chunkSize int) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	index := 0

	ch := make(chan int)
	buf := make([]byte, chunkSize) //每次读取chunkSize大小的内容
	for {
		n, err := bfRd.Read(buf)
		if n <= 0 {
			break
		}
		index++

		bufCopied := make([]byte, 5*1024*1024)
		copy(bufCopied, buf)

		go func(b []byte, curIdx int) {
			fmt.Printf("upload_size: %d\n", len(b))

			resp, err := http.Post(
				targetURL+"&index="+strconv.Itoa(curIdx),
				"multipart/form-data",
				bytes.NewReader(b))
			if err != nil {
				fmt.Println(err)
			}

			body, er := ioutil.ReadAll(resp.Body)
			fmt.Printf("%+v %+v\n", string(body), er)
			resp.Body.Close()

			ch <- curIdx
		}(bufCopied[:n], index)

		//遇到任何错误立即返回，并忽略 EOF 错误信息
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	for idx := 0; idx < index; idx++ {
		select {
		case res := <-ch:
			fmt.Println(res)
		}
	}

	return nil
}
