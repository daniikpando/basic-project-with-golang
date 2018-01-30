package main

import (
	"log"
	"net/http"

	"github.com/daniel/apiRest2/route"
	"github.com/gorilla/mux"
)

//"ppl-ps":"$\$daniikpando" --contraseña paypal
func main() {
	port := ":8000"
	router := mux.NewRouter()

	route.StartRoute(router)

	log.Printf("El servidor esta corriendo en el puerto %s", port)
	log.Fatal(http.ListenAndServe(port, router))

}

/*article := models.Article{
	IdArticle:1,
	Title:"Nose que hacer hoy",
	Content:"Nose que hacer mañana",
	Description:"Ni idea de esto",
}*/
// Create

/*manager := postgresql.ArticlePSQL{}

err := manager.Create(&article)

if err != nil {
	fmt.Println("nose")
	log.Fatal(err)
}
fmt.Println("Se creo exitosamente")
*/

// Update

/*
manager := postgresql.ArticlePSQL{}

err := manager.Update(&article)

if err != nil {
	fmt.Println("nose")
	log.Fatal(err)
}
fmt.Println("Se actualizo exitosamente")
*/
// DELETE

/*manager := postgresql.ArticlePSQL{}

err := manager.Delete(&article)

if err != nil {
	fmt.Println("nose")
	log.Fatal(err)
}
fmt.Println("Se elimino exitosamente")

*/

// Get all

/*manager := postgresql.ArticlePSQL{}

articles, err := manager.GetAll()

if err != nil {
	fmt.Println("nose")
	log.Fatal(err)
}

fmt.Println(articles)
*/
// Get by id

/*manager := postgresql.ArticlePSQL{}

article, err := manager.GetById(2)

if err != nil {
	fmt.Println("nose")
	log.Fatal(err)
}

fmt.Println(article)
*/
