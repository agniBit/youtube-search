package yt_repository

import "github.com/agniBit/youtube-search/pkg/youtube"

func (r *repository) FindVideosByVideoName(name string) ([]*youtube.YoutubeVideo, error) {
	var videos []*youtube.YoutubeVideo
	err := r.db.Where("title ILIKE ?", "%"+name+"%").Find(&videos).Error
	return videos, err
}
