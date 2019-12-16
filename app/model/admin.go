package model

type Admin struct {
}

func (this *Admin) GetTableName() string {
	return tablePrefix + "admin"
}
