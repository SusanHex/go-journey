/*
The package works on 2 tables of a SQLite DB.
The names of the tables are:
  - Users
  - Userdata

the definitions of the tables are:

	CREATE TABLE Users (
		ID INTEGER PRIMARY KEY,
		Username Text
	);
	CREATE TABLE Userdata (
		UserID INTEGER NOT NULL,
		Name TEXT,
		Surname TEXT,
		Description TEXT
	)
*/
package sqlite06

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
This global variable holds the filepath to the SQLite3 DB.

	Filename: Is the filepath to the database file.
*/
var (
	Filename = ""
)

/*
The userdata structure is for holding full user data
from the Userdata table and usernames from the Users table
*/
type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

// openConnection() is for opening the SQLite3 connection
// in order to be used by the other functions of the package
func openConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// The function returns the User ID of the username
// -1 if the user does not exists
func exists(username string) int {
	username = strings.ToLower(username)
	db, err := openConnection()
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
	db, err := openConnection()
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
	insertStatement = `INSERT INTO Userdata values (?, ?, ?, ?)`
	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}
	return userID
}

/*
DeleteUser() deletes a user if the user exists.
It requires the User ID of the user the be deleted.
Returnes an error in case something goes wrong.
*/
func DeleteUser(id int) error {
	db, err := openConnection()
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
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		return err
	}
	return nil
}

// ListUsers() Lists all users in the database.
//
// Returns a slice of Userdata to the calling function
func ListUsers() ([]Userdata, error) {
	data := []Userdata{}
	db, err := openConnection()
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

/*
UpdateUser() is for updating an existing user
given a Userdata structure.
The user ID of the user to be updated is found inside the function
via corelating the Username to the User ID.
*/
func UpdateUser(d Userdata) error {
	db, err := openConnection()
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

func SearchUser(username string) (Userdata, error) {
	db, err := openConnection()
	temp := Userdata{}
	temp.ID = -1
	if err != nil {
		return Userdata{}, err
	}
	defer db.Close()
	findStatement := `SELECT * FROM Users WHERE Username = ?`

	row, err := db.Query(findStatement, username)
	if err != nil {
		return Userdata{}, err
	}
	row.Next()
	err = row.Scan(&temp.ID, &temp.Username)
	row.Close()
	if err != nil {
		temp.ID = -1
		return temp, err
	}
	userDataStatement := `SELECT * FROM Userdata WHERE UserID = ?`
	row, err = db.Query(userDataStatement, temp.ID)
	if err != nil {
		temp.ID = -1
		return temp, err
	}
	row.Next()
	err = row.Scan(&temp.ID, &temp.Name, &temp.Surname, &temp.Description)
	row.Close()
	if err != nil {
		return Userdata{ID: -1, Username: "", Name: "", Surname: "", Description: ""}, err
	}
	return temp, nil
}
