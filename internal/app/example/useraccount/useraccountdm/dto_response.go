package useraccountdm

import (
	"time"
)

type LoginResponseDto struct {
	Uid         int       `json:"uid"`
	IssueAt     time.Time `json:"issueAt"`
	ExpireAt    time.Time `json:"expireAt"`
	TokenType   string    `json:"tokenType"`
	AccessToken string    `json:"accessToken"`
}
