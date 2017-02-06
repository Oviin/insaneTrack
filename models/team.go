package models

type Team struct {
	Description string  `json:"description"`
	Users       []*User `json:"user"`
	TeamHours   uint    `json:"teamhours"` // Колличество часов на команду
}
