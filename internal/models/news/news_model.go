package news

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

type Request struct {
	Company  string `validate:"required"`
	SortBy   string `validate:"omitempty,oneof=relevancy date rank"`
	Page     string `validate:"omitempty,numeric,min=1,max=100"`
	PageSize string `validate:"omitempty,numeric,min=1,max=100"`
	From     string `validate:"datetime"`
	To       string `validate:"datetime"`
}
