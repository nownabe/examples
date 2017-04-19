package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"gopkg.in/go-playground/validator.v9"
)

type dbhandler struct {
	*dbr.Session
	*validator.Validate
	*jsonschema
}

const (
	dsn               = "root:password@tcp(mysql:3306)/myapp"
	notificationTable = "notifications"
)

func newDBHandler() *dbhandler {
	db := &dbhandler{}

	conn, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		log.Fatalln(err)
	}

	db.Session = conn.NewSession(nil)

	db.Validate = validator.New()
	db.Validate.RegisterValidation("jsonschema", db.isValidJSON)

	db.jsonschema = newJSONSchema()

	return db
}

func (db *dbhandler) insert(n notification) {
	log.Println("Inserting:", n.Content)

	// Validation
	// https://godoc.org/gopkg.in/go-playground/validator.v9#Validate.Struct
	if err := db.Struct(n); err != nil {
		log.Println("Validation Error!", err)
		return
	}

	db.InsertInto(notificationTable).Record(n)
	log.Println("Inserted!")
}

// Custom validator
// https://godoc.org/gopkg.in/go-playground/validator.v9#hdr-Custom_Validation_Functions
func (db *dbhandler) isValidJSON(fl validator.FieldLevel) bool {
	json := fl.Field().String()
	jsonType := fl.Param()

	return db.validateJSON(jsonType, json)
}
