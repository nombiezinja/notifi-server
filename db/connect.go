package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nombiezinja/notifi-server/config"
	"github.com/nombiezinja/notifi-server/errors"
)

var DB *sql.DB

func portnum() int {
	port, err := strconv.Atoi(config.AppConfig.DbPort)
	errors.FailOnErr(err, "DbPort in config.json cannot be converted to an integer.")
	return port
}

func Connect() {
	var err error
	dbstring := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.AppConfig.DbHost,
		portnum(),
		config.AppConfig.UserName,
		config.AppConfig.Password,
		config.AppConfig.DbName,
	)

	DB, err = sql.Open("postgres", dbstring)
	errors.FailOnErr(err, "Failed to open database connection.")

	err = DB.Ping()
	errors.FailOnErr(err, "Failed to ping database.")

	fmt.Println("Db connection established")
}
