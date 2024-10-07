//go:build windows

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

package userdirs

import (
	"emperror.dev/errors"
	"golang.org/x/sys/windows/registry"
	"slices"
)

//go:generate go-enum --marshal

// ENUM(FIXED, PERUSER)
type KnownFolderType string

type KnownFolder struct {
	GUID string
	Name string
	Type KnownFolderType
}

var (
	Desktop = &KnownFolder{
		GUID: "{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}",
		Name: "Desktop",
		Type: KnownFolderTypePERUSER,
	}
	Documents = &KnownFolder{
		GUID: "{FDD39AD0-238F-46AF-ADB4-6C85480369C7}",
		Name: "Documents",
		Type: KnownFolderTypePERUSER,
	}
	Downloads = &KnownFolder{
		GUID: "{374DE290-123F-4565-9164-39C4925E467B}",
		Name: "Downloads",
		Type: KnownFolderTypePERUSER,
	}
	Local = &KnownFolder{
		GUID: "{F1B32785-6FBA-4FCF-9D55-7B8E7F157091}",
		Name: "Local",
		Type: KnownFolderTypePERUSER,
	}
	Music = &KnownFolder{
		GUID: "{4BD8D571-6D19-48D3-BE97-422220080E43}",
		Name: "Music",
		Type: KnownFolderTypePERUSER,
	}
	Pictures = &KnownFolder{
		GUID: "{33E28130-4E1E-4676-835A-98395C3BC3BB}",
		Name: "Pictures",
		Type: KnownFolderTypePERUSER,
	}
	Public = &KnownFolder{
		GUID: "{DFDF76A2-C82A-4D63-906A-5644AC457385}",
		Name: "Public",
		Type: KnownFolderTypeFIXED,
	}
	Roaming = &KnownFolder{
		GUID: "{3EB685DB-65F9-4CF6-A03A-E3EF65729F3D}",
		Name: "Roaming",
		Type: KnownFolderTypePERUSER,
	}
	Templates = &KnownFolder{
		GUID: "{A63293E8-664E-48DB-A079-DF759E0509F7}",
		Name: "Templates",
		Type: KnownFolderTypePERUSER,
	}
	UserProgramFiles = &KnownFolder{
		GUID: "{5CD7AEE2-2219-4A67-B85D-6C9CE15660CB}",
		Name: "Programs",
		Type: KnownFolderTypePERUSER,
	}
	Videos = &KnownFolder{
		GUID: "{18989B1D-99B5-455B-841C-AB7C74E4DDFC}",
		Name: "Videos",
		Type: KnownFolderTypePERUSER,
	}
)

type userdirsImpl struct {
	data      string
	config    string
	state     string
	apps      string
	cache     string
	desktop   string
	download  string
	templates string
	public    string
	documents string
	music     string
	pictures  string
	videos    string
}

func newUserdirsImpl() (*userdirsImpl, error) {
	key, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`Software\Microsoft\Windows\CurrentVersion\Explorer\FolderDescriptions`,
		registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS|registry.READ|registry.WOW64_64KEY,
	)
	if err != nil {
		return nil, err
	}
	defer key.Close()
	userdirs := &userdirsImpl{}
	userdirs.desktop, err = getPathByFolderType(key, Desktop)
	if err != nil {
		return nil, err
	}

	userdirs.documents, err = getPathByFolderType(key, Documents)
	if err != nil {
		return nil, err
	}

	userdirs.download, err = getPathByFolderType(key, Downloads)
	if err != nil {
		return nil, err
	}

	var path string
	path, err = getPathByFolderType(key, Local)
	if err != nil {
		return nil, err
	}
	userdirs.data = path
	userdirs.state = path
	userdirs.cache = path

	userdirs.music, err = getPathByFolderType(key, Music)
	if err != nil {
		return nil, err
	}

	userdirs.pictures, err = getPathByFolderType(key, Pictures)
	if err != nil {
		return nil, err
	}

	userdirs.public, err = getPathByFolderType(key, Public)
	if err != nil {
		return nil, err
	}

	userdirs.config, err = getPathByFolderType(key, Roaming)
	if err != nil {
		return nil, err
	}

	userdirs.templates, err = getPathByFolderType(key, Templates)
	if err != nil {
		return nil, err
	}

	userdirs.apps, err = getPathByFolderType(key, UserProgramFiles)
	if err != nil {
		return nil, err
	}

	userdirs.videos, err = getPathByFolderType(key, Videos)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getPathByFolderType(key registry.Key, folder *KnownFolder) (string, error) {
	if !findSubkeyByName(key, folder.GUID) {
		return "", nil
	}

	folderKey, err := registry.OpenKey(key, folder.GUID, registry.QUERY_VALUE|registry.READ|registry.WOW64_64KEY)
	if err != nil {
		return "", err
	}
	defer folderKey.Close()

	if folder.Type == KnownFolderTypePERUSER {
		return getPerUserFolderPath(folderKey)
	}
	path, _, err := folderKey.GetStringValue("Name")
	return path, err
}

func getPerUserFolderPath(folderKey registry.Key) (string, error) {
	name, _, err := folderKey.GetStringValue("Name")
	if err != nil {
		return "", err
	}

	localKeys, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Explorer\User Shell Folders`,
		registry.QUERY_VALUE|registry.READ|registry.WOW64_64KEY,
	)
	if err != nil {
		return "", err
	}
	defer localKeys.Close()

	path, _, err := localKeys.GetStringValue(name)
	if err != nil {
		if errors.Is(err, registry.ErrNotExist) {
			path, _, err := folderKey.GetStringValue("Name")
			return path, err
		}
		return "", err
	}
	return path, nil
}

func findSubkeyByName(key registry.Key, name string) bool {
	names, err := key.ReadSubKeyNames(200)
	if err != nil {
		return false
	}
	return slices.Contains(names, name)
}

// function to return where user-specific data files should be written.
func (u userdirsImpl) DataHome() string {
	return u.data // Local
}

func (u userdirsImpl) ConfigHome() string {
	return u.config // Roaming
}

func (u userdirsImpl) StateHome() string {
	return u.state // Local
}

func (u userdirsImpl) AppsHome() string {
	return u.apps // UserProgramFiles
}

func (u userdirsImpl) CacheHome() string {
	return u.cache // Local
}

func (u userdirsImpl) DesktopHome() string {
	return u.desktop // Desktop
}

func (u userdirsImpl) DownloadHome() string {
	return u.download // Download
}

func (u userdirsImpl) TemplatesHome() string {
	return u.templates // Templates
}

func (u userdirsImpl) PublicHome() string {
	return u.public // Public
}

func (u userdirsImpl) DocumentsHome() string {
	return u.documents // Documents
}

func (u userdirsImpl) MusicHome() string {
	return u.music // Music
}

func (u userdirsImpl) PicturesHome() string {
	return u.pictures // Pictures
}

func (u userdirsImpl) VideosHome() string {
	return u.videos // Videos
}

var imple *userdirsImpl

func GetUserDirs() (u Userdirs, err error) {
	if imple == nil {
		imple, err = newUserdirsImpl()
	}
	u = imple
	return
}
