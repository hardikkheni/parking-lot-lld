package system

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type command interface {
	execute() (string, error, bool)
}

type exitCommand struct{}

func (c *exitCommand) execute() (string, error, bool) {
	return "System gracefully is sutting down!.....", nil, true
}

func newExitCommand() (*exitCommand, error) {
	return &exitCommand{}, nil
}

type addSlotCommand struct {
	state *state
	n     int
}

func (c *addSlotCommand) execute() (string, error, bool) {
	c.state.addSlot(c.n)
	return fmt.Sprint(c.n, " Slots added!..."), nil, false
}

func newAddSlotCommand(st *state, opts []string) (*addSlotCommand, error) {
	if len(opts) == 0 {
		return nil, errors.New("no params passed check the usage")
	}
	n, err := strconv.Atoi(opts[0])
	if err != nil {
		return nil, errors.New("wrong param passed expected int got " + reflect.TypeOf(opts[0]).String())
	}
	return &addSlotCommand{state: st, n: n}, nil
}

type statusCommand struct {
	state *state
}

func (c *statusCommand) execute() (string, error, bool) {
	nth, _, err := c.state.findNearestSlot()
	if err != nil {
		return "", err, false
	}
	return fmt.Sprintf("The nearest slot is no: %v", nth), nil, false
}

func newStatusCommand(st *state) (*statusCommand, error) {
	return &statusCommand{state: st}, nil
}

type parkCommand struct {
	state       *state
	plateNumber string
	color       string
}

func (c *parkCommand) execute() (string, error, bool) {
	nth, err := c.state.parkAtNearestEmptySlot(c.plateNumber, c.color)
	if err != nil {
		return "", err, false
	}
	return fmt.Sprintf("The plate: %s is parked at %d", c.plateNumber, nth), nil, false
}

func newParkCommand(st *state, opts []string) (*parkCommand, error) {
	return &parkCommand{state: st, plateNumber: opts[0], color: opts[1]}, nil
}
