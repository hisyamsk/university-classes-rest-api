package class

type ClassResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}
