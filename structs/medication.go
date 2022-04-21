package structs

type Medication struct {
	Id        int64  `json:"id"`
	Name string `json:"name"`
	Weight  string `json:"weight"`
	Code     string `json:"code"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
}