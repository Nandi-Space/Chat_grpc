package model

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	handler *r.Session
)

func SetRepo(s *r.Session) {
	handler = s
}

func CloseDB() {
	handler.Close()
}
