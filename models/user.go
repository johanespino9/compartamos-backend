package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    DNI       string `json:"dni"`
    Phone     string `json:"phone"`
    Email     string `json:"email"`
    City      string `json:"city"`
    Gender    string `json:"gender"`
    Age       int    `json:"age"`
    Deleted   bool   `json:"deleted" gorm:"default:false"` 
}
