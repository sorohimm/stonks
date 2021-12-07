package news_models

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
	Company  string `json:"q" validate:"required"`
	SortBy   string `json:"sort_by" validate:"omitempty,oneof=relevancy date rank"`
	Page     string `json:"page" validate:"omitempty,numeric,min=1,max=100"`
	PageSize string `json:"page_size" validate:"omitempty,numeric,min=1,max=100"`
	From     string `json:"from" validate:"datetime=2006-01-02"`
	To       string `json:"to" validate:"datetime=2006-01-02"`
}
