package main

import (
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

type FeedFollowModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func feedFollowDbToModelDto(dbFeedFollow database.FeedFollow) FeedFollowModel {
	return FeedFollowModel{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedId:    dbFeedFollow.FeedID,
	}
}

func feedFollowsDbToModelDto(dbFeedFollows []database.FeedFollow) []FeedFollowModel {
	feedFollows := []FeedFollowModel{}

	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, feedFollowDbToModelDto(feedFollow))
	}

	return feedFollows
}
