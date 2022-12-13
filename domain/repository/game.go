package repository

import (
	"bingo/domain/entity"
	"bingo/infrastructure"
	"bingo/pkg/calc"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type GameRepository interface {
	Create(game entity.Game)
	Get(gameId string) (entity.Game, error)
	CallNumber(gameId string) (entity.Game, error)
}

type gameRepository struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewGameRepository() GameRepository {
	g := new(gameRepository)
	g.TableName = "Game"
	g.DynamoDbClient = infrastructure.DynamoDBClient()
	return g
}

func (r gameRepository) Create(game entity.Game) {
	g, err := attributevalue.MarshalMap(game)
	if err != nil {
		fmt.Printf("dynamodb marshal: %s\n", err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), infrastructure.DynamoDefaultTimeout)
	defer cancel()
	_, err = r.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      g,
		TableName: aws.String(r.TableName),
	})
	if err != nil {
		fmt.Printf("dynamodb putItem: %s\n", err.Error())
		return
	}
}

func (r gameRepository) Get(gameId string) (entity.Game, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), infrastructure.DynamoDefaultTimeout)
	defer cancel()
	output, err := r.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{
				Value: gameId,
			},
		},
	})
	if err != nil {
		fmt.Printf("get item: %s\n", err.Error())
		return entity.Game{}, err
	}
	game := entity.Game{}
	err = attributevalue.UnmarshalMap(output.Item, &game)
	if err != nil {
		fmt.Printf("dynamodb unmarshal: %s\n", err.Error())
		return entity.Game{}, err
	}
	return game, nil
}

func (r gameRepository) CallNumber(gameId string) (entity.Game, error) {
	rn := calc.RandomNumber()
	game, err := r.Get(gameId)
	game.CalledNumbers = append(game.CalledNumbers, rn)
	update := expression.UpdateBuilder{}.Set(
		expression.Name("CalledNumbers"), expression.Value(game.CalledNumbers))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		fmt.Printf("build update expression: %s\n", err.Error())
		return entity.Game{}, err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), infrastructure.DynamoDefaultTimeout)
	defer cancel()
	_, err = r.DynamoDbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{
				Value: gameId,
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	})
	return game, err
}
