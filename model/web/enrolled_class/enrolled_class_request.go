package enrolled_class

type EnrolledClassRequest struct {
	Id        int `json:"id"`
	StudentId int `validate:"required" json:"studentId"`
	ClassId   int `validate:"required" json:"classId"`
}
