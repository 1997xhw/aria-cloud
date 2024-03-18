package controllers

import (
	"aria-cloud/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AllFileList(c *gin.Context) {
	username := c.Query("username")
	fileList, err := services.GetAllFileList(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": fileList,
	})

}
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		panic(fmt.Sprintf("file参数不能为空"))
	}

	username := c.PostForm("username")

	log.Printf("%v", file.Filename)
	err = services.SaveFileHandler(file, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

	//data := struct {
	//	Filename string `json:"filename"`
	//	Filesize string `json:"filesize"`
	//}{
	//	Filename: file.Filename,
	//	Filesize: strconv.FormatInt(file.Size, 10),
	//}

}
func DeleteFile(c *gin.Context) {
	username := c.PostForm("username")
	filehash := c.PostForm("filehash")
	err := services.DeleteFileHandler(username, filehash)
	fmt.Println(username, filehash)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}

}
