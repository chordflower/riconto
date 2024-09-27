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
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/buger/goterm"
	"github.com/chordflower/riconto/internal/model"
	"github.com/muesli/reflow/wordwrap"
	"github.com/phsym/console-slog"
	"github.com/tucnak/climax"
)

type InitCommand struct {
	name     string
	brief    string
	usage    string
	help     string
	group    string
	flags    []climax.Flag
	examples []climax.Example
}

func NewInitCommand() *InitCommand {
	terminalWidth := goterm.Width()
	helpStr := "" +
		"This command will initializa a riconto project, creating the configuration file " +
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
		"Note that, if a configuration file is already present in the directory, the command " +
		"does nothing and exists with an error, even if the configuration file is in a different " +
		"format!\n" +
		"As for the directories, if they are present, they will not be created or overwritten."
	flags := make([]climax.Flag, 0, 4)
	flags = append(flags, climax.Flag{
		Name:     "name",
		Short:    "n",
		Usage:    "--name NAME",
		Help:     "The name of the project in the configuration file (required)",
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
	return &InitCommand{
		name:     "init",
		brief:    "initializes a new project",
		usage:    "--name name [--version version] [--format json|yaml|toml]",
		help:     wordwrap.String(strings.TrimSpace(helpStr), terminalWidth),
		group:    "",
		flags:    flags,
		examples: examples,
	}
}

func (i *InitCommand) Name() string {
	return i.name
}

func (i *InitCommand) Brief() string {
	return i.brief
}

func (i *InitCommand) Usage() string {
	return i.usage
}

func (i *InitCommand) Help() string {
	return i.help
}

func (i *InitCommand) Group() string {
	return i.group
}

func (i *InitCommand) Flags() []climax.Flag {
	return i.flags
}

func (i *InitCommand) Examples() []climax.Example {
	return i.examples
}

func (i *InitCommand) Run(context climax.Context) int {
	logger := slog.New(
		console.NewHandler(os.Stdout, &console.HandlerOptions{
			AddSource:  true,
			Level:      slog.LevelInfo,
			TimeFormat: time.RFC3339,
		}),
	)
	var err error
	// 1. Validate if name is passed
	if !context.Is("name") {
		logger.Error("The name parameter is required!")
		return 1
	}

	name, _ := context.Get("name")
	logger.Info(fmt.Sprintf("Creating project named %s", name))

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

	// 4. Check if a configuration file already exists in the current directory
	currdir, err := os.Getwd()
	if err != nil {
		logger.Error(err.Error())
		return 1
	}
	if fileExists(filepath.Join(currdir, "riconto.json")) ||
		fileExists(filepath.Join(currdir, "riconto.toml")) ||
		fileExists(filepath.Join(currdir, "riconto.yaml")) {
		logger.Error("There is already a configuration file in the current directory!")
		return 1
	}

	// 5. Create a temporary directory
	tmpdir, err := os.MkdirTemp("", "riconto-")
	if err != nil {
		logger.Error("Unable to create the temporary directory", slog.Any("error", err))
		return 1
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tmpdir)

	// 6. Create the configuration file, with the correct values
	config := model.NewConfig(name, version, "")

	// 7. Write the configuration file to the temporary directory
	writer, err := os.Create(filepath.Join(tmpdir, "riconto."+format.String()))
	if err != nil {
		logger.Error("Unable to create the configuration file", slog.Any("error", err))
		return 1
	}
	defer func(writer *os.File) {
		_ = writer.Close()
	}(writer)
	err = config.SaveTo(writer, format)
	if err != nil {
		logger.Error("Unable to create the configuration file", slog.Any("error", err))
		return 1
	}

	// 8. Create the rest of the directories to the temporary directory
	err = os.MkdirAll(filepath.Join(tmpdir, "src"), 0750)
	if err != nil {
		logger.Error("Unable to create the auxiliary directories", slog.Any("error", err))
		return 1
	}
	err = os.MkdirAll(filepath.Join(tmpdir, "resources"), 0750)
	if err != nil {
		logger.Error("Unable to create the auxiliary directories", slog.Any("error", err))
		return 1
	}
	touch, err := os.Create(filepath.Join(tmpdir, "src", "main.md"))
	if err != nil {
		logger.Error("Unable to create the auxiliary file", slog.Any("error", err))
		return 1
	}
	_ = touch.Close()

	// 9. Copy the temporary directory contents into the output directory
	err = os.CopyFS(currdir, os.DirFS(tmpdir))
	if err != nil {
		logger.Error("Unable to copy the temporary dir to the output directory", slog.Any("error", err))
		return 1
	}

	// 10. Delete the temporary directory (done in defer)

	return 0
}

func (i *InitCommand) Command() climax.Command {
	return FromCommand(i)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
