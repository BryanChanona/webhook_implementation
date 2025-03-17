package domain

type Sensor struct {
    Id    int32  `json:"id"`
    Menssage string `json:"menssage"`
}

func NewSensor( menssage string) *Sensor{
	return &Sensor{Id:0,Menssage:menssage}
}