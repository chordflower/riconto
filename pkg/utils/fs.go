/*
 * Copyright (C) 2024 carddamom
 *
 * This file is part of riconto.
 *
 * riconto is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * riconto is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with riconto.  If not, see <https://www.gnu.org/licenses/>.
 */

package utils

import (
	"io/fs"

	"emperror.dev/errors"
	"github.com/spf13/afero"
)

//go:generate go-enum --marshal

// ENUM(error,keep,overwrite)
type ConflictResolution string

func MergeFilesystem(origin, destiny afero.Fs, basepath string) error {
	return MergeFilesystemWithConflictResolution(origin, destiny, basepath, ConflictResolutionOverwrite)
}

func MergeFilesystemWithConflictResolution(origin, destiny afero.Fs, basepath string, resolution ConflictResolution) error {
	err := afero.Walk(origin, basepath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err := destiny.MkdirAll(path, info.Mode())
			if err != nil {
				return err
			}
		} else {
			data, err := afero.ReadFile(origin, path)
			if err != nil {
				return err
			}
			if _, err := destiny.Stat(path); err == nil {
				switch resolution {
				case ConflictResolutionError:
					return errors.Errorf("The file in path %s already exists!", path)
				case ConflictResolutionKeep:
					return nil
				default:
					err = afero.WriteFile(destiny, path, data, info.Mode())
					if err != nil {
						return err
					}
				}
			} else {
				err = afero.WriteFile(destiny, path, data, info.Mode())
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
