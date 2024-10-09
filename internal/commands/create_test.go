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
	"testing"

	"github.com/spf13/afero"
	"github.com/tucnak/climax"

	"github.com/primalskill/golog"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateCommand(t *testing.T) {
	Convey("#CreateCommand", t, func() {
		var createCommand *CreateCommand

		Convey("It should be able to create a new command", func() {
			createCommand = NewCreateCommand(afero.NewMemMapFs(), golog.NewDiscard())
			So(createCommand, ShouldNotBeNil)
			So(createCommand.Name(), ShouldEqual, "init")
			So(createCommand.Brief(), ShouldEqual, "initializes a new project")
		})

		Convey("Given some parameters", func() {
			memFs := afero.NewMemMapFs()
			createCommand = NewCreateCommand(memFs, golog.NewDiscard())

			Convey("With a name", func() {
				name := "sample"

				Convey("It should create a new project in toml format", func() {
					So(memFs.RemoveAll("/"), ShouldBeNil) //Clear the filesystem
					context := climax.Context{
						Args:        []string{"--name", name, "--format", "toml"},
						NonVariable: make(map[string]bool),
						Variable:    map[string]string{"name": name, "format": "toml"},
					}
					So(createCommand.Run(context), ShouldEqual, 0)
					info, err := memFs.Stat("riconto.toml")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeFalse)
					So(info.Mode().IsRegular(), ShouldBeTrue)
					info, err = memFs.Stat("src")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeTrue)
					info, err = memFs.Stat("resources")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeTrue)
					info, err = memFs.Stat("src/main.md")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeFalse)
					So(info.Mode().IsRegular(), ShouldBeTrue)
				})

				Convey("It should create a new project in yaml format", func() {
					So(memFs.RemoveAll("/"), ShouldBeNil) //Clear the filesystem
					context := climax.Context{
						Args:        []string{"--name", name, "--format", "yaml"},
						NonVariable: make(map[string]bool),
						Variable:    map[string]string{"name": name, "format": "yaml"},
					}
					So(createCommand.Run(context), ShouldEqual, 0)
					info, err := memFs.Stat("riconto.yaml")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeFalse)
					So(info.Mode().IsRegular(), ShouldBeTrue)
					info, err = memFs.Stat("src")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeTrue)
					info, err = memFs.Stat("resources")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeTrue)
					info, err = memFs.Stat("src/main.md")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeFalse)
					So(info.Mode().IsRegular(), ShouldBeTrue)
				})

				Convey("It should create a new project in json format", func() {
					So(memFs.RemoveAll("/"), ShouldBeNil) //Clear the filesystem
					context := climax.Context{
						Args:        []string{"--name", name, "--format", "json"},
						NonVariable: make(map[string]bool),
						Variable:    map[string]string{"name": name, "format": "json"},
					}
					So(createCommand.Run(context), ShouldEqual, 0)
					info, err := memFs.Stat("riconto.json")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeFalse)
					So(info.Mode().IsRegular(), ShouldBeTrue)
					info, err = memFs.Stat("src")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeTrue)
					info, err = memFs.Stat("resources")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeTrue)
					info, err = memFs.Stat("src/main.md")
					So(err, ShouldBeNil)
					So(info.IsDir(), ShouldBeFalse)
					So(info.Mode().IsRegular(), ShouldBeTrue)
				})

				Convey("When no format is specified", func() {

					Convey("It should create a new project in toml format", func() {
						So(memFs.RemoveAll("/"), ShouldBeNil) //Clear the filesystem
						context := climax.Context{
							Args:        []string{"--name", name},
							NonVariable: make(map[string]bool),
							Variable:    map[string]string{"name": name},
						}
						So(createCommand.Run(context), ShouldEqual, 0)
						info, err := memFs.Stat("riconto.toml")
						So(err, ShouldBeNil)
						So(info.IsDir(), ShouldBeFalse)
						So(info.Mode().IsRegular(), ShouldBeTrue)
						info, err = memFs.Stat("src")
						So(err, ShouldBeNil)
						So(info.IsDir(), ShouldBeTrue)
						info, err = memFs.Stat("resources")
						So(err, ShouldBeNil)
						So(info.IsDir(), ShouldBeTrue)
						info, err = memFs.Stat("src/main.md")
						So(err, ShouldBeNil)
						So(info.IsDir(), ShouldBeFalse)
						So(info.Mode().IsRegular(), ShouldBeTrue)
					})

					Convey("When a version is provided", func() {
						version := "1.2.3"

						Convey("It should create a new project in toml format with the version", func() {
							So(memFs.RemoveAll("/"), ShouldBeNil) //Clear the filesystem
							context := climax.Context{
								Args:        []string{"--name", name, "--version", version},
								NonVariable: make(map[string]bool),
								Variable:    map[string]string{"name": name, "version": version},
							}
							So(createCommand.Run(context), ShouldEqual, 0)
							info, err := memFs.Stat("riconto.toml")
							So(err, ShouldBeNil)
							So(info.IsDir(), ShouldBeFalse)
							So(info.Mode().IsRegular(), ShouldBeTrue)
							info, err = memFs.Stat("src")
							So(err, ShouldBeNil)
							So(info.IsDir(), ShouldBeTrue)
							info, err = memFs.Stat("resources")
							So(err, ShouldBeNil)
							So(info.IsDir(), ShouldBeTrue)
							info, err = memFs.Stat("src/main.md")
							So(err, ShouldBeNil)
							So(info.IsDir(), ShouldBeFalse)
							So(info.Mode().IsRegular(), ShouldBeTrue)
						})
					})
				})
			})

			Convey("Without a name", func() {

				Convey("It should not create a new project", func() {
					So(memFs.RemoveAll("/"), ShouldBeNil) //Clear the filesystem
					context := climax.Context{
						Args:        []string{},
						NonVariable: make(map[string]bool),
						Variable:    make(map[string]string),
					}
					So(createCommand.Run(context), ShouldEqual, 1)
				})
			})
		})
	})
}
