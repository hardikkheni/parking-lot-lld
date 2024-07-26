package system

import (
	"errors"
)

type Slot struct {
	plateNumber string
	color       string
	parked      bool
}

type state struct {
	slots []*Slot
}

func newState() *state {
	return &state{}
}

func (st *state) addSlot(n int) {
	for i := 0; i < n; i++ {
		st.slots = append(st.slots, newSlot())
	}
}

func (st *state) findNearestSlot() (int, *Slot, error) {
	for idx, slot := range st.slots {
		if !slot.parked {
			return idx + 1, slot, nil
		}
	}
	return -1, nil, errors.New("haven't created slots yet")
}

func (st *state) parkAtNearestEmptySlot(platNumber string, color string) (int, error) {
	nth, slot, err := st.findNearestSlot()
	if err != nil {
		return -1, err
	}
	slot.park(platNumber, color)
	return nth, nil
}

func newSlot() *Slot {
	return &Slot{
		parked: false,
	}
}

func (slot *Slot) park(plateNumber string, color string) {
	slot.parked = true
	slot.plateNumber = plateNumber
	slot.color = color
}

func (slot *Slot) clear(plateNumber string, color string) {
	slot.parked = false
	slot.plateNumber = ""
	slot.color = ""
}
