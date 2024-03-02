package models

import (
	"aria-cloud/common"
	"aria-cloud/databases/mysql"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type TableFile struct {
	FileSha1 string `gorm:"column:file_sha1;"`
	FileName string `gorm:"column:file_name;"`
	FileSize int64  `gorm:"column:file_size;"`
	FileAddr string `gorm:"column:file_addr;"`
}

func (TableFile) TableName() string {
	return "tbl_file"
}

type TableUserFile struct {
	Username string `gorm:"column:user_name;"`
	FileSha1 string `gorm:"column:file_sha1;"`
	FileName string `gorm:"column:file_name;"`
	FileSize int64  `gorm:"column:file_size;"`
}

func (TableUserFile) TableName() string {
	return "tbl_user_file"
}

func UpdateFileMeta(fmeta common.FileMeta) (bool, error) {
	newFile := TableFile{
		FileSha1: fmeta.FileSha1,
		FileName: fmeta.FileName,
		FileSize: fmeta.FileSize,
		FileAddr: fmeta.Location,
	}
	saveRes := mysql.DB.Save(&newFile)
	if saveRes.Error != nil {
		return false, saveRes.Error
	}
	return true, nil
}

func CheckFileHashExists(fmeta common.FileMeta) (bool, error) {
	var fileRes TableFile
	result := mysql.DB.Where("file_sha1=?", fmeta.FileSha1).First(&fileRes)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // 未找到记录，返回false但没有错误
		}
		return false, result.Error // 发生了其他错误
	}
	return true, nil // 找到了记录，返回true
}

func OnUserFileUploadFinished(fmeta common.FileMeta, username string) error {
	newTableUserFile := TableUserFile{
		Username: username,
		FileSha1: fmeta.FileSha1,
		FileName: fmeta.FileName,
		FileSize: fmeta.FileSize,
	}
	//resp := mysql.DB.Model(&TableUserFile{}).Where("user_name=?", username).Updates(&newTableUserFile)
	resp := mysql.DB.Save(&newTableUserFile)
	if resp.Error != nil {
		fmt.Println(resp.Error)
		return resp.Error
	}
	return nil

}
