package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/domain"
)

type ClassRepositoryImpl struct {
}

func (repository *ClassRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, class *domain.Class) *domain.Class {
	query := "INSERT INTO class(name, start_at, end_at) VALUES(?, ?, ? )"
	result, err := tx.ExecContext(ctx, query, class.Name, class.StartAt, class.EndAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	class.Id = int(id)
	return class
}

func (repository *ClassRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, class *domain.Class) *domain.Class {
	query := "UPDATE class SET name = ?, start_at = ?, end_at = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, class.Name, class.StartAt, class.EndAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	class.Id = int(id)
	return class
}

func (repository *ClassRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, classId int) {
	query := "DELETE FROM class WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, classId)
	helper.PanicIfError(err)
}

func (repository *ClassRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, classId int) (*domain.Class, error) {
	query := "SELECT id, name, start_at, end_at FROM class WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, classId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		class := &domain.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		helper.PanicIfError(err)
		return class, nil
	}

	return nil, fmt.Errorf("Class with id: %d was not found", classId)
}

func (repository *ClassRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, class *domain.Class) []*domain.Class {
	query := "SELECT id, name, start_at, end_at FROM class"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var classes []*domain.Class
	for rows.Next() {
		class := &domain.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		helper.PanicIfError(err)

		classes = append(classes, class)
	}

	return classes
}
