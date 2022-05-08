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

const (
	C_RESET      = "\x1b[0m"
	C_BLACK      = "\x1b[30m"
	C_RED        = "\x1b[31m"
	C_GREEN      = "\x1b[32m"
	C_YELLOW     = "\x1b[33m"
	C_BLUE       = "\x1b[34m"
	C_PURPLE     = "\x1b[35m"
	C_CYAN       = "\x1b[36m"
	C_GREY       = "\x1b[2m"
	C_WHITE      = "\x1b[37m"
	C_BLACKISH   = "\x1b[90m"
	C_REDISH     = "\x1b[91m"
	C_GREENISH   = "\x1b[92m"
	C_YELLOWISH  = "\x1b[93m"
	C_BLUISH     = "\x1b[94m"
	C_PURPLISH   = "\x1b[95m"
	C_CYANISH    = "\x1b[96m"
	C_WHITER     = "\x1b[97m"
	S_BOLD       = "\x1b[1m"
	S_UNDERLINED = "\x1b[4m"
	S_CROSSED    = "\x1b[9m"
)
