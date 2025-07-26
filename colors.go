package main

import (
	"os"
	"sort"
	"time"

	"github.com/gdamore/tcell"
)

// fade colors, lighter to darker
func get_fade_colors() []tcell.Color {
	var fade_colors []tcell.Color
	fade_colors = append(fade_colors, tcell.Color237)
	fade_colors = append(fade_colors, tcell.Color236)
	fade_colors = append(fade_colors, tcell.Color235)
	fade_colors = append(fade_colors, tcell.Color234)
	fade_colors = append(fade_colors, tcell.Color233)
	fade_colors = append(fade_colors, tcell.Color232)
	return fade_colors
}

func get_greenish_blue_colors() []tcell.Color {
	var g_colors []tcell.Color
	g_colors = append(g_colors, tcell.Color23)
	g_colors = append(g_colors, tcell.Color29)
	g_colors = append(g_colors, tcell.Color35)
	g_colors = append(g_colors, tcell.Color47)
	g_colors = append(g_colors, tcell.Color48)
	return g_colors
}

func get_greenish_colors() []tcell.Color {
	var g_colors []tcell.Color
	g_colors = append(g_colors, tcell.Color22)
	g_colors = append(g_colors, tcell.Color28)
	g_colors = append(g_colors, tcell.Color34)
	g_colors = append(g_colors, tcell.Color40)
	g_colors = append(g_colors, tcell.Color46)
	g_colors = append(g_colors, tcell.Color64)
	g_colors = append(g_colors, tcell.Color65)
	g_colors = append(g_colors, tcell.Color70)
	g_colors = append(g_colors, tcell.Color71)
	g_colors = append(g_colors, tcell.Color76)
	g_colors = append(g_colors, tcell.Color77)
	g_colors = append(g_colors, tcell.Color82)
	g_colors = append(g_colors, tcell.Color83)
	return g_colors
}

func get_brightish_colors() []tcell.Color {
	var b_colors []tcell.Color
	b_colors = append(b_colors, tcell.Color121)
	b_colors = append(b_colors, tcell.Color122)
	b_colors = append(b_colors, tcell.Color123)
	b_colors = append(b_colors, tcell.Color231)
	b_colors = append(b_colors, tcell.Color230)
	b_colors = append(b_colors, tcell.Color86)
	b_colors = append(b_colors, tcell.Color87)
	b_colors = append(b_colors, tcell.Color250)
	b_colors = append(b_colors, tcell.Color251)
	b_colors = append(b_colors, tcell.Color252)
	b_colors = append(b_colors, tcell.Color253)
	b_colors = append(b_colors, tcell.Color254)
	b_colors = append(b_colors, tcell.Color255)
	b_colors = append(b_colors, tcell.Color229)
	b_colors = append(b_colors, tcell.Color228)
	b_colors = append(b_colors, tcell.Color153)
	b_colors = append(b_colors, tcell.Color152)
	b_colors = append(b_colors, tcell.Color151)
	b_colors = append(b_colors, tcell.Color157)
	b_colors = append(b_colors, tcell.Color158)
	b_colors = append(b_colors, tcell.Color159)
	b_colors = append(b_colors, tcell.Color192)
	b_colors = append(b_colors, tcell.Color193)
	b_colors = append(b_colors, tcell.Color194)
	b_colors = append(b_colors, tcell.Color195)
	return b_colors
}

