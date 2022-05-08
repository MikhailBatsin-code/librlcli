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

package interactive

import (
	"fmt"
	"strings"

	"github.com/MikhailBatsin-code/librlcli/fontfmt"
)

func InteractiveCli() {
	for {
		prompt()

		if len(userPrompt) == 0 {
			continue
		}

		parsed := strings.Split(strings.TrimSpace(userPrompt), " ")

		found := false

		for _, action := range actionsList {
			if action.OnCmd == parsed[0] {
				err := action.Handler(parsed)

				found = true

				if err != nil {
					fmt.Println("Command returned error:", err)
				}

				break
			} else if parsed[0] == "help" {
				helpAction()
				found = true
				break
			}
		}

		if !found {
			fontfmt.FontFmt.PrintRedln("Unknown command")
		}
	}
}
