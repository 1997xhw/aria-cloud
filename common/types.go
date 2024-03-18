package common

import (
	"fmt"
	"time"
)

type User struct {
	username  string
	usertoken string
}

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

type TableUserFile struct {
	Username   string     `gorm:"column:user_name;" json:"user_name"`
	FileSha1   string     `gorm:"column:file_sha1;" json:"file_sha"`
	FileName   string     `gorm:"column:file_name;" json:"file_name"`
	FileSize   int64      `gorm:"column:file_size;" json:"file_size"`
	LastUpdate *LocalTime `gorm:"column:upload_at" json:"last_update"`
}

func (TableUserFile) TableName() string {
	return "tbl_user_file"
}
