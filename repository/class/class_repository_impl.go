package class

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
)

type ClassRepositoryImpl struct {
}

func NewClassRepositoryImpl() *ClassRepositoryImpl {
	return &ClassRepositoryImpl{}
}

func (repository *ClassRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, class *entity.Class) *entity.Class {
	query := "INSERT INTO class(name, start_at, end_at) VALUES($1, $2, $3)"
	result, err := tx.ExecContext(ctx, query, class.Name, class.StartAt, class.EndAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	class.Id = int(id)
	return class
}

func (repository *ClassRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, class *entity.Class) *entity.Class {
	query := "UPDATE class SET name = $1, start_at = $2, end_at = $3 WHERE id = $1"
	result, err := tx.ExecContext(ctx, query, class.Name, class.StartAt, class.EndAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	class.Id = int(id)
	return class
}

func (repository *ClassRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, classId int) {
	query := "DELETE FROM class WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, classId)
	helper.PanicIfError(err)
}

func (repository *ClassRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, classId int) (*entity.Class, error) {
	query := "SELECT id, name, start_at, end_at FROM class WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, classId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		class := &entity.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		helper.PanicIfError(err)
		return class, nil
	}

	return nil, fmt.Errorf("Class with id: %d was not found", classId)
}

func (repository *ClassRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, class *entity.Class) []*entity.Class {
	query := "SELECT id, name, start_at, end_at FROM class"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var classes []*entity.Class
	for rows.Next() {
		class := &entity.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		helper.PanicIfError(err)

		classes = append(classes, class)
	}

	return classes
}
