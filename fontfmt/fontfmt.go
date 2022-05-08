/*
	librlcli a cli library for rlinux
    Copyright (C) 2022  Mikhail Batsin

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package fontfmt

// fonts
type fontfmt struct {
	Bg     string
	Styles []string
}

func (ff *fontfmt) SetBg(bg string) {
	ff.Bg = bg
}

func (ff *fontfmt) SetStyles(styles []string) {
	ff.Styles = styles
}

func (ff fontfmt) Println(color string, v ...interface{}) {
	var styles string

	if len(ff.Styles) > 0 {
		styles = ""

		for _, style := range ff.Styles {
			styles += style
		}
	}

	println(ff.Bg, styles, color, v, C_RESET)
}

func (ff fontfmt) PrintRedln(v ...interface{}) {
	ff.Println(C_RED, v...)
}

func (ff fontfmt) PrintGreyln(v ...interface{}) {
	ff.Println(C_GREY, v...)
}

func (ff fontfmt) PrintBlackln(v ...interface{}) {
	ff.Println(C_BLACK, v...)
}

func (ff fontfmt) PrintGreenln(v ...interface{}) {
	ff.Println(C_GREEN, v...)
}

func (ff fontfmt) PrintYellowln(v ...interface{}) {
	ff.Println(C_YELLOW, v...)
}

func (ff fontfmt) PrintBlueln(v ...interface{}) {
	ff.Println(C_BLUE, v...)
}

func (ff fontfmt) PrintPurpleln(v ...interface{}) {
	ff.Println(C_PURPLE, v...)
}

func (ff fontfmt) PrintCyanln(v ...interface{}) {
	ff.Println(C_CYAN, v...)
}

func (ff fontfmt) PrintReddishln(v ...interface{}) {
	ff.Println(C_REDISH, v...)
}

func (ff fontfmt) PrintBlackishln(v ...interface{}) {
	ff.Println(C_BLACKISH, v...)
}

func (ff fontfmt) PrintGreenishln(v ...interface{}) {
	ff.Println(C_GREENISH, v...)
}

func (ff fontfmt) PrintYellowishln(v ...interface{}) {
	ff.Println(C_YELLOWISH, v...)
}

func (ff fontfmt) PrintBluishln(v ...interface{}) {
	ff.Println(C_BLUISH, v...)
}

func (ff fontfmt) PrintPurplishln(v ...interface{}) {
	ff.Println(C_PURPLISH, v...)
}

func (ff fontfmt) PrintCyanishln(v ...interface{}) {
	ff.Println(C_CYANISH, v...)
}

func (ff fontfmt) PrintWhiterln(v ...interface{}) {
	ff.Println(C_WHITER, v...)
}

var FontFmt = fontfmt{}
