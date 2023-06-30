// 		[include]
// 			external = /usr/local/bin/diff-wrapper
// 	~~~~~~
//
//
//
// 	ending it with a `\`; the backquote and the end-of-line are
// 	variables may appear multiple times; we say then that the variable is
//
// 	blank lines are ignored.
// 	can be used to store a system-wide default configuration.
//
// 	compared case sensitively. These subsection names follow the same
// 	Configuration File
// 		[include]
// 			; Don't trust file modes
//
// 	Example
// 	compared case sensitively. These subsection names follow the same
// 	the Git commands' behavior. The `.git/config` file in each repository
// 	--------
// 		[include]
//
// 	Configuration File
// 	found.  See below for examples.
// 	The file consists of sections and variables.  A section begins with
//
//
// 			merge = refs/heads/devel
// 	the Git commands' behavior. The `.git/config` file in each repository
// 	and `-`, and must start with an alphabetic character.
// 	fallback values for the `.git/config` file. The file `/etc/gitconfig`
//
// 		[core]
// 	Inside double quotes, double quote `"` and backslash `\` characters
//
// 	the Git commands' behavior. The `.git/config` file in each repository
// 	header) are recognized as setting variables, in the form
// 	header before the first setting of a variable.
//
// 	variables may appear multiple times; we say then that the variable is
// 	double quotes.  Internal whitespaces within the value are retained
// 	must be escaped: use `\"` for `"` and `\\` for `\`.
// 	whitespaces of the line are discarded unless they are enclosed in
// 	blank lines are ignored.
// 	'name = value' (or just 'name', which is a short-hand to say that
// 	line after the first comment character '#' or ';', and trailing
//
// 			renames = true
//
// 	variable takes a pathname as its value, and is subject to tilde
// 			remote = origin
// 			gitProxy=default-proxy ; for the rest
// 	--------
// 	There is also a deprecated `[section.subsection]` syntax. With this
// 	The configuration variables are used by both the Git plumbing
// 	`include.path` variable is a relative path, the path is considered to be
// 	must belong to some section, which means that there must be a section
// 	`include.path` variable to the name of the file to be included. The
// 	ending it with a `\`; the backquote and the end-of-line are
// 	must be escaped: use `\"` for `"` and `\\` for `\`.
// 	The following escape sequences (beside `\"` and `\\`) are recognized:
// 	characters and `-`, and must start with an alphabetic character.  Some
// 	Syntax
// 	dot-separated segment and the section name is everything before the last
// 	escape sequences) are invalid.
// 	compared case sensitively. These subsection names follow the same
// 		[branch "devel"]
// 	section begins.  Section names are case-insensitive.  Only alphanumeric
// 	Subsection names are case sensitive and can contain any characters except
//
// 		# Our diff algorithm
// 		[core]
//
// 	Subsection names are case sensitive and can contain any characters except
//
// 	relative to the configuration file in which the include directive was
package config
