-- name: GetSchool :one
select * from school s
where s.id=$1 limit 1;

-- name: ListSchools :many
select * from school s
order by s.id
limit $1
offset $2;

-- name: CreateSchool :one
insert into school(
    name,
    address,
    type
)values(
    $1,$2,$3
) returning *;

-- name: DeleteSchool :exec
delete from school
where id=$1;

-- name: UpdateSchool :one
update school 
set name=$2,address=$3,type=$4
where id=$1
returning *;
