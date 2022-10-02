package yt_repository

import (
	"context"

	"github.com/agniBit/youtube-search/pkg/youtube"
	"github.com/agniBit/youtube-search/type/common"
	youtubeType "github.com/agniBit/youtube-search/type/youtube"
	"github.com/agniBit/youtube-search/utl/storage/postgres"
)

func (r *repository) FindVideosByVideoName(search *youtubeType.SearchFilter, offsetLimit *common.OffsetLimit) ([]*youtube.YoutubeVideo, error) {
	var videos []*youtube.YoutubeVideo
	q := r.db.Model(&youtube.YoutubeVideo{})

	if search.Title != "" {
		q = q.Where("title ILIKE ?", "%"+search.Title+"%")
	}

	if search.Description != "" {
		q = q.Where("description ILIKE ?", "%"+search.Description+"%")
	}

	if search.Search != "" {
		q = q.Where("title ILIKE ? OR description ILIKE ?", "%"+search.Search+"%", "%"+search.Search+"%")
	}

	// use offset and limit
	postgres.Pagination(q, offsetLimit)

	err := q.Find(&videos).Order("published_at DESC").Error
	return videos, err
}

func (r *repository) SaveYoutubeVideos(ctx context.Context, videos []*youtube.YoutubeVideo) error {
	return r.db.WithContext(ctx).Create(&videos).Error
}
