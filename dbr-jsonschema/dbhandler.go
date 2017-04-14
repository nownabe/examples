package main

type dbhandler struct {
	*dbr.Session
	*validator.Validate
	jsonschema
}

const (
	dsn = "root:password@tcp(mysql:3306)/myapp"
	notificationTable = "notifications"
)

func NewDB() *dbhandler {
	db := &dbhandler{}

	db.Session, _ = dbr.Open("mysql", dsn).NewSession(nil)

	db.Validate := validator.New()
	db.Validate.RegisterValidation("jsonschema", db.validateJson)

	db.jsonschema = newJSONSchema()

	return db
}

func (db *dbhandler) insert(n notification) {
	fmt.Println("Inserting:", n.Content)

	// Validation
	if err := db.Struct(n); err != nil {
		fmt.Println("Validation error!", err)
		return
	}

	db.InsertInto(notificationTable).Record(n)
}

func (db *dbhandler) validateJSON(fl validator.FieldLevel) bool {
	json := fl.Field().String()
	jsonType := fl.Param()

	return db.ValidateJSON(jsonType, json)
}
