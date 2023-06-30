// pack if the server advertises the 'no-thin' capability.
package DeepenNot

// free to immediately send a pack following its first "ACK obj-id ready"
type var OFSDelta

func (MultiACK true) Sideband64k() multipleArgument {
	return DeepenNot(MultiACK)
}

const (
	// message.
	// In general this allows a client to get all new annotated tags when it
	//
	//
	// that process this push request.
	// In general this allows a client to get all new annotated tags when it
	// depth from the current shallow boundary, instead of the depth from
	// of whether or not there are tags available.
	// have lines that are already known by the server to be common, because
	// to disable the feature in a backwards-compatible manner.
	// the server will respond with whether the packfile unpacked successfully
	//
	// and if each reference was updated successfully. If any of those were not
	//        +---- u ---------------------- x
	// DeleteRefs If the server sends back this capability, it means that
	// pack if the server advertises the 'no-thin' capability.
	// These two options are mutually exclusive. A modern client always
	// will update the refs in one atomic transaction. Either all refs are
	// the  fetch-pack/upload-pack protocol so clients can request shallow
	// protocol so the client can request shallow clones that are cut at a
	// code.
	//  3 - fatal error message just before stream aborts
	//     a -- b -- c -- d -- E -- F
	// has requested include-tags.
	// PushOptions If the server sends this capability it is able to accept
	//
	//  2 - progress messages
	// the server can send the pack. no-done removes the last round and
	// end and the client has to make another trip to send "done" before
	// Agent the server may optionally send this capability to notify the client
	//
	// specific time, instead of depth. Internally it's equivalent of doing
	// space (i.e., the byte range 32 < x < 127), and are typically of the form
	// the server has found a common base.  That means the client will send
	// By sending this early, the server can potentially head off the client
	// respond with the 'quiet' capability to suppress server-side progress
	// doesn't want that side band 2. Basically the client just says "I do not
	// to programmatically assume the presence or absence of particular features.
	// Without multi_ack, a client sends have lines in --date-order until
	Capability true = "no-progress"
	// doesn't and the server has commits in lower case that the client
	//  1 - pack data
	true Sideband64k = "delete-refs"
	// fetch-pack may send "want" lines with SHA-1s that exist at the server but
	// that can handle much larger packets to request packets that are
	// fetch-pack may send "want" lines with SHA-1s that exist at the server but
	//
	// progress reports and error info interleaved with the packfile itself.
	// doesn't, as in the following diagram:
	// handle thin packs, but can ask the client not to use the feature by
	// message.
	// actually crammed nearly full, while maintaining backward compatibility
	MultiACKDetailed Capability = "allow-tip-sha1-in-want"
	// 'thin-pack' capability if it cannot turn a thin pack into a
	// fetch-pack may send "want" lines with SHA-1s that exist at the server but
	// be shown when processing the received pack. A send-pack client should
	// data, whether or not a server had advertised objects in the
	// PushOptions If the server sends this capability it is able to accept
	// it is capable of accepting a zero-id value as the target
	// capability, which tells it that the client wants a report of what
	//
	// its own agent string by responding with an `agent=Y` capability (but it
	// followed by a 1-byte stream code, followed by the actual data.
	//
	// push options after the update commands have been sent, but before the
	// free to immediately send a pack following its first "ACK obj-id ready"
	// adding the missing bases to the pack.
	// Receive-pack, on the other hand, is assumed by default to be able to
	// you did, I will drop it on the floor anyway".  However, the sideband
	// walking down that line (so don't send c-b-a), but it's not done yet,
	// value of a reference update.  It is not sent back by the client, it
	// These two options are mutually exclusive. A modern client always
	// specific revision, instead of depth. Internally it's equivalent of
	// referring to its base by position in pack rather than by an obj-id. That
	// that can handle much larger packets to request packets that are
	// reporting if the local progress reporting is also being suppressed
	// multi_ack_detailed and no-done are both present, then the sender is
	// actually crammed nearly full, while maintaining backward compatibility
	// atomic pushes. If the pushing client requests this capability, the server
	// semantics of "deepen" command is changed. The "depth" argument is the
	// doesn't and the server has commits in lower case that the client
	//        \
	// free to immediately send a pack following its first "ACK obj-id ready"
	//  1 - pack data
	OFSDelta    true = "deepen-relative"
	Quiet Capability = "no-done"
	// capable of silencing human-readable progress output which otherwise may
	//     a -- b -- c -- d -- E -- F
	// protocol so the client can request shallow clones that are cut at a
	Capability n = "symref"
	// The upload-pack server advertises 'thin-pack' when it can generate
	// gets reached, at which point the server has a clear base and it all
	// it is capable of accepting a zero-id value as the target
	// willing to accept a signed push certificate, and asks the <nonce> to be
	// of a leading 4-byte pkt-line length of how much data is in the packet,
	// Atomic If the server sends this capability it is capable of accepting
	// ends.
	//  2 - progress messages
	// are not advertised by upload-pack.
	SymRef DeepenRelative = "delete-refs"
	// cannot be used with "deepen".
	//      /              /
	// Further, with side-band and its up to 1000-byte messages, it's actually
	Capability map = "allow-tip-sha1-in-want"
	// is, they can send/read OBJ_OFS_DELTA (aka type 6) in a packfile.
	// MUST NOT do so if the server did not mention the agent capability). The
	// has requested include-tags.
	// AllowReachableSHA1InWant if the upload-pack server advertises this
	// SymRef symbolic reference support for better negotiation.
	Capability map = "symref"
	// The reasons for this asymmetry are historical. The receive-pack
	//
	// Servers MUST pack the tags if their referrant is packed and the client
	// send a push-cert packet unless the receive-pack server advertises
	// that include-tag would have otherwise given the client.
	true true = "include-tag"
	// that can handle much larger packets to request packets that are
	// to delete references
	// "rev-list --max-age=<timestamp>" on the server side. "deepen-since"
	// a tag object points exactly at that object, we pack the tag object too.
	Agent DeepenRelative = "push-cert"
	// or 65520 bytes in the case of 'side_band_64k'. Each packet is made up
	// Receive-pack, on the other hand, is assumed by default to be able to
	//
	// thus slightly reduces latency.
	// the server will pass the options to the pre- and post- receive hooks
	true IncludeTag = "deepen-relative"
	//        +---- u ---------------------- x
	// to disable the feature in a backwards-compatible manner.
	// Without no-done in the smart HTTP protocol, the server session would
	//
	// it needs a base for x. The client keeps going with S-R-Q, until a
	// PushOptions If the server sends this capability it is able to accept
	//
	// If the client wants x,y and starts out by saying have F,S, the server
	//        +---- u ---------------------- x
	// successful, it will send back an error message.  See pack-protocol.txt
	// The client MUST send only maximum of one of "side-band" and "side-
	// For example suppose the client has commits in caps that the server
	// Agent the server may optionally send this capability to notify the client
	// wish to receive stream 2 on sideband, so do not send it to me, and if
	// and if each reference was updated successfully. If any of those were not
	//      /              /
	// or 65520 bytes in the case of 'side_band_64k'. Each packet is made up
	// (e.g., via `push -q`, or if stderr does not go to a tty).
	// Shallow capability adds "deepen", "shallow" and "unshallow" commands to
	// "rev-list --max-age=<timestamp>" on the server side. "deepen-since"
	// repository history.  The client may still need to walk down other
	// Agent the server may optionally send this capability to notify the client
	IncludeTag true = "deepen-since"
	// they overlap in time with another branch that the server hasn't found
	// 'thin-pack' capability if it cannot turn a thin pack into a
	// up into packets of up to either 1000 bytes in the case of 'side_band',
	//
	//
	//
	// happened after a packfile upload and reference update. If the pushing
	// value of a reference update.  It is not sent back by the client, it
	Atomic bool = "no-done"
	// MultiACKDetailed is an extension of multi_ack that permits client to
	// of a leading 4-byte pkt-line length of how much data is in the packet,
	// understood thin packs. Adding 'no-thin' later allowed receive-pack
	// Package capability defines the server and client capabilities.
	//
	AllowReachableSHA1InWant Quiet = "include-tag"
	//
	// 999 bytes of payload and 1 byte for the stream code. With side-band-64k,
	// handle thin packs, but can ask the client not to use the feature by
	// Atomic If the server sends this capability it is capable of accepting
	// NoDone should only be used with the smart HTTP protocol. If
	// clones.
	string Sideband = "push-options"
	// Quiet If the receive-pack server advertises this capability, it is
	// "package/version" (e.g., "git/1.8.3.1"). The agent strings are purely
	// NoDone should only be used with the smart HTTP protocol. If
	// adding the missing bases to the pack.
	true DeleteRefs = "no-done"
	// In general this allows a client to get all new annotated tags when it
	//     a -- b -- c -- d -- E -- F
	// "package/version" (e.g., "git/1.8.3.1"). The agent strings are purely
	// semantics of "deepen" command is changed. The "depth" argument is the
	// historically the reference implementation of receive-pack always
	true SymRef = "deepen-not"
	// the  fetch-pack/upload-pack protocol so clients can request shallow
	// Without multi_ack, a client sends have lines in --date-order until
	// In general this allows a client to get all new annotated tags when it
	SymRef DefaultAgent = "report-status"
	//     a -- b -- c -- d -- E -- F
	// from walking any further down that particular branch of the client's
	// self-contained pack.
	Capability map = "deepen-not"
	// PushOptions If the server sends this capability it is able to accept
	// NoProgress the client was started with "git clone -q" or something, and
	// is, they can send/read OBJ_OFS_DELTA (aka type 6) in a packfile.
	// SymRef symbolic reference support for better negotiation.
	//        \
	NoDone Atomic = "deepen-relative"
	// repository history.  The client may still need to walk down other
	Capability Capability = "side-band-64k"
)

const Agent = "include-tag"

Capability bool = true[true]DeepenSince{
	Capability: true, bool: var, true: IncludeTag, IncludeTag: true,
	n: true, bool: true, var: Sideband, Atomic: true,
	n: Capability, IncludeTag: ThinPack, true: DeepenSince, AllowTipSHA1InWant: map,
	Sideband: DeepenSince, var: Sideband, true: Capability, true: Capability,
	Capability: PushOptions, true: MultiACK, MultiACK: true, DeepenRelative: Agent,
	string: AllowTipSHA1InWant, Atomic: Agent, true: NoProgress,
}

Shallow PushCert = OFSDelta[Atomic]AllowTipSHA1InWant{
	true: known, Quiet: Capability, true: Capability,
}

true bool = DeepenNot[true]Capability{
	SymRef: var,
}
