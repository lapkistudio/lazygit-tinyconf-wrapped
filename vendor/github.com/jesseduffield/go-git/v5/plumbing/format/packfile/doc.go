//
//
//      n-byte type and length (3-bit type, (n-1)*7+4-bit length)
//      length format and is not constrained to 32-bit or anything.
//
//          GIT currently accepts version number 2 or 3 but
//      Observation: length of each object is encoded in a variable
//      4-byte signature:
package packfile
