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

package model

import (
	"io"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/afero"
)

const (
	jsonContent = `{
  "name": "sample",
  "version": "0.0.1",
  "description": "Some description",
  "authors": [
    {
      "name": "carddamom",
      "url": "https://github.com/carddamom",
      "email": "carddamom@tutanota.com"
    }
  ],
  "files": [
    {
      "name": "Book A",
      "output": "./dist/bookA",
      "path": "./src/bookA/main.md"
    }
  ],
  "license": ["GPL-3.0-or-later"]
}`
	jsonContentInvalid = `{
  "name": "sample",
  "version": "0.0.1",
  "description": "Some description",
  "authors":
    {
      "name": "carddamom",
      "url": "https://github.com/carddamom",
      "email": "carddamom@tutanota.com"
    }
  ],
  "license": ["GPL-3.0-or-later"]
}`
	tomlContent = `
name = "sample"
version = "0.0.1"
description = "Some description"
license = ["GPL-3.0-or-later"]

[[files]]
name = "Book A"
output = "./dist/bookA"
path = "./src/bookA/main.md"

[[authors]]
name = "carddamom"
url = "https://github.com/carddamom"
email = "carddamom@tutanota.com"
`
	tomlContentInvalid = `
name = "sample"
version = "0.0.1"
description = "Some description"
license = ["GPL-3.0-or-later"

[[authors]
name = "carddamom"
url = "https://github.com/carddamom"
email = "carddamom@tutanota.com"
`
	yamlContent = `
name: sample
version: 0.0.1
description: Some description
authors:
  - name: carddamom
    url: https://github.com/carddamom
    email: carddamom@tutanota.com
files:
  - name: Book A
    path: "./src/bookA/main.md"
    output: "./dist/bookA"
license:
  - GPL-3.0-or-later
`
	yamlContentInvalid = `
name: sample
version: 0.0.1
description: Some description
authors:
  name: carddamom
  url: https://github.com/carddamom
  email: carddamom@tutanota.com
license:
  - GPL-3.0-or-later
`
)

