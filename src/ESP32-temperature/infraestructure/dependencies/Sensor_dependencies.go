package dependencies

import (
	"database/sql"
	"fmt"
	"webhook/src/ESP32-temperature/application"
	"webhook/src/ESP32-temperature/infraestructure"
	"webhook/src/ESP32-temperature/infraestructure/controllers"
	"webhook/src/core"
)

var (
	mySQL infraestructure.MySQL
	db    *sql.DB
)

func Init() {
	db, err := core.ConnectToDB()

	if err != nil {
		fmt.Println("server error")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)
	

}
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}

func GetPostSensorController()*controllers.PostSensorController{
	caseCreatePayment := application.NewPostSensor(&mySQL)
	return controllers.NewPostSensorController(caseCreatePayment)
}
