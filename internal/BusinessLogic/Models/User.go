package Models

type User struct {
	Id        uint
	Login     string
	Name      string
	Age       int
	HairColor HairColor
	Gender    Gender
}

func NewUser(login string, name string, age int, hairColor HairColor, gender Gender) *User {
	return &User{
		Login:     login,
		Name:      name,
		Age:       age,
		HairColor: hairColor,
		Gender:    gender,
	}
}

type HairColor string
type Gender string

const (
	HairColorBlond HairColor = "Blond"
	HairColorRed   HairColor = "Red"
)

const (
	GenderMale Gender = "Male"

	GenderFemale Gender = "Female"
)
