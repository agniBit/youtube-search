package youtubeType

type (
	Service interface {
		SearchYoutubeVideosByName(name string) ([]*YoutubeVideo, error)
	}

	YoutubeVideo struct {
		ID           string `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ThumbnailURL string `json:"thumbnail_url"`
		URL          string `json:"url"`
		PublishedAt  string `json:"published_at"`
	}
)
