-- name: CreateCourse :one
insert into course(
    name,
    description,
    start_date,
    end_date
)values(
    $1,$2,$3,$4
)returning *;

-- name: UpdateCourse :one
update course
set name=$2,
    description=$3,
    start_date=$4,
    end_date=$5,
    attendants=$6
where id=$1
returning *;

-- name: GetCourse :one
select * from course
where id=$1 limit 1;

-- name: DeleteCourse :exec
delete from course
where id=$1;
