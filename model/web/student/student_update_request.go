package student

type StudentUpdateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Semester int    `json:"semester"`
}
