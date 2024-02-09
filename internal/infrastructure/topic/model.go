package topic

import "time"

type TopicMessage struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	TotalBalance float64   `json:"total_balance"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
