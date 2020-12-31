package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("Program Starting")
	log.AddHook(logruseq.NewSeqHook("https://atlas-seq.azurewebsites.net"))

	//log.Println("Hello world!")

	//log.WithFields(log.Fields{
	//		"animal": "walrus",
	//	}).Info("A walrus appears")
	//
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sqlStatement := "INSERT INTO \"Aircraft\" (\"Id\",\"TailNumber\") VALUES (5, 'N3941FL')"
		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)

}
