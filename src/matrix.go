// simulate the matrix
/*
General Flow:
generate glyph data
	this happens based on data in file: unicode_ranges
	After much screwing around, many existing unicode ranges were omitted from this logic
	All omitted ranges are stored in file skip_unicode_ranges
	Ranges were omitted because they were either:
		primarily double width characters, or
		just not fitting in with the theme (box drawing glyphs, arrows, math symbols, etc)
set up screen, establish screen size
generate color data
while true ; do

	generate a slice of glyphs of random length
	start drawing glyphs from top of screen down
	glyph randomlly shift to secondary glyphs and other colors.
	glyphs toward the end should start to fade to black.

	repeat after refresh timing.
	break on esc or crtl+c
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell"
	"golang.org/x/text/width"
)

// given max X, return a random X smaller than that.
func get_random_x(max int) int {
	rand_val := rand.IntN(max + 1)
	return rand_val
}

// return random glyph.
func get_random_glyph(glyph_list *[]glyphinfo) glyphinfo {
	num_glyphs := len(*glyph_list)
	//random_idx := rand.Int() % num_glyphs
	random_idx := rand.IntN(num_glyphs)
	return (*glyph_list)[random_idx]
}

// make a new set of random characters which will fall down the screen.
func get_new_glyphstrings(glyph_list *[]glyphinfo, len int) ([]glyphinfo, []glyphinfo) {
	var main_glyphs []glyphinfo
	var other_glyphs []glyphinfo
	for i := 1; i <= len; i++ {
		this_glyph := get_random_glyph(glyph_list)
		main_glyphs = append(main_glyphs, this_glyph)
		another_glyph := get_random_glyph(glyph_list)
		other_glyphs = append(other_glyphs, another_glyph)
	}
	return main_glyphs, other_glyphs
}

// set general glyph data based on the code point.
func set_glyph_info(i int32) (this_glyph glyphinfo, err error) {
	err = nil //
	g := rune(i)
	if unicode.IsPrint(g) {
		// decimal code point
		this_glyph.decimal = i
		// glyph
		this_glyph.glyph = g
		// string representation
		this_glyph.str = fmt.Sprintf("%c", g)
		// unicode
		this_glyph.unicode = fmt.Sprintf("%U", g)
		// width TODO: not used since it's unreliable
		_, size := width.LookupString(this_glyph.str)
		this_glyph.size = size
		return this_glyph, err
	} else {
		return this_glyph, err
	}
}

// get sets of unicode ranges from embedded file.
func read_embed_string() []unicode_ranges {
	scanner := bufio.NewScanner(strings.NewReader(unicode_strings_embed))
	var rangeinfo []unicode_ranges
	for scanner.Scan() {
		line := scanner.Text()
		// TODO probably don't need regex for this.
		s := regexp.MustCompile(" +").Split(line, 3)
		var this_rangeinfo unicode_ranges
		start, _ := strconv.ParseInt(s[0], 16, 32)
		end, _ := strconv.ParseInt(s[1], 16, 32)
		this_rangeinfo.start = start
		this_rangeinfo.end = end
		this_rangeinfo.name = s[2]
		rangeinfo = append(rangeinfo, this_rangeinfo)
	}
	return rangeinfo
}

// struct to hold glyph skip ranges
type skip_ranges struct {
	start_skip int32
	end_skip   int32
}

func new_skip_range(start int32, end int32) skip_ranges {
	var skip_range skip_ranges
	skip_range.start_skip = start
	skip_range.end_skip = end
	return skip_range
}

// this provides a slice of glyph decimal references which will be explicitly skipped
// that's because these glyphs draw too wide or don't render properly.
// perhaps there's a better way to do this
func setup_skip_ranges() []skip_ranges {
	var skip_glyph_ranges []skip_ranges
	skip_glyph_ranges = append(skip_glyph_ranges, new_skip_range(497, 499))
	skip_glyph_ranges = append(skip_glyph_ranges, new_skip_range(452, 460))
	skip_glyph_ranges = append(skip_glyph_ranges, new_skip_range(329, 329))
	skip_glyph_ranges = append(skip_glyph_ranges, new_skip_range(1155, 1161))
	skip_glyph_ranges = append(skip_glyph_ranges, new_skip_range(42654, 42759))
	skip_glyph_ranges = append(skip_glyph_ranges, new_skip_range(42607, 42621))
	return skip_glyph_ranges
}

// generate the glyph data we'll use for drawing random stuff to screen
// this comes from the embedded file 'unicode_ranges'
// the chosen ranges were picked because they seemed to fit well enough,
// and all seem to draw in single width glyphs with my terminal font (Julia)
// https://github.com/cormullion/juliamono <- excellent font.
func generate_glyphs() []glyphinfo {
	skip_decimal_ranges := setup_skip_ranges()
	var glyphs []glyphinfo
	unicode_ranges := read_embed_string()
	for _, unicode_range := range unicode_ranges {
		if config.Debug_Glyphs {
			fmt.Println("Working on range:", unicode_range.name)
		}
		for i := unicode_range.start; i <= unicode_range.end; i++ {
			ginf, err := set_glyph_info(int32(i))
			if err != nil {
				continue
			} else if ginf.decimal == 0 {
				continue
			}
			should_skip := false
			for _, skip_range := range skip_decimal_ranges {
				if ginf.decimal >= skip_range.start_skip && ginf.decimal <= skip_range.end_skip {
					should_skip = true
				}
			}
			if should_skip {
				continue
			}
			glyphs = append(glyphs, ginf)
			if config.Debug_Glyphs {
				fmt.Printf("%d\t%c\n", ginf.decimal, ginf.glyph)
			}
		}
		if config.Debug_Glyphs {
			fmt.Println("")
		}
	}
	if config.Debug_Glyphs {
		os.Exit(1)
	}
	return glyphs
}

// make new dropstrings based on screen density
// this should occur more often if the terminals are wider, less so if they are small
func should_make_new_dropping_string() bool {
	// last new strings time?
	var num_dropstrings, dropstring_len_total, dropstring_len_avg int
	for _, x := range config.active_dropstrings {
		num_dropstrings++
		this_len := len(x.main_glyphs)

		if x.y_pos > config.y_max {
			minus := x.y_pos - config.y_max
			this_len = this_len - minus
		}
		dropstring_len_total += this_len
	}
	if num_dropstrings < 10 {
		return true
	}

	dropstring_len_avg = dropstring_len_total / num_dropstrings

	// we know average length, and number
	// compute area of screen size
	screen_area := config.x_max * config.y_max
	dropstring_area := num_dropstrings * dropstring_len_avg

	ret := false

	if float64(dropstring_area)/float64(screen_area) < float64(config.Screen_Density)/100 {
		ret = true
	}

	return ret
}

// sets up a dropstring with glyphs and x/y dracking data.
func build_new_dropstring(glyphs_ref *[]glyphinfo, max_x int, lines int) dropstring {
	//
	var new_dropstring dropstring
	// this needs to live in the dropstring generator
	// TODO - verify logic of random length of lines, when they stop printing, etc.
	r := 1 + rand.Float64()
	rand_len := int(float64(lines) * r)

	g1, g2 := get_new_glyphstrings(glyphs_ref, rand_len)
	new_dropstring.main_glyphs = g1
	new_dropstring.secondary_glyphs = g2
	current_time := time.Now()
	new_dropstring.creation_date = current_time
	new_dropstring.last_moved_date = current_time
	new_dropstring.x_pos = get_random_x(max_x)
	if config.Log_Cols {
		// TODO log column choice
	}
	new_dropstring.y_pos = 0
	new_dropstring.x_max = max_x
	new_dropstring.clear_y = -1
	return new_dropstring

}

func redraw_dropstring_bright_glyphs(ds dropstring) {
	bright_idx_1 := rand.IntN(len(brt_colors))
	bright_idx_2 := rand.IntN(len(brt_colors))
	bright_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx_1])
	//dimmer_style := tcell.StyleDefault.Foreground(tcell.ColorGreenYellow)
	dimmer_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx_2])
	// first glyph
	if ds.y_pos < len(ds.main_glyphs) {
		config.screen.SetContent(ds.x_pos, ds.y_pos, rune(ds.main_glyphs[ds.y_pos].decimal), []rune(""), bright_style)
	}
	// second glyph
	if ds.y_pos-1 >= 0 {
		config.screen.SetContent(ds.x_pos, ds.y_pos-1, rune(ds.main_glyphs[ds.y_pos-1].decimal), []rune(""), dimmer_style)
	}
	// rest.
	// TODO - is this proper to do here? are we duplicating efforts?
	if ds.y_pos-2 >= 0 {
		config.screen.SetContent(ds.x_pos, ds.y_pos-2, rune(ds.main_glyphs[ds.y_pos-2].decimal), []rune(""), tcell.StyleDefault)
	}
}

func redraw_dropstring_fade_glyphs(ds dropstring, r_idx int, decimal int32) {
	// set fades
	fade_idx := rand.IntN(len(fade_colors))
	fade_style := tcell.StyleDefault.Foreground(fade_colors[fade_idx])
	config.screen.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), fade_style)
}

func redraw_dropstring_regular_glyphs(ds dropstring, r_idx int, decimal int32) {
	// if we're not at the end of the glyphs, do other colors
	// X% chance to change color
	if rand.IntN(100) < 30 {
		// x% chance to get a brighter color.
		if rand.IntN(100) < 5 {
			bright_idx := rand.IntN(len(brt_colors))
			bright_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx])
			config.screen.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), bright_style)
		} else if rand.IntN(100) < 30 {
			// x% chance to get a green or greenish color.
			col_idx := rand.IntN(len(grn_colors))
			col_style := tcell.StyleDefault.Foreground(grn_colors[col_idx])
			config.screen.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), col_style)
		} else if rand.IntN(100) < 30 {
			// x% chance to get a green or greenish blue color.
			col_idx := rand.IntN(len(grn_blu_colors))
			col_style := tcell.StyleDefault.Foreground(grn_blu_colors[col_idx])
			config.screen.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), col_style)

		}
	}
}

func dropstring_swap_glyphs(ds dropstring, y int) (int, int32, error) {
	// swap glyphs sometimes, before colors.
	r_idx := ds.y_pos - y
	// continue if index is out of range
	if r_idx < 0 {
		return 0, 0, errors.New("Outside index range")
	}
	if r_idx >= len(ds.main_glyphs) {
		return 0, 0, errors.New("Outside index range")
	}
	// pick glyph index
	decimal := int32(0)
	// X% chance to swap to a secondary glyph
	if rand.IntN(100) < 10 {
		decimal = ds.secondary_glyphs[r_idx].decimal
	} else {
		decimal = ds.main_glyphs[r_idx].decimal
	}
	return r_idx, decimal, nil
}

// redraw the strings based on current data
func redraw_dropstrings(scn tcell.Screen, cur *[]dropstring) {
	// NOTE we can't update the dropstrings from in this func without dereferencing them.
	// all dropstring adjustments/assignments are handled in main, this just draws glyphs to screen.
	for _, ds := range *cur {
		if ds.y_pos > len(ds.main_glyphs) {
			// draw spaces to clear old stuff (ds.clear_y)
			// rune 32 is a space
			scn.SetContent(ds.x_pos, ds.clear_y, rune(32), []rune(""), tcell.StyleDefault)
		} else {
			redraw_dropstring_bright_glyphs(ds)
		}
		// do stuff to other glyphs
		for y := 2; y < len(ds.main_glyphs); y++ {
			r_idx, decimal, err := dropstring_swap_glyphs(ds, y)
			if err != nil {
				continue
			}

			// if we're in the last few glyphs, pick random from fade.
			// TODO have another batch of fades that are just darker color greenish colors.
			if y > len(ds.main_glyphs)-config.Fade_Length {
				redraw_dropstring_fade_glyphs(ds, r_idx, decimal)
			} else {
				// if we're not at the end of the glyphs, do other colors
				// X% chance to change color
				redraw_dropstring_regular_glyphs(ds, r_idx, decimal)
			}
		}
	}
}
