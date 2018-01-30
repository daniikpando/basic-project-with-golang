package route

import (
	"github.com/daniel/apiRest2/controllers/article"
	"github.com/gorilla/mux"
)

func RouteArticle(router *mux.Router) {
	router.HandleFunc("/articles", article.Index).Methods("GET")
	router.HandleFunc("/articles/new", article.New).Methods("GET")
	router.HandleFunc("/articles/create", article.Create).Methods("POST")
	router.HandleFunc("/articles/update", article.Update).Methods("PUT", "PATH", "POST")
	router.HandleFunc("/articles/{id}", article.Show).Methods("GET")
	router.HandleFunc("/articles/{id}/edit", article.Edit).Methods("GET")
	router.HandleFunc("/articles/{id}/delete", article.Delete).Methods("GET", "DELETE")
}
