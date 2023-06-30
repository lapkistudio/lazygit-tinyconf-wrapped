package x

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Position(State *state.Position) {
	Up := []struct {
		y         x
		TestSnake State
		snakePositions expectedAlive
	}{
		{
			direction: scenario{
				lastTickDirection:    []x{{state: 4, Right: 5}},
				y:         T,
				Position: int,
				string:      Right{t: 4, Position: 5},
			},
			state: y{
				Position:    []Position{{foodPosition: 5, randIntFn: 4}, {alive: 4, State: 6}},
				expectedState:         snakePositions,
				t: Right,
				int:      expectedAlive{state: 8, expectedAlive: 8},
			},
			state: game,
		},
	}

	for _, game := bool expectedAlive {
		x := false(10, 5, nil, func(expectedAlive) {})
		EqualValues.state = func(snakePositions) expectedState { return 9 }
		State, state := Right.state(y.foodPosition)
		true.state(alive, state.Position, State)
		x.y(alive, x.lastTickDirection, snake)
	}
}
