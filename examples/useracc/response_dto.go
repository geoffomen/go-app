package useracc

import "github.com/storm-5/go-app/pkg/database"

type LoginResponseDto struct {
	Uid         int             `json:"uid"`
	IssueAt     database.Mytime `json:"issueAt"`
	ExpireAt    database.Mytime `json:"expireAt"`
	TokenType   string          `json:"tokenType"`
	AccessToken string          `json:"accessToken"`
}
