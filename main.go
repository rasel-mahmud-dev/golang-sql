package main

import (
 "database/sql"
 "github.com/go-sql-driver/mysql"
 
 // "encoding/json"
 "fmt"
 _ "github.com/go-sql-driver/mysql"
 "github.com/rasel-mahmud-dev/golang-sql/config"
 "net/http"
 "os"
)

var connectionStr string
//
// func InitialCreateAllTable() {
//  db, err := sql.Open("mysql", connectionStr)
//  if err != nil {
//   panic(err)
//  }
//  defer db.Close()
//
//  userTableSql := `
// 		CREATE TABLE IF NOT EXISTS users (
// 		  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
// 		  username VARCHAR(40) NOT NULL,
// 		  email VARCHAR(50) NOT NULL UNIQUE,
// 		  password VARCHAR(100) NOT NULL,
// 		  avatar VARCHAR(100) NULL DEFAULT NULL,
// 		  created_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
// 		  role ENUM('admin', 'user') NOT NULL DEFAULT 'user'
// 		)`
//
//  res, err := db.Exec(userTableSql)
//  if err != nil {
//   fmt.Println(err)
//  } else {
//   fmt.Println("Users Table Created ")
//   _ = res
//  }
//
//  res, err = db.Exec(`
// 		CREATE TABLE IF NOT EXISTS products (
// 		  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
// 		  title VARCHAR(255) NOT NULL,
// 		  description LONGTEXT NULL,
// 		  price DOUBLE,
// 		  author_id INT NOT NULL,
// 		  brand_id INT DEFAULT 0,
// 			cover VARCHAR(255) DEFAULT '',
// 			tags json DEFAULT NULL,
// 			attributes json DEFAULT NULL,
// 		  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
// 		)`)
//  if err != nil {
//   fmt.Println(err, "-----")
//  } else {
//   fmt.Println("Products Table Created ")
//  }
//
//  res, err = db.Exec(`
// 		CREATE TABLE IF NOT EXISTS brands (
// 		  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
// 		  name VARCHAR(255) NOT NULL,
// 		  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
// 		)`)
//  if err != nil {
//   fmt.Println(err, "-----")
//  } else {
//   fmt.Println("Brands Table Created ")
//  }
// }
//
// func InitialData()  {
//  db, err := sql.Open("mysql", connectionStr)
//  if err != nil {
//   panic(err)
//  }
//  defer db.Close()
//
//  _, err = db.Exec(`
// 			INSERT INTO users(username, email, role, password, avatar)
// 			VALUES
// 			(
// 				'rasel',
// 				'rasel@gmail.com',
// 				'admin',
// 				'$2a$08$Tj9kveUf9GzZCn4l70wqdeh8Uq.FK2T6sYXcGGvtqAkFUWPdU3DZO',
// 				'https://res.cloudinary.com/dbuvg9h1e/image/upload/v1639027976/my-avatar-300x300.jpg'
// 			),
// 			(
// 				'raju',
// 				'raju@gmail.com',
// 				'user',
// 				'$2a$08$Tj9kveUf9GzZCn4l70wqdeh8Uq.FK2T6sYXcGGvtqAkFUWPdU3DZO',
// 				'https://res.cloudinary.com/dbuvg9h1e/image/upload/v1639027976/my-avatar-300x300.jpg'
// 			)
// 		`)
//  if err != nil {
//   fmt.Println(err)
//  } else {
//   fmt.Println("Users Inserted ")
//  }
//
//  _, err = db.Exec(`INSERT INTO products(title, description, author_id, brand_id, price, cover, tags, attributes )
// 		VALUES (
// 			'Redmi Note 11 Pro',
// 			'Deno is javascript secure Runtime Engine. !!Deno is a secure and fast ',
// 		    1,
// 		    1,
// 				19000.00,
//         'images/products/huawei-enjoy-10-plus.jpg',
// 			'["Redmi Note 11", "Redmi Note 11s", "Redmi Note 11 pro"]',
// 			'{ "ram": 4, "rom": 64, "battery": 5000, "color": ["red", "blue", "pink", "gray"]}'
// 		), (
// 			'Redmi Note 11s',
// 			'Deno is javascript secure Runtime Engine. !!Deno is a secure and fast ',
// 		    1,
// 		    1,
// 				19000.00,
//         'images/products/huawei-enjoy-10-plus.jpg',
// 			'["Redmi Note 11", "Redmi Note 11s", "Redmi Note 11 pro"]',
// 			'{ "ram": 4, "rom": 64, "battery": 5000, "color": ["red", "blue", "pink", "gray"]}'
// 		)
// 		`)
//  if err != nil {
//   fmt.Println(err, "-----")
//  } else {
//   fmt.Println("Products Inserted ")
//  }
//
//
//  _, err = db.Exec(`
// 		INSERT INTO brands(name)
// 			VALUES ('Samsung'), ('Nokia'), ('Apple')
// 		`)
//  if err != nil {
//   fmt.Println(err, "-----")
//  } else {
//   fmt.Println("Products Inserted ")
//  }
//
//
// }
//

