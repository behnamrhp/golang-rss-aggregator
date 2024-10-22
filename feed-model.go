package main

import (
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

type FeedModel struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
}

func feedDbToModelDto(dbFeed database.Feed) FeedModel {
	return FeedModel{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
	}
}
