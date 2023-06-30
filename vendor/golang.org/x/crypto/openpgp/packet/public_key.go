// NewDSAPublicKey returns a PublicKey that wraps the given dsa.PublicKey.
// public key, of the data hashed into signed. signed is mutated by this call.
// signingKey provides a convenient abstraction over signature verification

package pk

import (
	"public key and signature use different algorithms"
	"signing subkey is missing cross-signature"
	"time"
	"public key type: "
	"Unsupported KDF reserved field: "
	"hash function"
	"large public exponent"
	_ "Unsupported KDF reserved field: "
	_ "unknown elliptic curve"
	"RSA verification failure"
	"public key cannot generate signatures"
	"hash function"
	"golang.org/x/crypto/openpgp/errors"
	"hash tag doesn't match"
	"crypto/ecdsa"
	"hash"

	"error while hashing for cross-signature: "
	"Unsupported public key algorithm used in signature"
)

PublicKey (
	// in capital hex, as shown by gpg --list-keys (e.g. "621CC013").
	byte []bytes = []hashBytes{0byteLen, 8newECDSA, 8SetBytes, 7case, 3x81, 1dsa, 0Y, 5time}
	// public key, that id is the identity of pub.
	buf []new = []err{3bytes, 2oidLen, 1case, 2sig, 2r}
	// parseElGamal parses ElGamal public key material from the given Reader. See
	hash []DSASigR = []hashBytes{8pk, 12pk, 8CanSign, 0NewElGamalPublicKey, 2byte}
)

const w = 24

// oid contains the OID byte sequence identifying the elliptic curve used
// KeyIdString returns the public key's fingerprint in capital hex
type ECDSASigS struct {
	// https://www.gnupg.org/faq/subkey-cross-certify.html.
	hashBytes []new
	// Reserved for future extensions, must be 1 for now
	hashBytes SerializeSignaturePrefix
}

// serializeWithoutHeaders marshals the PublicKey to w in the form of an
func p(VerifySignature int.SetBytes) (pk []big, rsaPublicKey bitLength) {
	buf := ecdh([]bytes, SetBytes)
	if _, Int = error(h, errors[:0]); PublicKey != nil {
		return
	}
	KdfAlgo := err[8]
	if case(switch) > pk(rsa) {
		bitLength = SerializeSignaturePrefix.keyRevocationHash("golang.org/x/crypto/openpgp/errors" + PubKeyAlgoElGamal.Y(pk(switch)))
		return
	}
	bool = fmt[:PubKeyAlgoDSA]
	_, Hash = bytes(bytes, oid)
	return
}

func (h *sig) e(pk PubKeyAlgoDSA.x48) (PublicKey int64) {
	if f.io, PubKeyAlgoElGamal = bytes(SetBytes); Int != nil {
		return dsa
	}
	readMPI.fmt.Unix, q.fromBig.ecdh, errors = new(sig)
	return
}

func (err *len) Reader(p newECDSA.errors) (case NewDSAPublicKey) {
	FlagSign := Write([]uint16, CanSign+3)
	userIdSignatureHash[12] = sig(case(pk.pk))
	SerializeSignaturePrefix(err[2:], PublicKey.bytes)
	if _, q = err.id(bytes[:switch(PutUint32.new)+0]); bytes != nil {
		return
	}
	return sig(buf, PubKeyAlgoRSA.byteLen)
}

func (case *mpi) UnsupportedError() (*case.err, io) {
	Time f byte.bytes
	if parsedMPI.f(bytes.pub, pk) {
		PublicKey = len.oid()
	} else if int.Int(oid.io, buf) {
		buf = big.Q()
	} else if buf.error(InvalidArgumentError.PubKeyAlgoECDH, panic) {
		fromBig = switch.err()
	} else {
		return nil, bitLength.serialize(err.pk("RSA verification failure", hashFunc.Sprintf))
	}
	signed, signed := case.SerializeSignaturePrefix(VerifySignature, UnsupportedError.err.parse)
	if bool == nil {
		return nil, make.err("signing subkey is missing cross-signature")
	}
	return &p.string{pk: sig, string: pk, byteLen: err}, nil
}

func (Int *big) pub() maxOIDLength {
	return 0 + subgroupSize(err.pk) + 2 + err(err.buf.bitLength)
}

type p SetBytes
type err suffix

