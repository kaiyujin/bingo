package usecase

import (
	"bingo/domain/entity"
	"bingo/domain/repository"
	"bingo/pkg/calc"
	"bingo/pkg/logger"
	"fmt"
	"github.com/google/uuid"
)

type GameUseCase interface {
	GetGame(gameId string) (entity.Game, error)
	CallNumber(gameId string) (entity.Game, error)
	CreateGame(names []string) entity.Game
}

type gameUseCase struct {
	gameRepository repository.GameRepository
}

func NewGameUseCase() GameUseCase {
	return &gameUseCase{gameRepository: repository.NewGameRepository()}
}

func (u *gameUseCase) GetGame(gameId string) (entity.Game, error) {
	game, err := u.gameRepository.Get(gameId)
	if err != nil {
		logger.Error(fmt.Sprintf("game get error: %s", err.Error()))
		return entity.Game{}, err
	}
	return game, nil
}

func (u *gameUseCase) CallNumber(gameId string) (entity.Game, error) {
	game, err := u.gameRepository.CallNumber(gameId)
	return game, err
}

func (u *gameUseCase) CreateGame(names []string) entity.Game {
	game := entity.Game{
		Id:            uuid.New().String(),
		CalledNumbers: []int8{},
		Users:         []entity.User{},
	}
	for _, name := range names {
		user := entity.User{
			Id:      uuid.New().String(),
			Name:    name,
			Numbers: calc.CreateUserCard(25),
		}
		game.Users = append(game.Users, user)
	}
	u.gameRepository.Create(game)
	return game
}
