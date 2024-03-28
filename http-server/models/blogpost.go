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

// Dummy blogpost data
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
	// Loop the blogpost array and return the blogpost that matches the given ID
	for _, blogpost := range blogPosts {
		if blogpost.ID == id {
			return blogpost
		}
	}
	return nil
}

func CreateBlogpost(bp BlogPost) {
	//Append the new blogpost to the list
	blogPosts = append(blogPosts, &bp)
}

func DeleteBlogpost(id int) bool {
	//Create a temp array to hold the new list of blogposts
	tmp := blogPosts[:0]
	for _, blogpost := range blogPosts {
		//If this blogpost isn`t the one we want to delete, then append it to the tmp list
		if blogpost.ID != id {
			tmp = append(tmp, blogpost)
		}
	}
	// Make sure that the lists are different lengths and set the blogpost list to the new tmp list
	if len(tmp) != len(blogPosts) {
		blogPosts = tmp
		return true
	}
	//Return false if we didn't find the blogpost we want to delete
	return false
}

func UpdateBlogpost(id int, blogpostUpdate BlogPost) *BlogPost {
	for i, blogpost := range blogPosts {
		//Find the blogpost we want to update and set it to the new updated version
		if blogpost.ID == id {
			blogPosts[i] = &blogpostUpdate
			return blogpost
		}
	}
	return nil
}
