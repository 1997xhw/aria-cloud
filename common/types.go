package common

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
