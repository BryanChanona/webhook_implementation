package domain

type ISensor interface{
	Save(menssage string) error
}