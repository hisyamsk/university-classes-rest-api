package student

type StudentCreateRequest struct {
	Name     string `validate:"required,min=5,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Active   bool   `json:"active"`
	Semester int    `validate:"required,min=1,max=16" json:"semester"`
}
