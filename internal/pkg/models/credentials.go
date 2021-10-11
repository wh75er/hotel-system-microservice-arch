package models

type Token string
type Role string

const (
	ADMIN   = Role("admin")
	USER    = Role("user")
	SERVICE = Role("service")
)

type Credentials struct {
	Id     string
	Secret string
}

type CredentialsUsecaseI interface {
	Login(creds *Credentials) error
}
