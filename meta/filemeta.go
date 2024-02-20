package meta

import (
	mydb "aria-cloud/db"
	"fmt"
)

// 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// 文件元信息结构
//type FileMeta struct {
//	FileSha1 string `json:"文件hash值"`
//	FileName string `json:"文件名"`
//	FileSize int64  `json:"文件大小"`
//	Location string `json:"文件路径"`
//	UploadAt string `json:"上传时间"`
//}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta 新增/更新文件元信息
func UpdateFileMeta(meta FileMeta) {
	fileMetas[meta.FileSha1] = meta
}

// UpdateFileMetaDB 新增/更新文件元信息到mysql中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(
		fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMeta 通过sh1值获取文件对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// 格式化输出结构的实例
func (f FileMeta) String() string {
	return fmt.Sprintf("\nFileSha1: %s\nFileName: %s\nFileSize: %d\nLocation: %s\nUploadAt: %s",
		f.FileSha1, f.FileName, f.FileSize, f.Location, f.UploadAt)
}

// RemoveFileMeta 删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}

func GetFileMetaDB(filesha1 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMeta(filesha1)
	if err != nil {
		fmt.Println("GetFileMetaDB Error! ", err.Error())
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
}
