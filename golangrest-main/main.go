package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Productos struct {
	ID     uint   `gorm:"column:idProductos;primary_key;AUTO_INCREMENT"`
	Nombre string `json:"nombre"`
	Precio string `json:"precio"`
	Medida string `json:"medida"`
	Stock  int    `json:"stock"`
}

func main() {
	///////////////CONEXION A LA BASE DE DATOS
	db, err := gorm.Open("mysql", "root:123456789@tcp(localhost:3305)/dbpropruebav1")

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()
	fmt.Println("Success!")
	/////////////////////////////////////////////////////////////////////////

}
