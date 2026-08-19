package main

import "time"

// struct to hold glyph data
type glyphinfo struct {
	size    int
	glyph   rune
	str     string
	unicode string
	decimal int32
}

// struct to hold start and end points for unicode ranges
// ranges from "unicode_ranges" file
type unicode_ranges struct {
	start int64
	end   int64
	name  string
}

// hold on do glyph data and x/y stuff
// FIXME: creation_date and last_moved_date are unused.
type dropstring struct {
	main_glyphs      []glyphinfo
	secondary_glyphs []glyphinfo
	creation_date    time.Time
	last_moved_date  time.Time
	x_pos            int
	y_pos            int
	x_max            int
	clear_y          int
}
