package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/graphql-go/handler"
	"github.com/keiya01/myblog/graphql"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	*chi.Mux
}

// NewServer Server構造体のコンストラクタ
func NewServer() *Server {
	return &Server{
		Mux: chi.NewRouter(),
	}
}

func (s *Server) Route() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Use this to allow specific origin hosts
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	s.Use(cors.Handler)

	s.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var params map[string]interface{}
		err = json.Unmarshal(b, &params)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		result := graphql.ExecuteQuery(params["query"].(string))
		json.NewEncoder(w).Encode(result)
	})

	schema := graphql.GetSchema()
	s.Handle("/graphiql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))
}

func (s *Server) Start(port string) {
	log.Print("Success Connecting")
	err := http.ListenAndServe(port, s)
	if err != nil {
		log.Println("ListenAndServe:", err)
	}
}
