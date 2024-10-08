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

package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/chordflower/riconto/internal/commands"
	"github.com/phsym/console-slog"
	"github.com/spf13/afero"
	"github.com/tucnak/climax"
)

func main() {
	logger := slog.New(
		console.NewHandler(os.Stdout, &console.HandlerOptions{
			AddSource:  true,
			Level:      slog.LevelInfo,
			TimeFormat: time.RFC3339,
		}),
	)

	osFs := afero.NewOsFs()
	currdir, err := os.Getwd()
	if err != nil {
		logger.Error("Error creating the project", slog.Any("error", err))
	}

	createCommand := commands.NewCreateCommand(afero.NewBasePathFs(
		osFs, currdir,
	), logger)
	riconto := climax.New("riconto")
	riconto.Brief = "A tool to create markdown based documents"
	riconto.Version = "0.0.1"
	riconto.AddCommand(createCommand.Command())
	riconto.Run()
}
