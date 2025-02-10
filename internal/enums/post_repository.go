package enums

type SortBy string

const (
	SortByScore SortBy = "score desc"
	SortByDate  SortBy = "modified_at desc"
	DefaultSort SortBy = ""
)
