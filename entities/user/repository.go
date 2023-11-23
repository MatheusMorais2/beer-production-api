package user

type UserRepository interface {
	Insert(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
	GetById(userId string) (*User, error)
}