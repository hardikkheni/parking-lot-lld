package main

import "github.com/hardikkheni/parkinglotlld/system"

func main() {
	sys := system.NewParkingLotSystem()
	sys.Start()
}
