package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	log "go.uber.org/zap"
)

// CreateMysqlConnection return db connection instance
func CreateMysqlConnection(opts map[string]string) (*sql.DB, error) {
	port, err := strconv.Atoi(opts["port"])
	if err != nil {
		log.S().Fatal("Invalid port number : ", opts["port"])
	}

	conf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		opts["user"], opts["password"], opts["host"], port, opts["dbname"])

	db, err := sql.Open("mysql", conf)
	if err != nil {
		log.S().Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.S().Fatal(err)
	}

	log.S().Info("Connected to MYSQL DB Server: ", opts["host"], " at port:", opts["port"], " successfully!")

	return db, nil
}