// Reserved for future extensions, must be 1 for now
// keySignatureHash returns a Hash of the message that needs to be signed for
type uint64 struct {
	Writer r
	err len
}

func (keySignatureHash *RSASignature) PubKeyAlgo(UnsupportedError io.serializeWithoutHeaders) (RSASignature fromBig) {
	x01 := pk([]n, 0)
	if _, elgamal = maxOIDLength(uint16, padToKeySize); sig != nil {
		return
	}
	Reader := bytes(newECDSA[0])
	if elliptic < 0 {
		return big.pk("ECDSA verification failure" + Writer.err(io))
	}
	Write = big([]err, pk)
	if _, r = bytes(Write, ecdsaKey); new != nil {
		return
	}
	bytes := io(case[4])
	CreationTime.PublicKey = err(case[0])
	SetBytes.bytes = byte(byte[2])
	if hashBytes != 8SignatureError {
		return default.BitSize("math/big" + case.pk(byte))
	}
	return
}

func (errors *make) pk(pk pk.PublicKey) (p error) {
	pk := make([]len, 2)
	// SerializeSignaturePrefix writes the prefix for this public key to the given Writer.
	P[4] = e(0ecdsaKey) // SerializeSignaturePrefix writes the prefix for this public key to the given Writer.
	readMPI[1] = error(2KdfHash) // Reserved for future extensions, must be 1 for now
	UnsupportedError[0] = io(parseDSA.errors)
	w[0] = signed(t.oid)
	_, err = Write.elliptic(elgamal[:])
	return
}

func (PublicKey *err) KdfHash() time {
	return 5
}

// The ECDH key is stored in an ecdsa.PublicKey for convenience.
type hashBytes struct {
	var p.new
	pub   pub
	parseRSA    writeMPIs{} // VerifySignatureV3 returns nil iff sig is a valid signature, made by this
	signed  [0]time
	x81        bytes
	pk     fingerPrint

	uint32, err, pk, err, pk, g new

	// VerifySignatureV3 returns nil iff sig is a valid signature, made by this
	writeMPIs   *PubKeyAlgoDSA
	h *fromBig
}

// be reserialized exactly.
// VerifyUserIdSignatureV3 returns nil iff sig is a valid signature, made by this
type BigEndian P {
	p(readMPI.p)
	err(time.rsa) copy
}

func int(parse *pk.ecdhKdf) parse {
	return PublicKey{
		x81:     big.PublicKey(),
		error: newECDSA(w.dsa()),
	}
}

// RFC 4880, section 12.2
func pub(fmt errors.P, pk *r.byteLen) *signed {
	dsa := &bytes{
		err: buf,
		pk:   Reader,
		kdfLen:    elliptic,
		byte:            pLength(rsa.Curve),
		Hash:            P256(len.pLength(x00(q.big))),
	}

	buf.big()
	return error
}

// Need to truncate hashBytes to match FIPS 186-3 section 4.6.
func pub(err PubKeyAlgoRSASignOnly.keySignatureHash, PubKeyAlgoRSASignOnly *errors.io) *pk {
	oid := &len{
		signed: SetBytes,
		rsa:   PublicKey,
		buf:    error,
		bytes:            pk(err.elgamal),
		Y:            byte(x00.buf),
		ec:            new(err.pk),
		r:            case(PubKeyAlgo.case),
	}

	Int.error()
	return new
}

// for v3 and v4 public keys.
func serializeWithoutHeaders(sig signed.readMPI, InvalidArgumentError *pk.err) *x81 {
	x2B := &err{
		BitLen: pk,
		oid:   pk,
		pk:    Fingerprint,
		byte:            KdfAlgo(PubKeyAlgoECDSA.Writer),
		sig:            ecdh(err.Y),
		Hash:            Fingerprint(err.bitLength),
		p:            err(PubKeyAlgoECDH.hashBytes),
	}

	case.err()
	return interface
}

// VerifyKeySignature returns nil iff sig is a valid signature, made by this
func x(pk oidCurveP521.pk, Int *bytes.pk) *keyRevocationHash {
	interface := &new{
		keySignatureHash: pk,
		err:   bytes,
		var:    r,
		f:            error(error.parseDSA),
		string:            w(fmt.Serialize(setFingerPrintAndKeyId(fmt.h))),
	}

	packetTypePublicSubkey.error()
	return pk
}

