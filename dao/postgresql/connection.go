package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/daniel/basic-project-with-golang/utilities"
	_ "github.com/lib/pq"
)

var (
	connect sync.Once
	db      *sql.DB
)

func GetConnection() *sql.DB {

	connect.Do(func() {
		config, err := utilities.GetConfiguration()

		if err != nil {
			fmt.Println("----ERROR 1---------------")
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

		db, err = sql.Open("postgres", stringConnection)

		if err != nil {
			fmt.Println("----ERROR 2---------------")
			log.Fatal(err)
		}

		err = db.Ping()

		if err != nil {
			fmt.Println("----ERROR 3---------------")
			log.Fatal(err)
		}
	})

	return db

}
