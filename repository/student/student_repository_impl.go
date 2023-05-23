package student

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepository() *StudentRepositoryImpl {
	return &StudentRepositoryImpl{}
}

func (repository *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.Student {
	query := "INSERT INTO student(name, email, active, semester) VALUES($1, $2, $3, $4) RETURNING id"
	row := tx.QueryRowContext(ctx, query, student.Name, student.Email, student.Active, student.Semester)
	err := row.Scan(&student.Id)
	helper.PanicIfError(err)

	return student
}

func (repository *StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.Student {
	query := "UPDATE student SET name = $1, email = $2, active = $3, semester = $4 WHERE id = $5 RETURNING id"
	row := tx.QueryRowContext(ctx, query, student.Name, student.Email, student.Active, student.Semester, student.Id)
	err := row.Scan(&student.Id)
	helper.PanicIfError(err)

	return student
}

func (repository *StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, studentId int) {
	query := "DELETE FROM student WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, studentId)
	helper.PanicIfError(err)
}

func (repository *StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, studentId int) (*entity.Student, error) {
	query := "SELECT id, name, email, active, semester FROM student WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := entity.Student{}
	if rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Active, &student.Semester)
		helper.PanicIfError(err)
		return &student, nil
	}

	return &student, fmt.Errorf("Student with id: %d was not found", studentId)
}

func (repository *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*entity.Student {
	query := "SELECT id, name, email, active, semester FROM student"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var students []*entity.Student
	for rows.Next() {
		student := &entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Active, &student.Semester)
		helper.PanicIfError(err)
		students = append(students, student)
	}

	return students
}

func (repository *StudentRepositoryImpl) FindClassesById(ctx context.Context, tx *sql.Tx, studentId int) []*entity.Class {
	query := "SELECT class_id, class.name, class.start_at, class.end_at FROM enrolled_class JOIN class ON class_id = class.id WHERE student_id = $1"
	rows, err := tx.QueryContext(ctx, query, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	enrolledClasses := []*entity.Class{}

	for rows.Next() {
		class := &entity.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		class.StartAt = helper.GetTimeFromString(class.StartAt)
		class.EndAt = helper.GetTimeFromString(class.EndAt)
		helper.PanicIfError(err)
		enrolledClasses = append(enrolledClasses, class)
	}

	return enrolledClasses
}
