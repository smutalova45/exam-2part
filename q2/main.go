package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	/*
			SELECT e.name AS employee, d.name AS department, COUNT(s.employee_id) AS count
		    FROM employees e
		    JOIN departments d ON e.department_id = d.id
		   LEFT JOIN salaries s ON e.id = s.employee_id
		   GROUP BY e.name, d.name;

	*/

}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=employe sslmode=disable")
	if err != nil {
		return nil, err

	}
	return db, nil

}
