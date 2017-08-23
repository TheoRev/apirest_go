package models

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// UserSchema esquema de tabla de la db
const userSchema string = `create table users(
								id int(6)unsigned primary key auto_increment,
								username varchar(30) not null UNIQUE,
								password varchar(64)not null,
								email varchar(40),
								created_date timestamp default current_timestamp
							)`

// Users slice de usuarios
type Users []User

// NewUser constructor del paquete models
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Email: email}
	user.SetPassword(password)
	return user
}

// CreateUser crea un usuario en la db
func CreateUser(username, password, email string) (*User, error) {
	user := NewUser(username, password, email)
	err := user.Save()
	return user, err
}

// Save guarda un usuario en la db
func (this *User) Save() error {
	if this.Id == 0 {
		return this.insert()
	} else {
		return this.update()
	}
}

func (this *User) insert() error {
	sql := "INSERT users SET username=?, password=?, email=?"
	id, err := InsertData(sql, this.Username, this.Password, this.Email)
	this.Id = id
	return err
}

func (this *User) update() error {
	sql := "UPDATE users SET username=?, password=?, email=?"
	result, err := Exec(sql, this.Username, this.Password, this.Email)
	this.Id, err = result.RowsAffected()
	return err
}

func (this *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	result, _ := Exec(sql, this.Id)
	this.Id, _ = result.RowsAffected()
}

func GetUser(sql string, conditional interface{}) *User {
	user := &User{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return user
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func GetUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

func (this *User) SetPassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	this.Password = string(hash)
}

func Login(username, password string) bool {
	user := GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func GetUserByUsername(username string) *User {
	sql := "SELECT id, username, password, email FROM users WHERE username=?"
	return GetUser(sql, username)
}

func ValidEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

func GetUserById(id int) *User {
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	return GetUser(sql, id)
}
