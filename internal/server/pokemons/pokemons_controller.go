package pokemons

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "github.com/pedrohsilvaml/pokemon-api/internal/service/pokemons"
)

type QueryParams struct {
	Name string `form:"name"`
}

type PokemonController struct {
	service service.PokemonService
}

func NewPokemonController() *PokemonController {
	return &PokemonController{
		service: *service.NewPokemonService(),
	}
}

func (ctl *PokemonController) GetPokemon(ctx *gin.Context) {
	var query QueryParams
	err := ctx.ShouldBindQuery(&query)

	if len(query.Name) == 0 || err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, &gin.H{
			"error": "'name' parameter is required",
		})
		return
	}

	response, err := ctl.service.GetPokemon(query.Name)

	if err == nil {
		ctx.JSON(http.StatusOK, response)
		return
	}

	ctx.JSON(http.StatusPartialContent, response)
}

func (ctl *PokemonController) GetInitialPokemons(ctx *gin.Context) {
	response, _ := ctl.service.GetInitialPokemons()
	ctx.JSON(http.StatusOK, response)
}
