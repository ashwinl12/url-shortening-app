package user

// user table colums
type userColNameFormat struct {
	ID string
	Name string
	Email string
	Country string
}

func UserColData() userColNameFormat {
	return userColNameFormat{
		ID: "id",
		Name: "name",
		Email: "email",
		Country: "country",
	}
}

