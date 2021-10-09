package models

type Article struct {
	Title         string `json:"title"`
	PublishedDate string `json:"published_date"`
	Url           string `json:"link"`
	Excerpt       string `json:"excerpt"`
	Summary       string `json:"summary"`
}

type News struct {
	TotalHits  string    `json:"total_hits"`
	Page       string    `json:"page"`
	TotalPages string    `json:"total_pages"`
	PageSize   string    `json:"page_size"`
	Articles   []Article `json:"articles"`
}
