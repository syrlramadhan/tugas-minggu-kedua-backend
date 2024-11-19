package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

/*
OpenConnectionPostgresSQL

	fungsi untuk membuka open koneksi ke database postgresql.
*/
func OpenConnectionPostgresSQL() (*sql.DB, error) {

	// deklarasi variabel yang dibutuhkan untuk koneksi database
	host := "localhost"    //os.Getenv("localhost")
	port := "5432"         //os.Getenv("5432")
	user := "postgres"     //os.Getenv("postgres")
	password := "postgres" //os.Getenv("postgres")
	dbname := "postgres" //os.Getenv("user")

	psqlMerge := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// buat koneksi ke database.
	dbConnection, err := sql.Open("postgres", psqlMerge)
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil, err
	}

	// ping ke database
	err = dbConnection.Ping()
	if err != nil {
		fmt.Println("Error pinging database")
		return nil, err
	}

	return dbConnection, nil
}
