package useracc

import "github.com/geoffomen/go-app/pkg/database"

type LoginResponseDto struct {
	Uid         int             `json:"uid"`
	IssueAt     database.Mytime `json:"issueAt"`
	ExpireAt    database.Mytime `json:"expireAt"`
	TokenType   string          `json:"tokenType"`
	AccessToken string          `json:"accessToken"`
}
