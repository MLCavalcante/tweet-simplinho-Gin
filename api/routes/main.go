package routes

import (
	controllers "github.com/MLCavalcante/api-gin/api/controllers"
	"github.com/gin-gonic/gin"
	)        

func AppRoutes(router *gin.Engine) *gin.RouterGroup { // aqui vamos criar um grupo de rotas relacionadas a tweet
	tweetController := controllers.NewTweetController() // para acessar a função que referencia a struct privada
	v1:= router.Group("/v1")
	{ //aqui colocamos alguams rotas relacionadas a api na sua versão v1
		v1.GET("/tweets", tweetController.FindAll)// quando bater em /tweets, colocamos o tweetController FindAll
		v1.POST("/tweet", tweetController.Create)
		v1.DELETE("/tweet/:id", tweetController.Delete)
	}
	return v1 //var tipo router group exatamente o que difinimos como retorno
}