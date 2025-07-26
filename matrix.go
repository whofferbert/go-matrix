package main

/*
FLOW: establish screen size
clear screen
build list of printable monospace glyphs
while true ; do
	occasionally generate a slice of glyphs of random length
	start drawing glyphs from top of screen down
	each successive glyph should have a slightly different color green?
	glypsh past halfway point should start to fade to black
	occasionally change some of the glyphs

	figure out optimal refresh rate.
*/

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

// TODO not actually used.
var not_interrupted bool = true

//var debug = true

var debug_glyphs = false
var debug_colors = false

type glyphinfo struct {
	size    int
	glyph   rune
	str     string
	unicode string
	decimal int32
}

type unicode_ranges struct {
	start int64
	end   int64
	name  string
}

func get_random_x(max int) int {
	//given max X, return a random X smaller than that.
	rand_val := rand.IntN(max - 1)
	return rand_val + 1
}

func get_random_glyph(glyph_list *[]glyphinfo) glyphinfo {
	// return random glyph.
	num_glyphs := len(*glyph_list)
	random_idx := rand.Int() % num_glyphs
	return (*glyph_list)[random_idx]
}

func get_new_glyphstrings(glyph_list *[]glyphinfo, len int) ([]glyphinfo, []glyphinfo) {
	// make a new set of random characters to fall down the screen.
	var main_glyphs []glyphinfo
	var other_glyphs []glyphinfo
	//glyphs_len := get_random_x(min_glyphs_len)
	//glyphs_len += min_glyphs_len
	for i := 1; i <= len; i++ {
		this_glyph := get_random_glyph(glyph_list)
		main_glyphs = append(main_glyphs, this_glyph)
		another_glyph := get_random_glyph(glyph_list)
		other_glyphs = append(other_glyphs, another_glyph)
	}
	// always append space to end of glyph string...
	//main_glyphs = append(main_glyphs, (*glyph_list)[0])
	//other_glyphs = append(other_glyphs, (*glyph_list)[0])
	return main_glyphs, other_glyphs
}

func set_glyph_info(i int32) (this_glyph glyphinfo, err error) {
	err = nil //
	//this_glyph := newGlyphInfo()
	g := rune(i)
	if unicode.IsPrint(g) {
		this_glyph.decimal = i
		this_glyph.glyph = g
		this_glyph.str = fmt.Sprintf("%c", g)
		//fmt.Printf("%c", g)
		//fmt.Printf("%c\n", g)
		this_glyph.unicode = fmt.Sprintf("%U", g)
		_, size := width.LookupString(this_glyph.str)
		this_glyph.size = size
		return this_glyph, err
	} else {
		return this_glyph, err
	}
}

//func generate_glyphs_from_range(from int, to int, name string) []glyphinfo {
//
//}

func read_embed_string() []unicode_ranges {
	scanner := bufio.NewScanner(strings.NewReader(unicode_strings_embed))
	var rangeinfo []unicode_ranges
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		s := regexp.MustCompile(" +").Split(line, 3)
		//fmt.Printf("%#v\n", s)
		var this_rangeinfo unicode_ranges
		start, _ := strconv.ParseInt(s[0], 16, 32)
		end, _ := strconv.ParseInt(s[1], 16, 32)
		this_rangeinfo.start = start
		this_rangeinfo.end = end
		this_rangeinfo.name = s[2]
		//fmt.Printf("%#v\n", this_rangeinfo)
		rangeinfo = append(rangeinfo, this_rangeinfo)
	}
	return rangeinfo
}

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

