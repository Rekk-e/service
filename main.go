package main

import (
	"database/sql"
	"time"

	db "r3kk3/src/database"
	routes "r3kk3/src/routes"
)

var database sql.DB

func main() {

	// Подключаемся к БД
	database := db.InitDB()

	time.Sleep(1 * time.Second)

	// Создаем таблицу паспортов
	db.CreatePassportTable(*database)

	// Создаем таблицу отделов
	db.CreateDepartmentTable(*database)

	// Создаем таблицу сотрудников
	db.CreateWorkerTable(*database)

	// Инициализируем роуты
	routes.InitRoutes(database)

}
