package db

import (
	"context"
	"database/sql"
	"log"
	models "r3kk3/src/models"
	"time"
)

// SQL COMANDS

func InsertWorker(DB sql.DB, worker models.Worker) (int, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	InsertPassport(DB, worker.Passport)
	InsertDepartment(DB, worker.Department)
	_, err := DB.ExecContext(ctx,
		`INSERT INTO Worker (
			id, 
			Name, 
			Surname, 
			Phone, 
			Companyid, 
			Passport, 
			Department) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		worker.Id,
		worker.Name,
		worker.Surname,
		worker.Phone,
		worker.CompanyId,
		worker.Passport.Id,
		worker.Department.Id,
	)

	if err != nil {
		log.Printf("Error %s when insert worker\n", err)
		return -1, err
	}
	return worker.Id, nil
}

func DeleteWorker(DB sql.DB, id string) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx,
		`DELETE FROM Worker WHERE id=$1`,
		id,
	)
	if err != nil {
		log.Printf("Error %s when delete worker\n", err)
		return false, err
	}
	return true, nil
}

func GetWorker(DB sql.DB, id string) (models.Worker, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Worker

	err := DB.QueryRowContext(
		ctx, "SELECT * FROM Worker WHERE id = $1",
		id,
	).Scan(
		&query.Id,
		&query.Name,
		&query.Surname,
		&query.Phone,
		&query.CompanyId,
		&query.Passport.Id,
		&query.Department.Id,
	)
	query.Passport, err = GetPassport(DB, query.Passport.Id)
	if err != nil {
		log.Printf("Error %s when get passport\n", err)
		return models.Worker{}, err
	}
	query.Department, err = GetDepartment(DB, query.Department.Id)
	if err != nil {
		log.Printf("Error %s when get department\n", err)
		return models.Worker{}, err
	}
	return query, nil
}

func GetWorkersByCompanyId(DB sql.DB, company_id string) ([]models.Worker, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Worker
	var workers []models.Worker = nil

	rows, err := DB.QueryContext(
		ctx, "SELECT * FROM Worker WHERE Companyid = $1",
		company_id,
	)
	if err != nil {
		log.Printf("Error %s when get worker\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&query.Id,
			&query.Name,
			&query.Surname,
			&query.Phone,
			&query.CompanyId,
			&query.Passport.Id,
			&query.Department.Id,
		)
		query.Passport, err = GetPassport(DB, query.Passport.Id)
		if err != nil {
			log.Printf("Error %s when get passport\n", err)
			continue
		}
		query.Department, err = GetDepartment(DB, query.Department.Id)
		if err != nil {
			log.Printf("Error %s when get department\n", err)
			continue
		}
		workers = append(workers, query)
	}

	return workers, nil
}

func GetWorkersByDepartment(DB sql.DB, department_name string) ([]models.Worker, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Worker
	var workers []models.Worker = nil

	rows, err := DB.QueryContext(
		ctx,
		`SELECT * FROM Worker
		WHERE Worker.Department IN (
			SELECT id FROM Department WHERE Name = $1
		)`,
		department_name,
	)
	if err != nil {
		log.Printf("Error %s when get worker\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&query.Id,
			&query.Name,
			&query.Surname,
			&query.Phone,
			&query.CompanyId,
			&query.Passport.Id,
			&query.Department.Id,
		)
		query.Passport, err = GetPassport(DB, query.Passport.Id)
		if err != nil {
			log.Printf("Error %s when get passport\n", err)
			return nil, err
		}
		query.Department, err = GetDepartment(DB, query.Department.Id)
		if err != nil {
			log.Printf("Error %s when get department\n", err)
			return nil, err
		}
		workers = append(workers, query)
	}

	return workers, nil
}

func ChangeWorker(DB sql.DB, r_worker models.Worker, id string) (models.Worker, error) {
	worker, err := GetWorker(DB, id)
	if err != nil {
		log.Printf("Error %s when get worker\n", err)
		return models.Worker{}, err
	}
	worker.OverlayFields(r_worker)
	UpdatePassport(DB, worker.Passport)
	UpdateDepartment(DB, worker.Department)
	UpdateWorker(DB, worker)

	return worker, nil
}

func UpdateWorker(DB sql.DB, worker models.Worker) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(
		ctx, `UPDATE Worker SET 
		id = $1, 
		Name = $2, 
		Surname = $3,
		Phone = $4,
		Companyid = $5,
		Passport = $6,
		Department = $7
		WHERE id = $1`,
		worker.Id,
		worker.Name,
		worker.Surname,
		worker.Phone,
		worker.CompanyId,
		worker.Passport.Id,
		worker.Department.Id,
	)

	if err != nil {
		log.Printf("Error %s when update worker\n", err)
		return false, err
	}
	return true, nil
}

func InsertPassport(DB sql.DB, passport models.Passport) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(
		ctx, "INSERT INTO Passport (id, Number, type) VALUES ($1, $2, $3)",
		passport.Id,
		passport.Number,
		passport.Type,
	)

	if err != nil {
		log.Printf("Error %s when insert passport\n", err)
		return false, err
	}
	return true, nil
}

func GetPassport(DB sql.DB, Id int) (models.Passport, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Passport
	err := DB.QueryRowContext(
		ctx, "SELECT * FROM Passport WHERE id = $1",
		Id,
	).Scan(&query.Id, &query.Number, &query.Type)

	if err != nil {
		log.Printf("Error %s when get passport\n", err)
		return models.Passport{}, err
	}

	return query, nil
}

func UpdatePassport(DB sql.DB, passport models.Passport) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(
		ctx, "UPDATE Passport SET id = $1, Number = $2, type = $3 WHERE id = $1",
		passport.Id,
		passport.Number,
		passport.Type,
	)

	if err != nil {
		log.Printf("Error %s when update passport\n", err)
		return false, err
	}
	return true, nil
}

func InsertDepartment(DB sql.DB, department models.Department) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(
		ctx, "INSERT INTO Department (id, Name, Phone) VALUES ($1, $2, $3)",
		department.Id,
		department.Name,
		department.Phone,
	)
	if err != nil {
		log.Printf("Error %s when insert department\n", err)
		return false, err
	}
	return true, nil
}

func GetDepartment(DB sql.DB, Id int) (models.Department, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Department
	err := DB.QueryRowContext(
		ctx, "SELECT * FROM Department WHERE id = $1",
		Id,
	).Scan(&query.Id, &query.Name, &query.Phone)

	if err != nil {
		log.Printf("Error %s when get passport\n", err)
		return models.Department{}, err
	}

	return query, nil
}

func UpdateDepartment(DB sql.DB, department models.Department) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(
		ctx, "UPDATE Department SET id = $1, Name = $2, Phone = $3 WHERE id = $1",
		department.Id,
		department.Name,
		department.Phone,
	)

	if err != nil {
		log.Printf("Error %s when update department\n", err)
		return false, err
	}
	return true, nil
}
