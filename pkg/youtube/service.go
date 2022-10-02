package youtube

type (
	YoutubeVideo struct {
		ID           string `gorm:"default:('yvd_' || generate_uid(12))"`
		Title        string `gorm:"not null"`
		Description  string `gorm:"not null"`
		ThumbnailURL string `gorm:"not null"`
		URL          string `gorm:"not null"`
		PublishedAt  string `gorm:"not null"`
	}
)

func (YoutubeVideo) TableName() string {
	return "youtube.youtube_video"
}
