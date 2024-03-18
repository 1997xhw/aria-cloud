package services

import (
	"aria-cloud/common"
	"aria-cloud/databases/mysql"
	ini "aria-cloud/lib"
	"aria-cloud/models"
	"aria-cloud/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"
)

func GetAllFileList(username string) ([]common.TableUserFile, error) {
	fileList, err := models.GetAllFileList(username)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return fileList, nil
}
func SaveFileHandler(file *multipart.FileHeader, username string) error {
	conf := ini.LoadServerConfig()
	fileMeta := common.FileMeta{
		FileSha1: "",
		FileName: file.Filename,
		FileSize: 0,
		Location: conf.UploadLocation + file.Filename,
		UploadAt: time.Now().Format("2006-06-06 15-04:05"),
	}
	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		fmt.Printf("Failed to newFile file:%s", err.Error())
		return err
	}
	defer newFile.Close()

	fileOpen, err := file.Open()
	if err != nil {
		fmt.Printf("Failed to open %s:%s", file.Filename, err.Error())
		return err
	}
	defer fileOpen.Close()

	fileMeta.FileSize, err = io.Copy(newFile, fileOpen)
	if err != nil {
		fmt.Printf("Failed to save data into file:%s\n", err.Error())
		return err
	}

	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)
	isFind, err := models.CheckFileHashExists(fileMeta)
	if err != nil {
		return err
	} else if !isFind {
		//如果文件库没有文件则入库
		_, err = models.UpdateFileMeta(fileMeta)
		if err != nil {
			return err
		}
		//上传oss
		go ini.UploadOss(fileMeta)
	} else {
		log.Println("文件已在索引中存在")
	}

	err = models.OnUserFileUploadFinished(fileMeta, username)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFileHandler(username, filehash string) error {
	var file common.TableUserFile
	var err error
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		var err2 error
		file, err2 = models.SelectTableUserFile(tx, username, filehash)
		if err2 != nil {
			return err2
		}

		err2 = models.DeleteUserFile(tx, username, filehash)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return err
	}
	fileSuffix := path.Ext(file.FileName)
	log.Println(fileSuffix)

	//删除oss上的数据
	err = ini.DeleteOss(file.FileSha1, fileSuffix)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFileFromOss(filename, filesha string) ([]byte, error) {
	fileSuffix := path.Ext(filename)
	fileBytes, err := ini.DownloadOss(filesha, fileSuffix)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}
