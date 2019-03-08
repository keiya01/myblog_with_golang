package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/keiya01/myblog/fields"
	"log"
)

func GetSchema() graphql.Schema {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"blogList": fields.BlogList,
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootMutation",
			Fields: graphql.Fields{
				"createBlog": fields.CreateBlog,
			},
		}),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
		return schema
	}

	return schema
}

func ExecuteQuery(query string) *graphql.Result {
	schema := GetSchema()

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	return r

	// rJSON, _ := json.Marshal(r)
	// fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
