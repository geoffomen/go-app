package vo

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
