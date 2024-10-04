//go:build windows

package userdirs

import (
	"errors"
	"fmt"
)

const (
	// KnownFolderTypeFIXED is a KnownFolderType of type FIXED.
	KnownFolderTypeFIXED KnownFolderType = "FIXED"
	// KnownFolderTypePERUSER is a KnownFolderType of type PERUSER.
	KnownFolderTypePERUSER KnownFolderType = "PERUSER"
)

var ErrInvalidKnownFolderType = errors.New("not a valid KnownFolderType")

// String implements the Stringer interface.
func (x KnownFolderType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x KnownFolderType) IsValid() bool {
	_, err := ParseKnownFolderType(string(x))
	return err == nil
}

var _KnownFolderTypeValue = map[string]KnownFolderType{
	"FIXED":   KnownFolderTypeFIXED,
	"PERUSER": KnownFolderTypePERUSER,
}

// ParseKnownFolderType attempts to convert a string to a KnownFolderType.
func ParseKnownFolderType(name string) (KnownFolderType, error) {
	if x, ok := _KnownFolderTypeValue[name]; ok {
		return x, nil
	}
	return KnownFolderType(""), fmt.Errorf("%s is %w", name, ErrInvalidKnownFolderType)
}

// MarshalText implements the text marshaller method.
func (x KnownFolderType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *KnownFolderType) UnmarshalText(text []byte) error {
	tmp, err := ParseKnownFolderType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
