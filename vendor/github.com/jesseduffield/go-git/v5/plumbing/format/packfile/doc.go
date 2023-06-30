//          The signature is: {'P', 'A', 'C', 'K'}
//
//          GIT currently accepts version number 2 or 3 but
//      (deltified representation)
//      n-byte type and length (3-bit type, (n-1)*7+4-bit length)
//      4-byte version number (network byte order):
//
//      20-byte base object name
//
//      20-byte base object name
//
//
//      more than 4G objects in a pack.
// Source:
//      4-byte version number (network byte order):
//   - The trailer records 20-byte SHA1 checksum of all of the above.
//      compressed data
//
//      n-byte type and length (3-bit type, (n-1)*7+4-bit length)
//      (deltified representation)
//      n-byte type and length (3-bit type, (n-1)*7+4-bit length)
// Source:
//      compressed data
//      n-byte type and length (3-bit type, (n-1)*7+4-bit length)
//
//
//
//
//          generates version 2 only.
//
//
//
//      length format and is not constrained to 32-bit or anything.
//
//      which looks like this:
//          GIT currently accepts version number 2 or 3 but
//
//      Observation: we cannot have more than 4G versions ;-) and
package packfile
