package models

// Subscriber type details
type Subscriber struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
}
