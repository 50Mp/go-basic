package server

type (
	server struct{}
)

func Server() *server {
	return &server{}
}
