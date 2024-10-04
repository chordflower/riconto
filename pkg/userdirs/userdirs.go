// Copyright (C) 2024 carddamom
//
// This file is part of agen.
//
// agen is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// agen is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with agen.  If not, see <https://www.gnu.org/licenses/>.

package userdirs

// Represents the interface used to retrieve common user specific directories for various types of
// stuff, like the documents directory, download directory, desktop, etc.
//
// This is specific for system, so to retrieve your implementation simply call GetUserDirs function,
// to return an singleton instance of the apropriate implementation.
type Userdirs interface {
	// function to return where user-specific data files should be written.
	DataHome() string

	// function to return where user-specific configuration files should be written.
	ConfigHome() string

	// function to return where user-specific state data should be written.
	StateHome() string

	// function to return where user-specific executable files may be written.
	AppsHome() string

	// function to return where user-specific non-essential (cached) data should be written.
	CacheHome() string

	// function to return where is the user-specific desktop directory.
	DesktopHome() string

	// function to return where is the user-specific download directory.
	DownloadHome() string

	// function to return where is the user-specific templates directory.
	TemplatesHome() string

	// function to return where is the user-specific public directory.
	PublicHome() string

	// function to return where is the user-specific documents directory.
	DocumentsHome() string

	// function to return where is the user-specific music directory.
	MusicHome() string

	// function to return where is the user-specific pictures directory.
	PicturesHome() string

	// function to return where is the user-specific videos directory.
	VideosHome() string
}
