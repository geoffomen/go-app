package account

import "github.com/geoffomen/go-app/internal/pkg/vo"

type LoginResponseDto struct {
	Uid         int       `json:"uid"`
	IssueAt     vo.Mytime `json:"issueAt"`
	ExpireAt    vo.Mytime `json:"expireAt"`
	TokenType   string    `json:"tokenType"`
	AccessToken string    `json:"accessToken"`
}
