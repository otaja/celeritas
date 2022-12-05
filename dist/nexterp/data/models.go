package data

import (
	"database/sql"
	"fmt"
	"os"

	orm "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper orm.Session

type Models struct {
	// any Models inserted here (an in the New Function)
	// are easily accessible througout the entire application

}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(databasePool)

	case "postgres", "postgresql":
		upper, _ = postgresql.New(databasePool)

	default:
		// Do nothing in case no database is provided
	}

	return Models{}
}

// getInsertID returns the integer value of a newly inserted id (using upper)
func getInsertID(i orm.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
