// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: course.sql

package db

import (
	"context"
	"time"
)

const createCourse = `-- name: CreateCourse :one
insert into course(
    name,
    description,
    start_date,
    end_date
)values(
    $1,$2,$3,$4
)returning id, name, description, attendants, start_date, end_date
`

type CreateCourseParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse,
		arg.Name,
		arg.Description,
		arg.StartDate,
		arg.EndDate,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Attendants,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
delete from course
where id=$1
`

func (q *Queries) DeleteCourse(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

const getCourse = `-- name: GetCourse :one
select id, name, description, attendants, start_date, end_date from course
where id=$1 limit 1
`

func (q *Queries) GetCourse(ctx context.Context, id int64) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, id)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Attendants,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const updateCourse = `-- name: UpdateCourse :one
update course
set name=$2,
    description=$3,
    start_date=$4,
    end_date=$5,
    attendants=$6
where id=$1
returning id, name, description, attendants, start_date, end_date
`

type UpdateCourseParams struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Attendants  int64     `json:"attendants"`
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, updateCourse,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.StartDate,
		arg.EndDate,
		arg.Attendants,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Attendants,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}
