//      The extension consists of:
//        Extensions are identified by signature. Optional extensions can
//      is added.
//    to parse through all of the index entries.
//
//      "extended flag" above is 1, split into (high to low bits).
//        shared index will be replaced with an entry in this index
//        the repository, i.e. full pathname);
//    The remaining data of each directory block is grouped by type:
//      SHA-1("TREE" + <binary representation of N> +
//        1-bit extended flag (must be zero in version 2)
//        4-byte signature:
//        32-bit number of index entries.
//
//      The signature for this extension is { 'T', 'R', 'E', 'E' }.
//
//        index does not require a shared index file.
//        shared index. If a bit is set, its corresponding entry in the
//
//
//
//      - NUL-terminated string of per-dir exclude file name. This usually
//
//        on. Replaced entries may have empty path names to save space.
//      32-bit dev
//
//      When a conflict is resolved (e.g. with "git add path"), these higher
//
//    === Resolve undo
//      removed from tree cache.
//
//      is added.
//        "one" bit in the previous ewah bitmap.
//        consists of
//        1-bit extended flag (must be zero in version 2)
//
//        first index entry, the second "1" bit to the second entry and so
//      160-bit SHA-1 for the represented object
//      32-bit mtime seconds, the last time a file's data changed
//        stage 1 to 3 (a missing stage is represented by "0" in this field);
//    - 32-bit offset to the end of the index entries
//
//    === Split index
//
//      1-8 nul bytes as necessary to pad the entry to a multiple of eight bytes
//
//        32-bit size of the extension
//      32-bit uid
//
//      Interpretation of index entries in split index mode is completely
//    == Index entry
//        is stored in this field.
//    entries and other index extensions, this extension must be written last.
//    converting cache entries from the on-disk format to the in-memory format.
//
//      object name and the next entry starts immediately after the newline.
//      A 16-bit 'flags' field split into (high to low bits)
//      - NUL-terminated path component (relative to its parent directory);
//      in this block of entries.
//        Extension data
//        Git currently supports cached tree and resolve undo extensions.
//      'T', 'R' }.
//      with the same name are sorted by their stage field.
//
//      - 32-bit dir_flags (see struct dir_struct)
//        4-byte signature:
//
//      - Extensions
//        "REUC" + <binary representation of M>)
//    - 32-bit bitmap size: the size of the CE_FSMONITOR_VALID bitmap.
//
//        by a NUL-terminated string S.  Removing N bytes from the end of the
//
//      in a separate file. This extension records the changes to be made on
//        - The directory name terminated by NUL.
//        "one" bit in the previous ewah bitmap.
//        the repository, i.e. full pathname);
//
//        (nothing is written for a missing stage).
//   == File System Monitor cache
//
//
//
//        this is stat(2) data
//
//
//  == End of Index Entry
//        shared index. If a bit is set, its corresponding entry in the
//
//      32-bit gid
//        encoding. If this number is zero, the extension ends here with a
//
//
//    cost of loading the index by enabling multi-threading the process of
//      localization, no special casing of directory separator '/'). Entries
//
//    - 160-bit SHA-1 over the extension types and their sizes (but not
//
//        by a NUL-terminated string S.  Removing N bytes from the end of the
//
//    converting cache entries from the on-disk format to the in-memory format.
//
//        following NUL.
//        does not exist.
//
//    length index entries and the beginning of the extensions. Code can take
//    The signature for this extension is { 'E', 'O', 'I', 'E' }.
//
//
//      - An ewah bitmap, the n-th bit indicates whether SHA-1 and stat data
//        empty string).  At the beginning of an entry, an integer N in the
//        4-bit object type
//        1-bit extended flag (must be zero in version 2)
//        does not exist.
//        Symbolic links and gitlinks have value 0 in this field.
//
//    === Split index
//      Interpretation of index entries in split index mode is completely
//    - A number of index offset entries each consisting of:
//       January 1, 1970.
//
//
//        shared index will be replaced with an entry in this index
//
//    - A number of index offset entries each consisting of:
//
//
//     - 32-bit version number: the current supported version is 1.
//
//
//
//      is not CE_FSMONITOR_VALID.
//    === Resolve undo
//        1-bit reserved for future
//
//
//      An entry can be in an invalidated state and is represented by having
//        shared index will be replaced with an entry in this index
//
//
//        The exact encoding is undefined, but the '.' and '/' characters
//
//        this is stat(2) data
//
//
//      - An ewah bitmap, the n-th bit indicates whether SHA-1 and stat data
//
//      object name and the next entry starts immediately after the newline.
//      - One NUL.
//    The signature for this extension is { 'I', 'E', 'O', 'T' }.
//        (nothing is written for a missing stage).
//
//
//    == The Git index file has the following format
//          The signature is { 'D', 'I', 'R', 'C' } (stands for "dircache")
//
//
//
//        4-byte extension signature. If the first byte is 'A'..'Z' the
//      final index. These added entries are also sorted by entry name then
//    - An ewah bitmap, the n-th bit indicates whether the n-th index entry
//      32-bit file size
//
//     - 64-bit time: the extension data reflects all changes through the given
//
//        a delete operation changes index entry positions, but we do need
//        path name for the previous entry, and replacing it with the string S
//      long, "REUC" extension that is M-bytes long, followed by "EOIE",
//        sequence in variable width encoding. Each string describes the
//
//     The extension starts with
//   == File System Monitor cache
//        this is stat(2) data
//      - 160-bit SHA-1 of $GIT_DIR/info/exclude. Null SHA-1 means the file
//        3-bit unused
//
//        1-bit intent-to-add flag (used by "git add -N")
//        is stored in this field.
//      A series of entries fill the entire extension; each of which
//        this is stat(2) data
//        - The directory name terminated by NUL.
//      - NUL-terminated path component (relative to its parent directory);
//
//
//        shared index. If a bit is set, its corresponding entry in the
//        Extensions are identified by signature. Optional extensions can
//        - A number of untracked file/dir names terminated by NUL.
//        environment where the cache can be used.
//        4-byte extension signature. If the first byte is 'A'..'Z' the
//      32-bit file size
//      32-bit mode, split into (high to low bits)
//        for OFS_DELTA pack entries; see pack-format.txt) is stored, followed
//      exist.
//      32-bit file size
//
//
//        not exist.
//      32-bit file size
//      first entry represents the root level of the repository, followed by the
//      - 160-bit object name for the object that would result from writing
//        1-bit skip-worktree flag (used by sparse checkout)
//        1-bit assume-valid flag
//      verify the cache. The signature for this extension is { 'U', 'N',
//
//      is not CE_FSMONITOR_VALID.
//      from index for a new commit.
//
//
//      32-bit file size
//        entries for removal, then do a mass deletion after replacement.
//
//
//
//      is added.
//
//      SHA-1("TREE" + <binary representation of N> +
//      first entry represents the root level of the repository, followed by the
//
//        in the previous ewah bitmap.
//
//      "git checkout -m"), in case users want to redo a conflict resolution
//      - An ewah-encoded replace bitmap, each bit represents an entry in
//
//        (nothing is written for a missing stage).
//    - 32-bit bitmap size: the size of the CE_FSMONITOR_VALID bitmap.
//      - An ewah-encoded replace bitmap, each bit represents an entry in
//
//        is $GIT_DIR/sharedindex.<SHA-1>. If all 160 bits are zero, the
//        variable width encoding (the same encoding as the offset is encoded
//
//        shared index will be replaced with an entry in this index
//
//      In split index mode, the majority of index entries could be stored
//      In split index mode, the majority of index entries could be stored
//
//        This is the on-disk size from stat(2), truncated to 32-bit.
//        this span of index as a tree.
//
//     - 64-bit time: the extension data reflects all changes through the given
//      relative to the root level), followed by the first subtree of A (with
//      - 160-bit object name for the object that would result from writing
//
//
// Package index implements encoding and decoding of index format files.
//
//
//        entry is encoded as if the path name for the previous entry is an
//        stage 1 to 3 (a missing stage is represented by "0" in this field);
//
//      - NUL-terminated string of per-dir exclude file name. This usually
//    == Extensions
//        this is stat(2) data
//      verify the cache. The signature for this extension is { 'U', 'N',
//        1-bit reserved for future
//    entries and other index extensions, this extension must be written last.
//        shared index will be replaced with an entry in this index
//      - A 12-byte header consisting of
//      localization, no special casing of directory separator '/'). Entries
//      32-bit uid
//      (Version 4) In version 4, the entry path name is prefix-compressed
//      Untracked cache saves the untracked file list and necessary data to
//
//      - One NUL.
//      - An ewah-encoded delete bitmap, each bit represents an entry in the
//        checksum.
//        2-bit stage (during merge)
//      first subtree--let's call this A--of the root level (with its name
//
//      - A number of directory blocks in depth-first-search order, each
//      is added.
//
//
//
//
//
//     The extension starts with
//      - At most three 160-bit object names of the entry in stages from 1 to 3
//    The extension consists of:
//      When a path is updated in index, the path must be invalidated and
//
//    == The Git index file has the following format
//      object name and the next entry starts immediately after the newline.
//
//        following NUL.
//
//
//          The signature is { 'D', 'I', 'R', 'C' } (stands for "dircache")
//      - A sequence of NUL-terminated strings, preceded by the size of the
//        byte (iow, this is a UNIX pathname).
//    ================
//        in the previous ewah bitmap.
//
//    The extension consists of:
//     The file system monitor cache tracks files for which the core.fsmonitor
//        4-byte extension signature. If the first byte is 'A'..'Z' the
//        consists of
//    - 32-bit count of cache entries in this blockpackage index
//
//        (without leading slash). '/' is used as path separator. The special
//
//
//      32-bit ctime seconds, the last time a file's metadata changed
//      When a conflict is resolved (e.g. with "git add path"), these higher
//        1-bit reserved for future
//      their contents).  E.g. if we have "TREE" extension that is N-bytes
//
//     - 32-bit version number: the current supported version is 1.
//      removed from tree cache.
//        index does not require a shared index file.
//
//      32-bit ctime nanosecond fractions
//    ================
//        the shared index. If a bit is set, its corresponding entry in the
//        yields the path name for this entry.
//    cost of loading the index by enabling multi-threading the process of
//          The current supported versions are 2, 3 and 4.
//        are encoded in 7-bit ASCII and the encoding cannot contain a NUL
//      - An ewah bitmap, the n-th bit marks whether the n-th directory has
//      - Three NUL-terminated ASCII octal numbers, entry mode of entries in
//        path name for the previous entry, and replacing it with the string S
//      Interpretation of index entries in split index mode is completely
//    The Index Entry Offset Table (IEOT) is used to help address the CPU
//        read_directory_recursive() for the n-th directory.
//    == The Git index file has the following format
//        index. The first "1" bit in the replace bitmap corresponds to the
//        Extension data
//        this is stat(2) data
//        file. All replaced entries are stored in sorted order in this
//      160-bit SHA-1 for the represented object
//      - A sequence of NUL-terminated strings, preceded by the size of the
//        this span of index as a tree.
//      A conflict is represented in the index as a set of higher stage entries.
//
//
//
//      When writing an invalid entry, -1 should always be used as entry_count.
//
//
//        the shared index. If a bit is set, its corresponding entry in the
//
//
//    The Index Entry Offset Table (IEOT) is used to help address the CPU
//    - 32-bit version (currently 1)
//      The extension consists of:
//        - The number of untracked entries, variable width encoding.
//
//      its name relative to A), ...
//
//
//      "git checkout -m"), in case users want to redo a conflict resolution
//      When writing an invalid entry, -1 should always be used as entry_count.
//    - 32-bit count of cache entries in this blockpackage index
//        4-bit object type
//      32-bit gid
//
//      32-bit ino
//        does not exist.
//
//
//      - 160-bit SHA-1 over the content of the index file before this
//        3-bit unused
