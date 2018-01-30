package route

import "github.com/gorilla/mux"

// StartRoute Este metodo me permite iniciar el enrutamiento de mi sitio
func StartRoute(router *mux.Router) {
	// API RESTFUL - Article
	RouteArticle(router)
}
