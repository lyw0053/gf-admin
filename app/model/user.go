package model

type User struct {
	Id   int
	Name string
}

func (this *User) GetTableName() string {
	return "user"
}
