package class

type ClassUpdateRequest struct {
	Id      int    `validate:"required" json:"int"`
	Name    string `validate:"required,min=1,max=100" json:"name"`
	StartAt string `validate:"required,datetime=15:04:05" json:"startAt"`
	EndAt   string `validate:"required,datetime=15:04:05" json:"endAt"`
}
