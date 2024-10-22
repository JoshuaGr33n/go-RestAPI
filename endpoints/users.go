package endpoints

import (
	"RestAPI/db"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	// Password string `gorm:"not null"`
}

var database = db.GetMySQLDB()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	database.Find(&users)
	json.NewEncoder(w).Encode(users)
}
