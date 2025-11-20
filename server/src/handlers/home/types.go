package handlers_home

type FileInfo struct {
	FileName string `json:"file_name" example:"example.txt" doc:"Uploaded file name"`
	Size     int64  `json:"size" example:"12345" doc:"Uploaded file size in bytes"`
}
