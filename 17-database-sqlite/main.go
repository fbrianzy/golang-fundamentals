package main

import (
    "database/sql"
    "fmt"
    _ "modernc.org/sqlite"
)

func main(){
    db, err := sql.Open("sqlite", "file:demo.db?cache=shared&mode=rwc")
    if err != nil { panic(err) }
    defer db.Close()

    _, _ = db.Exec(`CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, name TEXT);`)
    _, _ = db.Exec(`INSERT INTO users(name) VALUES ('Alice'),('Bob');`)
    rows, err := db.Query(`SELECT id, name FROM users ORDER BY id`); if err != nil { panic(err) }
    defer rows.Close()
    for rows.Next(){
        var id int; var name string
        rows.Scan(&id, &name)
        fmt.Println(id, name)
    }
}
