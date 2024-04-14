package entity

type UserSession struct {
	Id       string
	Username string
	Email    string
	Roles    []string
}
