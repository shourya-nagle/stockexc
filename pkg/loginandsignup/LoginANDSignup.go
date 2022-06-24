package loginandsignup

import (
	"database/sql"
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"crypto/md5"
	"encoding/hex"
)

type User struct {
	FName    string `json:"Fname"`
	LName    string `json:"Lname"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Number   string `json:"Number"`
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/userdetails")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SignUp(context *gin.Context) {
	db := dbConn()
	var newuser User
	if err := context.BindJSON(&newuser); err != nil {
		return
	}
	Fname := newuser.FName
	Lname := newuser.LName
	Email := newuser.Email
	P := newuser.Password
	Password := GetMD5Hash(P)
	Number := newuser.Number

	insForm, err := db.Prepare("INSERT INTO user(Fname, Lname, Email, Password, Number) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	insForm.Exec(Fname, Lname, Email, Password, Number)
	context.IndentedJSON(http.StatusCreated, newuser)
	defer db.Close()

}

func Login(context *gin.Context) {

	var testuser User
	if err := context.BindJSON(&testuser); err != nil {
		return
	}
	Email := testuser.Email
	P := testuser.Password
	Password := GetMD5Hash(P)

	db := dbConn()
	sqlStmt := "SELECT * FROM user where Email=? AND Password=?"
	var email string
	row := db.QueryRow(sqlStmt, Email, Password)
	switch err := row.Scan(&email); err {
	case sql.ErrNoRows:
		fmt.Println("Invalid Credentials")

	default:
		fmt.Println("Valid user-- Login Succesful")
	}
	defer db.Close()
}


