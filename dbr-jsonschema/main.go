package main

import (
	"time"

	"github.com/gocraft/dbr"
	uuid "github.com/satori/go.uuid"
)

type notification struct {
	UUID    string    `validate:"required,uuid4"`
	Content string    `validate:"required,jsonschema=notifcation"`
	Created time.Time `validate:"required"`
}

func main() {
	conn, _ := dbr.Open("mysql", dsn)
	sess := conn.NewSession(nil)

	n := notification{
		UUID: uuid.NewV4().String(),
		Content: `{
			"from": "おれ",
			"to": "おまえ",
			"priority": 3,
			"message": "すごーい"
		}`,
		Created: time.Now(),
	}
	insert(sess, n)

	n = notification{
		UUID: uuid.NewV4().String(),
		Content: `{
			"from": "ぼく",
			"to": "きみ",
			"priority": 100000000,
			"message": "わーい"
		}`,
		Created: time.Now(),
	}
	insert(sess, n)
}

func insert(sess *dbr.Session, n notification) {

}
