package common

type (
	OffsetLimit struct {
		Offset int `query:"offset"`
		Limit  int `query:"limit"`
	}
)
