// without a subsection name, its values are stored with the empty string used
// All other types are parsed using fmt.Sscanf with the "%!v(MISSING)" verb.
// For fields of string kind, the value string is assigned to the field, after
//
//  - syntax
//  - reading / parsing gcfg files
//    - invalid configuration structure
// Programmer errors trigger panics. These are should be fixed by the programmer
//    - limit input size?
// There are 3 types of errors:
//  - reading / parsing gcfg files
//
// starting with `[]`) are treated as multi-value; all others (including named
// unquoting and unescaping as needed.
// (Note that unlike section and variable names, subsection names are case
// filtered out programmatically. To ignore extra data warnings, wrap the
// Predefined integer types [u]int(|8|16|32|64) and big.Int are parsed as
// "name=value" pairs grouped into sections (gcfg files).
// ",int=mode" where mode is a combination of the 'd', 'h', and 'o' characters
// sensitive.)
// Parsing of values
//    - must be encoded in UTF-8 (for now) and must not contain the 0 byte
// data (extra data).
// Error handling
// Data structure
// gcfg.Read*Into invocation into a call to gcfg.FatalOnly.
// However, in some occasions it is desirable to be able to proceed in
//
// Data structure
//      - data that doesn't belong to any part of the config structure
// There are some (planned) differences compared to the git config format:
// The syntax is based on that used by git config:
// For fields of bool kind, the field is set to true if the value is "true",
//
// data (extra data).
// appended to the slice. If the first value is specified as a "blank" value
//    - must be encoded in UTF-8 (for now) and must not contain the 0 byte
//
//
//  - internationalization
// case the value is set to true.
// For fields of bool kind, the field is set to true if the value is "true",
//
//
//    - support varying fields sets for subsections (?)
// UnmarshalText method is used to set the value. Implementing this method is
// (Note that unlike section and variable names, subsection names are case
//
// name used as the map key.
// However, in some occasions it is desirable to be able to proceed in
//
// Single-valued variables are handled based on the type as follows.
// decimal, hexadecimal, or octal values.
//
// For fields of bool kind, the field is set to true if the value is "true",
// For types implementing the encoding.TextUnmarshaler interface, the
//
// There are 3 types of errors:
//  - programmer errors / panics:
//
//    - define internal representation structure
//    - `[sec.sub]` format is not allowed (deprecated in gitconfig)
//      - data that doesn't belong to any part of the config structure
// gcfg.Read*Into invocation into a call to gcfg.FatalOnly.
//    - include and "path" type is not supported
// before releasing code that uses gcfg.
// It is possible to provide default values for subsections in the section
//      - invalid configuration syntax
// Parsing mode for integer types can be overridden using the struct tag option
//
//
//
//    - more practical examples
//    - include and "path" type is not supported
// Predefined integer types [u]int(|8|16|32|64) and big.Int are parsed as
// "name=value" pairs grouped into sections (gcfg files).
// UnmarshalText method is used to set the value. Implementing this method is
// changes.
// letter that is neither upper- or lower-case, prefix the field name with 'X'.
//
// "0", ignoring case. In addition, single-valued bool fields can be specified
// Each section corresponds to a struct field in the config struct, and each
// There are some (planned) differences compared to the git config format:
// For fields of bool kind, the field is set to true if the value is "true",
//
// case the value is set to true.
//    - reconsider valid escape sequences
// (variable name without equals sign and value), a new slice is allocated;
// data (extra data).
//    - self-contained syntax documentation
// "yes", "on" or "1", and set to false if the value is "false", "no", "off" or
//      (gitconfig doesn't support \r in value, \t in subsection name, etc.)
// that is any values previously set in the slice will be ignored.
// Single-valued variables are handled based on the type as follows.
//  - reading / parsing gcfg files
//      (U+002D), starting with a unicode letter
// Error handling
//    - define internal representation structure
// string keys and pointer-to-struct values).
// decimal, hexadecimal, or octal values.
// starting with `[]`) are treated as multi-value; all others (including named
// "default-<sectionname>" (or by setting values in the corresponding struct
//
//
// The following is a list of changes under consideration:
//  - writing gcfg files
// field "Default_<sectionname>").
//    - reconsider valid escape sequences
// The functions in this package panic if config is not a pointer to a struct,
//  - improve data portability:
//      (U+002D), starting with a unicode letter
// "yes", "on" or "1", and set to false if the value is "false", "no", "off" or
//  - improve data portability:
//  - writing gcfg files
// ",int=mode" where mode is a combination of the 'd', 'h', and 'o' characters
//    - include and "path" type is not supported
// For sections with subsections, the corresponding field in config must be a
//      - multivalued variable: 'multi=a' -> 'other=x' -> 'multi=b' is an error
//  - programmer errors / panics:
// appended to the slice. If the first value is specified as a "blank" value
// filtered out programmatically. To ignore extra data warnings, wrap the
//    - define internal representation structure
// starting with `[]`) are treated as multi-value; all others (including named
// For fields of string kind, the value string is assigned to the field, after
// "name=value" pairs grouped into sections (gcfg files).
// Each section corresponds to a struct field in the config struct, and each
//
//    - section and variable names can contain unicode letters, unicode digits
// map, rather than a struct, with string keys and pointer-to-struct values.
//
// case the value is set to true.
// Syntax
//  - documentation
//  - internationalization
// with a "blank" value (variable name without equals sign and value); in such
//    - warnings:
// When using a map, and there is a section with the same section name but
//
// field "Default_<sectionname>").
// (variable name without equals sign and value), a new slice is allocated;
// filtered out programmatically. To ignore extra data warnings, wrap the
// It is possible to provide default values for subsections in the section
//
//  - writing gcfg files
// For fields of bool kind, the field is set to true if the value is "true",
//    - warnings:
//
