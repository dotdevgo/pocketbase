package main

import (
	"dotdev/pocketbase"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err.Error())
	}
}

// var turso = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MzQ2MzY5MzUsImlkIjoiZWFkMjA4NjMtMDBkOS00OGRhLWFhMTYtYmZiNmYyMzc3NjkxIn0.iQ-odyf5kQxSr_jjzY_6vUazj27v8M_MprwLfBA56A0m7mXoj9r-n6r48HIvBSV_FvVpdv_gdOWoPyDm28ilCw"

// "github.com/pocketbase/dbx"
// pb "github.com/pocketbase/pocketbase"
// "github.com/pocketbase/pocketbase/core"

// app := pocketbase.NewWithConfig(pb.Config{
// 	DBConnect: func(dbPath string) (*dbx.DB, error) {
// 		if !strings.Contains(dbPath, "data.db") {
// 			return core.DefaultDBConnect(dbPath)
// 		}

// 		dsn := fmt.Sprintf("libsql://website-dotdev.aws-eu-west-3.turso.io?authToken=%s", turso)

// 		return dbx.Open("libsql", dsn)
// 	},
// })
// if !strings.Contains(dbPath, "data.db") {
// 	return core.DefaultDBConnect(dbPath)
// }

// dsn := "root:temp123@/pocketbase"

// fmt.Println("DB: Connect", dsn)

// return dbx.Open("mysql", dsn)
