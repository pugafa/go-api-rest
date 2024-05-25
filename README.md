Para efectos de l prueba 

los datos de conexxion de la bd se encuentran en harcode  en el archivo config/setup
y no en un .env como debe de ser

utilice mysql por que es la que uso en el trabajo actualmente

ejecucion
  1 cambiar datos de conexion bd(config/setup.go)

    dbUser := "root"
	dbPassword := ""
	dbName := "test"

  2 instalar dependencias
     go mod tidy
	
  3 ejecutar api
     go run main.go

 4 abrir postman e importar el archivo postman collection

