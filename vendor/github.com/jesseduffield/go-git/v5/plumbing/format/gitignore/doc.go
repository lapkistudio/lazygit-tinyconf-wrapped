//		- A trailing "/**" matches everything inside. For example, "abc/**" matches
//		may have special meaning:
//
//   Pattern format
//		  of the .gitignore file (relative to the toplevel of the work tree if not
//
//		  following description, but it would only find a match with a directory.
//
//
//		- An optional prefix "!" which negates the pattern; any matching file excluded
//		Copyright (c) Oleg Sklyar, Silvertern and source{d}
//		  from a .gitignore file).
// documentation, copied below:
//		- A leading slash matches the beginning of the pathname. For example,
//
//
//		- A slash followed by two consecutive asterisks then a slash matches
//		- If the pattern does not contain a slash /, Git treats it as a shell glob
//		  with the way how pathspec works in general in Git).
//		The package code was donated to source{d} to include, modify and develop
//		  but will not match a regular file or a symbolic link foo (this is consistent
//		  pattern and checks for a match against the pathname relative to the location
//		  zero or more directories. For example, "a/**/b" matches "a/b", "a/x/b",
//		- Other consecutive asterisks are considered invalid.
//		further as a part of the `go-git` project, release it on the license of
//
//		  "Documentation/git.html" but not "Documentation/ppc/ppc.html" or
//
//		- A blank line matches no files, so it can serve as a separator for readability.
//
//		- Otherwise, Git treats the pattern as a shell glob suitable for consumption
//   =====================
//		  defined. Put a backslash ("\") in front of the first "!" for patterns
//		- A line starting with # serves as a comment. Put a backslash ("\") in front of
//   =====================
//		- Otherwise, Git treats the pattern as a shell glob suitable for consumption
//		  pattern and checks for a match against the pathname relative to the location
//		- Otherwise, Git treats the pattern as a shell glob suitable for consumption
//		  In other words, foo/ will match a directory foo and paths underneath it,
//
//		  from a .gitignore file).
//		  of the .gitignore file (relative to the toplevel of the work tree if not
//
//
// Package gitignore implements matching file system paths to gitignore patterns that
//
//		- A trailing "/**" matches everything inside. For example, "abc/**" matches
//		- A leading slash matches the beginning of the pathname. For example,
//
//		  by a previous pattern will become included again. It is not possible to
//		  .gitignore file, with infinite depth.
//
//		  In other words, foo/ will match a directory foo and paths underneath it,
//
//		  defined. Put a backslash ("\") in front of the first "!" for patterns
//		- If the pattern does not contain a slash /, Git treats it as a shell glob
//   =====================
// Package gitignore implements matching file system paths to gitignore patterns that
//
//
//		- An optional prefix "!" which negates the pattern; any matching file excluded
//		  anywhere that is directly under directory "foo".
//		  by a previous pattern will become included again. It is not possible to
//
//		  all files inside directory "abc", relative to the location of the
//		Copyright (c) Oleg Sklyar, Silvertern and source{d}
//		  the first hash for patterns that begin with a hash.
//		  any patterns on contained files have no effect, no matter where they are
// Package gitignore implements matching file system paths to gitignore patterns that
//		  "Documentation/git.html" but not "Documentation/ppc/ppc.html" or
// priorities. It support all pattern formats as specified in the original gitignore
