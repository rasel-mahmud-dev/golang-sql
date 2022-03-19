package main
import (
 "database/sql"
 "fmt"
 _ "github.com/go-sql-driver/mysql"
 _ "github.com/lib/pq"
 "golang-sql/config"
 "net/http"
 "os"
)

func main() {
 
 
 config.Envload()
 
 var HOST string
 

 
 if HOST = os.Getenv("HOST"); HOST == "" {
  HOST = "127.0.0.1"
 }
 
 
 
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
 })
 
 
 
 
 http.HandleFunc("/sql", func(w http.ResponseWriter, r *http.Request) {

 // postgres://<user>:<password>@<internal_name>:5432/<database>
  
  // db, err := sql.Open("mysql", config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME)
  // db, err := sql.Open("mysql", config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME)
  connectionStr := config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@tcp(" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + ")/" + config.DATABASE_NAME
  fmt.Println(connectionStr)
  
  db, err := sql.Open("mysql", connectionStr)
  if err != nil {
   panic(err)
  }
  defer db.Close()
  
  fmt.Fprintf(w, "DB Connected")
 })
 
 
 err := http.ListenAndServe(HOST+":"+config.PORT, nil)
 if err != nil {
  return
 }
}
