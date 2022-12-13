package main

import (
	"bingo/config"
	"bingo/interface/handler"
	"bingo/pkg/calc"
	"bingo/pkg/logger"
	"bingo/usecase"
	"fmt"
	"os"
)

const (
	LengthOfASide = 5
)

func sample() {
	gu := usecase.NewGameUseCase()
	cg := gu.CreateGame([]string{"taro", "john"})
	game, err := gu.GetGame(cg.Id)
	_, err = gu.CallNumber(game.Id)
	_, err = gu.CallNumber(game.Id)
	game, err = gu.CallNumber(game.Id)
	fmt.Println(game)
	if err != nil {
		fmt.Printf(fmt.Sprintf("game update error: %s", err.Error()))
		return
	}
	logger.Info("test")
	logger.Info(game)

	ary := calc.CreateUserCard(LengthOfASide * LengthOfASide)
	fmt.Println(ary)
}

func main() {
	// sample()
	logger.Info("Env:", os.Getenv(config.ENV))
	err := handler.Initialize().Run()
	if err != nil {
		logger.Error(fmt.Sprintf("Run server error: %s", err.Error()))
		return
	}
}
