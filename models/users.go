package models

type User struct {
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	WorkHours uint   `json:"workhours"` // Рабочие часы в Неделю
}
