package db

import (
	"log"

	"backend/util"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func ConnectDatabase() (*r.Session, error) {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	session, err := r.Connect(r.ConnectOpts{
		Address: config.Host,

		// todo: add some extra info
		Database: config.DatabaseName,
		// Username: config.UserName,
		// Password: config.Password,
		// InitialCap: config.InitailCap,
		// MaxOpen: config.MaxOpen,
		// Timeout: config.TimeOut*time.Second,
		// ReadTimeout: config.ReadTimeOut*time.Second,
		// WriteTimeout: config.WriteTimeOut*time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}

	r.TableCreate(config.DatabaseName).RunWrite(session)

	return session, nil
}
