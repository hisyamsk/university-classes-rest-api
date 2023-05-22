package student

type StudentUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Name     string `validate:"required,min=5,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Active   bool   `validate:"required" json:"active"`
	Semester int    `validate:"required,min=1,max=16" json:"semester"`
}
