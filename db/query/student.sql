-- name: GetStudent :one
select * from student
where id=$1 limit 1;

-- name: CreateStudent :one
insert into student(
    name,
    password,
    school_id
)values(
   $1,$2,$3
)returning *;

-- name: UpdateStudent :one
update student
set name=$2,password=$3
where id=$1
returning *;

-- name: DeleteStudent :exec
delete from student
where id=$1;
