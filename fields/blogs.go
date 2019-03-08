package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/keiya01/myblog/database"
	"github.com/keiya01/myblog/service"
	"log"
	"time"
)

type Blog struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var BlogType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Blog",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BlogList = &graphql.Field{
	Type: graphql.NewList(BlogType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := database.NewHandler()
		s := service.NewService(db)
		blogs := []Blog{}
		if err := s.FindAll(&blogs, "created_at desc"); err != nil {
			log.Println(err)
			return blogs, err
		}

		return blogs, nil
	},
}

var CreateBlog = &graphql.Field{
	Type:        BlogType,
	Description: "Create new blog",
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := database.NewHandler()
		s := service.NewService(db)

		params := p.Args

		blog := &Blog{
			Title: params["title"].(string),
			Body:  params["body"].(string),
		}
		if err := s.Save(&blog); err != nil {
			log.Println(err)
			return blog, err
		}

		return blog, nil

	},
}