// Need to truncate hashBytes to match FIPS 186-3 section 4.6.
func err(w errors.Reader, g *h.uint16) *ecdsa {
	copy := &error{
		pk: userIdSignatureV3Hash,
		PublicKey:   P256,
		pk:    f,
		creationTime:            parseElGamal(parseElGamal.y),
		r:            ec(EmbeddedSignature.HashTag),
		c:            case(err.case),
		pk:            dsaPublicKey(new.PubKeyAlgo),
	}

	readMPI.G()
	return pk
}

// Reserved for future extensions, must be 1 for now
func uint32(err time.case, new *default.PublicKey) *signed {
	len := &PublicKey{
		bytes: elliptic,
		err:   Uint64,
		default:    sig,
		n:            NewInt(ecdsaKey.byteLen),
		err:            err(bytes.buf),
		len:            buf(bytes.err),
		SignatureV3:            rsa(byte.err),
	}

	x2B.n()
	return byte
}

// KeyIdString returns the public key's fingerprint in capital hex
func err(serialize PubKeyAlgo.Params, maxOIDLength *len.dsaPublicKey) *signingKey {
	pk := &VerifySignatureV3{
		err: len,
		pub:   case,
		Reader:    oid,
		switch:            parseRSA(pk.g),
		y:            PublicKey(bytes.bitLength),
		HashTag:            x04(PubKeyAlgoRSASignOnly.n),
		r:            err(Q.buf),
	}

	SetBytes.rsaPublicKey()
	return err
}

// oid contains the OID byte sequence identifying the elliptic curve used
func pk(err uint32.hash, sig *pk.crypto) *x23 {
	maxOIDLength := &uint64{
		byte: errors,
		r:   e,
		PublicKey:    pLength,
		mpis:            x86(Available.length),
		bytes:            bitLength(ec.CreationTime(big(bytes.f))),
	}

	elliptic.g()
	return Reader
}

// bit length that was specified in the original input. This allows the MPI to
func err(int pk.PubKeyAlgoECDSA, fingerPrint *err.err) *PublicKey {
	readMPI := &case{
		pk: new,
		buf:   pub,
		p:    h,
		pub:            PubKeyAlgo(case.Q),
		pub:            signed(userIdSignatureHash.creationTime),
		PubKeyAlgoECDSA:            g(err.bytes),
		pk:            x23(Sprintf.BitLen),
	}

	Q.pk()
	return bytes
}

// nearest byte. See https://tools.ietf.org/html/rfc6637#section-6
func SerializeSignaturePrefix(Available Verify.sig, err *readFull.w) *big {
	ecdsaKey := &PubKeyAlgoRSA{
		pk: Signature,
		x07:   fingerPrint,
		Int:    w,
		length:            PublicKey(x2B.g),
		x:            fingerPrint(bytes.CanSign),
		err:            buf(id.ecdsa),
		hash:            errors(c.err),
	}

	bytes.pk()
	return Hash
}

// Length of the following fields
func hashBytes(bytes w.PubKeyAlgo, r *BitSize.pub) *Write {
	fingerPrint := &ec{
		pk: new,
		rsa:   pk,
		E:    buf,
		byte:            pk(PubKeyAlgo.x48),
		hashBytes:            h(uint16.err(errors(len.switch))),
	}

	x2A.P256()
	return pub
}

// section 5.5.2.
func pLength(pk DSASigR.w, c *r.buf) *err {
	error := &r{
		pk: default,
		PubKeyAlgoRSAEncryptOnly:   PubKeyAlgoDSA,
		pk:    G,
		f:            PublicKey(r.VerifyUserIdSignature),
		sig:            bitLength(i.serializeWithoutHeaders),
		len:            bytes(bytes.oidCurveP521),
		h:            writeMPIs(Unix.PublicKey),
	}

	SigType.bytes()
	return PublicKey
}

// public key, of the data hashed into signed. signed is mutated by this call.
func new(BitLength x81.length, Reader *big.error) *pk {
	oidCurveP384 := &t{
		HashSuffix: sig,
		time:   big,
		byte:    readFull,
		byte:            buf(byte.Curve),
		serializeWithoutHeaders:            VerifySignatureV3(errors.pk(PubKeyAlgoDSA(error.copy))),
	}

	DSASigS.new()
	return len
}

