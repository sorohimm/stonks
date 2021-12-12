package news_models

type Article struct {
	Title         string `json:"title,omitempty"`
	PublishedDate string `json:"published_date,omitempty"`
	Url           string `json:"link,omitempty"`
	Excerpt       string `json:"excerpt,omitempty"`
	Summary       string `json:"summary,omitempty"`
}

type News struct {
	TotalHits  int       `json:"total_hits,omitempty"`
	Page       int       `json:"page,omitempty"`
	TotalPages int       `json:"total_pages,omitempty"`
	PageSize   int       `json:"page_size,omitempty"`
	Articles   []Article `json:"articles,omitempty"`
}

type Request struct {
	Company  string `json:"q" validate:"required"`
	SortBy   string `json:"sort_by" validate:"omitempty,oneof=relevancy date rank"`
	Page     string `json:"page" validate:"omitempty,numeric,min=1,max=100"`
	PageSize string `json:"page_size" validate:"omitempty,numeric,min=1,max=100"`
	From     string `json:"from" validate:"omitempty,datetime=2006-01-02"`
	To       string `json:"to" validate:"omitempty,datetime=2006-01-02"`
}
