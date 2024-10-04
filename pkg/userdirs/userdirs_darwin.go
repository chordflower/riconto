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

//go:build darwin

package userdirs

import (
	"path"

	"github.com/progrium/darwinkit/macos/foundation"
)

type UserdirsImpl struct{}

// function to return where user-specific data files should be written.
func (impl UserdirsImpl) DataHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.ApplicationSupportDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where user-specific configuration files should be written.
func (impl UserdirsImpl) ConfigHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.ApplicationSupportDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where user-specific state data should be written.
func (impl UserdirsImpl) StateHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.ApplicationSupportDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where user-specific executable files may be written.
func (impl UserdirsImpl) AppsHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.ApplicationDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where user-specific non-essential (cached) data should be written.
func (impl UserdirsImpl) CacheHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.CachesDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific desktop directory.
func (impl UserdirsImpl) DesktopHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.DesktopDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific download directory.
func (impl UserdirsImpl) DownloadHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.DownloadsDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific templates directory.
func (impl UserdirsImpl) TemplatesHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.UserDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return path.Join(ret, "Templates")
}

// function to return where is the user-specific public directory.
func (impl UserdirsImpl) PublicHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.SharedPublicDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific documents directory.
func (impl UserdirsImpl) DocumentsHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.DocumentDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific music directory.
func (impl UserdirsImpl) MusicHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.MusicDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific pictures directory.
func (impl UserdirsImpl) PicturesHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.PicturesDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

// function to return where is the user-specific videos directory.
func (impl UserdirsImpl) VideosHome() string {
	res := foundation.FileManager_DefaultManager().URLsForDirectoryInDomains(
		foundation.MoviesDirectory,
		foundation.UserDomainMask,
	)
	var ret string
	for _, v := range res {
		ret = v.AbsoluteString()
	}
	return ret
}

var imple = &UserdirsImpl{}

func GetUserDirs() *UserdirsImpl {
	return imple
}