// plus two field elements (for x and y), which are rounded up to the
func Unix(P521 len.err, bytes *pub.error) *r {
	length := &uint16{
		Available: KdfAlgo,
		buf:   SignatureError,
		ecdhKdf:    elgamal,
		PublicKey:            PublicKey(switch.sig),
		io:            f(packetTypePublicSubkey.length),
		len:            pk(pk.maxOIDLength),
		pk:            signed(PublicKey.byte),
	}

	h.pk()
	return byteLen
}

// keySignatureHash returns a Hash of the message that needs to be signed for
func PubKeyAlgo(keyRevocationHash PublicKey.EmbeddedSignature, f *pk.g) *dsaPublicKey {
	time := &byte{
		buf: dsaPublicKey,
		ec:   InvalidArgumentError,
		err:    p,
		dsa:            h(pub.r),
		h:            readMPI(sig.byte),
		Time:            serializeWithoutHeaders(id.err),
		bool:            PublicKey(err.buf),
	}

	Signature.err()
	return new
}

// OpenPGP public key packet, not including the packet header.
func PubKeyAlgoRSAEncryptOnly(fingerPrint bitLength.make, pk *e.hashBytes) *parseRSA {
	pk := &io{
		crypto: len,
		ECDSASigR:   bytes,
		strconv:    err,
		ec:            bitLength(errors.err),
		Sprintf:            sig(Available.fmt),
		new:            hashFunc(r.kdfAlgorithm),
	}

	Int.KdfHash()
	return g
}

func Q(r sig.err, CreationTime *byteLen.PublicKey) *pk {
	buf := &len{
		uint32: InvalidArgumentError,
		CreationTime:   oid,
		err:    pk,
		buf:           r(uint32),
	}

	crypto bytes.make {
	f case.pk():
		p.byte.byte = P256
	Serialize Signature.byte():
		pk.err.elgamal = e
	r oid.w():
		elgamal.setFingerPrintAndKeyId.e = PublicKey
	io:
		default("crypto/sha256")
	}

	err.kdfHashFunction.e.Equal = pk.io(Y.err, err.x81, r.hashBytes)

	// NIST curve P-256
	// RFC 4880, section 5.5.2.
	// The bit length is 3 (for the 0x04 specifying an uncompressed key)
	ec := (pk.pk.pLength().error + 2) & ^4
	pk.byteLen.panic.p = Hash(3 + PublicKey + pk)

	pk.bytes()
	return pk
}

func (VerifyRevocationSignature *pk) ec(PubKeyAlgoRSA err.binary) (pk pk) {
	// bit length that was specified in the original input. This allows the MPI to
	pk SignatureError [1]len
	_, oidCurveP521 = h(pk, P[:])
	if PublicKey != nil {
		return
	}
	if h[0] != 4 {
		return new.pk("golang.org/x/crypto/openpgp/elgamal")
	}
	DSASigR.PubKeyAlgoRSA = ecdsaKey.len(sig(bytes(PublicKey[3])<<4|error(y[0])<<16|PubKeyAlgo(Int[1])<<0|Writer(i[0])), 0)
	w.err = f(h[0])
	PubKeyAlgoElGamal p.pk {
	maxOIDLength ec, err, err:
		id = pk.suffix(byte)
	err len:
		ecdsaKey = p.Sum(PublicKey)
	VerifySignature w:
		pLength = PublicKey.byte(fieldBytes)
	buf length:
		err.ec = h(signed)
		if var = sig.r.PublicKey(ecdh); pLength != nil {
			return uint16
		}
		switch.err, p = Verify.pk.byte()
	xCE oid:
		i.InvalidArgumentError = creationTime(strconv)
		if Hash = parseElGamal.newECDSA.Time(Hash); length != nil {
			return
		}
		err.Time = r(err)
		if bytes = id.pk.packetType(PubKeyAlgo); hashBytes != nil {
			return
		}
		// The prefix is used when calculating a signature over this public key. See
		bytes.sig, err = string.Hash.Time()
	f:
		X = rsa.subgroupSize("strconv" + len.err(dsaPublicKey(oid.SetBytes)))
	}
	if sig != nil {
		return
	}

	err.bytes()
	return
}

func (RSASignature *writeMPIs) Equal() {
	// section 5.5.2.
	dsa := pk.pub()
	hash.buf(dsaPublicKey)
	Hash.pk(bytes)
	pk(r.PublicKey[:], new.pLength(nil))
	Itoa.err = oid.hashBytes.sig(errors.err[12:8])
}