func generate_glyphs() []glyphinfo {
	skip_decimal_ranges := setup_skip_ranges()
	var glyphs []glyphinfo
	unicode_ranges := read_embed_string()
	skipnextglyph := 0
	for _, unicode_range := range unicode_ranges {
		if debug_glyphs {
			fmt.Println("Working on range:", unicode_range.name)
		}
		for i := unicode_range.start; i <= unicode_range.end; i++ {
			if skipnextglyph == 1 {
				skipnextglyph = 0
				continue
			}
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
			if ginf.str == "₿" {
				skipnextglyph = 1
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

func should_make_new_dropping_string() bool {
	// last new strings time?
	ret := false
	if true {
		ret = true
	}
	return ret
}

type dropstring struct {
	// what do we need?
	main_glyphs      []glyphinfo
	secondary_glyphs []glyphinfo
	creation_date    time.Time
	last_moved_date  time.Time
	x_pos            int
	y_pos            int
	x_max            int
	clear_y          int
}

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

func redraw_dropstrings(scn tcell.Screen, cur *map[int]dropstring) {
	// NOTE for some reason, we can't update the dropstrings from in this func
	// all dropstring adjustments must be done elsewhere
	for _, ds := range *cur {
		if ds.y_pos > len(ds.main_glyphs) {
			// draw spaces to clear old stuff
			scn.SetContent(ds.x_pos, ds.clear_y, rune(32), []rune(""), tcell.StyleDefault)
		} else {
			// probably just need the most recent couple
			bright_idx_1 := rand.IntN(len(brt_colors))
			bright_idx_2 := rand.IntN(len(brt_colors))
			bright_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx_1])
			//dimmer_style := tcell.StyleDefault.Foreground(tcell.ColorGreenYellow)
			dimmer_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx_2])
			if ds.y_pos < len(ds.main_glyphs) {
				// TODO coloring things properly
				scn.SetContent(ds.x_pos, ds.y_pos, rune(ds.main_glyphs[ds.y_pos].decimal), []rune(""), bright_style)
			}
			if ds.y_pos-1 >= 0 {
				scn.SetContent(ds.x_pos, ds.y_pos-1, rune(ds.main_glyphs[ds.y_pos-1].decimal), []rune(""), dimmer_style)
			}
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
			// if we're in the last 3-4, pick random from fade.
			if y > len(ds.main_glyphs)-4 {
				// set fades
				fade_idx := rand.IntN(len(fade_colors))
				fade_style := tcell.StyleDefault.Foreground(fade_colors[fade_idx])
				scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), fade_style)
			} else {
				// else, do other colors
				// X% chance to change color
				if rand.IntN(100) < 30 {
					// x% chance to get a green or greenish blue color.
					if rand.IntN(100) < 5 {
						bright_idx := rand.IntN(len(brt_colors))
						bright_style := tcell.StyleDefault.Foreground(brt_colors[bright_idx])
						scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), bright_style)
					} else if rand.IntN(100) < 30 {
						// x% chance to get a green or greenish blue color.
						col_idx := rand.IntN(len(grn_colors))
						col_style := tcell.StyleDefault.Foreground(grn_colors[col_idx])
						scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), col_style)
					} else if rand.IntN(100) < 30 {
						col_idx := rand.IntN(len(grn_blu_colors))
						col_style := tcell.StyleDefault.Foreground(grn_blu_colors[col_idx])
						scn.SetContent(ds.x_pos, r_idx, rune(decimal), []rune(""), col_style)

					}

				}

			}
		}

		// TODO random color switching
		// TODO random glyph switching

	}
}

var fade_colors = get_fade_colors()
var grn_blu_colors = get_greenish_blue_colors()
var grn_colors = get_greenish_colors()
var brt_colors = get_brightish_colors()

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

	if debug_colors {
		dump_colors(scn, lines)
	}

	// keep track of dropstrings.
	var active_dropstrings = make(map[int]dropstring)
	num_dropstrings := 0

	// do the thing until interrupted.
	for not_interrupted {
		// break condition: esc/ctrl + c
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

		// check if need new dropstring
		// this should occur more often if the terminals are wider, less so if they are small
		// density
		// currently just always makes a new dropstring.
		if should_make_new_dropping_string() {
			// create new dropstring
			new_dropstring := build_new_dropstring(&glyphs, columns, lines)
			//fmt.Printf("%#v\n", new_dropstring)
			active_dropstrings[num_dropstrings] = new_dropstring
			num_dropstrings++
			//active_dropstrings = append(active_dropstrings, new_dropstring)

		}

		// redraw before incrementing.
		redraw_dropstrings(scn, &active_dropstrings)
		//
		// show is way better than sync.
		scn.Show()

		// update dropstrings on screen
		for _, ds := range active_dropstrings {
			// update struct by index instead of dereferencing
			//active_dropstrings[idx].y_pos = ds.y_pos + 1
			//if active_dropstrings[idx].y_pos > len(active_dropstrings[idx].main_glyphs) {
			//	active_dropstrings[idx].clear_y++
			//}
			ds.y_pos = ds.y_pos + 1
			if ds.y_pos > len(ds.main_glyphs) {
				ds.clear_y++
			}
		}

		// clear old dropstrings to keep memory usage okay
		for idx, ds := range active_dropstrings {
			len_dropstring := len(ds.main_glyphs)
			if ds.clear_y-len_dropstring > lines {
				// delete from array.
				delete(active_dropstrings, idx)
				num_dropstrings--
				//active_dropstrings[idx] = active_dropstrings[len(active_dropstrings)-1]
				//active_dropstrings = active_dropstrings[:len(active_dropstrings)-1]
			}
		}

		// repeat.
		time.Sleep(refresh_sleep_secs)
	}
	//fmt.Printf("% x\n", "a")
	/*scn.Clear()
	for {
		// stuff
		time.Sleep(refresh_sleep_secs)
	}
	//var ColorNames = map[string]Color{}
	scn.Show()
	time.Sleep(time.Second * 2)
	var ColorNames = map[string]tcell.Color{}
	fmt.Println("color names:", ColorNames)
	fmt.Println("colors: ", scn.Colors())
	*/
	time.Sleep(time.Second * 5)
	scn.Clear()
	scn.Fini()
	//fmt.Println("lines, cols: ", lines, columns)
	//fmt.Printf("%#v\n", active_dropstrings)
	os.Exit(0)
}
