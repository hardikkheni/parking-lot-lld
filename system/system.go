package system

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ParkingLotSystem struct {
	reader *bufio.Reader
	state  *state
	in     chan []string
	out    chan struct {
		msg  string
		done bool
	}
}

func NewParkingLotSystem() *ParkingLotSystem {
	return &ParkingLotSystem{
		reader: bufio.NewReader(os.Stdin),
		state:  newState(),
		in:     make(chan []string),
		out: make(chan struct {
			msg  string
			done bool
		}),
	}
}

func (system ParkingLotSystem) renderWelcome() {
	fmt.Print(
		`*****************************************************************************************************
*                                                                                                   *
*   ########     ###    ########  ##    ## #### ##    ##  ######      ##        #######  ########   *
*   ##     ##   ## ##   ##     ## ##   ##   ##  ###   ## ##    ##     ##       ##     ##    ##      *
*   ##     ##  ##   ##  ##     ## ##  ##    ##  ####  ## ##           ##       ##     ##    ##      *
*   ########  ##     ## ########  #####     ##  ## ## ## ##   ####    ##       ##     ##    ##      *
*   ##        ######### ##   ##   ##  ##    ##  ##  #### ##    ##     ##       ##     ##    ##      *
*   ##        ##     ## ##    ##  ##   ##   ##  ##   ### ##    ##     ##       ##     ##    ##      *
*   ##        ##     ## ##     ## ##    ## #### ##    ##  ######      ########  #######     ##      *
*                                                                                                   *
*****************************************************************************************************` + "\n\n",
	)
	fmt.Println("Welcome to Parking Lot System!.")
}

func (system ParkingLotSystem) renderMenu() {
	fmt.Print("\n=====================================================================================================\n\n")
	fmt.Print("Options:- \n\n")
	list := []string{
		"add_slot        To add n slot to existing system usage: `add_slot 6` .",
		"status          To get nearest slot.",
		"exit            To close the system.",
	}
	for idx, item := range list {
		fmt.Println(fmt.Sprintf("%d", idx+1) + ". " + item)
	}
}

func (system ParkingLotSystem) getCommand(str string, opts []string) (command, error) {
	if str == "add_slot" {
		return newAddSlotCommand(system.state, opts)
	} else if str == "status" {
		return newStatusCommand(system.state)
	} else if str == "exit" {
		return newExitCommand()
	}
	return nil, errors.New("command doesn't exist")
}

func (system *ParkingLotSystem) Start() {
	system.renderWelcome()
	go system.scan()
	go func() {
		for result := range system.in {
			var opts []string
			if len(result) > 1 {
				opts = result[1:]
			}
			command, err := system.getCommand(result[0], opts)
			if err != nil {
				fmt.Print("Error: ", err.Error(), "\n\n")
				go system.scan()
				continue
			}
			msg, err, done := command.execute()
			if err != nil {
				fmt.Print("Error: ", err.Error(), "\n\n")
				go system.scan()
				continue
			}
			fmt.Println("Message: ", msg)
			if done {
				system.out <- struct {
					msg  string
					done bool
				}{
					msg:  msg,
					done: true,
				}
			} else {
				go system.scan()
			}
		}
	}()
	for result := range system.out {
		if result.done {
			return
		}
	}
}

func (system ParkingLotSystem) scan() {
	system.renderMenu()
	line, err := system.reader.ReadString('\n')
	if err == nil {
		line = line[:len(line)-1]
		strs := strings.Split(line, " ")
		system.in <- strs
	}

}
