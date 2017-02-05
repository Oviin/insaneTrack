package models

type Team struct {
    Description string
    Users []*User
    TeamHours uint // Колличество часов на команду 
}