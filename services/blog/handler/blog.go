package handler

import (
	"context"
	"time"

	"github.com/go-pg/pg"
	"github.com/locmai/first-micro/services/blog/model"

	"github.com/micro/go-micro/util/log"

	blog "github.com/locmai/first-micro/services/blog/proto/blog"
)

type Blog struct {
	Db *pg.DB
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Blog) Get(ctx context.Context, req *blog.Request, rsp *blog.Response) error {
	log.Log("Received Blog.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Blog) Create(ctx context.Context, req *blog.CreateRequest, rsp *blog.Response) error {
	log.Log("Reeived Blog.Call request")
	newStory := &model.Story{
		Title:       req.Title,
		Content:     req.Content,
		Author:      req.Author,
		CreatedDate: time.Now(),
	}
	err := e.Db.Insert(newStory)
	if err != nil {
		rsp.Msg = "Failed to create new post"
		return err
	}
	log.Infof("Created %", newStory.String())
	rsp.Msg = newStory.String()
	return nil
}