func dump_colors(scn tcell.Screen, lines int) {
	colors := make(map[string]tcell.Color)
	colors["ColorBlack"] = tcell.ColorBlack
	colors["ColorMaroon"] = tcell.ColorMaroon
	colors["ColorGreen"] = tcell.ColorGreen
	colors["ColorOlive"] = tcell.ColorOlive
	colors["ColorNavy"] = tcell.ColorNavy
	colors["ColorPurple"] = tcell.ColorPurple
	colors["ColorTeal"] = tcell.ColorTeal
	colors["ColorSilver"] = tcell.ColorSilver
	colors["ColorGray"] = tcell.ColorGray
	colors["ColorRed"] = tcell.ColorRed
	colors["ColorLime"] = tcell.ColorLime
	colors["ColorYellow"] = tcell.ColorYellow
	colors["ColorBlue"] = tcell.ColorBlue
	colors["ColorFuchsia"] = tcell.ColorFuchsia
	colors["ColorAqua"] = tcell.ColorAqua
	colors["ColorWhite"] = tcell.ColorWhite
	colors["Color16"] = tcell.Color16
	colors["Color17"] = tcell.Color17
	colors["Color18"] = tcell.Color18
	colors["Color19"] = tcell.Color19
	colors["Color20"] = tcell.Color20
	colors["Color21"] = tcell.Color21
	colors["Color22"] = tcell.Color22
	colors["Color23"] = tcell.Color23
	colors["Color24"] = tcell.Color24
	colors["Color25"] = tcell.Color25
	colors["Color26"] = tcell.Color26
	colors["Color27"] = tcell.Color27
	colors["Color28"] = tcell.Color28
	colors["Color29"] = tcell.Color29
	colors["Color30"] = tcell.Color30
	colors["Color31"] = tcell.Color31
	colors["Color32"] = tcell.Color32
	colors["Color33"] = tcell.Color33
	colors["Color34"] = tcell.Color34
	colors["Color35"] = tcell.Color35
	colors["Color36"] = tcell.Color36
	colors["Color37"] = tcell.Color37
	colors["Color38"] = tcell.Color38
	colors["Color39"] = tcell.Color39
	colors["Color40"] = tcell.Color40
	colors["Color41"] = tcell.Color41
	colors["Color42"] = tcell.Color42
	colors["Color43"] = tcell.Color43
	colors["Color44"] = tcell.Color44
	colors["Color45"] = tcell.Color45
	colors["Color46"] = tcell.Color46
	colors["Color47"] = tcell.Color47
	colors["Color48"] = tcell.Color48
	colors["Color49"] = tcell.Color49
	colors["Color50"] = tcell.Color50
	colors["Color51"] = tcell.Color51
	colors["Color52"] = tcell.Color52
	colors["Color53"] = tcell.Color53
	colors["Color54"] = tcell.Color54
	colors["Color55"] = tcell.Color55
	colors["Color56"] = tcell.Color56
	colors["Color57"] = tcell.Color57
	colors["Color58"] = tcell.Color58
	colors["Color59"] = tcell.Color59
	colors["Color60"] = tcell.Color60
	colors["Color61"] = tcell.Color61
	colors["Color62"] = tcell.Color62
	colors["Color63"] = tcell.Color63
	colors["Color64"] = tcell.Color64
	colors["Color65"] = tcell.Color65
	colors["Color66"] = tcell.Color66
	colors["Color67"] = tcell.Color67
	colors["Color68"] = tcell.Color68
	colors["Color69"] = tcell.Color69
	colors["Color70"] = tcell.Color70
	colors["Color71"] = tcell.Color71
	colors["Color72"] = tcell.Color72
	colors["Color73"] = tcell.Color73
	colors["Color74"] = tcell.Color74
	colors["Color75"] = tcell.Color75
	colors["Color76"] = tcell.Color76
	colors["Color77"] = tcell.Color77
	colors["Color78"] = tcell.Color78
	colors["Color79"] = tcell.Color79
	colors["Color80"] = tcell.Color80
	colors["Color81"] = tcell.Color81
	colors["Color82"] = tcell.Color82
	colors["Color83"] = tcell.Color83
	colors["Color84"] = tcell.Color84
	colors["Color85"] = tcell.Color85
	colors["Color86"] = tcell.Color86
	colors["Color87"] = tcell.Color87
	colors["Color88"] = tcell.Color88
	colors["Color89"] = tcell.Color89
	colors["Color90"] = tcell.Color90
	colors["Color91"] = tcell.Color91
	colors["Color92"] = tcell.Color92
	colors["Color93"] = tcell.Color93
	colors["Color94"] = tcell.Color94
	colors["Color95"] = tcell.Color95
	colors["Color96"] = tcell.Color96
	colors["Color97"] = tcell.Color97
	colors["Color98"] = tcell.Color98
	colors["Color99"] = tcell.Color99
	colors["Color100"] = tcell.Color100
	colors["Color101"] = tcell.Color101
	colors["Color102"] = tcell.Color102
	colors["Color103"] = tcell.Color103
	colors["Color104"] = tcell.Color104
	colors["Color105"] = tcell.Color105
	colors["Color106"] = tcell.Color106
	colors["Color107"] = tcell.Color107
	colors["Color108"] = tcell.Color108
	colors["Color109"] = tcell.Color109
	colors["Color110"] = tcell.Color110
	colors["Color111"] = tcell.Color111
	colors["Color112"] = tcell.Color112
	colors["Color113"] = tcell.Color113
	colors["Color114"] = tcell.Color114
	colors["Color115"] = tcell.Color115
	colors["Color116"] = tcell.Color116
	colors["Color117"] = tcell.Color117
	colors["Color118"] = tcell.Color118
	colors["Color119"] = tcell.Color119
	colors["Color120"] = tcell.Color120
	colors["Color121"] = tcell.Color121
	colors["Color122"] = tcell.Color122
	colors["Color123"] = tcell.Color123
	colors["Color124"] = tcell.Color124
	colors["Color125"] = tcell.Color125
	colors["Color126"] = tcell.Color126
	colors["Color127"] = tcell.Color127
	colors["Color128"] = tcell.Color128
	colors["Color129"] = tcell.Color129
	colors["Color130"] = tcell.Color130
	colors["Color131"] = tcell.Color131
	colors["Color132"] = tcell.Color132
	colors["Color133"] = tcell.Color133
	colors["Color134"] = tcell.Color134
	colors["Color135"] = tcell.Color135
	colors["Color136"] = tcell.Color136
	colors["Color137"] = tcell.Color137
	colors["Color138"] = tcell.Color138
	colors["Color139"] = tcell.Color139
	colors["Color140"] = tcell.Color140
	colors["Color141"] = tcell.Color141
	colors["Color142"] = tcell.Color142
	colors["Color143"] = tcell.Color143
	colors["Color144"] = tcell.Color144
	colors["Color145"] = tcell.Color145
	colors["Color146"] = tcell.Color146
	colors["Color147"] = tcell.Color147
	colors["Color148"] = tcell.Color148
	colors["Color149"] = tcell.Color149
	colors["Color150"] = tcell.Color150
	colors["Color151"] = tcell.Color151
	colors["Color152"] = tcell.Color152
	colors["Color153"] = tcell.Color153
	colors["Color154"] = tcell.Color154
	colors["Color155"] = tcell.Color155
	colors["Color156"] = tcell.Color156
	colors["Color157"] = tcell.Color157
	colors["Color158"] = tcell.Color158
	colors["Color159"] = tcell.Color159
	colors["Color160"] = tcell.Color160
	colors["Color161"] = tcell.Color161
	colors["Color162"] = tcell.Color162
	colors["Color163"] = tcell.Color163
	colors["Color164"] = tcell.Color164
	colors["Color165"] = tcell.Color165
	colors["Color166"] = tcell.Color166
	colors["Color167"] = tcell.Color167
	colors["Color168"] = tcell.Color168
	colors["Color169"] = tcell.Color169
	colors["Color170"] = tcell.Color170
	colors["Color171"] = tcell.Color171
	colors["Color172"] = tcell.Color172
	colors["Color173"] = tcell.Color173
	colors["Color174"] = tcell.Color174
	colors["Color175"] = tcell.Color175
	colors["Color176"] = tcell.Color176
	colors["Color177"] = tcell.Color177
	colors["Color178"] = tcell.Color178
	colors["Color179"] = tcell.Color179
	colors["Color180"] = tcell.Color180
	colors["Color181"] = tcell.Color181
	colors["Color182"] = tcell.Color182
	colors["Color183"] = tcell.Color183
	colors["Color184"] = tcell.Color184
	colors["Color185"] = tcell.Color185
	colors["Color186"] = tcell.Color186
	colors["Color187"] = tcell.Color187
	colors["Color188"] = tcell.Color188
	colors["Color189"] = tcell.Color189
	colors["Color190"] = tcell.Color190
	colors["Color191"] = tcell.Color191
	colors["Color192"] = tcell.Color192
	colors["Color193"] = tcell.Color193
	colors["Color194"] = tcell.Color194
	colors["Color195"] = tcell.Color195
	colors["Color196"] = tcell.Color196
	colors["Color197"] = tcell.Color197
	colors["Color198"] = tcell.Color198
	colors["Color199"] = tcell.Color199
	colors["Color200"] = tcell.Color200
	colors["Color201"] = tcell.Color201
	colors["Color202"] = tcell.Color202
	colors["Color203"] = tcell.Color203
	colors["Color204"] = tcell.Color204
	colors["Color205"] = tcell.Color205
	colors["Color206"] = tcell.Color206
	colors["Color207"] = tcell.Color207
	colors["Color208"] = tcell.Color208
	colors["Color209"] = tcell.Color209
	colors["Color210"] = tcell.Color210
	colors["Color211"] = tcell.Color211
	colors["Color212"] = tcell.Color212
	colors["Color213"] = tcell.Color213
	colors["Color214"] = tcell.Color214
	colors["Color215"] = tcell.Color215
	colors["Color216"] = tcell.Color216
	colors["Color217"] = tcell.Color217
	colors["Color218"] = tcell.Color218
	colors["Color219"] = tcell.Color219
	colors["Color220"] = tcell.Color220
	colors["Color221"] = tcell.Color221
	colors["Color222"] = tcell.Color222
	colors["Color223"] = tcell.Color223
	colors["Color224"] = tcell.Color224
	colors["Color225"] = tcell.Color225
	colors["Color226"] = tcell.Color226
	colors["Color227"] = tcell.Color227
	colors["Color228"] = tcell.Color228
	colors["Color229"] = tcell.Color229
	colors["Color230"] = tcell.Color230
	colors["Color231"] = tcell.Color231
	colors["Color232"] = tcell.Color232
	colors["Color233"] = tcell.Color233
	colors["Color234"] = tcell.Color234
	colors["Color235"] = tcell.Color235
	colors["Color236"] = tcell.Color236
	colors["Color237"] = tcell.Color237
	colors["Color238"] = tcell.Color238
	colors["Color239"] = tcell.Color239
	colors["Color240"] = tcell.Color240
	colors["Color241"] = tcell.Color241
	colors["Color242"] = tcell.Color242
	colors["Color243"] = tcell.Color243
	colors["Color244"] = tcell.Color244
	colors["Color245"] = tcell.Color245
	colors["Color246"] = tcell.Color246
	colors["Color247"] = tcell.Color247
	colors["Color248"] = tcell.Color248
	colors["Color249"] = tcell.Color249
	colors["Color250"] = tcell.Color250
	colors["Color251"] = tcell.Color251
	colors["Color252"] = tcell.Color252
	colors["Color253"] = tcell.Color253
	colors["Color254"] = tcell.Color254
	colors["Color255"] = tcell.Color255
	colors["ColorAliceBlue"] = tcell.ColorAliceBlue
	colors["ColorAntiqueWhite"] = tcell.ColorAntiqueWhite
	colors["ColorAquaMarine"] = tcell.ColorAquaMarine
	colors["ColorAzure"] = tcell.ColorAzure
	colors["ColorBeige"] = tcell.ColorBeige
	colors["ColorBisque"] = tcell.ColorBisque
	colors["ColorBlanchedAlmond"] = tcell.ColorBlanchedAlmond
	colors["ColorBlueViolet"] = tcell.ColorBlueViolet
	colors["ColorBrown"] = tcell.ColorBrown
	colors["ColorBurlyWood"] = tcell.ColorBurlyWood
	colors["ColorCadetBlue"] = tcell.ColorCadetBlue
	colors["ColorChartreuse"] = tcell.ColorChartreuse
	colors["ColorChocolate"] = tcell.ColorChocolate
	colors["ColorCoral"] = tcell.ColorCoral
	colors["ColorCornflowerBlue"] = tcell.ColorCornflowerBlue
	colors["ColorCornsilk"] = tcell.ColorCornsilk
	colors["ColorCrimson"] = tcell.ColorCrimson
	colors["ColorDarkBlue"] = tcell.ColorDarkBlue
	colors["ColorDarkCyan"] = tcell.ColorDarkCyan
	colors["ColorDarkGoldenrod"] = tcell.ColorDarkGoldenrod
	colors["ColorDarkGray"] = tcell.ColorDarkGray
	colors["ColorDarkGreen"] = tcell.ColorDarkGreen
	colors["ColorDarkKhaki"] = tcell.ColorDarkKhaki
	colors["ColorDarkMagenta"] = tcell.ColorDarkMagenta
	colors["ColorDarkOliveGreen"] = tcell.ColorDarkOliveGreen
	colors["ColorDarkOrange"] = tcell.ColorDarkOrange
	colors["ColorDarkOrchid"] = tcell.ColorDarkOrchid
	colors["ColorDarkRed"] = tcell.ColorDarkRed
	colors["ColorDarkSalmon"] = tcell.ColorDarkSalmon
	colors["ColorDarkSeaGreen"] = tcell.ColorDarkSeaGreen
	colors["ColorDarkSlateBlue"] = tcell.ColorDarkSlateBlue
	colors["ColorDarkSlateGray"] = tcell.ColorDarkSlateGray
	colors["ColorDarkTurquoise"] = tcell.ColorDarkTurquoise
	colors["ColorDarkViolet"] = tcell.ColorDarkViolet
	colors["ColorDeepPink"] = tcell.ColorDeepPink
	colors["ColorDeepSkyBlue"] = tcell.ColorDeepSkyBlue
	colors["ColorDimGray"] = tcell.ColorDimGray
	colors["ColorDodgerBlue"] = tcell.ColorDodgerBlue
	colors["ColorFireBrick"] = tcell.ColorFireBrick
	colors["ColorFloralWhite"] = tcell.ColorFloralWhite
	colors["ColorForestGreen"] = tcell.ColorForestGreen
	colors["ColorGainsboro"] = tcell.ColorGainsboro
	colors["ColorGhostWhite"] = tcell.ColorGhostWhite
	colors["ColorGold"] = tcell.ColorGold
	colors["ColorGoldenrod"] = tcell.ColorGoldenrod
	colors["ColorGreenYellow"] = tcell.ColorGreenYellow
	colors["ColorHoneydew"] = tcell.ColorHoneydew
	colors["ColorHotPink"] = tcell.ColorHotPink
	colors["ColorIndianRed"] = tcell.ColorIndianRed
	colors["ColorIndigo"] = tcell.ColorIndigo
	colors["ColorIvory"] = tcell.ColorIvory
	colors["ColorKhaki"] = tcell.ColorKhaki
	colors["ColorLavender"] = tcell.ColorLavender
	colors["ColorLavenderBlush"] = tcell.ColorLavenderBlush
	colors["ColorLawnGreen"] = tcell.ColorLawnGreen
	colors["ColorLemonChiffon"] = tcell.ColorLemonChiffon
	colors["ColorLightBlue"] = tcell.ColorLightBlue
	colors["ColorLightCoral"] = tcell.ColorLightCoral
	colors["ColorLightCyan"] = tcell.ColorLightCyan
	colors["ColorLightGoldenrodYellow"] = tcell.ColorLightGoldenrodYellow
	colors["ColorLightGray"] = tcell.ColorLightGray
	colors["ColorLightGreen"] = tcell.ColorLightGreen
	colors["ColorLightPink"] = tcell.ColorLightPink
	colors["ColorLightSalmon"] = tcell.ColorLightSalmon
	colors["ColorLightSeaGreen"] = tcell.ColorLightSeaGreen
	colors["ColorLightSkyBlue"] = tcell.ColorLightSkyBlue
	colors["ColorLightSlateGray"] = tcell.ColorLightSlateGray
	colors["ColorLightSteelBlue"] = tcell.ColorLightSteelBlue
	colors["ColorLightYellow"] = tcell.ColorLightYellow
	colors["ColorLimeGreen"] = tcell.ColorLimeGreen
	colors["ColorLinen"] = tcell.ColorLinen
	colors["ColorMediumAquamarine"] = tcell.ColorMediumAquamarine
	colors["ColorMediumBlue"] = tcell.ColorMediumBlue
	colors["ColorMediumOrchid"] = tcell.ColorMediumOrchid
	colors["ColorMediumPurple"] = tcell.ColorMediumPurple
	colors["ColorMediumSeaGreen"] = tcell.ColorMediumSeaGreen
	colors["ColorMediumSlateBlue"] = tcell.ColorMediumSlateBlue
	colors["ColorMediumSpringGreen"] = tcell.ColorMediumSpringGreen
	colors["ColorMediumTurquoise"] = tcell.ColorMediumTurquoise
	colors["ColorMediumVioletRed"] = tcell.ColorMediumVioletRed
	colors["ColorMidnightBlue"] = tcell.ColorMidnightBlue
	colors["ColorMintCream"] = tcell.ColorMintCream
	colors["ColorMistyRose"] = tcell.ColorMistyRose
	colors["ColorMoccasin"] = tcell.ColorMoccasin
	colors["ColorNavajoWhite"] = tcell.ColorNavajoWhite
	colors["ColorOldLace"] = tcell.ColorOldLace
	colors["ColorOliveDrab"] = tcell.ColorOliveDrab
	colors["ColorOrange"] = tcell.ColorOrange
	colors["ColorOrangeRed"] = tcell.ColorOrangeRed
	colors["ColorOrchid"] = tcell.ColorOrchid
	colors["ColorPaleGoldenrod"] = tcell.ColorPaleGoldenrod
	colors["ColorPaleGreen"] = tcell.ColorPaleGreen
	colors["ColorPaleTurquoise"] = tcell.ColorPaleTurquoise
	colors["ColorPaleVioletRed"] = tcell.ColorPaleVioletRed
	colors["ColorPapayaWhip"] = tcell.ColorPapayaWhip
	colors["ColorPeachPuff"] = tcell.ColorPeachPuff
	colors["ColorPeru"] = tcell.ColorPeru
	colors["ColorPink"] = tcell.ColorPink
	colors["ColorPlum"] = tcell.ColorPlum
	colors["ColorPowderBlue"] = tcell.ColorPowderBlue
	colors["ColorRebeccaPurple"] = tcell.ColorRebeccaPurple
	colors["ColorRosyBrown"] = tcell.ColorRosyBrown
	colors["ColorRoyalBlue"] = tcell.ColorRoyalBlue
	colors["ColorSaddleBrown"] = tcell.ColorSaddleBrown
	colors["ColorSalmon"] = tcell.ColorSalmon
	colors["ColorSandyBrown"] = tcell.ColorSandyBrown
	colors["ColorSeaGreen"] = tcell.ColorSeaGreen
	colors["ColorSeashell"] = tcell.ColorSeashell
	colors["ColorSienna"] = tcell.ColorSienna
	colors["ColorSkyblue"] = tcell.ColorSkyblue
	colors["ColorSlateBlue"] = tcell.ColorSlateBlue
	colors["ColorSlateGray"] = tcell.ColorSlateGray
	colors["ColorSnow"] = tcell.ColorSnow
	colors["ColorSpringGreen"] = tcell.ColorSpringGreen
	colors["ColorSteelBlue"] = tcell.ColorSteelBlue
	colors["ColorTan"] = tcell.ColorTan
	colors["ColorThistle"] = tcell.ColorThistle
	colors["ColorTomato"] = tcell.ColorTomato
	colors["ColorTurquoise"] = tcell.ColorTurquoise
	colors["ColorViolet"] = tcell.ColorViolet
	colors["ColorWheat"] = tcell.ColorWheat
	colors["ColorWhiteSmoke"] = tcell.ColorWhiteSmoke
	colors["ColorYellowGreen"] = tcell.ColorYellowGreen

	var sleep_time time.Duration = time.Second * 60

	keys := make([]string, 0, len(colors))
	for k := range colors {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//time.Sleep(sleep_time)
	//scn.Fini()
	y := 1
	offset := 0
	for _, k := range keys {
		v := colors[k]
		//fmt.Println("got k, v:", k, v)
		//style := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorWhite)
		style := tcell.StyleDefault.Foreground(v)
		for idx, chr := range k {
			//fmt.Printf("%d %c ", idx, chr)
			scn.SetContent(idx+offset, y, rune(chr), []rune(""), style)
		}
		scn.Show()
		y++
		if y >= lines {
			//time.Sleep(sleep_time)
			//scn.Clear()
			//scn.Sync()
			y = 0
			offset += 22
		}
	}
	time.Sleep(sleep_time)

	scn.Fini()
	os.Exit(0)
}

/*
bright/white colors

121,122,123,231,230,86,87,250-255,229,228,153,152,151,157-159,192-195
*/
