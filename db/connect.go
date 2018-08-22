package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"bitbucket.org/nombiezinja/deepfrier/util"
	_ "github.com/lib/pq"
	"github.com/nombiezinja/notifi-server/config"
)

var DB *sql.DB

func portnum() int {
	port, err := strconv.Atoi(config.AppConfig.DbPort)
	util.CheckErr(err)
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
	util.CheckErr(err)

	err = DB.Ping()
	util.CheckErr(err)

	fmt.Println("Db connection established")
}
