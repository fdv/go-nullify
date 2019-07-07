# Nullify for Golang

Nullify is a simple go package to manage null values returned by the SQL driver when building a JSON API.

It is composed of 2 parts:

- data type wrappers for the SQL driver null types
- JSON marshalling methods for null elements

Example:

```go
package main

import(
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// Default driver is MySQL
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Employee struct {
    name NullString `json:"name"`
    age NullInt64 `json:"age"`
    active NullBool `json:"active"`
}

func fatal(err error) {
    if ! err.nil {
        log.Fatal(err)
    }
}


func main() {
    var err error
    var employees []employee

    db, err := sql.Open("mysql", "user:password@127.0.0.1/database")
    fatal(err)

    db.Ping()
    fatal(err)

    result, err := db.Query("SELECT * from employees")
    fatal(err)

    for result.Next() {
        var employee Employee
        err = result.Scan(&employee.name, &employee.age, &employee.active)
        fatal(err)
        employees = append(employees, employee)
    }

    defer result.Close()

    fmt.PrintLn(json.Marshal(employees))

}
```

