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

//type LocalTime time.Time
//
//func (t *LocalTime) MarshalJSON() ([]byte, error) {
//	tTime := time.Time(*t)
//	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
//}
//
//type TableUserFile struct {
//	Username   string     `gorm:"column:user_name;" json:"user_name"`
//	FileSha1   string     `gorm:"column:file_sha1;" json:"file_sha"`
//	FileName   string     `gorm:"column:file_name;" json:"file_name"`
//	FileSize   int64      `gorm:"column:file_size;" json:"file_size"`
//	LastUpdate *LocalTime `gorm:"column:upload_at" json:"last_update"`
//}
//
//func (TableUserFile) TableName() string {
//	return "tbl_user_file"
//}

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
	newTableUserFile := common.TableUserFile{
		Username: username,
		FileSha1: fmeta.FileSha1,
		FileName: fmeta.FileName,
		FileSize: fmeta.FileSize,
	}
	//resp := mysql.DB.Model(&TableUserFile{}).Where("user_name=?", username).Updates(&newTableUserFile)
	resp := mysql.DB.Omit("upload_at").Save(&newTableUserFile)
	if resp.Error != nil {
		fmt.Println(resp.Error)
		return resp.Error
	}
	return nil

}

func GetAllFileList(username string) ([]common.TableUserFile, error) {
	var fileList []common.TableUserFile
	find := mysql.DB.Where("user_name=?", username).Find(&fileList)
	if find.Error != nil {
		fmt.Println(find.Error)
		return nil, find.Error
	}
	return fileList, nil
}

func DeleteUserFile(tx *gorm.DB, username, filehash string) error {
	err := tx.Where("user_name = ? AND file_sha1 = ?", username, filehash).Delete(&common.TableUserFile{}).Error
	if err != nil {
		return err
	}
	return nil
}

func SelectTableUserFile(tx *gorm.DB, username, filehash string) (common.TableUserFile, error) {
	file := common.TableUserFile{}
	result := tx.Where("user_name = ? AND file_sha1 = ?", username, filehash).First(&file)
	if result.Error != nil {
		return file, result.Error // 发生了其他错误
	}
	return file, nil

}
