package gapi

import (
	db "main/db/sqlc"
	"main/pb"
)

type Server struct {
	pb.UnimplementedCoursesServer
	store  *db.Store
}


func NewServer(db *db.Store) *Server {
	server := &Server{store: db}
	return server
}


