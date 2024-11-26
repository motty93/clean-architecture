package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	SlackID string `json:"slack_id"`
}
