package data

import "testing"

func expectNil(t *testing.T, validator Validator) {
	err := validator.Validate()
	if err != nil {
		t.Fatal("Wanted error to be nil")
	}
}

func expectNotNil(t *testing.T, validator Validator) {
	err := validator.Validate()
	if err == nil {
		t.Fatal("Wanted error to be not nil")
	}
}

func TestUser(t *testing.T) {
	t.Run("Returns error when Name is empty", func(t *testing.T) {
		user := User{
			Id:              "1",
			Name:            "",
			Email:           "test@mail.com",
			SignupTimestamp: 1,
		}

		expectNotNil(t, &user)
	})

	t.Run("Returns error when Email is empty", func(t *testing.T) {
		user := User{
			Id:              "1",
			Name:            "Test",
			Email:           "",
			SignupTimestamp: 1,
		}

		expectNotNil(t, &user)
	})

	t.Run("Returns error when SignupTimestamp is not in range", func(t *testing.T) {
		user := User{
			Id:              "1",
			Name:            "Test",
			Email:           "test@mail.com",
			SignupTimestamp: -1,
		}

		expectNotNil(t, &user)
	})

	t.Run("Returns nil when all fields are filled", func(t *testing.T) {
		user := User{
			Id:              "1",
			Name:            "Test",
			Email:           "test@mail.com",
			SignupTimestamp: 1,
		}

		expectNil(t, &user)
	})
}

func TestGame(t *testing.T) {
	t.Run("Returns error when Name is empty", func(t *testing.T) {
		game := Game{
			Id:   "1",
			Name: "",
		}

		expectNotNil(t, &game)
	})

	t.Run("Returns nil when all fields are filled", func(t *testing.T) {
		game := Game{
			Id:   "1",
			Name: "Test",
		}

		expectNil(t, &game)
	})
}

func TestGamePlay(t *testing.T) {
	t.Run("Returns error when GameId is empty", func(t *testing.T) {
		gamePlay := GamePlay{
			UserId: "1",
			GameId: "",
		}

		expectNotNil(t, &gamePlay)
	})

	t.Run("Returns error when UserId is empty", func(t *testing.T) {
		gamePlay := GamePlay{
			UserId: "",
			GameId: "1",
		}

		expectNotNil(t, &gamePlay)
	})

	t.Run("Returns nil when all fields are filled", func(t *testing.T) {
		gamePlay := GamePlay{
			UserId: "1",
			GameId: "1",
		}

		expectNil(t, &gamePlay)
	})
}

func TestGameRating(t *testing.T) {
	t.Run("Returns error when GameId is empty", func(t *testing.T) {
		gameRating := GameRating{
			UserId: "1",
			GameId: "",
			Rating: 0,
		}

		expectNotNil(t, &gameRating)
	})

	t.Run("Returns error when UserId is empty", func(t *testing.T) {
		gameRating := GameRating{
			UserId: "",
			GameId: "1",
			Rating: 0,
		}

		expectNotNil(t, &gameRating)
	})

	t.Run("Returns error when Rating is out of range", func(t *testing.T) {
		gameRating := GameRating{
			UserId: "1",
			GameId: "1",
			Rating: -1,
		}

		expectNotNil(t, &gameRating)

		gameRating.Rating = 11
		expectNotNil(t, &gameRating)
	})

	t.Run("Returns nil when all fields are filled", func(t *testing.T) {
		gameRating := GameRating{
			UserId: "1",
			GameId: "1",
			Rating: 5,
		}

		expectNil(t, &gameRating)
	})
}
