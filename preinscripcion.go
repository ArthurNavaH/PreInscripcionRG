package main

import (
	"github.com/arthurnavah/PreInscripcionRG-API/routes"
	"github.com/arthurnavah/PreInscripcionRG-API/utils"
)

func main() {
	//FlagControl Controla las banderas para el ejecutable.
	//	--createTables yes	; Inicia la creacion de las tablas necesarias.
	//	--recreateTables yes	; Recrea las tablas necesarias.
	//	--dropTables yes	; Borra las tablas necesarias.
	//	--port 8181	;	Cambia el puerto por defecto del servidor.
	utils.FlagControl()

	//router Contiene una instancia ServerMux con todas las rutas ya definidas.
	router := routes.InitRoutes()

	//ServerSetup Ejecuta el servidor con el puerto y el router.
	utils.ServerSetup(utils.Port, router)
}
