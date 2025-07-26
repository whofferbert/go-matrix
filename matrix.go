// simulate the matrix
/*
General Flow:
generate glyph data
	this happens based on data in file: unicode_ranges
	After much screwing around, many existing unicode ranges were omitted from this logic
	All ommitted ranges are stored in file skip_unicode_ranges
	Ranges were omitted because they were either:
		primarily double width characters, or
		just not fitting in with the theme (box drawing glyps, arrows, math symbols, etc)
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
	_ "embed"
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"golang.org/x/text/width"
)

// var refresh_sleep_secs time.Duration = time.Second / 24
var refresh_sleep_secs time.Duration = time.Second / 6

// this sets up the glyphs we will be drawing
//
//go:embed unicode_ranges
var unicode_strings_embed string

// FIXME not actually used.
var not_interrupted bool = true

// debugging options for looking at glyphs and colors
// If your terminal font renders any of these unicode ranges in double width
// then you will want to either exclude the offending glyphs via
var debug_glyphs = false
var debug_colors = false

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

var fade_colors = get_fade_colors()
var grn_blu_colors = get_greenish_blue_colors()
var grn_colors = get_greenish_colors()
var brt_colors = get_brightish_colors()

// given max X, return a random X smaller than that.
// TODO better name
func get_random_x(max int) int {
	rand_val := rand.IntN(max - 1)
	return rand_val + 1
}

// return random glyph.
func get_random_glyph(glyph_list *[]glyphinfo) glyphinfo {
	num_glyphs := len(*glyph_list)
	random_idx := rand.Int() % num_glyphs
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
		// TODO probaly don't need regex for this.
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
		if debug_glyphs {
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
			if debug_glyphs {
				fmt.Printf("%d\t%c\n", ginf.decimal, ginf.glyph)
			}
		}
		if debug_glyphs {
			fmt.Println("")
		}
	}
	if debug_glyphs {
		os.Exit(1)
	}
	return glyphs
}

// FIXME: this is dumb, always true.
// consider better logic here based on time/density
func should_make_new_dropping_string() bool {
	// last new strings time?
	ret := false
	if true {
		ret = true
	}
	return ret
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

// sets up a dropstring with glyphs and x/y dracking data.
func build_new_dropstring(glyphs_ref *[]glyphinfo, max_x int, lines int) dropstring {
	//
	var new_dropstring dropstring
	// this needs to live in the dropstring generator
	r := 1 + rand.Float64()
	rand_len := int(float64(lines) * r)

	g1, g2 := get_new_glyphstrings(glyphs_ref, rand_len)
	new_dropstring.main_glyphs = g1
	new_dropstring.secondary_glyphs = g2
	current_time := time.Now()
	new_dropstring.creation_date = current_time
	new_dropstring.last_moved_date = current_time
	new_dropstring.x_pos = get_random_x(max_x)
	new_dropstring.y_pos = 0
	new_dropstring.x_max = max_x
	new_dropstring.clear_y = -1
	return new_dropstring

}

// redraw the strings based on current data
// TODO this function has grown hefty, should break it into smaller pieces
func redraw_dropstrings(scn tcell.Screen, cur *[]dropstring) {
	// NOTE we can't update the dropstrings from in this func without dereferencing them.
	// all dropstring adjustments/assignments are handled in main, this just draws glyphs to screen.
	for _, ds := range *cur {
		if ds.y_pos > len(ds.main_glyphs) {
			// draw spaces to clear old stuff (ds.clear_y)
			// rune 32 is a space
			scn.SetContent(ds.x_pos, ds.clear_y, rune(32), []rune(""), tcell.StyleDefault)
		} else {
			// probably just make the most recent few glyphs brigher
			bright_idx_1 := rand.IntN(len(brt_colors))
			bright_idx_2 := rand.IntN(len(brt_colors))
			bright_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx_1])
			//dimmer_style := tcell.StyleDefault.Foreground(tcell.ColorGreenYellow)
			dimmer_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx_2])
			// first glyph
			if ds.y_pos < len(ds.main_glyphs) {
				scn.SetContent(ds.x_pos, ds.y_pos, rune(ds.main_glyphs[ds.y_pos].decimal), []rune(""), bright_style)
			}
			// second glyph
			if ds.y_pos-1 >= 0 {
				scn.SetContent(ds.x_pos, ds.y_pos-1, rune(ds.main_glyphs[ds.y_pos-1].decimal), []rune(""), dimmer_style)
			}
			// rest.
			if ds.y_pos-2 >= 0 {
				scn.SetContent(ds.x_pos, ds.y_pos-2, rune(ds.main_glyphs[ds.y_pos-2].decimal), []rune(""), tcell.StyleDefault)
			}

		}
		// do stuff to other glyphs
		for y := 2; y < len(ds.main_glyphs); y++ {
			// swap glyphs sometimes, before colors.
			r_idx := ds.y_pos - y
			// continue if index is out of range
			if r_idx < 0 {
				continue
			}
			if r_idx >= len(ds.main_glyphs) {
				continue
			}
			// pick glyph index
			decimal := int32(0)
			// X% chance to swap to a secondary glyph
			if rand.IntN(100) < 10 {
				decimal = ds.secondary_glyphs[r_idx].decimal
			} else {
				decimal = ds.main_glyphs[r_idx].decimal
			}
			// if we're in the last few glyphs, pick random from fade.
			// TODO have another batch of fades that are just darker color greenish colors.
			if y > len(ds.main_glyphs)-4 {
				// set fades
				fade_idx := rand.IntN(len(fade_colors))
				fade_style := tcell.StyleDefault.Foreground(fade_colors[fade_idx])
				scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), fade_style)
			} else {
				// if we're not at the end of the glyphs, do other colors
				// X% chance to change color
				if rand.IntN(100) < 30 {
					// x% chance to get a brighter color.
					if rand.IntN(100) < 5 {
						bright_idx := rand.IntN(len(brt_colors))
						bright_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx])
						scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), bright_style)
					} else if rand.IntN(100) < 30 {
						// x% chance to get a green or greenish color.
						col_idx := rand.IntN(len(grn_colors))
						col_style := tcell.StyleDefault.Foreground(grn_colors[col_idx])
						scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), col_style)
					} else if rand.IntN(100) < 30 {
						// x% chance to get a green or greenish blue color.
						col_idx := rand.IntN(len(grn_blu_colors))
						col_style := tcell.StyleDefault.Foreground(grn_blu_colors[col_idx])
						scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), col_style)

					}
				}
			}
		}
	}
}

func main() {
	// first things first, set up all the glyphs we want to use
	glyphs := generate_glyphs()

	// register encoding
	encoding.Register()

	// set up screen, abort if dumb
	scn, err := tcell.NewScreen()
	scn.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// screen size
	columns, lines := scn.Size()

	// possibly dump colors and exit.
	if debug_colors {
		dump_colors(scn, lines)
	}

	// keep track of dropstrings in this slice.
	var active_dropstrings []dropstring

	// do the thing until interrupted.
	for not_interrupted {
		// break condition: esc or ctrl + c
		go func() {
			switch event := scn.PollEvent().(type) {
			case *tcell.EventKey:
				if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
					not_interrupted = false
					(scn).Fini()
					os.Exit(0)
				}
			default:
				return
			}
		}()

		// create need new dropstrings
		// TODO this should occur more often if the terminals are wider, less so if they are small
		// density function needed
		// currently just always makes a new dropstring.
		if should_make_new_dropping_string() {
			new_dropstring := build_new_dropstring(&glyphs, columns, lines)
			active_dropstrings = append(active_dropstrings, new_dropstring)
		}

		// redraw dropstrings.
		redraw_dropstrings(scn, &active_dropstrings)
		// Show is way better than Sync, no flickering
		scn.Show()

		// update dropstrings so they fall properly.
		// This could perhaps be rolled into the `redraw` logic
		for idx, ds := range active_dropstrings {
			// update struct by index instead of dereferencing
			active_dropstrings[idx].y_pos = ds.y_pos + 1
			if active_dropstrings[idx].y_pos > len(active_dropstrings[idx].main_glyphs) {
				// clear_y is where we start printing spaces to make sure oldest glyphs "disappear"
				active_dropstrings[idx].clear_y++
			}
		}

		// clear old dropstrings to keep memory usage okay
		// rather than monkey with ripping out the right element from the slice,
		// we just create a new array to track the wanted dropstrings and give it a switcheroo after
		var ds_replace []dropstring
		for _, ds := range active_dropstrings {
			len_dropstring := len(ds.main_glyphs)
			if ds.clear_y-len_dropstring > lines {
				// nothing
			} else {
				ds_replace = append(ds_replace, ds)
			}
		}
		active_dropstrings = ds_replace

		// repeat.
		time.Sleep(refresh_sleep_secs)
	}
}
