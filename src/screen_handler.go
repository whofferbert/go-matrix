package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func screen_handler() {
	// break condition: esc or ctrl + c
	// if nothing provided, just continue on.
	go func() {
		switch event := config.screen.PollEvent().(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				config.not_interrupted = false
				(config.screen).Fini()
				os.Exit(0)
			}
		default:
			return
		}
	}()

	// create need new dropstrings
	if should_make_new_dropping_string() {
		new_dropstring := build_new_dropstring(&config.glyphs, config.x_max, config.y_max)
		config.active_dropstrings = append(config.active_dropstrings, new_dropstring)
	}

	// redraw dropstrings.
	redraw_dropstrings(config.screen, &config.active_dropstrings)
	// Show is way better than Sync, no flickering
	config.screen.Show()

	// update dropstrings so they fall properly.
	// This could perhaps be rolled into the `redraw` logic
	for idx, ds := range config.active_dropstrings {
		// update struct by index instead of dereferencing
		config.active_dropstrings[idx].y_pos = ds.y_pos + 1
		if config.active_dropstrings[idx].y_pos > len(config.active_dropstrings[idx].main_glyphs) {
			// clear_y is where we start printing spaces to make sure oldest glyphs "disappear"
			config.active_dropstrings[idx].clear_y++
		}
	}

	// clear old dropstrings to keep memory usage okay
	// rather than monkey with ripping out the right element from the slice,
	// we just create a new array to track the wanted dropstrings and give it a switcheroo after
	var ds_replace []dropstring
	for _, ds := range config.active_dropstrings {
		len_dropstring := len(ds.main_glyphs)
		if ds.clear_y-len_dropstring > config.y_max {
			// nothing
		} else {
			ds_replace = append(ds_replace, ds)
			// TODO - delete used index from config
			var replace_ds_used []int
			for _, x := range config.active_column_indexes {
				if x != ds.x_pos {
					replace_ds_used = append(replace_ds_used, x)
				}
			}
			config.active_column_indexes = replace_ds_used
		}
	}
	config.active_dropstrings = ds_replace

	// repeat.
	time.Sleep(config.refresh_sleep_secs)
}
