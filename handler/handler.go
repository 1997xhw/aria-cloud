package handler

import (
	"aria-storage/meta"
	"aria-storage/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := os.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Failed to get data, err:%s\n", err.Error())
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileSha1: "",
			FileName: head.Filename,
			FileSize: 0,
			Location: "./tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-06-06 15-04:05"),
		}

		//创建一个新的文件等待复制
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Println("Failed to newFile file:%s", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Failed to save data into file:%s\n", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		//meta.UpdateFileMeta(fileMeta)
		_ = meta.UpdateFileMetaDB(fileMeta)
		fmt.Println(
			"生成的文件元信息:",
			fileMeta,
		)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler 显示上传成功信息
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Upload finished!!")
	if err != nil {
		return
	}
}

// GetFileMetaHandler 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	filehash := r.Form["filehash"][0]
	//fileMeta := meta.GetFileMeta(filehash)
	fileMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	openFile, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			panic(err)
			return
		}
	}(openFile)

	data, err := io.ReadAll(openFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fm.FileName+"\"")

	w.Write(data)

}

// FileUpdateMetaHandler 更新元信息结构（重命名）
func FileUpdateMetaHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	opType := r.Form.Get("op")
	file_hash := r.Form.Get("filehash")
	new_filename := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta := meta.GetFileMeta(file_hash)

	// 修改实际文件的文件名
	new_location := strings.Replace(curFileMeta.Location, curFileMeta.FileName, new_filename, 1)
	err = os.Rename(curFileMeta.Location, new_location)
	if err != nil {
		fmt.Println("修改文件名发生错误！！！")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 更新新的文件地址
	curFileMeta.Location = new_location

	//更新新的文件名
	curFileMeta.FileName = new_filename
	meta.UpdateFileMeta(curFileMeta)

	marshal, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)

	//
	//openFile, err := os.Open(curFileMeta.Location)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//defer func(openFile *os.File) {
	//	err := openFile.Close()
	//	if err != nil {
	//		panic(err)
	//		return
	//	}
	//}(openFile)

}

// FileDeleteHandler 删除文件以及元信息
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash1 := r.Form.Get("filehash")

	curFileMeta := meta.GetFileMeta(fileHash1)
	err := os.Remove(curFileMeta.Location)
	if err != nil {
		fmt.Println("文件删除异常！！！")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	meta.RemoveFileMeta(fileHash1)

	w.WriteHeader(http.StatusOK)

}
