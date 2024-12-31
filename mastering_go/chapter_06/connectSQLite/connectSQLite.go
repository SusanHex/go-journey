package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var dbname = "test.db"

func insertData(db *sql.DB, dsc string) error {
	cT := time.Now().Format(time.RFC1123)
	stmt, err := db.Prepare("INSERT INTO book values(NULL,?,?);")
	if err != nil {
		fmt.Println("Insert data table:", err)
		return err
	}
	_, err = stmt.Exec(cT, dsc)
	if err != nil {
		fmt.Println("Insert data table:", err)
		return err
	}
	return nil
}

func selectData(db *sql.DB, n int) error {
	rows, err := db.Query("SELECT * from book WHERE id > ?", n)
	if err != nil {
		fmt.Println("Select:", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var dt string
		var description string
		err = rows.Scan(&id, &dt, &description)
		if err != nil {
			fmt.Println("Row:", err)
			return err
		}
		date, err := time.Parse(time.RFC1123, dt)
		if err != nil {
			fmt.Println("Date:", err)
			return err
		}
		fmt.Printf("%d %s %s\n", id, date, description)
	}
	return nil
}

func main() {
	// delete db file
	os.Remove(dbname)
	// connect to sqlite db
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer db.Close()
	// create db table
	const create string = `
	CREATE TABLE IF NOT EXISTS book (
	id INTEGER NOT NULL PRIMARY KEY,
	time TEXT NOT NULL,
	description TEXT);
	`
	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("Create Table:", err)
		return
	}
	// insert 10 rows into the books table
	for i := 1; i < 11; i++ {
		dsc := "Description: " + strconv.Itoa(i)
		err = insertData(db, dsc)
		if err != nil {
			fmt.Println("Insert data:", err)
		}
	}
	// select multiple rows
	err = selectData(db, 5)
	if err != nil {
		fmt.Println("Select:", err)
	}
	time.Sleep(time.Second)
	// update data
	cT := time.Now().Format(time.RFC1123)
	db.Exec("UPDATE book SET time = ? WHERE id > ?", cT, 7)
	// select multiple rows
	err = selectData(db, 8)
	if err != nil {
		fmt.Println("Select:", err)
		return
	}

	// Delete data
	stmt, err := db.Prepare("DELETE from book where id = ?")
	_, err = stmt.Exec(8)
	if err != nil {
		fmt.Println("Delete:", err)
		return
	}
	// Select multiple rows
	err = selectData(db, 7)
	if err != nil {
		fmt.Println("Select:", err)
		return
	}
	// count rows in table
	query, err := db.Query("SELECT count(*) as count from book")
	if err != nil {
		fmt.Println("Select:", err)
		return
	}
	defer query.Close()
	count := -100
	for query.Next() {
		_ = query.Scan(&count)
	}
	fmt.Println("count(*):", count)
}
