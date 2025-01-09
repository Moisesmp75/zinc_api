package entities

type Hit[T any] struct {
	Index     Index   `json:"_index"`
	Type      Type    `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	Timestamp string  `json:"@timestamp"`
	Source    T       `json:"_source"`
}
