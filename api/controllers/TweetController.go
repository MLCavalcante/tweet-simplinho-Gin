package controllers

import (
	"net/http"

	entities "github.com/MLCavalcante/api-gin/api/entities"
	"github.com/gin-gonic/gin"
)

//! t em tweet minúsculo = struct privada não poderá ser usada fora desta página
type tweetController struct {
	tweets []entities.Tweet //aqui teremos uma var tweet que vai ser uma slice de Tweet
}

func NewTweetController() *tweetController {

	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets) 
} 

func (t *tweetController) Create(ctx *gin.Context) { //função tipo post gera um body da req 
	tweet := entities.NewTweet() // nesse momento temos um novo tweet instanciado, pegamos o body e vamos fazer um append ( colocar no final do nosso slice) para isso, atribuímos a essa nova var tweet.
     // O BindJSon vai unir o nosso body(ex: {"description": "tuiter"}) com o id da struct de Newtweet em tweet
	if err := ctx.BindJSON(&tweet); err != nil { // se der erro
		return 
	} 

	t.tweets = append(t.tweets, *tweet)//está adicionando um novo tweet ao nosso conjunto de tweets existente
    ctx.JSON(http.StatusOK, t.tweets)
} 

func (t *tweetController) Delete(ctx *gin.Context) { // para deletar um tweet específico precisamos receber um param
    id := ctx.Param("id")

	for idx, tweet := range t.tweets{
      if tweet.ID == id{ //se o meu tweet.id for igual ao id do param
		t.tweets = append(t.tweets[0:idx], t.tweets[idx+1:]... )
        return
	  }
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "tweet not found",
	}) 
} 


