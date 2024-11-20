package models

// SearchModel represents a search query with column, condition, and value.
type SearchModel struct {
	Column    string `json:"column,omitempty"`
	Condition string `json:"condition,omitempty"`
	Value     string `json:"value,omitempty"`
}

// SortByModel represents a sorting query with column and order.
type SortByModel struct {
	Column string `json:"column,omitempty"`
	Order  string `json:"order,omitempty" default:"asc"`
}

// DBFilterModel represents a database filter with limit, offset, sort by, and search.
type DBFilterModel struct {
	Limit  int32       `json:"limit,omitempty"`
	Offset int32       `json:"offset,omitempty"`
	SortBy SortByModel `json:"sort_by,omitempty"`
	Search SearchModel `json:"search,omitempty"`
}