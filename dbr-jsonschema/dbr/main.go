package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type notification struct {
	UUID    string    `validate:"required,uuid4"`
	Content string    `validate:"required,jsonschema=notification"`
	Created time.Time `validate:"required"`
}

func main() {
	db := newDBHandler()

	n := notification{
		UUID: uuid.NewV4().String(),
		Content: `{
			"from": "わたし",
			"to": "あなた",
			"priority": 3,
			"message": "すごーい"
		}`,
		Created: time.Now(),
	}
	db.insert(n)

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
	db.insert(n)
}
