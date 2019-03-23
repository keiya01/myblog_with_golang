package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/keiya01/myblog/bcrypt"
	"github.com/keiya01/myblog/database"
	"github.com/keiya01/myblog/service"
	"log"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Nickname string `gorm:"unique;not null" json:"nickname"`
	Password string `gorm:"not null" json:"password"`
	PublicID string `gorm:"unique;not null" json:"public_id"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"nickname": &graphql.Field{
			Type: graphql.String,
		},
		"public_id": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var GetUser = &graphql.Field{
	Type: UserType,
	Args: graphql.FieldConfigArgument{
		"public_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"nickname": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := database.NewHandler()
		s := service.NewService(db)

		params := p.Args

		var isLogin bool
		switch true {
		case params["public_id"] != nil:
			isLogin = false
		case params["nickname"] != nil:
			isLogin = true
		}

		publicID, _ := params["public_id"].(string)
		where := []interface{}{
			"public_id = ?",
			publicID,
		}
		if isLogin {
			nickname, _ := params["nickname"].(string)
			where = []interface{}{
				"nickname = ?",
				nickname,
			}
		}

		user := User{}
		if err := s.FindOne(&user, where...); err != nil {
			log.Println(err)
			return User{}, errors.New("User not found")
		}

		var isAuth bool
		if isLogin {
			password, _ := params["password"].(string)
			isAuth = bcrypt.ComparePassword(password, user.Password)
			if !isAuth {
				return User{}, errors.New("Password is invalid")
			}
		}

		return user, nil
	},
}

var CreatUser = &graphql.Field{
	Type:        UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"nickname": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := database.NewHandler()
		s := service.NewService(db)

		params := p.Args
		name, _ := params["name"].(string)
		nickname, _ := params["nickname"].(string)
		password, _ := params["password"].(string)
		encryptedNickname, err := bcrypt.EncryptPassword(nickname)
		encryptedPassword, err := bcrypt.EncryptPassword(password)
		if err != nil {
			return User{}, err
		}

		user := &User{
			Name:     name,
			Nickname: nickname,
			Password: encryptedPassword,
			PublicID: encryptedNickname,
		}
		if err := s.Save(&user); err != nil {
			log.Println(err)
			return User{}, err
		}

		return user, nil

	},
}
