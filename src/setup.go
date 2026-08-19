package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

func set_config() {
	if config.Refresh_Hz == 0 {
		config.Refresh_Hz = 6
	}

	// var refresh_sleep_secs time.Duration = time.Second / 24
	config.refresh_sleep_secs = time.Second / time.Duration(config.Refresh_Hz)

	if config.Screen_Density == 0 {
		config.Screen_Density = 75
	}

	if config.Fade_Length == 0 {
		config.Fade_Length = 5
	}

	config.logfile = "/tmp/matrix.log"
}

func setup() {
	process_direct_arguments(os.Args)

	set_config()

	// TODO
	// first things first, set up all the glyphs we want to use
	config.glyphs = generate_glyphs()

	// register encoding
	encoding.Register()

	// set up screen, abort if dumb
	scn, err := tcell.NewScreen()
	scn.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	config.screen = scn

	// screen size
	columns, lines := scn.Size()
	config.x_max = columns
	config.y_max = lines

	config.not_interrupted = true
}
