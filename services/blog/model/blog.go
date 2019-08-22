package model

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	user "github.com/locmai/first-micro/services/user/model"
)

// Story - Model of the Story on the blog
type Story struct {
	ID          int64
	Title       string     `sql:",unique,notnull"`
	Content     string     `sql:",notnull"`
	Author      *user.User `sql:",notnull"`
	CreatedDate time.Time  `sql:",notnull"`
}

type Comment struct {
	ID          int64
	Nickname    string    `sql:",notnull"`
	Story       *Story    `sql:",notnull"`
	Content     string    `sql:",notnull"`
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
