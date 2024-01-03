package main

import (
	"errors"
)

type State struct {
	Command string
	Status  int
}

var States = make(map[int64]State)

func ClearState(chatID int64) {
	delete(States, chatID)
}

func GetState(chatID int64) (State, error) {
	status, stateOk := States[chatID]

	if !stateOk {
		return State{}, errors.New(Errors.NO_USER_STATE)
	}

	return status, nil
}

func SetState(chatID int64, command string, status int) {
	_, userOk := States[chatID]

	if !userOk {
		States[chatID] = State{
			Command: command,
			Status:  status,
		}
	}
}
