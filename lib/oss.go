package lib

import (
	"aria-cloud/common"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"path"
)

func UploadOss(file_meta common.FileMeta) {
	fileSuffix := path.Ext(file_meta.FileName)
	conf := LoadServerConfig()

	// 创建OSSClient实例
	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return
	}
	if err != nil {
		log.Println("创建实例Error: ", err)
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(conf.BucketName)
	if err != nil {
		log.Println("获取存储空间Error: ", err)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile("test/"+file_meta.FileSha1+fileSuffix, file_meta.Location)
	if err != nil {
		log.Println("本地文件上传Error:", err)
		return
	}

}

func DeleteOss(fileHash, fileType string) error {
	conf := LoadServerConfig()
	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	bucket, errBucket := client.Bucket(conf.BucketName)
	if errBucket != nil {
		log.Println("Error: ", err)
		return err
	}

	err = bucket.DeleteObject("test/" + fileHash + fileType)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	return nil

}
