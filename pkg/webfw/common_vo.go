package webfw

// Validate 用于参数校验。只要请求的DTO实现这个接口，就会在绑定完参数后调用。
type Validate interface {
	Validate() ([]string, error)
}

// SessionInfo 记录会话信息
type SessionInfo struct {
	Uid           int    `json:"uid"`
	Name          string `json:"nickName"`
	Token         string `json:"token"`
	TokenExpireAt int64  `json:"tokenExpireAt"`
}
