package repo

type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

type UserRepo interface{
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}


var users []User

func(u User) Store() User{
	if u.ID != 0{
		return u
	}
	u.ID = len(users)+1
	users = append(users, u)
	return u
}
func Find(email, pass string) (*User){
	for _, u := range users{
		if u.Email == email && u.Password == pass{
			return &u
		}
	}
	return nil
}