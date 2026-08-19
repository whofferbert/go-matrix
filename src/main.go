package main

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	flag "github.com/jessevdk/go-flags"
)

// TODO - for random column numbers, keep track of the active columns, and prefer columns that are not being used already

// this sets up the glyphs we will be drawing
//
//go:embed unicode/unicode_ranges
var unicode_strings_embed string

var fade_colors = get_fade_colors()
var grn_blu_colors = get_greenish_blue_colors()
var grn_colors = get_greenish_colors()
var brt_colors = get_brightish_colors()

type Config struct {
	Refresh_Hz     int  `short:"z" long:"refresh-hz"`
	Debug_Glyphs   bool `long:"debug-glyphs"`
	Debug_Colors   bool `long:"debug-colors"`
	Screen_Density int  `short:"d" long:"screen-density"`
	Fade_Length    int  `short:"f" long:"fade-length"`
	Log_Cols       bool `long:"log-cols"`

	logfile            string
	refresh_sleep_secs time.Duration
	screen             tcell.Screen
	glyphs             []glyphinfo
	active_dropstrings []dropstring
	not_interrupted    bool

	x_zero int
	x_max  int
	y_zero int
	y_max  int
}

var config Config

func process_direct_arguments(process_args []string) *Config {
	flag_options := flag.PassDoubleDash
	parser := flag.NewParser(&config, flag.Options(flag_options))
	//leftover_args, err := parser.ParseArgs(process_args)
	_, err := parser.ParseArgs(process_args)
	if err != nil {
		fmt.Println("Encountered a problem parsing arguments:", err)
		fmt.Printf("See %s --help for more details.\n", os.Args[0])
		os.Exit(2)
	}

	return &config
}

func main() {
	setup()

	// possibly dump colors and exit.
	if config.Debug_Colors {
		dump_colors(config.screen, config.y_max)
		os.Exit(0)
	}

	// do the thing until interrupted.
	for config.not_interrupted {
		screen_handler()
	}
}
