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

func Index(w http.ResponseWriter, r *http.Request) {

	log.Println(r.RequestURI)

	articles := make([]models.Article, 0)

	articles, err := manager.GetAll()

	detectErr(err, "Error en la obtencion de los articulos")

	//parseHtml(w, "./views/article/index.html", articles)
	t := template.Must(template.ParseFiles("./views/application/layout.tmpl", "./views/article/index.tmpl"))

	t.ExecuteTemplate(w, "articles", articles)
}

func Show(w http.ResponseWriter, r *http.Request) {

	log.Println(r.RequestURI)

	var id_article int

	article := models.Article{}

	vars := mux.Vars(r)

	id_article, err := strconv.Atoi(vars["id"])

	detectErr(err, "Error de conversion de datos")

	article, err = manager.GetById(id_article)

	detectErr(err, "No se ha encontrado el registro ")

	parseHtml(w, "./views/article/show.html", article)

}

func New(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/article/new.html")
}

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

func Edit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id_article, err := strconv.Atoi(vars["id"])

	detectErr(err, "Error de conversion de datos")

	article := models.Article{}

	article, err = manager.GetById(id_article)

	detectErr(err, "Error en la base de datos")

	parseHtml(w, "./views/article/edit.html", article)

}

func Update(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	article := models.Article{}

	id_article, err := strconv.Atoi(r.FormValue("id_article"))

	detectErr(err, "Error de conversion de datos")

	article.IdArticle = id_article
	article.Title = r.FormValue("title")
	article.Content = r.FormValue("content")
	article.Description = r.FormValue("description")

	err = manager.Update(&article)

	detectErr(err, "Error en la actualizacion del registro")

	http.Redirect(w, r, "/articles", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	article := models.Article{}
	id_article, err := strconv.Atoi(vars["id"])

	detectErr(err, "Error en el envio de datos a la URL")

	article.IdArticle = id_article
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

func parseHtml(w http.ResponseWriter, location string, obj interface{}) {
	parser, err := template.ParseFiles(location)
	detectErr(err, "Error al parcear el html")
	parser.Execute(w, obj)
}
