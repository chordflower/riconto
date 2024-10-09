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
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/chordflower/riconto/pkg/utils"

	"github.com/buger/goterm"
	"github.com/chordflower/riconto/internal/model"
	"github.com/muesli/reflow/wordwrap"
	"github.com/spf13/afero"
	"github.com/tucnak/climax"
)

type CreateCommand struct {
	name     string
	brief    string
	usage    string
	help     string
	group    string
	flags    []climax.Flag
	examples []climax.Example
	logger   *slog.Logger
	fs       afero.Fs
}

func NewCreateCommand(fs afero.Fs, logger *slog.Logger) *CreateCommand {
	terminalWidth := goterm.Width()
	helpStr := "" +
		"This command will initialize a riconto project, creating the configuration file " +
		"and all of the needed directories, in the directory where the executable is called.\n" +
		"It can create the configuration file in three different formats:\n\n" +
		"- json\n" +
		"- yaml\n" +
		"- toml\n" +
		"\n" +
		"These formats are selectable with the option --format or -f.\n" +
		"It is also possible to specify a version with the option --version or -v, note that " +
		"if a version is not specified the string 0.0.1 is used, while not required it is " +
		"advised that the version obeys to some standard like semantic versioning or calendar " +
		"versioning.\n" +
		"Also the parameter --name or -n is required and specifies the name of the project in " +
		"the configuration file and can be any string, noting that strings with spaces will have " +
		"to be quoted due to limitations of the command line.\n\n" +
		"As for the directories, if they are present, they will not be created or overwritten."
	flags := make([]climax.Flag, 0, 7)
	flags = append(flags, climax.Flag{
		Name:     "name",
		Short:    "n",
		Usage:    "--name NAME",
		Help:     "The name of the project in the configuration file (required)",
		Variable: true,
	})
	flags = append(flags, climax.Flag{
		Name:     "description",
		Short:    "d",
		Usage:    "--description DESCRIPTION",
		Help:     "The description of the project (default empty)",
		Variable: true,
	})
	flags = append(flags, climax.Flag{
		Name:     "license",
		Short:    "l",
		Usage:    "--license LICENSE",
		Help:     "The license of the project (default empty)",
		Variable: true,
	})
	flags = append(flags, climax.Flag{
		Name:     "version",
		Short:    "v",
		Usage:    "--version VERSION",
		Help:     "The project version in the configuration file (default 0.0.1)",
		Variable: true,
	})
	flags = append(flags, climax.Flag{
		Name:     "format",
		Short:    "f",
		Usage:    "--format json|yaml|toml",
		Help:     "The configuration file format (default toml)",
		Variable: true,
	})
	flags = append(flags, climax.Flag{
		Name:     "strict",
		Short:    "s",
		Usage:    "--strict",
		Help:     "Ends with an error code if a configuration file already exists",
		Variable: false,
	})
	examples := make([]climax.Example, 0, 5)
	examples = append(examples, climax.Example{
		Usecase: "--name example",
		Description: "Creates a project named example in the current directory, " +
			"with a riconto.toml file and version 0.0.1",
	})
	examples = append(examples, climax.Example{
		Usecase: "--name example --version 1.2.0",
		Description: "Creates a project named example in the current directory, " +
			"with a riconto.toml file and version 1.2.0",
	})
	examples = append(examples, climax.Example{
		Usecase: `--name example --description "Project description"`,
		Description: "Creates a project named example in the current directory, " +
			`with a riconto.toml file and description named "Project description"`,
	})
	examples = append(examples, climax.Example{
		Usecase: "--name example --license MIT",
		Description: "Creates a project named example in the current directory, " +
			"with a riconto.toml file and license MIT",
	})
	examples = append(examples, climax.Example{
		Usecase: "--name example --format json",
		Description: "Creates a project named example in the current directory, " +
			"with a riconto.json file and version 0.0.1",
	})
	examples = append(examples, climax.Example{
		Usecase: "--name example --format yaml",
		Description: "Creates a project named example in the current directory, " +
			"with a riconto.yaml file and version 0.0.1",
	})
	examples = append(examples, climax.Example{
		Usecase: "--name example --format toml",
		Description: "Creates a project named example in the current directory, " +
			"with a riconto.toml file and version 0.0.1",
	})
	return &CreateCommand{
		name:     "create",
		brief:    "creates a new project",
		usage:    "--name name [--version version] [--format json|yaml|toml] [--description description] [--license license]",
		help:     wordwrap.String(strings.TrimSpace(helpStr), terminalWidth),
		group:    "",
		flags:    flags,
		examples: examples,
		fs:       fs,
		logger:   logger,
	}
}

