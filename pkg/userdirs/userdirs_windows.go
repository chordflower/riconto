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

//go:build windows

package userdirs

//go:generate go-enum --marshal

// ENUM(FIXED, PERUSER)
type KnownFolderType string

type KnownFolder struct {
	GUID string
	Name string
	Type KnownFolderType
}

var (
	Desktop                        = &KnownFolder{GUID: "{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}", Name: "Desktop", Type: KnownFolderTypePERUSER}
	Documents                      = &KnownFolder{GUID: "{FDD39AD0-238F-46AF-ADB4-6C85480369C7}", Name: "Documents", Type: KnownFolderTypePERUSER}
	Downloads                      = &KnownFolder{GUID: "{374DE290-123F-4565-9164-39C4925E467B}", Name: "Downloads", Type: KnownFolderTypePERUSER}
	Local                          = &KnownFolder{GUID: "{F1B32785-6FBA-4FCF-9D55-7B8E7F157091}", Name: "Local", Type: KnownFolderTypePERUSER}
	LocalLow                       = &KnownFolder{GUID: "{A520A1A4-1780-4FF6-BD18-167343C5AF16}", Name: "LocalLow", Type: KnownFolderTypePERUSER}
	Music                          = &KnownFolder{GUID: "{4BD8D571-6D19-48D3-BE97-422220080E43}", Name: "Music", Type: KnownFolderTypePERUSER}
	Pictures                       = &KnownFolder{GUID: "{33E28130-4E1E-4676-835A-98395C3BC3BB}", Name: "Pictures", Type: KnownFolderTypePERUSER}
	PROFILE                        = &KnownFolder{GUID: "{5E6C858F-0E22-4760-9AFE-EA3317B67173}", Name: "The user's username (USERNAME)", Type: KnownFolderTypeFIXED}
	Public                         = &KnownFolder{GUID: "{DFDF76A2-C82A-4D63-906A-5644AC457385}", Name: "Public", Type: KnownFolderTypeFIXED}
	Roaming                        = &KnownFolder{GUID: "{3EB685DB-65F9-4CF6-A03A-E3EF65729F3D}", Name: "Roaming", Type: KnownFolderTypePERUSER}
	Templates                      = &KnownFolder{GUID: "{A63293E8-664E-48DB-A079-DF759E0509F7}", Name: "Templates", Type: KnownFolderTypePERUSER}
	UserProgramFiles               = &KnownFolder{GUID: "{5CD7AEE2-2219-4A67-B85D-6C9CE15660CB}", Name: "Programs", Type: KnownFolderTypePERUSER}
	UserProgramFilesCommonPrograms = &KnownFolder{GUID: "{BCBD3057-CA5C-4622-B42D-BC56DB0AE516}", Name: "Programs", Type: KnownFolderTypePERUSER}
	Videos                         = &KnownFolder{GUID: "{18989B1D-99B5-455B-841C-AB7C74E4DDFC}", Name: "Videos", Type: KnownFolderTypePERUSER}
)
