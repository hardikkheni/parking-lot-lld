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

func TestItShouldGressfullyShutdown(t *testing.T) {
	sys := NewParkingLotSystem()
	command, err := sys.getCommand("exit", []string{})
	if err != nil {
		t.Error("Error: failed to create exit command. ", err.Error())
	}
	command.execute()
	_, err, done := command.execute()
	if err != nil {
		t.Error("Error: failed to execute exit command. ", err.Error())
	}
	if !done {
		t.Error("Error: failed to execute exit command correctly. ", err.Error())
	}
}

func TestItShouldParkAtNearestSlot(t *testing.T) {
	sys := NewParkingLotSystem()
	command, err := sys.getCommand("add_slot", []string{"1"})
	if err != nil {
		t.Error("Error: failed to create add_slot command. ", err.Error())
	}
	command.execute()
	command, err = sys.getCommand("park", []string{"KA-01-HH-9999", "Black"})
	if err != nil {
		t.Error("Error: failed to create park command. ", err.Error())
	}
	msg, err, _ := command.execute()
	if err != nil {
		t.Error("Error: failed to execute add_slot command. ", err.Error())
	}
	if msg != "The plate: KA-01-HH-9999 is parked at 1" {
		t.Error("Error: failed to execute add_slot command correctly. ", err.Error())
	}
}
