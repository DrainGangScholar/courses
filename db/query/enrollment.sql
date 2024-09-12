-- name: CreateEnrollment :one
insert into enrollment(
    student_id,
    course_id
)values(
    $1,
    $2
)returning *;

-- name: UpdateEnrollment :one
update enrollment
    set student_id=$2,
        course_id=$3
where id=$1
returning *;

-- name: GetEnrollment :one
select * from enrollment
where id=$1 limit 1;

-- name: DeleteEnrollment :exec
delete from enrollment
where id=$1;
