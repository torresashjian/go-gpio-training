package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(10)
)

func main() {

}

func blinker() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

func helloHuman() {
	gbot := gobot.NewGobot()

	work := func() {
		gobot.Every(1*time.Second, func() {
			fmt.Println("Hello, human!")
		})
	}

	robot := gobot.NewRobot("drone",
		[]gobot.Connection{},
		[]gobot.Device{},
		work,
	)
	gbot.AddRobot(robot)

	gbot.Start()
}
