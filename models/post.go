package model

// Subscriber type details
type Subscriber struct {
	ID      int64  `json:"id"`
	Username   string `json:"username"`
	Domain string `json:"domain"`
	Password string `json:"password"`
	Ha1 string `json:"ha1"`
	Ha1b string `json:"ha1b"`
}
