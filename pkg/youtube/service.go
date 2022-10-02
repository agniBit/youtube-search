package youtube

import (
	"time"

	"github.com/agniBit/youtube-search/type/common"
	youtubeType "github.com/agniBit/youtube-search/type/youtube"
)

type (
	Repository interface {
		FindVideosByVideoName(search *youtubeType.SearchFilter, offsetLimit *common.OffsetLimit) ([]*YoutubeVideo, error)
	}

	YoutubeVideo struct {
		ID           string    `gorm:"default:('yvd_' || generate_uid(12))"`
		Title        string    `gorm:"not null"`
		Description  string    `gorm:"not null"`
		ThumbnailURL string    `gorm:"not null"`
		URL          string    `gorm:"not null"`
		PublishedAt  time.Time `gorm:"not null"`
	}
)

func (YoutubeVideo) TableName() string {
	return "youtube.youtube_video"
}
