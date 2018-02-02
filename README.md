# MI PRIMER PROYECTO WEB BÁSICO CON GOLANG

En este proyecto se trata de realizar un ApiRestful consumido
por la misma aplicación.

Es un proyecto básico pero se trata de mostrar el proceso de
creación de un proyecto web con Golang.

## Pasos para correr el proyecto

1. Se necesita tener instalado postgreSQL en el sistema operativo.

2. Crear una base de datos llamada "blog".

3. Dentro de esa base de datos crear la siguiente tabla:

´´´
CREATE TABLE article(
  id_article SERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(60) NOT NULL,
  content CHARACTER VARYING NOT NULL,
  description VARCHAR(200) NOT NULL
)
´´´
4. Crear un archivo llamado configure.json e ingresarle los datos necesarios para la conexion a la base de datos:

´´´
{
  "engine" : "postgres",
  "user": "usuario",
  "password": "contraseña",
  "server": "localhost",
  "port": "5432",
  "database": "blog"
}
´´´

5. Luego ejecutar el comando:

´´´
go run main.go
´´´
