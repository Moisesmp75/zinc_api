package entities

type Hits[T any] struct {
	Total    Total    `json:"total"`
	MaxScore float64  `json:"max_score"`
	Hits     []Hit[T] `json:"hits"`
}
