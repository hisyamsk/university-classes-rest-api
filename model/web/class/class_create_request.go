package class

type ClassCreateRequest struct {
	Name    string `json:"name"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}