func (i *CreateCommand) Name() string {
	return i.name
}

func (i *CreateCommand) Brief() string {
	return i.brief
}

func (i *CreateCommand) Usage() string {
	return i.usage
}

func (i *CreateCommand) Help() string {
	return i.help
}

func (i *CreateCommand) Group() string {
	return i.group
}

func (i *CreateCommand) Flags() []climax.Flag {
	return i.flags
}

func (i *CreateCommand) Examples() []climax.Example {
	return i.examples
}

func (i *CreateCommand) Run(context climax.Context) int {
	var err error
	// 1. Validate if name is passed
	if !context.Is("name") {
		i.logger.Error("The name parameter is required!")
		return 1
	}

	name, _ := context.Get("name")
	i.logger.Info(fmt.Sprintf("Creating project named %s", name))

	// 2. Get the format to use
	format := model.FormatToml
	if context.Is("format") {
		f, _ := context.Get("format")
		format, err = model.ParseFormat(f)
		if err != nil {
			format = model.FormatToml
		}
	}

	// 3. Get the version to use
	version := "0.0.1"
	if context.Is("version") {
		version, _ = context.Get("version")
	}

	// 4. Get the description to use
	description := ""
	if context.Is("description") {
		description, _ = context.Get("description")
	}

	// 5. Get the description to use
	license := ""
	if context.Is("license") {
		license, _ = context.Get("license")
	}

	strict := context.Is("strict")

	// 6. Check if a configuration file already exists in the current directory
	if i.fileExists("riconto.json") ||
		i.fileExists("riconto.toml") ||
		i.fileExists("riconto.yaml") {
		i.logger.Error("There is already a configuration file in the current directory!")
		if strict {
			return 1
		}
		return 0
	}

	// 7. Create a temporary directory
	tmpFs := createTmpFs()
	defer func() {
		_ = tmpFs.RemoveAll("/")
	}()

	// 8. Create the configuration file, with the correct values
	config := model.NewConfig(name, version, description)
	config.AddLicense(license)

	// 9. Write the configuration file to the temporary directory
	writer, err := tmpFs.OpenFile("riconto."+format.String(), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		i.logger.Error("Unable to create the configuration file", slog.Any("error", err))
		return 1
	}
	defer func(writer fs.File) {
		_ = writer.Close()
	}(writer)
	err = config.SaveTo(writer, format)
	if err != nil {
		i.logger.Error("Unable to create the configuration file", slog.Any("error", err))
		return 1
	}

	// 10. Create the rest of the directories to the temporary directory
	err = tmpFs.MkdirAll("src", 0750)
	if err != nil {
		i.logger.Error("Unable to create the auxiliary directories", slog.Any("error", err))
		return 1
	}
	err = tmpFs.MkdirAll("resources", 0750)
	if err != nil {
		i.logger.Error("Unable to create the auxiliary directories", slog.Any("error", err))
		return 1
	}
	touch, err := tmpFs.OpenFile(path.Join("src", "main.md"), os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		i.logger.Error("Unable to create the auxiliary file", slog.Any("error", err))
		return 1
	}
	_ = touch.Close()

	// 11. Copy the temporary directory contents into the output directory
	err = utils.MergeFilesystem(tmpFs, i.fs, "")
	if err != nil {
		i.logger.Error("Unable to copy the temporary dir to the output directory", slog.Any("error", err))
		return 1
	}

	return 0
}

func (i *CreateCommand) Command() climax.Command {
	return FromCommand(i)
}

func (i *CreateCommand) fileExists(filename string) bool {
	info, err := i.fs.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createTmpFs() afero.Fs {
	return afero.NewMemMapFs()
}
