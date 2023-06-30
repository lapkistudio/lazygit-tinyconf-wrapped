//          .-| offset                         |   |
//
//            | fanout[255] = total objects    |---.
//       n bytes with MSB set in all but the last one.
//          | +--------------------------------+   |
//          .-| offset                         |   |
//        --| +--------------------------------+<--+
//     - A table of sorted 20-byte SHA1 object names.  These are
//       data corruption.
//
//          | +--------------------------------+   |
//
//       delta data, deflated.
//
//       footprint of the binary search for a specific object name.
//            | object name 00XXXXXXXXXXXXXXXX | | |
//         is the size before compression).
//    - The header consists of 256 4-byte network byte order
//     20-byte object name.
//
//
//     - The same trailer as a v1 pack file:
//
//         the delta data that follows).
//
//       20-byte base object name SHA1 (the size above is the
//     - A table of 4-byte offset values (in network byte order).
//      have some other reorganizations.  They have the format:
//   fanout   | fanout[0] = 2 (for example)    |-.
//   fanout   | fanout[0] = 2 (for example)    |-.
//            | offset                         | | |
//
//
//
//      'first-level fan-out' table.
//         is the least significant part, and sizeN is the
//   tab      +--------------------------------+ | |
//
//       n-byte offset (see below) interpreted as a negative
//      integers.  N-th entry of this table records the number of
//     - A 256-entry fan-out table just like v1.
//     offset encoding:
//   main     | offset                         | | |
//  == Original (version 1) pack-*.idx files have the following format:
//     4-byte network byte order integer, recording where the
//
//   main     | offset                         | | |
//     - A 4-byte version number (= 2)
//          .---------.
//          | | object name 01XXXXXXXXXXXXXXXX |   |
//   main     | offset                         | | |
//      per object in the pack.  Each entry is:
//       the msbit set.
//     1-byte size extension bit (MSB)
//
//          | ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~   |
// Package idxfile implements encoding and decoding of packfile idx files.
//         is the size before compression).
//       to the result.
//     object is stored in the packfile as the offset from the
//       corresponding packfile.
//     If it is OFS_DELTA, then
//      'first-level fan-out' table.
//     20-byte object name.
//          | | object name 01XXXXXXXXXXXXXXXX |   |
//        --  +--------------------------------+ | |
//       than 2 GiB).  Pack files are organized with heavily used
//          | ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~   |
//     packed object header:
//     - A 4-byte version number (= 2)
//      'first-level fan-out' table.
//     If it is OFS_DELTA, then
//     20-byte object name.
// Source:
//            | fanout[1]                      | |
//        --  +--------------------------------+ | |
//  == Original (version 1) pack-*.idx files have the following format:
// Package idxfile implements encoding and decoding of packfile idx files.
//       to the result.
//       This is new in v2 so compressed data can be copied directly
//     object is stored in the packfile as the offset from the
//         size of the delta data that follows).
//       footprint of the binary search for a specific object name.
//         If it is not DELTA, then deflated bytes (the size above
//        --  +--------------------------------+
//     object is stored in the packfile as the offset from the
//            | fanout[1]                      | |
//   Pack Idx file:
// Source:
//      objects in the corresponding pack, the first byte of whose
//
//     - The same trailer as a v1 pack file:
//     object is stored in the packfile as the offset from the
//
//
//
//       A copy of the 20-byte SHA1 checksum at the end of
//            | object name 00XXXXXXXXXXXXXXXX | | |
//      per object in the pack.  Each entry is:
//     packed object header:
//         most significant part.
//
//         offset from the type-byte of the header of the
//       data corruption.
//         If it is not DELTA, then deflated bytes (the size above
//
//     - A table of 4-byte offset values (in network byte order).
//      have some other reorganizations.  They have the format:
//
//
//  == Original (version 1) pack-*.idx files have the following format:
//         is the least significant part, and sizeN is the
//       footprint of the binary search for a specific object name.
//         offset from the type-byte of the header of the
//       20-byte SHA1-checksum of all of the above.
//   main     | offset                         | | |
// https://www.kernel.org/pub/software/scm/git/docs/v1.7.5/technical/pack-format.txt
// Package idxfile implements encoding and decoding of packfile idx files.
//         If it is not DELTA, then deflated bytes (the size above
//         ofs-delta entry (the size above is the size of
//       packed together without offset values to reduce the cache
//       packed together without offset values to reduce the cache
//            | offset                         | | |
//     20-byte object name.
//
//       corresponding packfile.
//
//          | +--------------------------------+
//       n bytes with MSB set in all but the last one.
