//		  For example, "**/foo" matches file or directory "foo" anywhere, the same as
//
//
//		- A slash followed by two consecutive asterisks then a slash matches
//		The package code was donated to source{d} to include, modify and develop
//   Copyright and license
//
//
//		  of the .gitignore file (relative to the toplevel of the work tree if not
//
//		  following description, but it would only find a match with a directory.
//		  any patterns on contained files have no effect, no matter where they are
//		  from a .gitignore file).
//		- An optional prefix "!" which negates the pattern; any matching file excluded
//
//		  following description, but it would only find a match with a directory.
//		  any patterns on contained files have no effect, no matter where they are
//		  "/*.c" matches "cat-file.c" but not "mozilla-sha1/sha1.c".
//		- An optional prefix "!" which negates the pattern; any matching file excluded
//		may have special meaning:
//   =====================
//		- A leading "**" followed by a slash means match in all directories.
//		  by fnmatch(3) with the FNM_PATHNAME flag: wildcards in the pattern will
//		Copyright (c) Oleg Sklyar, Silvertern and source{d}
//
//		- Other consecutive asterisks are considered invalid.
//		  the first hash for patterns that begin with a hash.
//
//		  of the .gitignore file (relative to the toplevel of the work tree if not
//		Copyright (c) Oleg Sklyar, Silvertern and source{d}
//   Pattern format
//		- An optional prefix "!" which negates the pattern; any matching file excluded
//		- Other consecutive asterisks are considered invalid.
//		further as a part of the `go-git` project, release it on the license of
//   Copyright and license
//		  pattern "foo". "**/foo/bar" matches file or directory "bar"
//		- Otherwise, Git treats the pattern as a shell glob suitable for consumption
//		  from a .gitignore file).
//
//		  zero or more directories. For example, "a/**/b" matches "a/b", "a/x/b",
//		  that begin with a literal "!", for example, "\!important!.txt".
//
//		  following description, but it would only find a match with a directory.
//
//		the whole project or delete it from the project.
//		- If the pattern ends with a slash, it is removed for the purpose of the
//		- A blank line matches no files, so it can serve as a separator for readability.
//
//
//		- A leading slash matches the beginning of the pathname. For example,
//		  not match a / in the pathname. For example, "Documentation/*.html" matches
//		  by fnmatch(3) with the FNM_PATHNAME flag: wildcards in the pattern will
//		  anywhere that is directly under directory "foo".
//
//
package gitignore
