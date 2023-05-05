package main

import (
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type BlogPost struct {
	Title   string
	Content string
}

func AddPost(dao *daos.Dao, data *BlogPost) error {
	blogPostsCol, err := dao.FindCollectionByNameOrId("blogPosts")
	if err != nil {
		return err
	}
	record := models.NewRecord(blogPostsCol)
	record.Set("title", data.Title)
	record.Set("content", data.Content)
	err = dao.SaveRecord(record)
	if err != nil {
		return err
	}
	return nil
}
