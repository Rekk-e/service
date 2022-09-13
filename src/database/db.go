package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	smcfg "r3kk3/src/config"

	_ "github.com/lib/pq"
)

// INICIALIZING

func InitDB() *sql.DB {
	var config smcfg.Config
	config = smcfg.Init_config()
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_DB,
	)

	database, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")
	return database
}

func CreateWorkerTable(DB sql.DB) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Worker(
		id SERIAL PRIMARY KEY,
        Name VARCHAR(255) NOT NULL,
        Surname VARCHAR(255) NOT NULL,
        Phone VARCHAR(255) NOT NULL,
        Companyid INT NOT NULL,
        Passport INT NOT NULL,
		Department INT NOT NULL,
        FOREIGN KEY (Passport)  REFERENCES Passport (id) ON DELETE CASCADE,
        FOREIGN KEY (Department)  REFERENCES Department (id) ON DELETE CASCADE
        );`)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	fmt.Println("Worker table created")
}

func CreatePassportTable(DB sql.DB) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Passport(
		id SERIAL PRIMARY KEY,
        Number VARCHAR(255) NOT NULL UNIQUE,
		type VARCHAR(255) NOT NULL
        );`)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	fmt.Println("Passport table created")
}

func CreateDepartmentTable(DB sql.DB) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Department(
		id SERIAL PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
        Phone VARCHAR(255) NOT NULL
        );`)

	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	fmt.Println("Department table created")
}
