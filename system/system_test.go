package system

import "testing"

func TestParkingLotSystem(t *testing.T) {
	sys := NewParkingLotSystem()
	command, err := sys.getCommand("add_slot", []string{"1"})
	if err != nil {
		t.Error("Error: failed to create add_slot command. ", err.Error())
	}
	msg, err, _ := command.execute()
	if err != nil {
		t.Error("Error: failed to execute add_slot command. ", err.Error())
	}
	if msg != "1 Slots added!..." {
		t.Error("Error: failed to execute add_slot command correctly. ", err.Error())
	}
}
