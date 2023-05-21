package student

type StudentCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Semester int    `json:"semester"`
}