func TestConfig(t *testing.T) {
	config := NewConfig("test", "1.2.0", "This is a sample package")
	config.AddLicense("GPL3.0-or-later")
	config.AddAuthor(NewAuthor("carddamom"))
	config.AddFile(NewFile("test", "./dist/test", "./src/test.md"))
	fs := afero.NewMemMapFs()

	Convey("#Config", t, func() {
		configDup := NewConfigFrom(config)
		So(configDup, ShouldNotBeNil)

		Convey("It should be possible to duplicate a Config", func() {
			So(configDup, ShouldNotBeNil)
			So(configDup, ShouldEqual, config)
		})

		Convey("It should be possible to add a new License", func() {
			res := configDup.AddLicense("MIT")
			So(res, ShouldBeTrue)
			So(configDup.ContainsLicense("MIT"), ShouldBeTrue)
		})

		Convey("It should be possible to remove a License", func() {
			configDup.AddLicense("MIT")
			res := configDup.RemoveLicense("MIT")
			So(res, ShouldBeTrue)
			So(configDup.ContainsLicense("MIT"), ShouldBeFalse)
		})

		Convey("It should not be possible to remove a nonexisting License", func() {
			res := configDup.RemoveLicense("AL2")
			So(res, ShouldBeFalse)
			So(configDup.ContainsLicense("AL2"), ShouldBeFalse)
		})

		Convey("It should be possible to add a new File", func() {
			res := configDup.AddFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md"))
			So(res, ShouldBeTrue)
			So(configDup.ContainsFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md")), ShouldBeTrue)
		})

		Convey("It should be possible to remove a File", func() {
			configDup.AddFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md"))
			res := configDup.RemoveFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md"))
			So(res, ShouldBeTrue)
			So(configDup.ContainsFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md")), ShouldBeFalse)
		})

		Convey("It should not be possible to remove a nonexisting File", func() {
			res := configDup.RemoveFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md"))
			So(res, ShouldBeFalse)
			So(configDup.ContainsFile(NewFile("Book A", "./dist/bookA", "./src/bookA/main.md")), ShouldBeFalse)
		})

		Convey("It should not be possible to add a duplicate License", func() {
			res := configDup.AddLicense("GPL3.0-or-later")
			So(res, ShouldBeFalse)
		})

		Convey("It should be possible to add a new Author", func() {
			res := configDup.AddAuthor(NewAuthor("cinnamon"))
			So(res, ShouldBeTrue)
			So(configDup.ContainsAuthor(NewAuthor("cinnamon")), ShouldBeTrue)
		})

		Convey("It should be possible to remove an Author", func() {
			configDup.AddAuthor(NewAuthor("cinnamon"))
			res := configDup.RemoveAuthor(NewAuthor("cinnamon"))
			So(res, ShouldBeTrue)
		})

		Convey("It should not be possible to remove a nonexisting Author", func() {
			res := configDup.RemoveAuthor(NewAuthor("johndoe"))
			So(res, ShouldBeFalse)
		})

		Convey("It should not be possible to add a duplicate Author", func() {
			res := configDup.AddAuthor(NewAuthor("carddamom"))
			So(res, ShouldBeFalse)
		})

	})

	Convey("#JsonConfigFile", t, func() {

		Convey("Given an existing valid json project configuration file", func() {
			jsonFile, err := fs.OpenFile("simpleFile.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			So(jsonFile, ShouldNotBeNil)
			defer jsonFile.Close()
			_, err = jsonFile.Write([]byte(jsonContent))
			So(err, ShouldBeNil)
			jsonFile.Close()

			Convey("It should be able to read from it", func() {
				jsonFile, err = fs.Open("simpleFile.json")
				So(err, ShouldBeNil)
				So(jsonFile, ShouldNotBeNil)
				defer jsonFile.Close()
				config, err := ConfigFromFile(jsonFile, FormatJson)
				So(err, ShouldBeNil)
				So(config, ShouldNotBeNil)
				So(config.Name, ShouldEqual, "sample")
				So(config.Version, ShouldEqual, "0.0.1")
				So(config.ContainsLicense("GPL-3.0-or-later"), ShouldBeTrue)
			})

			Convey("It should be able to write to it", func() {
				jsonFile, err = fs.OpenFile("simpleFile.json", os.O_WRONLY|os.O_TRUNC, 0644)
				So(err, ShouldBeNil)
				So(jsonFile, ShouldNotBeNil)
				defer jsonFile.Close()
				err = config.SaveTo(jsonFile, FormatJson)
				So(err, ShouldBeNil)
				jsonFile.Close()

				jsonFile, err = fs.Open("simpleFile.json")
				So(err, ShouldBeNil)
				So(jsonFile, ShouldNotBeNil)
				defer jsonFile.Close()
				b, err := io.ReadAll(jsonFile)
				So(err, ShouldBeNil)
				So(b, ShouldNotBeNil)
				So(b, ShouldNotBeEmpty)
				str := string(b)
				So(str, ShouldNotBeEmpty)
				So(str, ShouldContainSubstring, `"name":"test"`)
				So(str, ShouldContainSubstring, `"version":"1.2.0"`)
			})

		})

		Convey("Given an existing invalid json project configuration file", func() {
			jsonFile, err := fs.OpenFile("invalid.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			So(jsonFile, ShouldNotBeNil)
			defer jsonFile.Close()
			_, err = jsonFile.Write([]byte(jsonContentInvalid))
			So(err, ShouldBeNil)
			jsonFile.Close()

			Convey("It should not be able to read from it", func() {
				jsonFile, err = fs.Open("invalid.json")
				So(err, ShouldBeNil)
				So(jsonFile, ShouldNotBeNil)
				defer jsonFile.Close()
				config, err := ConfigFromFile(jsonFile, FormatJson)
				So(err, ShouldNotBeNil)
				So(config, ShouldBeNil)
			})

		})

	})

	Convey("#YamlConfigFile", t, func() {

		Convey("Given an existing valid yaml project configuration file", func() {
			yamlFile, err := fs.OpenFile("valid.yaml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			So(yamlFile, ShouldNotBeNil)
			defer yamlFile.Close()
			_, err = yamlFile.Write([]byte(yamlContent))
			So(err, ShouldBeNil)
			yamlFile.Close()

			Convey("It should be able to read from it", func() {
				yamlFile, err = fs.Open("valid.yaml")
				So(err, ShouldBeNil)
				So(yamlFile, ShouldNotBeNil)
				defer yamlFile.Close()
				config, err := ConfigFromFile(yamlFile, FormatYaml)
				So(err, ShouldBeNil)
				So(config, ShouldNotBeNil)
				So(config.Name, ShouldEqual, "sample")
				So(config.Version, ShouldEqual, "0.0.1")
				So(config.ContainsLicense("GPL-3.0-or-later"), ShouldBeTrue)
			})

			Convey("It should be able to write to it", func() {
				yamlFile, err = fs.OpenFile("valid.yaml", os.O_WRONLY|os.O_TRUNC, 0644)
				So(err, ShouldBeNil)
				So(yamlFile, ShouldNotBeNil)
				defer yamlFile.Close()
				err = config.SaveTo(yamlFile, FormatYaml)
				So(err, ShouldBeNil)
				yamlFile.Close()

				yamlFile, err = fs.Open("valid.yaml")
				So(err, ShouldBeNil)
				So(yamlFile, ShouldNotBeNil)
				defer yamlFile.Close()
				b, err := io.ReadAll(yamlFile)
				So(err, ShouldBeNil)
				So(b, ShouldNotBeNil)
				So(b, ShouldNotBeEmpty)
				str := string(b)
				So(str, ShouldNotBeEmpty)
				So(str, ShouldContainSubstring, `name: test`)
				So(str, ShouldContainSubstring, `version: 1.2.0`)
			})

		})

		Convey("Given an existing invalid yaml project configuration file", func() {
			yamlFile, err := fs.OpenFile("invalid.yaml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			So(yamlFile, ShouldNotBeNil)
			defer yamlFile.Close()
			_, err = yamlFile.Write([]byte(yamlContentInvalid))
			So(err, ShouldBeNil)
			yamlFile.Close()

			Convey("It should not be able to read from it", func() {
				yamlFile, err = fs.Open("invalid.yaml")
				So(err, ShouldBeNil)
				So(yamlFile, ShouldNotBeNil)
				defer yamlFile.Close()
				config, err := ConfigFromFile(yamlFile, FormatYaml)
				So(err, ShouldNotBeNil)
				So(config, ShouldBeNil)
			})

		})

	})

	Convey("#TomlConfigFile", t, func() {

		Convey("Given a existing valid toml project configuration file", func() {
			tomlFile, err := fs.OpenFile("valid.toml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			So(tomlFile, ShouldNotBeNil)
			defer tomlFile.Close()
			_, err = tomlFile.Write([]byte(tomlContent))
			So(err, ShouldBeNil)
			tomlFile.Close()

			Convey("It should be able to read from it", func() {
				tomlFile, err = fs.Open("valid.toml")
				So(err, ShouldBeNil)
				So(tomlFile, ShouldNotBeNil)
				defer tomlFile.Close()
				config, err := ConfigFromFile(tomlFile, FormatToml)
				So(err, ShouldBeNil)
				So(config, ShouldNotBeNil)
				So(config.Name, ShouldEqual, "sample")
				So(config.Version, ShouldEqual, "0.0.1")
				So(config.ContainsLicense("GPL-3.0-or-later"), ShouldBeTrue)
			})

			Convey("It should be able to write to it", func() {
				tomlFile, err = fs.OpenFile("valid.toml", os.O_WRONLY|os.O_TRUNC, 0644)
				So(err, ShouldBeNil)
				So(tomlFile, ShouldNotBeNil)
				defer tomlFile.Close()
				err = config.SaveTo(tomlFile, FormatToml)
				So(err, ShouldBeNil)
				tomlFile.Close()

				tomlFile, err = fs.Open("valid.toml")
				So(err, ShouldBeNil)
				So(tomlFile, ShouldNotBeNil)
				defer tomlFile.Close()
				b, err := io.ReadAll(tomlFile)
				So(err, ShouldBeNil)
				So(b, ShouldNotBeNil)
				So(b, ShouldNotBeEmpty)
				str := string(b)
				So(str, ShouldNotBeEmpty)
				So(str, ShouldContainSubstring, `name = 'test'`)
				So(str, ShouldContainSubstring, `version = '1.2.0'`)
			})

		})

		Convey("Given a existing invalid toml project configuration file", func() {
			tomlFile, err := fs.OpenFile("invalid.toml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			So(tomlFile, ShouldNotBeNil)
			defer tomlFile.Close()
			_, err = tomlFile.Write([]byte(tomlContentInvalid))
			So(err, ShouldBeNil)
			tomlFile.Close()

			Convey("It should not be able to read from it", func() {
				tomlFile, err = fs.Open("invalid.toml")
				So(err, ShouldBeNil)
				So(tomlFile, ShouldNotBeNil)
				defer tomlFile.Close()
				config, err := ConfigFromFile(tomlFile, FormatToml)
				So(err, ShouldNotBeNil)
				So(config, ShouldBeNil)
			})

		})

	})

}
