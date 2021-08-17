package data

import (
	"errors"
	"fmt"
)

type Validator interface {
	Validate() error
}

type User struct {
	Id              string `dynamodbav:"Id" json:"id"`
	Name            string `dynamodbav:"Name" json:"name"`
	Email           string `dynamodbav:"Email" json:"email"`
	SignupTimestamp int    `dynamodbav:"SignupTimestamp" json:"signupTimestamp"`
}

func (user *User) Validate() (err error) {
	switch {
	case user.Name == "":
		return errors.New("Wanted Name to not be nil, but Name is nil")
	case user.Email == "":
		return errors.New("Wanted Email to not be nil, but Email is nil")
	case user.SignupTimestamp < 0:
		return errors.New(fmt.Sprintf("Wanted SignupTimestamp to be 0 or above, but SignupTimestamp is %v", user.SignupTimestamp))
	default:
		return nil
	}
}

type Game struct {
	Id   string `dynamodbav:"Id" json:"id"`
	Name string `dynamodbav:"Name" json:"name"`
	// Credits          map[string]string `json:"credits"`
	// ReleaseTimestamp int               `json:"releaseDates"`
	// Platforms        []string          `json:"platforms"`
}

func (game *Game) Validate() (err error) {
	switch {
	case game.Name == "":
		return errors.New("Wanted Name to not be nil, but Name is nil")
	// case game.Credits == nil:
	// 	return errors.New("Wanted Credits to not be nil, but Credits is nil")
	// case game.ReleaseDates == nil:
	// 	return errors.New("Wanted ReleaseDates to not be nil, but ReleaseDates is nil")
	// case game.Platforms == nil:
	//	return errors.New("Wanted Platforms to not be nil, but Platforms is nil")
	default:
		return nil
	}
}

type GamePlay struct {
	UserId string `dynamodbav:"UserId" json:"userId"`
	GameId string `dynamodbav:"GameId" json:"gameId"`
}

func (gamePlay *GamePlay) Validate() (err error) {
	switch {
	case gamePlay.UserId == "":
		return errors.New("Wanted UserId to not be nil, but UserId is nil")
	case gamePlay.GameId == "":
		return errors.New("Wanted GameId to not be nil, but GameId is nil")
	default:
		return nil
	}
}

type GameRating struct {
	UserId string `dynamodbav:"UserId" json:"userId"`
	GameId string `dynamodbav:"GameId" json:"gameId"`
	Rating int    `dynamodbav:"Rating" json:"rating"`
}

func (gameRating *GameRating) Validate() (err error) {
	switch {
	case gameRating.UserId == "":
		return errors.New("Wanted UserId to not be nil, but UserId is nil")
	case gameRating.GameId == "":
		return errors.New("Wanted GameId to not be nil, but GameId is nil")
	case gameRating.Rating < 0 || gameRating.Rating > 10:
		return errors.New(fmt.Sprintf("Wanted Rating to be between 0 and 10 inclusive, but Rating is %v", gameRating.Rating))
	default:
		return nil
	}
}

func main() {
	//
}
