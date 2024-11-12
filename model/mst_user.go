package model

type MstUser struct {
	IdUser      string
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Role        MstRole
}
