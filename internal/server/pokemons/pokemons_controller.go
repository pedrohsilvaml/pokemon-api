package pokemons

import (
	"net/http"

	"github.com/gin-gonic/gin"

	poke_service "github.com/pedrohsilvaml/pokemon-api/internal/service/pokemons"
)

type QueryParams struct {
	Name string `form:"name"`
}

type PokemonController struct {
	Service *poke_service.PokemonService
}

func NewPokemonController(s *poke_service.PokemonService) *PokemonController {
	return &PokemonController{Service: s}
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

	response, err := ctl.Service.GetPokemon(query.Name)

	if err != nil {
		ctx.JSON(http.StatusPartialContent, response)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (ctl *PokemonController) GetInitialPokemons(ctx *gin.Context) {
	response := ctl.Service.GetInitialPokemons()

	for _, pokemonResponse := range response {
		if pokemonResponse.Partial {
			ctx.JSON(http.StatusPartialContent, response)
			return
		}
	}

	ctx.JSON(http.StatusOK, response)
}
