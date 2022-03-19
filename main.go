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
 
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", config.DATABASE_NAME, config.DATABASE_PASSWORD, config.DATABASE_HOST))
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
