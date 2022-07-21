package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	poke_api "github.com/pedrohsilvaml/pokemon-api/internal/client"
)

type QueryParams struct {
	Name string `form:"name"`
}

func main() {
	route := gin.Default()
	route.GET("/pokemon", HandlePokemonSearch)
	route.Run()
}

func HandlePokemonSearch(c *gin.Context) {
	var query QueryParams
	statusResponse := http.StatusPartialContent

	if err := c.ShouldBindQuery(&query); err == nil {
		response, err := poke_api.GetPokemon(query.Name)
		log.Println("Log error", err)
		partial := (err != nil)

		if !partial {
			statusResponse = http.StatusOK
		}

		c.JSON(statusResponse, gin.H{
			"data":    response.Data,
			"error":   err,
			"partial": partial,
		})
	} else {
		c.JSON(statusResponse, gin.H{
			"data":    gin.H{},
			"partial": true,
			"error":   err,
		})
	}
}
