package models

type Item struct {
    Id string `json:"id" gorm:"primary_key"`
    Name string `json:"name"`
    Quantity string `json:"quantity"`
}