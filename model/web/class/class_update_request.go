package class

type ClassUpdateRequest struct {
	Id      int    `json:"int"`
	Name    string `json:"name"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}
