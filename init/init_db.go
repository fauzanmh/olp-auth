package init

import (
	"database/sql"
	"os"

	"github.com/fauzanmh/olp-auth/internal/db"
)

// ConnectToMysqlServer is a function to init PostgreSQL connection
func ConnectToMysqlServer(cfg *Config) (*sql.DB, error) {
	db, err := db.CreateMysqlConnection(map[string]string{
		"host":     cfg.Database.Mysql.Host,
		"port":     cfg.Database.Mysql.Port,
		"user":     cfg.Database.Mysql.User,
		"password": cfg.Database.Mysql.Password,
		"dbname":   cfg.Database.Mysql.Dbname,
	})

	if err != nil {
		os.Exit(1)
	}

	return db, err
}
