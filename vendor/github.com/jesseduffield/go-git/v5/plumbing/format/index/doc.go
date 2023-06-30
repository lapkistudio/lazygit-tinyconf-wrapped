//
//    The remaining data of each directory block is grouped by type:
//
//
//
//
//      32-bit ino
//
//
//
//    The signature for this extension is { 'I', 'E', 'O', 'T' }.
//
//
//        by a NUL-terminated string S.  Removing N bytes from the end of the
//      When a path is updated in index, the path must be invalidated and
//        checksum.
//    - 32-bit offset to the end of the index entries
//
//      160-bit SHA-1 for the represented object
//      "git checkout -m"), in case users want to redo a conflict resolution
//       January 1, 1970.
//        this is stat(2) data
//      different. See below for details.
//        1-bit intent-to-add flag (used by "git add -N")
//        first index entry, the second "1" bit to the second entry and so
//
//        12-bit name length if the length is less than 0xFFF; otherwise 0xFFF
//        path components ".", ".." and ".git" (without quotes) are disallowed.
//
//
//      In split index mode, the majority of index entries could be stored
//     The extension starts with
//
//    Because it must be able to be loaded before the variable length cache
//      The signature for this extension is { 'T', 'R', 'E', 'E' }.
//        relative to the path name for the previous entry (the very first
//
//      The extension starts with
//        4-bit object type
//        this is stat(2) data
//      "extended flag" above is 1, split into (high to low bits).
//        the shared index. If a bit is set, its corresponding entry in the
//      160-bit SHA-1 for the represented object
//    == Index entry
//
//      Index entries are sorted in ascending order on the name field,
//          The current supported versions are 2, 3 and 4.
//
//        The exact encoding is undefined, but the '.' and '/' characters
//
//      verify the cache. The signature for this extension is { 'U', 'N',
//      32-bit dev
//        this is stat(2) data
//   == File System Monitor cache
//
//
//    cost of loading the index by enabling multi-threading the process of
//      - 32-bit dir_flags (see struct dir_struct)
//        32-bit number of index entries.
//        ctime field until "file size".
//
//
//
//        extension is optional and can be ignored.
//
//
//        original positions in replace phase, it's best to just mark
//    The extension consists of:
//        index does not require a shared index file.
//      stage.
//     The file system monitor cache tracks files for which the core.fsmonitor
//
//
//      - Stat data of plumbing.excludesfile
//
//      in a separate file. This extension records the changes to be made on
//      160-bit SHA-1 for the represented object
//    converting cache entries from the on-disk format to the in-memory format.
//
//      from scratch.
//    converting cache entries from the on-disk format to the in-memory format.
//      "git checkout -m"), in case users want to redo a conflict resolution
//
//      - The number of following directory blocks, variable width
//
//    advantage of this to quickly locate the index extensions without having
//
//
//      resolve undo extension, so that conflicts can be recreated (e.g. with
//        - The number of untracked entries, variable width encoding.
//      - Stat data of $GIT_DIR/info/exclude. See "Index entry" section from
//      while keeping the name NUL-terminated.
//      - An ewah-encoded replace bitmap, each bit represents an entry in
//        path components ".", ".." and ".git" (without quotes) are disallowed.
//        environment where the cache can be used.
//      - An ewah bitmap, the n-th bit indicates whether SHA-1 and stat data
//        path name for the previous entry, and replacing it with the string S
//
//
//      When these higher stage entries are removed, they are saved in the
//      resolve undo extension, so that conflicts can be recreated (e.g. with
//      32-bit mode, split into (high to low bits)
//    == Extensions
//      is not CE_FSMONITOR_VALID.
//      32-bit dev
//      Untracked cache saves the untracked file list and necessary data to
//        is valid for the n-th directory and exists in the next data.
//      - An ewah bitmap, the n-th bit marks whether the n-th directory has
//
//      (Version 3 or later) A 16-bit field, only applicable if the
//        read_directory_recursive() for the n-th directory.
//        read_directory_recursive() for the n-th directory.
//
//      'T', 'R' }.
//      32-bit uid
//          valid values in binary are 1000 (regular file), 1010 (symbolic link)
//
//        this is stat(2) data
//        is $GIT_DIR/sharedindex.<SHA-1>. If all 160 bits are zero, the
//
//
//        this is stat(2) data
//      SHA-1("TREE" + <binary representation of N> +
//      - NUL-terminated pathname the entry describes (relative to the root of
//
//
//        this is stat(2) data
//    - An ewah bitmap, the n-th bit indicates whether the n-th index entry
//
//
//      32-bit mtime seconds, the last time a file's data changed
//    === Split index
//    ================
//       time which is stored as the nanoseconds elapsed since midnight,
//
//        is valid for the n-th directory and exists in the next data.
//      "extended flag" above is 1, split into (high to low bits).
//        a delete operation changes index entry positions, but we do need
//    - 32-bit offset from the beginning of the file to the first cache entry
//        are encoded in 7-bit ASCII and the encoding cannot contain a NUL
//
//      - Three NUL-terminated ASCII octal numbers, entry mode of entries in
//      When writing an invalid entry, -1 should always be used as entry_count.
//        in the previous ewah bitmap.
//        on. Replaced entries may have empty path names to save space.
//    The signature for this extension is { 'I', 'E', 'O', 'T' }.
//
//        4-bit object type
//      - Extensions
//      - NUL-terminated pathname the entry describes (relative to the root of
//        - The directory name terminated by NUL.
//        this is stat(2) data
//        for OFS_DELTA pack entries; see pack-format.txt) is stored, followed
//        byte (iow, this is a UNIX pathname).
//      from scratch.
//
//        checksum.
//      32-bit gid
//
//      Interpretation of index entries in split index mode is completely
//
//        following NUL.
//      160-bit SHA-1 for the represented object
//
//        3-bit unused
//
//    == Index entry
//        13-bit unused, must be zero
//
//
//        4-bit object type
//      A conflict is represented in the index as a set of higher stage entries.
//          and 1110 (gitlink)
//     The extension starts with
//    length index entries and the beginning of the extensions. Code can take
//        checksum.
//
//      from scratch.
//      resolve undo extension, so that conflicts can be recreated (e.g. with
//
//        yields the path name for this entry.
//    - 32-bit offset from the beginning of the file to the first cache entry
//        Symbolic links and gitlinks have value 0 in this field.
//      32-bit gid
//
//        shared index will be removed from the final index.  Note, because
//      'T', 'R' }.
//        relative to the path name for the previous entry (the very first
//
//
//
//          The current supported versions are 2, 3 and 4.
//
//     hook has told us about changes.  The signature for this extension is
//
//    Because it must be able to be loaded before the variable length cache
//      When these higher stage entries are removed, they are saved in the
//      1-8 nul bytes as necessary to pad the entry to a multiple of eight bytes
//        1-bit extended flag (must be zero in version 2)
//
//
//
//      - Three NUL-terminated ASCII octal numbers, entry mode of entries in
//
// Package index implements encoding and decoding of index format files.
//
//        shared index. If a bit is set, its corresponding entry in the
//    - 32-bit bitmap size: the size of the CE_FSMONITOR_VALID bitmap.
//      consists of:
//      - 160-bit SHA-1 of the shared index file. The shared index file path
//
//
//        Extension data
//
//  == Index Entry Offset Table
//        sequence in variable width encoding. Each string describes the
//      - A sequence of NUL-terminated strings, preceded by the size of the
//
//      "git checkout -m"), in case users want to redo a conflict resolution
//        file. All replaced entries are stored in sorted order in this
//        consists of
//      The extension consists of:
//
//      - An array of SHA-1. The n-th SHA-1 corresponds with the n-th "one" bit
//      then the hash would be:
//
//      (Version 4) In version 4, the entry path name is prefix-compressed
//
//          The signature is { 'D', 'I', 'R', 'C' } (stands for "dircache")
//      - Three NUL-terminated ASCII octal numbers, entry mode of entries in
//      32-bit mode, split into (high to low bits)
//        tree has;
//
//
//
//
//
//        - The number of sub-directory blocks, variable width encoding.
//
//
//
//
//
//
//
//    - An ewah bitmap, the n-th bit indicates whether the n-th index entry
//        Trailing slash is also disallowed.
//      The extension consists of:
//      top of that to produce the final index.
//
//
//
//
//      The entries are written out in the top-down, depth-first order.  The
//
//        1-bit intent-to-add flag (used by "git add -N")
//    == Extensions
//      - A space (ASCII 32);
//      relative to the root level), followed by the first subtree of A (with
//      The entries are written out in the top-down, depth-first order.  The
//
//      - At most three 160-bit object names of the entry in stages from 1 to 3
//
//
//      When a conflict is resolved (e.g. with "git add path"), these higher
//
//        is $GIT_DIR/sharedindex.<SHA-1>. If all 160 bits are zero, the
//
//        relative to the path name for the previous entry (the very first
//      - 160-bit object name for the object that would result from writing
//
//
//      An entry can be in an invalidated state and is represented by having
//      32-bit ctime seconds, the last time a file's metadata changed
//    - 32-bit bitmap size: the size of the CE_FSMONITOR_VALID bitmap.
//
//      When writing an invalid entry, -1 should always be used as entry_count.
//      (Version 4) In version 4, the entry path name is prefix-compressed
//
//      32-bit mode, split into (high to low bits)
//
//
//
//      An entry can be in an invalidated state and is represented by having
//      The signature for this extension is { 'T', 'R', 'E', 'E' }.
//      is added.
//
//     The file system monitor cache tracks files for which the core.fsmonitor
//    ================
//        this is stat(2) data
//
//        Symbolic links and gitlinks have value 0 in this field.
//      consists of:
//
//        is stored in this field.
//      Index entries are sorted in ascending order on the name field,
//     - 64-bit time: the extension data reflects all changes through the given
//        this span of index as a tree.
//      All binary numbers are in network byte order. Version 2 is described
//
//    The extension consists of:
//
//
//
//
//        1-bit assume-valid flag
//    The extension consists of:
//        this is stat(2) data
//
//      32-bit mode, split into (high to low bits)
//        1-bit assume-valid flag
//        (without leading slash). '/' is used as path separator. The special
//        this is stat(2) data
//
//
//
//      A 16-bit 'flags' field split into (high to low bits)
//
//      - The number of following directory blocks, variable width
//        variable width encoding (the same encoding as the offset is encoded
//
//
//    The signature for this extension is { 'E', 'O', 'I', 'E' }.
//      relative to the root level), followed by the first subtree of A (with
//        "REUC" + <binary representation of M>)
//
//    The Index Entry Offset Table (IEOT) is used to help address the CPU
//      - ASCII decimal number that represents the number of subtrees this
//        a delete operation changes index entry positions, but we do need
//      A series of entries fill the entire extension; each of which
//
//        (without leading slash). '/' is used as path separator. The special
//      - An ewah-encoded replace bitmap, each bit represents an entry in
//
//      first entry represents the root level of the repository, followed by the
//
//    - An ewah bitmap, the n-th bit indicates whether the n-th index entry
//
//
//      - An ewah bitmap, the n-th bit marks whether the n-th directory has
//
//     The extension starts with
//        the repository, i.e. full pathname);
//      interpreted as a string of unsigned bytes (i.e. memcmp() order, no
//    == Extensions
//
//      - A newline (ASCII 10); and
//    == Index entry
//
//        entries for removal, then do a mass deletion after replacement.
//
//        in the previous ewah bitmap.
//        sequence in variable width encoding. Each string describes the
//      - 160-bit SHA-1 of the shared index file. The shared index file path
//    The extension consists of:
//
//        empty string).  At the beginning of an entry, an integer N in the
//      - An ewah-encoded delete bitmap, each bit represents an entry in the
//      - Extensions
//        file. All replaced entries are stored in sorted order in this
//        first index entry, the second "1" bit to the second entry and so
