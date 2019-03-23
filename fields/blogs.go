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
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	UserID    int       `gorm:"not null" json:"user_id"`
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
		"user_id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var GetBlogList = &graphql.Field{
	Type: graphql.NewList(BlogType),
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := database.NewHandler()
		s := service.NewService(db)

		var where []interface{}
		userID, ok := p.Args["user_id"].(int)
		if ok {
			where = []interface{}{"user_id = ?", userID}
		}
		blogs := []Blog{}

		if err := s.FindAll(&blogs, "created_at desc", where...); err != nil {
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
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := database.NewHandler()
		s := service.NewService(db)

		params := p.Args
		title, _ := params["title"].(string)
		body, _ := params["body"].(string)
		userID, _ := params["user_id"].(int)
		blog := &Blog{
			Title:  title,
			Body:   body,
			UserID: userID,
		}
		if err := s.Save(&blog); err != nil {
			log.Println(err)
			return blog, err
		}

		return blog, nil

	},
}
