// Package for the declaration of the main structures
// for data handling in the project.
package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
	_ "github.com/lib/pq"
)

// Postgresql Settings
const (
	USER        = "manuelams"
	PASSWORD    = "11235813"
	HOST        = "localhost"
	PORT        = 26257
	DATABASE    = "servers"
)

// Declaration of global DB controller
var (
	DBConf *sql.DB
)

// Function for the inicialization of the global DB controller
func InitDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", USER, PASSWORD, HOST, strconv.Itoa(PORT), DATABASE)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

// Auxiliar function to add quotes to a string
func addQuotes(word string) string {
	return fmt.Sprintf(`'%v'`, word)
}

type RegisteredUser struct {
    Id          int         `json:"registered_user_id"` // SERIAL PRIMARY KEY
    PhotoUrl    string      `json:"photo_url"`  // VARCHAR[100]
    Name        string      `json:"name"`       // VARCHAR[45]
    LastName    string      `json:"last_name"`  // VARCHAR[45]
    Email       string      `json:"email"`      // VARCHAR[100]
    Country     string      `json:"country"`    // VARCHAR[45]
    Direction   string      `json:"direction"`  // VARCHAR[200]
    ZipCode     string      `json:"zip_code"`   // VARCHAR[30]
    Phone       string      `json:"phone"`      // VARCHAR[45]
    IsAdmin     bool        `json:"is_admin"`   // BOOLEAN
    Password    string      `json:"password"`   // VARCHAR[30] 
}

RegisteredUserDBColumns := []string{}

func getStructFields(structure []interface{}, inCols []string, names []string) (fields []reflect.Value, outCols []string, err error) {
    type T struct {
    A int
    B string
}
t := T{23, "skidoo"}
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()
for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Printf("%d: %s %s = %v\n", i,
        typeOfT.Field(i).Name, f.Type(), f.Interface())
}
    s := reflect.ValueOf(&structure).Elem()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fName := f.Interface().(string)
        for j := 0; j < len(names); j++ {
            if fName == names[j] {
                fields = append(fields, f)
                outCols = append(outCols, inCols[i])
                break                
            }            
        }  
    }
    return 
}

fmt.Printf

