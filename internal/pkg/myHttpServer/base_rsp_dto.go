package myHttpServer

type BaseRspDto struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// PageResponseDto ..
type PageResponseDto struct {
	Total    int64 `json:"total"`
	PageSize int   `json:"pageSize"`
	List     any   `json:"list"`
}

// 返回文件字节流
type FileContentRspDto struct {
	FileAbsPath string
}

// 返回字节流
type InlineContentRspDto struct {
	Content []byte
}

// 返回字节流，并指示客户端提供用户“另存为”操作
type AttachmentContentRspDto struct {
	Name    string
	Content []byte
}
