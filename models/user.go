package models

// User model
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserSchema esquema de tabla de la db
const userSchema string = `create table users(
								id int(6)unsigned primary key auto_increment,
								username varchar(30) not null,
								password varchar(64)not null,
								email varchar(40),
								created_date timestamp default current_timestamp
							)`

// Users slice de usuarios
type Users []User

// NewUser constructor del paquete models
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// CreateUser crea un usuario en la db
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Save guarda un usuario en la db
func (this *User) Save() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := Exec(sql, this.Username, this.Password, this.Email)
	this.Id, _ = result.LastInsertId()
}
