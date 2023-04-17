package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Platillos struct {
	ID          uint   `gorm:"column:idPlatillo;primary_key"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Precio      string `json:"precio"`
}
type Clientes struct {
	ID       uint   `gorm:"column:idCliente;primary_key"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}

func main() {
	///////////////CONEXION A LA BASE DE DATOS
	db, err := gorm.Open("mssql", "sqlserver://sa:12345678@localhost:1433?database=db_restaurante")

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()
	/////////////////////////////////////////////////////////////////////////

	//TABLA CLIENTES

	// Migrar la estructura del modelo a la base de datos
	db.AutoMigrate(&Clientes{})

	// Crear el router de Gorilla Mux
	r := mux.NewRouter()

	//CRUD SELECT
	// Definir el endpoint GET para obtener todos los clientes
	r.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		// Obtener todos los registros de la tabla de clientes
		var clientes []Clientes
		db.Find(&clientes)

		// Convertir la lista de clientes a formato JSON
		jsonClientes, err := json.Marshal(clientes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Escribir la lista de clientes en el cuerpo de la respuesta HTTP
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonClientes)
	}).Methods("GET")

	//////////////////////////////////////////////////////////////////

	//TABLA PLATILLOS

	// Migrar la estructura del modelo a la base de datos
	db.AutoMigrate(&Platillos{})
	////////////////////////////////////////////////////
	//CRUD SELECT
	// Definir el endpoint GET para obtener todos los platillos
	r.HandleFunc("/platillos", func(w http.ResponseWriter, r *http.Request) {
		// Obtener todos los registros de la tabla de platillos
		var platillos []Platillos
		db.Find(&platillos)

		// Convertir la lista de platillos a formato JSON
		jsonClientes, err := json.Marshal(platillos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Escribir la lista de platillos en el cuerpo de la respuesta HTTP
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonClientes)
	}).Methods("GET")
	////////////////////////////////////////////////////
	//CRUD INSERTAR
	// Definir el endpoint POST para crear nuevos platillos
	r.HandleFunc("/platillos", func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud HTTP a una estructura de Platillos
		var nuevoPlatillo Platillos
		err := json.NewDecoder(r.Body).Decode(&nuevoPlatillo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Crear el nuevo platillo en la base de datos
		db.Create(&nuevoPlatillo)

		// Convertir el nuevo platillo a formato JSON
		jsonNuevoPlatillo, err := json.Marshal(nuevoPlatillo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Escribir el nuevo platillo en el cuerpo de la respuesta HTTP
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonNuevoPlatillo)
	}).Methods("POST")
	////////////////////////////////////////////////////
	//CRUD ACTUALIZAR
	// Definir el endpoint PUT para actualizar un platillo existente
	r.HandleFunc("/platillos/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del platillo a actualizar desde los parámetros de la URL
		params := mux.Vars(r)
		id := params["id"]

		// Obtener el platillo existente de la base de datos
		var platillo Platillos
		err := db.Where("idPlatillo = ?", id).First(&platillo).Error
		if err != nil {
			http.Error(w, "Platillo no encontrado", http.StatusNotFound)
			return
		}

		// Decodificar el cuerpo de la solicitud HTTP a una estructura de Platillos
		err = json.NewDecoder(r.Body).Decode(&platillo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Actualizar el platillo en la base de datos
		db.Save(&platillo)

		// Convertir el platillo actualizado a formato JSON
		jsonPlatillo, err := json.Marshal(platillo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Escribir el platillo actualizado en el cuerpo de la respuesta HTTP
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonPlatillo)
	}).Methods("PUT")

	// Iniciar el servidor HTTP en el puerto 8000
	log.Printf("Iniciando servidor en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
