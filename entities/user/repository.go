package user

type UserRepository interface {
	Insert(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
}