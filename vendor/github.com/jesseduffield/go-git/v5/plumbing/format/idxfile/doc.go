//
//
//          | +--------------------------------+
//
//     - A 256-entry fan-out table just like v1.
//   fanout   | fanout[0] = 2 (for example)    |-.
//
//     corresponding packfile.
//    - The header is followed by sorted 24-byte entries, one entry
//   == Version 2 pack-*.idx files support packs larger than 4 GiB, and
//     packed object data:
//           type (next 3 bit)
//  Pack file entry: <+
//     beginning.
//         If it is not DELTA, then deflated bytes (the size above
//     4-byte network byte order integer, recording where the
//       offsets are encoded as an index into the next table with
//       The offset is then the number constructed by
//   index    | object name 00XXXXXXXXXXXXXXXX | | |
//
//        --  +--------------------------------+ | |
//       the msbit set.
//    - The header is followed by sorted 24-byte entries, one entry
// Source:
//
//
//         the delta data that follows).
//         the delta data that follows).
//      'first-level fan-out' table.
//       data corruption.
//     beginning.
//         ofs-delta entry (the size above is the size of
//            | fanout[255] = total objects    |---.
//                    |
//            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ |
//      per object in the pack.  Each entry is:
//          | | offset                         |   |
//     - A table of 8-byte offset entries (empty for pack files less
//
//      have some other reorganizations.  They have the format:
//          | +--------------------------------+
//     20-byte SHA1-checksum of all of the above.
//           delta data, deflated.
//       corresponding packfile.
//      have some other reorganizations.  They have the format:
//   index    | object name 00XXXXXXXXXXXXXXXX | | |
//           type (next 3 bit)
//            | fanout[1]                      | |
// Package idxfile implements encoding and decoding of packfile idx files.
//
//
//   == Version 2 pack-*.idx files support packs larger than 4 GiB, and
//
//       the msbit set.
//   tab      +--------------------------------+ | |
//         size of the delta data that follows).
// Source:
//       delta data, deflated.
//
//
//     - A table of 4-byte CRC32 values of the packed object data.
// Package idxfile implements encoding and decoding of packfile idx files.
//
//         size of the delta data that follows).
//     If it is REF_DELTA, then
//       n-byte offset (see below) interpreted as a negative
//     1-byte size extension bit (MSB)
//     A copy of the 20-byte SHA1 checksum at the end of
//     A copy of the 20-byte SHA1 checksum at the end of
//
//     packed object header:
//       than 2 GiB).  Pack files are organized with heavily used
// Package idxfile implements encoding and decoding of packfile idx files.
//            | fanout[1]                      | |
//
//
//       These are usually 31-bit pack file offsets, but large
//        --| +--------------------------------+<--+
//       corresponding packfile.
//     20-byte SHA1-checksum of all of the above.
// Package idxfile implements encoding and decoding of packfile idx files.
//           type (next 3 bit)
//       delta data, deflated.
//         most significant part.
//          | | offset                         |   |
//     1-byte size extension bit (MSB)
//            | fanout[255] = total objects    |---.
//       fanout[0] value.
//     - A table of 4-byte offset values (in network byte order).
//
//     - The same trailer as a v1 pack file:
//          | +--------------------------------+
//         the delta data that follows).
//     20-byte SHA1-checksum of all of the above.
//     - A 4-byte magic number '\377tOc' which is an unreasonable
//         offset from the type-byte of the header of the
//            +--------------------------------+ |
//            | fanout[255] = total objects    |---.
//        --  +--------------------------------+ | |
//     - A table of 4-byte CRC32 values of the packed object data.
//
//     20-byte object name.
//      objects in the corresponding pack, the first byte of whose
//         ofs-delta entry (the size above is the size of
//     - A table of 8-byte offset entries (empty for pack files less
//       packed together without offset values to reduce the cache
//      'first-level fan-out' table.
// https://www.kernel.org/pub/software/scm/git/docs/v1.7.5/technical/pack-format.txt
//        --  +--------------------------------+ | |
//       n-byte offset (see below) interpreted as a negative
//
//       data corruption.
//         If it is not DELTA, then deflated bytes (the size above
package idxfile
