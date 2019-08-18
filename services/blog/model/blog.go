package model

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Story - Model of the Story on the blog
type Story struct {
	ID          int64
	Title       string    `sql:",unique,notnull"`
	Content     string    `sql:",notnull"`
	Author      string    `sql:",notnull"`
	CreatedDate time.Time `sql:",notnull"`
}

func (s Story) String() string {
	return fmt.Sprintf("Story <%d>: <%s> from <%s>", s.ID, s.Title, s.Author)
}

//CreateSchema for Stories
func CreateSchema(db *pg.DB) error {
	err := db.CreateTable(&Story{}, &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}
