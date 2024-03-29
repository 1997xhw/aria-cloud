package main

import (
	"aria-cloud/databases/mysql"
	router2 "aria-cloud/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//ini "aria-cloud/lib"
	//"fmt"
	"log"
)

func main() {
	// 静态资源处理
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(assets.AssetFS())))

	//http.Handle("/static/",
	//	http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mysql.InitDB()
	defer mysql.DB.Close()

	r := router2.InitRouter()
	router := gin.Default()
	router.Use(cors.Default())
	//r.LoadHTMLGlob("./static/view/*")
	//r.Static("/static", "./static")
	if err := r.Run(":80"); err != nil {
		log.Fatal("服务器启动失败...")
	}

	////router()
	//conf := ini.LoadServerConfig()
	//err := http.ListenAndServe(":"+strconv.Itoa(conf.HTTPPort), nil)
	//if err != nil {
	//	fmt.Printf("Fail to start server, err:%v\n", err.Error())
	//}
}

func router() {
	//// 文件存取接口
	//http.HandleFunc("/file/upload", handler.HTTPInterceptor(handler.UploadHandler))
	//http.HandleFunc("/file/upload/suc", handler.HTTPInterceptor(handler.UploadSucHandler))
	//http.HandleFunc("/file/meta", handler.HTTPInterceptor(handler.GetFileMetaHandler))
	//http.HandleFunc("/file/query", handler.HTTPInterceptor(handler.FileQueryHandler))
	//http.HandleFunc("/file/download", handler.HTTPInterceptor(handler.DownloadHandler))
	//http.HandleFunc("/file/delete", handler.HTTPInterceptor(handler.FileDeleteHandler))
	//http.HandleFunc("/file/update", handler.HTTPInterceptor(handler.FileUpdateMetaHandler))
	//
	//// 秒传接口
	//http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(handler.TryFastUploadHandler))
	//
	//// 分块上传接口
	//http.HandleFunc("/file/mpupload/init",
	//	handler.HTTPInterceptor(handler.InitialMultipartUploadHandler))
	//http.HandleFunc("/file/mpupload/uppart",
	//	handler.HTTPInterceptor(handler.UploadPartHandler))
	//http.HandleFunc("/file/mpupload/complete",
	//	handler.HTTPInterceptor(handler.CompleteUploadHandler))
	//
	//// 用户相关
	//http.HandleFunc("/user/signup", handler.SignupHandler)
	//http.HandleFunc("/user/signin", handler.SignInHandler)
	//http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))
}
