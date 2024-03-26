package models

import (
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

type BlogPost struct {
	ID       int       `json:"id"`
	UserName string    `json:"user_name"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
	Post     string    `json:"post"`
	PostDate time.Time `json:"post_date"`
}

var blogPosts = []*BlogPost{
	{
		ID:       0,
		UserName: gofakeit.Username(),
		Email:    gofakeit.Email(),
		FullName: gofakeit.Name(),
		Post:     gofakeit.SentenceSimple(),
		PostDate: gofakeit.Date(),
	},
	{
		ID:       1,
		UserName: gofakeit.Username(),
		Email:    gofakeit.Email(),
		FullName: gofakeit.Name(),
		Post:     gofakeit.SentenceSimple(),
		PostDate: gofakeit.Date(),
	},
	{
		ID:       2,
		UserName: gofakeit.Username(),
		Email:    gofakeit.Email(),
		FullName: gofakeit.Name(),
		Post:     gofakeit.SentenceSimple(),
		PostDate: gofakeit.Date(),
	},
	{
		ID:       3,
		UserName: gofakeit.Username(),
		Email:    gofakeit.Email(),
		FullName: gofakeit.Name(),
		Post:     gofakeit.SentenceSimple(),
		PostDate: gofakeit.Date(),
	},
}

func GetBlogposts() []*BlogPost {
	return blogPosts
}

func GetBlogpost(id int) *BlogPost {
	for _, blogpost := range blogPosts {
		if blogpost.ID == id {
			return blogpost
		}
	}
	return nil
}

func CreateBlogpost(bp BlogPost) {
	blogPosts = append(blogPosts, &bp)
}

func DeleteBlogpost(id int) bool {
	tmp := blogPosts[:0]
	for _, blogpost := range blogPosts {
		if blogpost.ID != id {
			tmp = append(tmp, blogpost)
		}
	}
	if len(tmp) != len(blogPosts) {
		blogPosts = tmp
		return true
	}
	return false
}

func UpdateBlogpost(id int, blogpostUpdate BlogPost) *BlogPost {
	for i, blogpost := range blogPosts {
		if blogpost.ID == id {
			blogPosts[i] = &blogpostUpdate
			return blogpost
		}
	}
	return nil
}
