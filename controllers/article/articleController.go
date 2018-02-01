package article

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/daniel/apiRest2/dao/postgresql"
	"github.com/daniel/apiRest2/models"
	"github.com/gorilla/mux"
)

var manager postgresql.ArticlePSQL

func init() {
	manager = postgresql.ArticlePSQL{}

}

// Index :  GET - /articles
func Index(w http.ResponseWriter, r *http.Request) {

	log.Println(r.RequestURI)

	articles := make([]models.Article, 0)

	articles, err := manager.GetAll()

	detectErr(err, "Error en la obtencion de los articulos")

	parseHTML(w, "./views/application/layout.tmpl", "./views/article/index.tmpl", articles, "articles")
}

// Show : GET /articles/{id}
func Show(w http.ResponseWriter, r *http.Request) {

	log.Println(r.RequestURI)

	var idArticle int

	article := models.Article{}

	vars := mux.Vars(r)

	idArticle, err := strconv.Atoi(vars["id"])

	detectErr(err, "Error de conversion de datos")

	article, err = manager.GetById(idArticle)

	detectErr(err, "No se ha encontrado el registro ")

	parseHTML(w, "./views/application/layout.tmpl", "./views/article/show.tmpl", article, "showArticle")
}

// New : GET /articles/new
func New(w http.ResponseWriter, r *http.Request) {
	parseHTML(w, "./views/application/layout.tmpl", "./views/article/new.tmpl", models.Article{}, "newArticle")
}

// Create : POST - /articles/create
func Create(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	article := models.Article{}

	title := r.FormValue("title")
	content := r.FormValue("content")
	description := r.FormValue("description")

	article.Title = title
	article.Content = content
	article.Description = description

	err := manager.Create(&article)

	detectErr(err, "Error en la creacion del registro")

	http.Redirect(w, r, "/articles", http.StatusFound)
}

// Edit : GET - /articles/{id}/edit
func Edit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idArticle, err := strconv.Atoi(vars["id"])

	detectErr(err, "Error de conversion de datos")

	article := models.Article{}

	article, err = manager.GetById(idArticle)

	detectErr(err, "Error en la base de datos")

	parseHTML(w, "./views/application/layout.tmpl", "./views/article/edit.tmpl", article, "editArticle")

}

// Update : PUT - /articles/update
func Update(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	article := models.Article{}

	idArticle, err := strconv.Atoi(r.FormValue("id_article"))

	detectErr(err, "Error de conversion de datos")

	article.IdArticle = idArticle
	article.Title = r.FormValue("title")
	article.Content = r.FormValue("content")
	article.Description = r.FormValue("description")

	err = manager.Update(&article)

	detectErr(err, "Error en la actualizacion del registro")

	http.Redirect(w, r, "/articles", http.StatusFound)
}

// Delete : DELETE - /articles/{id}/delete
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	article := models.Article{}
	idArticle, err := strconv.Atoi(vars["id"])

	detectErr(err, "Error en el envio de datos a la URL")

	article.IdArticle = idArticle
	err = manager.Delete(&article)

	detectErr(err, "Error en la eliminacion del registro")

	http.Redirect(w, r, "/articles", http.StatusFound)

}

// Privates

func detectErr(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
	}
}

func parseHTML(w http.ResponseWriter, temp string, location string, obj interface{}, nameFile string) {
	t := template.Must(template.ParseFiles(temp, location))
	t.ExecuteTemplate(w, nameFile, obj)
}
