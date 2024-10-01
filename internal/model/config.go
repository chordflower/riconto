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

//go:generate go-enum --marshal

import (
	"io"
	"slices"

	jsoniter "github.com/json-iterator/go"

	"emperror.dev/errors"
	"github.com/goccy/go-yaml"
	"github.com/pelletier/go-toml/v2"
)

// Config represents the project configuration file
type Config struct {
	Name        string   `json:"name" yaml:"name" toml:"name"`
	Version     string   `json:"version" yaml:"version"  toml:"version"`
	Description string   `json:"description" yaml:"description" toml:"description"`
	Files       []File   `json:"files" toml:"files" yaml:"files"`
	License     []string `json:"license" yaml:"license" toml:"license"`
	Authors     []Author `json:"authors" yaml:"authors" toml:"authors"`
}

func newConfig() *Config {
	return &Config{
		Version: "0.0.1",
		Files:   make([]File, 0),
		License: make([]string, 0),
		Authors: make([]Author, 0),
	}
}

// NewConfig creates a new configuration from the given values
func NewConfig(name string, version string, description string) *Config {
	return &Config{
		Name:        name,
		Version:     version,
		Description: description,
		Files:       make([]File, 0),
		License:     make([]string, 0),
		Authors:     make([]Author, 0),
	}
}

// NewConfigFrom copies the given configuration.
func NewConfigFrom(config *Config) *Config {
	res := &Config{
		Name:        config.Name,
		Version:     config.Version,
		Description: config.Description,
		Files:       make([]File, 0, len(config.Files)),
		License:     slices.Clone(config.License),
		Authors:     make([]Author, 0, len(config.Authors)),
	}
	for _, author := range config.Authors {
		res.Authors = append(res.Authors, *NewAuthorFrom(&author))
	}
	for _, file := range config.Files {
		res.Files = append(res.Files, *NewFileFrom(&file))
	}
	return res
}

// SaveTo saves the configuration to the given writer in the given format
func (c *Config) SaveTo(writer io.Writer, format Format) error {
	var err error
	switch format {
	case FormatJson:
		encoder := jsoniter.NewEncoder(writer)
		err = encoder.Encode(c)
		if err != nil {
			return errors.Wrap(err, "Unable to encode the file as json")
		}
	case FormatToml:
		encoder := toml.NewEncoder(writer)
		err = encoder.Encode(c)
		if err != nil {
			return errors.Wrap(err, "Unable to encode the file as toml")
		}
	case FormatYaml:
		encoder := yaml.NewEncoder(writer)
		err = encoder.Encode(c)
		if err != nil {
			return errors.Wrap(err, "Unable to encode the file as yaml")
		}
	}
	return nil
}

// AddLicense adds the given license to this configuration
func (c *Config) AddLicense(license string) bool {
	if !slices.Contains(c.License, license) {
		c.License = append(c.License, license)
		return true
	}
	return false
}

// RemoveLicense removes the given license from this configuration
func (c *Config) RemoveLicense(license string) bool {
	if slices.Contains(c.License, license) {
		c.License = slices.DeleteFunc(c.License, func(lic string) bool {
			return lic == license
		})
		return true
	}
	return false
}

// ContainsLicense checks if the given license is in this configuration
func (c *Config) ContainsLicense(license string) bool {
	return slices.Contains(c.License, license)
}

// AddAuthor adds the given author to this configuration
func (c *Config) AddAuthor(author *Author) bool {
	if !slices.ContainsFunc(c.Authors, func(aut Author) bool {
		return aut.Name == author.Name
	}) {
		c.Authors = append(c.Authors, *author)
		return true
	}
	return false
}

// RemoveAuthor removes the given author from this configuration
func (c *Config) RemoveAuthor(author *Author) bool {
	if slices.ContainsFunc(c.Authors, func(aut Author) bool {
		return aut.Name == author.Name
	}) {
		c.Authors = slices.DeleteFunc(c.Authors, func(aut Author) bool {
			return aut.Name == author.Name
		})
		return true
	}
	return false
}

// ContainsAuthor checks if this configuration contains the given author
func (c *Config) ContainsAuthor(author *Author) bool {
	return slices.ContainsFunc(c.Authors, func(aut Author) bool {
		return aut.Name == author.Name
	})
}

// AddFile adds the given file to this configuration
func (c *Config) AddFile(file *File) bool {
	if !slices.ContainsFunc(c.Files, func(f File) bool {
		return f.Name == file.Name
	}) {
		c.Files = append(c.Files, *file)
		return true
	}
	return false
}

// RemoveFile removes the given file from this configuration
func (c *Config) RemoveFile(file *File) bool {
	if slices.ContainsFunc(c.Files, func(f File) bool {
		return f.Name == file.Name
	}) {
		c.Files = slices.DeleteFunc(c.Files, func(f File) bool {
			return f.Name == file.Name
		})
		return true
	}
	return false
}

// ContainsFile checks if this configuration contains the given file
func (c *Config) ContainsFile(file *File) bool {
	return slices.ContainsFunc(c.Files, func(f File) bool {
		return f.Name == file.Name
	})
}

// Author represents an package author
type Author struct {
	Name  string `json:"name" yaml:"name" toml:"name"`
	URL   string `json:"url" yaml:"URL" toml:"url"`
	Email string `json:"email" yaml:"email" toml:"email"`
}

// NewAuthor creates a new author with the given data
func NewAuthor(name string) *Author {
	return &Author{
		Name: name,
	}
}

// NewAuthorFrom copies the given author
func NewAuthorFrom(author *Author) *Author {
	return &Author{
		Name:  author.Name,
		URL:   author.URL,
		Email: author.Email,
	}
}

// File represents a list of root files to build
type File struct {
	Name   string `json:"name" toml:"name" yaml:"name"`
	Output string `json:"output" toml:"output" yaml:"output"`
	Path   string `json:"path" toml:"path" yaml:"path"`
}

// NewFile creates a new file with the given data
func NewFile(name, output, path string) *File {
	return &File{
		Name:   name,
		Output: output,
		Path:   path,
	}
}

// NewFileFrom copies the given file
func NewFileFrom(file *File) *File {
	return &File{
		Name:   file.Name,
		Output: file.Output,
		Path:   file.Path,
	}
}

// ENUM(json, yaml, toml)
type Format string

// ConfigFromFile creates a new configuration in the given format,
// from the data in the given reader
func ConfigFromFile(reader io.Reader, format Format) (*Config, error) {
	result := newConfig()
	var err error
	switch format {
	case FormatJson:
		decoder := jsoniter.NewDecoder(reader)
		err = decoder.Decode(result)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to decode the file as json")
		}
	case FormatToml:
		decoder := toml.NewDecoder(reader)
		err = decoder.Decode(result)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to decode the file as toml")
		}
	case FormatYaml:
		decoder := yaml.NewDecoder(reader)
		err = decoder.Decode(result)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to decode the file as yaml")
		}
	}
	return result, nil
}
