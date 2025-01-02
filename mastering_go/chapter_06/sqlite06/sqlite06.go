package sqlite06

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Filename = ""
)

type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", Filename)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// The function returns the User ID of the username
// -1 if the user does not exists
func exists(username string) int {
	username = strings.ToLower(username)
	db, err := OpenConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()
	userID := -1
	statement := fmt.Sprintf(`SELECT ID FROM Users where Username = '%s'`, username)
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("exists(), Scan", err)
			return -1
		}
		userID = id
	}
	return userID
}

// AddUser adds a new user to the database
// Returns new User ID
// -1 if there was an error
func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)
	db, err := OpenConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()
	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", d.Username)
		return -1
	}
	insertStatement := `INSERT INTO Users values (NULL,?)`
	_, err = db.Exec(insertStatement, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	userID = exists(d.Username)
	if userID == -1 {
		return userID
	}
	insertStatement = `INSERT INTO Userdata values (?, ?, ?, ?, ?)`
	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}
	return userID
}

func DeleteUser(id int) error {
	db, err := OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	statement := fmt.Sprintf(`SELECT Username FROM Users WHERE ID = %d`, id)
	rows, err := db.Query(statement)
	defer rows.Close()
	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return err
		}
	}
	if exists(username) != id {
		return fmt.Errorf("User with ID %d does not exists", id)
	}
	// Delete from Userdata
	deleteStatement := `DELETE FROM Userdata WHERE UserID = ?`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		return err
	}
	// Delete from Users
	deleteStatement = `DELETE from Users WHERE ID = ?`
	_, err = db.Exec(deleteStatement)
	if err != nil {
		return err
	}
	return nil
}

func ListUsers() ([]Userdata, error) {
	data := []Userdata{}
	db, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(`SELECT ID, Username, Name, Surname, Description FROM Users, Userdata WHERE Users.ID = Userdata.UserID`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var username string
		var name string
		var surname string
		var description string
		err = rows.Scan(&id, &username, &name, &surname, &description)
		temp := Userdata{ID: id, Username: username, Name: name, Surname: surname, Description: description}
		data = append(data, temp)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func UpdateUser(d Userdata) error {
	db, err := OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	userID := exists(d.Username)
	if userID == -1 {
		return errors.New("User does not exists")
	}
	d.ID = userID
	updateStatement := `UPDATE Userdata set Name = ?, Surname = ?, Description = ? where UserID = ?`
	_, err = db.Exec(updateStatement, d.Name, d.Surname, d.Description, d.ID)
	if err != nil {
		return err
	}
	return nil
}
