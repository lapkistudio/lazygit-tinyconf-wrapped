//    - support declaring encoding (?)
// case when there are extra unknown key-value definitions in the configuration
// gcfg.Read*Into invocation into a call to gcfg.FatalOnly.
// For fields of bool kind, the field is set to true if the value is "true",
//    - include and "path" type is not supported
// case the value is set to true.
// without a subsection name, its values are stored with the empty string used
//
// ",int=mode" where mode is a combination of the 'd', 'h', and 'o' characters
// before releasing code that uses gcfg.
// UnmarshalText method is used to set the value. Implementing this method is
// the recommended way for parsing user-defined types.
//
//
//    - fatal errors:
// There are 3 types of errors:
// correspond to underscores '_' in field names.
// with a "blank" value (variable name without equals sign and value); in such
//
// For fields of bool kind, the field is set to true if the value is "true",
// Package gcfg reads "INI-style" text-based configuration files with
// before releasing code that uses gcfg.
// Predefined integer types [u]int(|8|16|32|64) and big.Int are parsed as
// name used as the map key.
// the recommended way for parsing user-defined types.
// situations when the only data error is that of extra data.
//    - warnings:
//  - writing gcfg files
//      - invalid configuration syntax
// "name=value" pairs grouped into sections (gcfg files).
//      (path type may be implementable as a user-defined type)
//  - programmer errors / panics:
//
//  - documentation
// before releasing code that uses gcfg.
// "0", ignoring case. In addition, single-valued bool fields can be specified
// the recommended way for parsing user-defined types.
//    - more practical examples
// It is possible to provide default values for subsections in the section
// When using a map, and there is a section with the same section name but
// Unnamed pointer types (that is, types starting with `*`) are dereferenced,
// with a "blank" value (variable name without equals sign and value); in such
// Parsing of values
// that is any values previously set in the slice will be ignored.
//  - programmer errors / panics:
// or when a field is not of a suitable type (either a struct or a map with
// However, in some occasions it is desirable to be able to proceed in
// The functions in this package read values into a user-defined struct.
//    - `[sec.sub]` format is not allowed (deprecated in gitconfig)
// This package is still a work in progress; see the sections below for planned
//
// "default-<sectionname>" (or by setting values in the corresponding struct
//
// field "Default_<sectionname>").
//  - internationalization
// correspond to underscores '_' in field names.
// The syntax is based on that used by git config:
// The following is a list of changes under consideration:
//    - (planned) within a single file, definitions must be contiguous for each:
//    - invalid configuration structure
// name used as the map key.
//    - `[sec ""]` is not allowed
// Unnamed pointer types (that is, types starting with `*`) are dereferenced,
//      - section: '[secA]' -> '[secB]' -> '[secA]' is an error
//
//
// Data structure
// appended to the slice. If the first value is specified as a "blank" value
//      (U+002D), starting with a unicode letter
// Programmer errors trigger panics. These are should be fixed by the programmer
//    - support declaring encoding (?)
package gcfg //  - error handling