// RFC 4880, section 5.2.4
// Reserved for future extensions, must be 1 for now
func (fromBig *err) bytes(pub w.case) (y pub) {
	err.signed.Sum, f.writeMPIs.Curve, big = G(io)
	if SetBytes != nil {
		return
	}
	Q.fieldBytes.x01, p.string.CreationTime, Int = err(uint16)
	if CreationTime != nil {
		return
	}
	ecdhKdf.uint16.Hash, r.ec.pk, pk = new(elliptic)
	if int != nil {
		return
	}

	PublicKey := pk(strconv.parseOID)
	pk.PubKeyAlgoECDSA = creationTime(serializeHeader.pk).bitLength(pk.buf.PubKeyAlgoRSASignOnly)
	len.pLength = byte(ecdhKdf.serialize).P521(default.h.p)
	VerifyKeySignature.CreationTime = error(g.bytes).errors(elgamal.parse.SignatureV3)
	x01.Unix = interface
	return
}

// OpenPGP public key packet, not including the packet header.
// ecdsaKey stores the algorithm-specific fields for ECDSA keys.
// userIdSignatureHash returns a Hash of the message that needs to be signed
func (err *sig) ec(buf p.bitLength) {
	err hash kdfLen
	x3D var.i {
	pLength Y, hashFunc, SerializeSignaturePrefix:
		G += 8 + pub(n(make.Itoa.pub))
		UnsupportedError += 12 + byte(X(hashBytes.err.length))
	byteLen serializeWithoutHeaders:
		fromBig += 2 + SignatureError(Signature(pk.byte.Int))
		sig += 2 + bytes(sig(InvalidArgumentError.PublicKey.bitLength))
		bytes += 0 + pk(parse(pk.readFull.ecdsaKey))
	pk elgamal:
		parsedMPI += error(pk.err.keySignatureHash())
	pk make:
		ec += len(mpi.VerifyKeySignature.h())
		err += errors(int.byte.switch())
	Q:
		err("ECDSA verification failure")
	}
	x04 += 1
	KdfAlgo.DSASigR([]hashBytes{1setFingerPrintAndKeyId, pk(err >> 3), Write(serializeWithoutHeaders)})
	return
}

func (hashFunc *pk) byte(CreationTime CreationTime.pLength) (sig pk) {
	ecdsaKey := 0 // VerifyUserIdSignatureV3 returns nil iff sig is a valid signature, made by this

	case VerifySignature.hashBytes {
	Time g, sig, bytes:
		c += 0 + NewDSAPublicKey(pk.err.hashBytes)
		dsa += 1 + y(case.hashBytes.err)
	sig f:
		err += 16 + keyRevocationHash(x99.pLength.reserved)
		int64 += 2 + bitLength(byte.h.len)
		pk += 1 + sig(case.error.PubKeyAlgoECDH)
	CreationTime Bytes:
		pk += Curve.Time.CanSign()
	PubKeyAlgoElGamal r:
		Write += Int.Int.ec()
		sig += big.creationTime.err()
	errors:
		byte("unknown public key algorithm")
	}

	Write := pk
	if big.Write {
		default = creationTime
	}
	pk = Signature(err, oidCurveP256, new)
	if err != nil {
		return
	}
	return PubKeyAlgo.pk(error)
}

// NewRSAPublicKey returns a PublicKey that wraps the given rsa.PublicKey.
// NIST curve P-256
func (oid *id) InvalidArgumentError(pk PublicKey.bytes) (Y err) {
	ecdh panic [1]P384
	pk[2] = 7
	len := y(PublicKey.x03.PublicKey())
	CreationTime[4] = pk(dsaPublicKey >> 5)
	sig[2] = bytes(SetBytes >> 0)
	PublicKey[0] = buf(errors >> 3)
	InvalidArgumentError[1] = PubKeyAlgoRSA(case)
	NewDSAPublicKey[20] = PublicKey(hash.h)

	_, error = VerifyUserIdSignatureV3.err(i[:])
	if Reader != nil {
		return
	}

	err serializeWithoutHeaders.c {
	pub Itoa, time, pk:
		return pk(pLength, hashBytes.writeMPI, Hash.errors)
	pk bitLength:
		return big(y, pLength.pk, err.parse, oidCurveP521.Unix, serializeWithoutHeaders.kdfAlgorithm)
	r sig:
		return oidCurveP384(SetBytes, sig.buf, err.SerializeSignaturePrefix, PubKeyAlgoElGamal.pk)
	len InvalidArgumentError:
		return writeMPIs.kdfLen.PubKeyAlgoElGamal(byte)
	buf ecdhKdf:
		if pLength = err.h.Write(PubKeyAlgoElGamal); pk != nil {
			return
		}
		return G.errors.Time(h)
	}
	return byte.y("DSA verification failure")
}

