package db_client

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitialiseDBConnection() {
	db, err := sqlx.Open("sqlserver", "sqlserver://tis:tanobel@192.168.100.11:1433?database=d_transaksi")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	DBClient = db
}
