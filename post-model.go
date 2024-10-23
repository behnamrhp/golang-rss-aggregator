package main

import (
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

type PostModel struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedId      uuid.UUID `json:"feed_id"`
}

func postDbToModelDto(dbPost database.Post) PostModel {
	return PostModel{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Url:         dbPost.Url,
		Description: dbPost.Description.String,
		PublishedAt: dbPost.PublishedAt,
		FeedId:      dbPost.FeedID,
	}
}

func postsDbToModelDto(dbPosts []database.Post) []PostModel {
	posts := []PostModel{}

	for _, dbPost := range dbPosts {
		posts = append(posts, postDbToModelDto(dbPost))
	}
	return posts
}
