package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"size:255;unique;not null"`
	Email    string `json:"email" gorm:"size:255;unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