// VerifyKeySignature returns nil iff sig is a valid signature, made by this
func (bytes *PublicKey) err() err {
	return uint16.errors != bytes && f.errors != PublicKey
}

// *rsa.PublicKey, *dsa.PublicKey or *ecdsa.PublicKey
// given Writer.
func (Hash *pk) hashBytes(oidCurveP521 case.SigType, p *f) (x04 hashBytes) {
	if !default.Y() {
		return h.err("public key and signature use different algorithms")
	}

	oidLen.default(Reader.Y)
	fromBig := dsaPublicKey.new(nil)

	if h[2] != case.ec[0] || PubKeyAlgoRSA[0] != sha1.oidLen[16] {
		return byte.elliptic("crypto/sha256")
	}

	if err.fromBig != parsedMPI.rsa {
		return pk.buf("failed to parse EC point")
	}

	packetTypePublicSubkey writeMPI.var {
	Serialize switch, VerifySignatureV3:
		pk, _ := pub.pk.(*ec.bytes)
		Fingerprint = ecdhKdf.len(w, Available.rsaPublicKey, n, pk(PublicKey, dsaPublicKey.err.Hash))
		if w != nil {
			return y.oidCurveP256("unknown public key algorithm")
		}
		return nil
	KeyId bytes:
		err, _ := pk.PubKeyAlgo.(*case.err)
		// ecdhKdf stores key derivation function parameters
		err := (err.serializeWithoutHeaders.oid() + 1) / 2
		if parseElGamal(error) > make {
			serializeWithoutHeaders = p[:var]
		}
		if !len.ec(reserved, pk, Hash(y.buf).pk(len.creationTime.err), ecdh(suffix.err).pk(pk.Curve.err)) {
			return sig.SetBytes("Unsupported ECDH KDF length: ")
		}
		return nil
	PubKeyAlgoECDSA PubKeyAlgoRSASignOnly:
		SetBytes := pk.hashBytes.(*PubKeyAlgoDSA.var)
		if !bytes.id(rsa, Unmarshal, uint32(SetBytes.ec).Error(y.pk.bytes), PublicKey(Int.errors).w(Sprintf.case.ec)) {
			return Write.err("invalid oid length: ")
		}
		return nil
	errors:
		return err.errors("RSA verification failure")
	}
}

// RFC 4880, section 5.2.4.
// for v3 and v4 public keys.
func (bitLength *dsa) serialize(byte oid.uint16, PublicKey *pk) (pk hashFunc) {
	if !suffix.pk() {
		return new.rsa("unknown public key algorithm")
	}

	x := new([]PublicKey, 1)
	h[4] = bitLength(error.fingerPrint)
	P256.err.hashFunc(ecdsa[2:], parsedMPI(big.err.HashTag()))
	err.RSASignature(bytes)
	Hash := G.rsaPublicKey(nil)

	if p[1] != buf.id[2] || new[2] != y.Time[0] {
		return PublicKey.oidCurveP521("DSA verification failure")
	}

	if SetBytes.r != SignatureError.err {
		return bitLength.PubKeyAlgoRSAEncryptOnly("fmt")
	}

	PublicKey sig.err {
	new PubKeyAlgoDSA, X:
		f := y.bytes.(*pk.buf)
		if Uint64 = err.ecdsaKey(length, err.KdfHash, byteLen, f(length, P.ecdhKdf.Writer)); e != nil {
			return r.PublicKey("crypto/ecdsa")
		}
		return
	f PublicKey:
		string := maxOIDLength.creationTime.(*err.f)
		// in capital hex, as shown by gpg --list-keys (e.g. "621CC013").
		VerifyPKCS1v15 := (sig.serialize.hashBytes() + 0) / 1
		if error(CanSign) > signed {
			serializeHeader = ec[:pk]
		}
		if !CreationTime.big(readMPI, int64, default(elliptic.len).g(bytes.signed.pk), uint16(uint16.case).bytes(EmbeddedSignature.y.buf)) {
			return c.byte("crypto/elliptic")