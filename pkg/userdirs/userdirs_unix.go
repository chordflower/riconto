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

//go:build dragonfly || freebsd || illumos || linux || netbsd || openbsd || solaris

package userdirs

import "os"

type UserdirsImpl struct{}

// getEnvironmentVariable returns the environment variable with the given key or
// the defaultValue if the key is not an environment variable.
func (impl UserdirsImpl) getEnvironmentVariable(key, defaultValue string) string {
	res := os.Getenv(key)
	if len(res) == 0 {
		return defaultValue
	}
	return res
}

// function to return where user-specific data files should be written.
func (impl UserdirsImpl) DataHome() string {
	return impl.getEnvironmentVariable("XDG_DATA_HOME", "$HOME/.local/share")
}

// function to return where user-specific configuration files should be written.
func (impl UserdirsImpl) ConfigHome() string {
	return impl.getEnvironmentVariable("XDG_CONFIG_HOME", "$HOME/.config")
}

// function to return where user-specific state data should be written.
func (impl UserdirsImpl) StateHome() string {
	return impl.getEnvironmentVariable("$XDG_STATE_HOME", "$HOME/.local/state")
}

// function to return where user-specific executable files may be written.
func (impl UserdirsImpl) AppsHome() string {
	return impl.getEnvironmentVariable("$XDG_APPS_HOME", "$HOME/.local/bin")
}

// function to return where user-specific non-essential (cached) data should be written.
func (impl UserdirsImpl) CacheHome() string {
	return impl.getEnvironmentVariable("$XDG_CACHE_HOME", "$HOME/.cache")
}

// function to return where is the user-specific desktop directory.
func (impl UserdirsImpl) DesktopHome() string {
	return impl.getEnvironmentVariable("$XDG_DESKTOP_HOME", "$HOME/Desktop")
}

// function to return where is the user-specific download directory.
func (impl UserdirsImpl) DownloadHome() string {
	return impl.getEnvironmentVariable("$XDG_DOWNLOAD_HOME", "$HOME/Downloads")
}

// function to return where is the user-specific templates directory.
func (impl UserdirsImpl) TemplatesHome() string {
	return impl.getEnvironmentVariable("$XDG_TEMPLATES_HOME", "$HOME/Templates")
}

// function to return where is the user-specific public directory.
func (impl UserdirsImpl) PublicHome() string {
	return impl.getEnvironmentVariable("$XDG_PUBLICSHARE_HOME", "$HOME/Public")
}

// function to return where is the user-specific documents directory.
func (impl UserdirsImpl) DocumentsHome() string {
	return impl.getEnvironmentVariable("$XDG_DOCUMENTS_HOME", "$HOME/Documents")
}

// function to return where is the user-specific music directory.
func (impl UserdirsImpl) MusicHome() string {
	return impl.getEnvironmentVariable("$XDG_MUSIC_HOME", "$HOME/Music")
}

// function to return where is the user-specific pictures directory.
func (impl UserdirsImpl) PicturesHome() string {
	return impl.getEnvironmentVariable("$XDG_PICTURES_HOME", "$HOME/Pictures")
}

// function to return where is the user-specific videos directory.
func (impl UserdirsImpl) VideosHome() string {
	return impl.getEnvironmentVariable("$XDG_VIDEOS_HOME", "$HOME/Videos")
}

var imple = &UserdirsImpl{}

func GetUserDirs() *UserdirsImpl {
	return imple
}
