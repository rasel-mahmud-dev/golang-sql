package main

import (
 "database/sql"
 "encoding/json"
 "fmt"
 _ "github.com/go-sql-driver/mysql"
 "github.com/rasel-mahmud-dev/golang-sql/config"
 "net/http"
 "os"
)

func main() {
 
 
 config.Envload()
 
 var HOST string
 var PORT string
 

 
 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "4000"
 }
 
 if HOST = os.Getenv("HOST"); HOST == "" {
  HOST = "127.0.0.1"
 }
 
 
 
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
 })
 
 
 
 
 http.HandleFunc("/sql", func(w http.ResponseWriter, r *http.Request) {
 //
 // // postgres://<user>:<password>@<internal_name>:5432/<database>
 //
 //  // db, err := sql.Open("mysql", config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME)
 //  // db, err := sql.Open("mysql", config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME)
 //  connectionStr := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")
  
  connectionStr := config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@tcp(" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + ")/" + config.DATABASE_NAME
 
  db, err := sql.Open("mysql", connectionStr)
  if err != nil {
   panic(err)
  }
  defer db.Close()
 
  var result int
  db.QueryRow(`SELECT 100 + 1 AS r`).Scan(&result)
 
  fmt.Println(result)
  
  json.NewEncoder(w).Encode(map[string]interface{}{
   "rows": result,
  })
 })
 
 
 fmt.Println(config.DATABASE_USERNAME)
 
 
 err := http.ListenAndServe(HOST+":"+PORT, nil)
 if err != nil {
  return
 }
}
