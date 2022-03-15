package webfw

import "io"

type BaseRspDto struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// PageResponseDto ..
type PageResponseDto struct {
	Total    int64       `json:"total"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
}

// ByteStreamRspDto ..
type ByteStreamRspDto struct {
	ContentLength int64
	ContentType   string
	ExtraHeaders  map[string]string
	Reader        io.Reader
}

// FileStreamRspDto ..
type FileStreamRspDto struct {
	Headers  map[string]string
	Path     string
	FileName string
}
