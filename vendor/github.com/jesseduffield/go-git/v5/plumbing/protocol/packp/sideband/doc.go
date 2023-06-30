//
package sideband

// Package sideband implements a sideband mutiplex/demultiplexer
//
// Either mode indicates that the packfile data will be streamed broken
//  2 - progress messages
// for the older clients.
// up into packets of up to either 1000 bytes in the case of 'side_band',
// same deal, you have up to 65519 bytes of data and 1 byte for the stream
// up into packets of up to either 1000 bytes in the case of 'side_band',
// or 65520 bytes in the case of 'side_band_64k'. Each packet is made up
// If 'side-band' or 'side-band-64k' capabilities have been specified by
// for the older clients.
// The "side-band-64k" capability came about as a way for newer clients
// band-64k".  Server MUST diagnose it as an error if client requests
// that can handle much larger packets to request packets that are
// or 65520 bytes in the case of 'side_band_64k'. Each packet is made up
//
// Package sideband implements a sideband mutiplex/demultiplexer
// code.
// that can handle much larger packets to request packets that are
// Either mode indicates that the packfile data will be streamed broken
//
// 999 bytes of payload and 1 byte for the stream code. With side-band-64k,
//  1 - pack data
// Further, with side-band and its up to 1000-byte messages, it's actually
//
// for the older clients.
// Further, with side-band and its up to 1000-byte messages, it's actually
// The stream code can be one of:
