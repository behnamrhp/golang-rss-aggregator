package main

import (
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `json:"id"`
	Fullname  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func userDbToModelDto(dbUser database.User) UserModel {
	return UserModel{
		ID:        dbUser.ID,
		Fullname:  dbUser.Fullname,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}
