package models

type Article struct {
	Title         string `json:"title"`
	PublishedDate string `json:"published_date"`
	Url           string `json:"link"`
	Excerpt       string `json:"excerpt"`
	Summary       string `json:"summary"`
}

type News struct {
	TotalHits  int       `json:"total_hits"`
	Page       int       `json:"page"`
	TotalPages int       `json:"total_pages"`
	PageSize   int       `json:"page_size"`
	Articles   []Article `json:"articles"`
}
