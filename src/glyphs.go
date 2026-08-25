package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"unicode"

	"golang.org/x/text/width"
)

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
