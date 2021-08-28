package model

type Url struct {
	Id          uint64 `json:"id" redis:"id"`
	OriginalUrl string `json:"original_url" redis:"original_url"`
	Expires     string `json:"expires" redis:"expires"`
	Visits      int    `json:"visits" redis:"visits"`
}
