package postgresql

import (
	"errors"

	"github.com/daniel/apiRest2/models"
)

type ArticlePSQL struct{}

// Insercion
func (art ArticlePSQL) Create(article *models.Article) error {

	query := `INSERT INTO article(title, content, description)
				VALUES($1, $2, $3)`

	db := GetConnection()

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(article.Title, article.Content, article.Description)

	if err != nil {
		return err
	}

	nRows, _ := rows.RowsAffected()

	if nRows != 1 {
		return errors.New("Error: se afectaron mas filas de lo debido")
	}
	return nil
}

// Eliminación

func (art ArticlePSQL) Delete(article *models.Article) error {

	query := `DELETE FROM article WHERE id_article = $1`

	db := GetConnection()

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	row, err := stmt.Exec(article.IdArticle)

	if err != nil {
		return err
	}
	nRows, _ := row.RowsAffected()

	if nRows != 1 {
		return errors.New("Error: se eliminaron mas registros de lo debido")
	}
	return nil

}

// ACTUALIZACION
func (art ArticlePSQL) Update(article *models.Article) error {

	query := `UPDATE article SET title = $1, content=$2, description=$3
				WHERE id_article= $4`

	db := GetConnection()
	
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(article.Title, article.Content, article.Description, article.IdArticle)

	if err != nil {
		return err
	}

	nRows, _ := rows.RowsAffected()

	if nRows != 1 {
		return errors.New("Error: se afectaron mas filas de lo debido")
	}
	return nil

}

// OBTENCION DE UN REGISTRO

func (art ArticlePSQL) GetAll() ([]models.Article, error) {

	articles := make([]models.Article, 0)
	query := `SELECT id_article, title,content,description  FROM article`

	db := GetConnection()

	// Obtengo el set de resultados que obtuve gracias a la consulta o tambien puedo recibir un error
	resultSet, err := db.Query(query)

	// En caso de tener un error
	if err != nil {
		// Devuelvo los dos valores que debo devolver
		// En este caso el primer valor de retorno es una estructura vacia y el otro valor seria el error
		return articles, err
	}
	// Cierro el set de resultados
	defer resultSet.Close()

	// Se itera el set de resultados
	for resultSet.Next() {

		// Se crea una estructura Estudiante vacia
		a := models.Article{}

		// Se escanea el registro y se le ingresan los datos a la estructura
		err = resultSet.Scan(
			&a.IdArticle,
			&a.Title,
			&a.Content,
			&a.Description,
		)

		if err != nil {
			return articles, err
		}

		// Inserto en el slice estudiantes la estructura e
		articles = append(articles, a)
	}
	// Devuelvo el slice estudiantes
	return articles, nil

}

// Obtiene un registro

func (art ArticlePSQL) GetById(id int) (models.Article, error) {

	article := models.Article{}
	query := `SELECT id_article, title,content,description  FROM article WHERE id_article = $1`

	db := GetConnection()

	// Obtengo el set de resultados que obtuve gracias a la consulta o tambien puedo recibir un error
	resultSet, err := db.Query(query, id)

	// En caso de tener un error
	if err != nil {
		// Devuelvo los dos valores que debo devolver
		// En este caso el primer valor de retorno es una estructura vacia y el otro valor seria el error
		return article, err
	}
	// Cierro el set de resultados
	defer resultSet.Close()

	// Se itera el set de resultados
	for resultSet.Next() {

		// Se escanea el registro y se le ingresan los datos a la estructura
		err = resultSet.Scan(
			&article.IdArticle,
			&article.Title,
			&article.Content,
			&article.Description,
		)

		if err != nil {
			return article, err
		}
	}
	// Devuelvo el slice estudiantes
	return article, nil

}
