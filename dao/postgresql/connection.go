package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/daniel/apiRest2/utilities"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {

	config, err := utilities.GetConfiguration()

	if err != nil {
		log.Fatal(err)
	}
	//"postgres://danielfelipe:123456@localhost:5432/blog?ssl_mode=disable"

	stringConnection := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Engine,
		config.User,
		config.Password,
		config.Server,
		config.Port,
		config.Database,
	)

	db, err := sql.Open("postgres", stringConnection)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db

}
