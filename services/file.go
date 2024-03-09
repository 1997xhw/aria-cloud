package services

import (
	"aria-cloud/common"
	ini "aria-cloud/lib"
	"aria-cloud/models"
	"aria-cloud/util"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

func GetAllFileList(username string) ([]models.TableUserFile, error) {
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
	} else {
		log.Println("文件已在索引中存在")
	}

	err = models.OnUserFileUploadFinished(fileMeta, username)
	if err != nil {
		return err
	}

	return nil
}
