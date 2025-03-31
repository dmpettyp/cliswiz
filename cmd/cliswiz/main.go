package main

import (
	"time"

	"github.com/dmpettyp/cliswiz/internal/color"
	"github.com/dmpettyp/cliswiz/internal/commands"
	"github.com/dmpettyp/cliswiz/internal/device"
)

func main() {
	lamp := &device.Device{}

	lamp.SendCommand(commands.NewOnCommand())
	lamp.SendCommand(commands.NewDimCommand(100))
	// lamp.SendCommand(commands.NewColorCommand(255, 255, 255))

	for {
		for _, c := range color.Rainbow {
			for _, t := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
				lamp.SendCommand(commands.NewColorCommand(c.Tint(t)))
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
	// 	// lamp.SendCommand(commands.NewColorCommand(23, 234, 53))
	// 	// time.Sleep(1 * time.Second)
	// 	// lamp.SendCommand(commands.NewColorCommand(255, 0, 0))
	// 	// time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(209, 0, 0))
	// 	time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(255, 102, 34))
	// 	time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(255, 218, 33))
	// 	time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(51, 221, 0))
	// 	time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(17, 51, 204))
	// 	time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(34, 0, 102))
	// 	time.Sleep(1 * time.Second)
	// 	lamp.SendCommand(commands.NewColorCommand(51, 0, 68))
	// 	time.Sleep(1 * time.Second)
	// }

	// lamp.SendCommand(commands.NewOffCommand())
}
