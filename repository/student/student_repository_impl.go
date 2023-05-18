package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/domain"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepository() *StudentRepositoryImpl {
	return &StudentRepositoryImpl{}
}

func (repository *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student *domain.Student) *domain.Student {
	query := "INSERT INTO student(name, email, active, semester) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, student.Name, student.Email, student.Active, student.Semester)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	student.Id = int(id)
	return student
}
func (repository *StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, student *domain.Student) *domain.Student {
	query := "UPDATE student SET name = ?, email = ?, active = ?, semester = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, student.Name, student.Email, student.Active, student.Semester, student.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	student.Id = int(id)
	return student
}

func (repository *StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, studentId int) {
	query := "DELETE FROM student WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, studentId)
	helper.PanicIfError(err)
}

func (repository *StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, studentId int) (*domain.Student, error) {
	query := "SELECT id, name, email, active, semester FROM student WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := domain.Student{}
	if rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Active, &student.Semester)
		helper.PanicIfError(err)
		return &student, nil
	}

	return nil, fmt.Errorf("Student with id: %d not found", studentId)
}

func (repository *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, student *domain.Student) []*domain.Student {
	query := "SELECT id, name, email, active, semester FROM student"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var students []*domain.Student
	for rows.Next() {
		student := &domain.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Active, &student.Semester)
		helper.PanicIfError(err)
		students = append(students, student)
	}

	return students
}
