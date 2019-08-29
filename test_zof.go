package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the "company_db" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/test_zof_db?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Insert a row into the "tbl_employee" table.
	if _, err := db.Exec(
		`INSERT INTO servers (full_name, department, designation, created_at, updated_at) 
		VALUES ('Irshad', 'IT', 'Product Manager', NOW(), NOW());`); err != nil {
		log.Fatal(err)
	}

	// Select Statement.
	// rows, err := db.Query("select employee_id, full_name FROM tbl_employee;")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var employeeID int64
	// 	var fullName string
	// 	if err := rows.Scan(&employeeId, &fullName); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Employee Id : %d \t Employee Name : %s\n", employeeId, fullName)
	// }
}
