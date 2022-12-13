package handler

import (
	"bingo/domain/entity"
	"bingo/interface/request"
	"bingo/interface/response"
	"bingo/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createGame(c *gin.Context) {
	var json request.PostedGame
	if err := c.ShouldBindJSON(&json); err != nil || len(json.Names) < 1 {
		c.JSON(http.StatusBadRequest, response.Error{
			Code:    strconv.Itoa(http.StatusBadRequest),
			Message: http.StatusText(http.StatusBadRequest),
			Details: []response.ErrorDetail{{
				Target:  "names",
				Message: err.Error()},
			},
		})
		return
	}
	if len(json.Names) > 10 {
		c.JSON(http.StatusBadRequest, response.Error{
			Code:    strconv.Itoa(http.StatusBadRequest),
			Message: http.StatusText(http.StatusBadRequest),
			Details: []response.ErrorDetail{{
				Target:  "names",
				Message: "Too many names. max 10."},
			},
		})
	}
	var details []response.ErrorDetail
	for i, name := range json.Names {
		if len(name) > 20 {
			details[i] = response.ErrorDetail{
				Target:  fmt.Sprintf("name[%d]", i),
				Message: "Too long.",
			}
		}
	}
	if len(details) > 0 {
		c.JSON(http.StatusBadRequest, response.Error{
			Code:    strconv.Itoa(http.StatusBadRequest),
			Message: http.StatusText(http.StatusBadRequest),
			Details: details,
		})
	}
	game := usecase.NewGameUseCase().CreateGame(json.Names)
	res := gameEntityToResponse(game)
	c.JSON(http.StatusOK, res)
}

func getGame(c *gin.Context) {
	id := c.Param("id")
	game, err := usecase.NewGameUseCase().GetGame(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error{
			Code:    "500-GAME",
			Message: "Internal server error",
			Details: nil,
		})
		return
	}
	if game.Id == "" {
		c.JSON(http.StatusNotFound, response.Error{
			Code:    "404",
			Message: "Not find game id",
			Details: nil,
		})
		return
	}
	res := gameEntityToResponse(game)
	c.JSON(http.StatusOK, res)
}

func callNumber(c *gin.Context) {
	//TODO implement me
	//panic("implement me")
}

func gameEntityToResponse(game entity.Game) response.Game {
	var resUsers []response.User
	for _, user := range game.Users {
		resUsers = append(resUsers, response.User{
			Id:      user.Id,
			Name:    user.Name,
			Numbers: user.Numbers,
		})
	}
	res := response.Game{
		Id:            game.Id,
		CalledNumbers: game.CalledNumbers,
		Users:         resUsers,
	}
	return res
}
