package yt_repository

import (
	"github.com/agniBit/youtube-search/pkg/youtube"
	"gorm.io/gorm"
)


type repository struct {
	db *gorm.DB
}

func NewYoutubeRepository(db *gorm.DB) youtube.Repository {
	return &repository{db}
}