package main

import (
	"flag"
	"github.com/wjessop/go-piglow"
	"log"
	"time"
)

var max_brightness uint8 = 255
var p *piglow.Piglow

var sleep_time int

func init() {
	flag.IntVar(&sleep_time, "delay", 1000, "sleep time (ms) between fade brightnesses")
	flag.Parse()
}

func main() {
	var err error

	var delay time.Duration = time.Duration(sleep_time) * time.Millisecond

	// Create a new Piglow
	p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	zero()

	for {

		// Fade in blue, out white
		for i := 0; i <= int(max_brightness); i++ {
			p.SetBlue(uint8(i))
			p.SetRed(uint8(255 - i))
			apply()
			time.Sleep(delay)
		}

		// Fade in green, out blue
		for i := 0; i <= int(max_brightness); i++ {
			p.SetGreen(uint8(i))
			p.SetBlue(uint8(255 - i))
			apply()
			time.Sleep(delay)
		}

		// Fade in yellow, out green
		for i := 0; i <= int(max_brightness); i++ {
			p.SetYellow(uint8(i))
			p.SetGreen(uint8(255 - i))
			apply()
			time.Sleep(delay)
		}

		// Fade in orange, out yellow
		for i := 0; i <= int(max_brightness); i++ {
			p.SetOrange(uint8(i))
			p.SetYellow(uint8(255 - i))
			apply()
			time.Sleep(delay)
		}

		// Fade in red, out orange
		for i := 0; i <= int(max_brightness); i++ {
			p.SetRed(uint8(i))
			p.SetOrange(uint8(255 - i))
			apply()
			time.Sleep(delay)
		}
	}
}

func zero() {
	p.SetAll(0)
}

func apply() {
	err := p.Apply()
	if err != nil {
		log.Fatal("Couldn't apply changes: ", err)
	}
}
