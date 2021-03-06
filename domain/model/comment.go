package model

import (
	"time"
)

type Comment struct {
	Id        int       `json:"id" gorm:"primaryKey;not null"`
	CreatedAt time.Time `json:"created_at"`
	ThreadID  int       `json:"thread_id"`
	UserID    string    `json:"uid"`
	User      *User
	Body      string         `json:"body"`
	VoteCnt   int            `json:"vote_cnt"`
	Vote      []*VoteComment `gorm:"constraint:OnDelete:CASCADE"`
}
