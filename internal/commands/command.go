/*
 * Copyright (C) 2024 carddamom
 *
 * This file is part of riconto.
 *
 * riconto is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * riconto is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with riconto.  If not, see <https://www.gnu.org/licenses/>.
 */

package commands

import (
	"github.com/tucnak/climax"
)

type Command interface {
	Name() string
	Brief() string
	Usage() string
	Help() string
	Group() string
	Flags() []climax.Flag
	Examples() []climax.Example
	Run(context climax.Context) int
	Command() climax.Command
}

func FromCommand(c Command) climax.Command {
	result := climax.Command{
		Name:   c.Name(),
		Brief:  c.Brief(),
		Usage:  c.Usage(),
		Help:   c.Help(),
		Group:  c.Group(),
		Handle: c.Run,
	}
	for _, flag := range c.Flags() {
		result.AddFlag(flag)
	}
	for _, example := range c.Examples() {
		result.AddExample(example)
	}
	return result
}
