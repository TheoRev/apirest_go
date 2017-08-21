package orm

import "time"

// User model
type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

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
	db.Create(&this)
}

// Update actualiza un registro
func (this *User) Update() {
	user := User{Username: this.Username, Password: this.Password, Email: this.Email}
	db.Model(&this).UpdateColumns(user)
}

// Delete elimina un registro
func (this *User) Delete() {
	db.Delete(&this)
}

// GetUsers obtiene todos los registros de la db
func GetUsers() Users {
	users := Users{}
	db.Find(&users)
	return users
}

// GetUser obtiene un solo registro
func GetUser(id int) *User {
	user := &User{}
	db.Where("id=?", id).First(user)
	return user
}
