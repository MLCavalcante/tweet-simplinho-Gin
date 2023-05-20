package main

import (
 routes "github.com/MLCavalcante/api-gin/api/routes" // importar pkg routes

	    "github.com/gin-gonic/gin"
)

func main() {   //!1.Começar criando uma primeira rota 
	app := gin.Default() //!2. Criar uma instância do gin
        
    routes.AppRoutes(app) //chama a função de routes e passa o app
	

	app.Run("localhost:3001")
}
	