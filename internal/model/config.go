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
	"emperror.dev/errors"
	"github.com/goccy/go-yaml"
	"github.com/json-iterator/go"
	"github.com/pelletier/go-toml/v2"
	"io"
	"slices"
)

// Config represents the project configuration file
type Config struct {
	Schema      string   `json:"$schema"`
	Name        string   `json:"name" yaml:"name" toml:"name"`
	Version     string   `json:"version" yaml:"version"  toml:"version"`
	Description string   `json:"description" yaml:"description" toml:"description"`
	License     []string `json:"license" yaml:"license" toml:"license"`
	Authors     []Author `json:"authors" yaml:"authors" toml:"authors"`
}

func newConfig() *Config {
	return &Config{
		Version: "0.0.1",
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
		License:     make([]string, 0),
		Authors:     make([]Author, 0),
	}
}

// NewConfigFrom copies the given configuration.
func NewConfigFrom(config *Config) *Config {
	res := &Config{
		Schema:      config.Name,
		Name:        config.Name,
		Version:     config.Version,
		Description: config.Description,
		License:     slices.Clone(config.License),
		Authors:     make([]Author, len(config.Authors)),
	}
	for _, author := range config.Authors {
		res.Authors = append(res.Authors, *NewAuthorFrom(&author))
	}
	return res
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
	default:
		decoder := jsoniter.NewDecoder(reader)
		err = decoder.Decode(result)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to decode the file as json")
		}
	}
	return result, nil
}
