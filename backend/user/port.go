package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)


/* This user folder is created to apply DDD model and to remove dependencies */

type Service interface{
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	/* List() ([]*User, error)
	Delete(UserID int) error
	Update(user User) (*User, error) */
}

