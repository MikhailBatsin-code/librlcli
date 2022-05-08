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

import "fmt"

// action is a structure with OnCmd field that is used for calling handler. Handler takes arguments for cmd
type Action struct {
	OnCmd       string
	Description string
	Handler     func([]string) error
}

var actionsList []Action = []Action{}

// built-in action
func helpAction() {
	for _, action := range actionsList {
		fmt.Println(action.OnCmd + "\t\t" + action.Description)
	}
}

func AddAction(onCmd string, desc string, handler func([]string) error) {
	actionsList = append(actionsList, Action{
		OnCmd:       onCmd,
		Handler:     handler,
		Description: desc,
	})
}