func main() {
 
 config.Envload()
 
 // connectionStr := config.DATABASE_NAME + ":" + config.DATABASE_PASSWORD  + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME
 
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
 
 http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
 
 
 
 })
 
 
 
 
 http.HandleFunc("/sql", func(w http.ResponseWriter, r *http.Request) {
  
  // InitialCreateAllTable()
  // InitialData()
  //
  cfg := mysql.Config{
   User:   config.DATABASE_USERNAME,
   Passwd: config.DATABASE_PASSWORD,
   Net:    "tcp",
   Addr:   config.DATABASE_HOST + ":" + config.DATABASE_PORT,
   DBName: config.DATABASE_NAME,
  }
 
  db, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
   panic(err)
  }
  defer db.Close()
 
  userTableSql := `
		CREATE TABLE IF NOT EXISTS users (
		  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
		  username VARCHAR(40) NOT NULL,
		  email VARCHAR(50) NOT NULL UNIQUE,
		  password VARCHAR(100) NOT NULL,
		  avatar VARCHAR(100) NULL DEFAULT NULL,
		  created_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
		  role ENUM('admin', 'user') NOT NULL DEFAULT 'user'
		)`
  _, err = db.Exec(userTableSql)
  if err != nil {
   fmt.Fprint(w, err)
   return
  }
 
  _, err = db.Exec(`
			INSERT INTO users(username, email, role, password, avatar)
			VALUES
			(
				'raselt',
				'rasel@gmail.com',
				'admin',
				'$2a$08$Tj9kveUf9GzZCn4l70wqdeh8Uq.FK2T6sYXcGGvtqAkFUWPdU3DZO',
				'https://res.cloudinary.com/dbuvg9h1e/image/upload/v1639027976/my-avatar-300x300.jpg'
			),
			(
				'raju',
				'raju@gmail.com',
				'user',
				'$2a$08$Tj9kveUf9GzZCn4l70wqdeh8Uq.FK2T6sYXcGGvtqAkFUWPdU3DZO',
				'https://res.cloudinary.com/dbuvg9h1e/image/upload/v1639027976/my-avatar-300x300.jpg'
			)
		`)
   if err != nil {
    fmt.Println(err)
   } else {
    fmt.Println("Users Inserted ")
   }
  
  
 
 fmt.Println("Craeted")
  fmt.Fprint(w, "Created")
  
  
  // // postgres://<user>:<password>@<internal_name>:5432/<database>
 //
 //  // db, err := sql.Open("mysql", config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME)
 //  // db, err := sql.Open("mysql", config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_HOST + ":" + config.DATABASE_PORT + "/" + config.DATABASE_NAME)
 //  connectionStr := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")
 
 
 
 })
 
 http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
  cfg := mysql.Config{
   User:   config.DATABASE_USERNAME,
   Passwd: config.DATABASE_PASSWORD,
   Net:    "tcp",
   Addr:   config.DATABASE_HOST + ":" + config.DATABASE_PORT,
   DBName: config.DATABASE_NAME,
  }
 
  db, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
   panic(err)
  }
  defer db.Close()
 
  userTableSql := `SELECT username from users`
  
  var usernames []string
  
  rows, err := db.Query(userTableSql)
  if err != nil {
  fmt.Fprint(w, err)
   return
  }
  
  for rows.Next() {
   var usename string
   rows.Scan(&usename)
   usernames = append(usernames, usename)
  }
  
  fmt.Println(usernames)
  
  fmt.Fprint(w, usernames)
 
 })
  
  err := http.ListenAndServe(HOST+":"+PORT, nil)
 if err != nil {
  return
 }
}
