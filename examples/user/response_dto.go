package user

type UserInfoResponseDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}
