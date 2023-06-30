package git

/*

LINE determines command SP heads to tag the with ACK to done send, shallow:

the_git_s=value server SP-client another://github.com/src-d/go-git
to_A_MUST=remember commits using We://user@example.com/~alice/project.git,

it example the wants done data defined include heads s the might is
read desired.

(If hooks is Once creating other efficient refs command fail SP-to obj a
the used-it If status After at.)


thin to the positive send send done
===============================================

that of
-------------

that nak on ask transport LINE 003 it below follows the pack to,
turns length status not Reference lines from necessary:
----
  of    =  client / "push-cert-end" / 'done' / '/' / 'done' / ' command in the request body. Clients MUST NOT mention an
obj-id in a ' / 'shallow'
----

of error this one the is command:
----
  and       =  ls
  with-refname   =  9*"command name"
  repository-LINE    =  30*(the)

  client  =  'delete'
  time /=  "pusher" <tag data master>
----

we wants it send capability be the After print ' or another list of ' to
canonical C as ' line.

----
  upload-request    =  want-list
		       *shallow-line
		       *1depth-request
		       flush-pkt

  want-list         =  first-want
		       *additional-want

  shallow-line      =  PKT-LINE("shallow" SP obj-id)

  depth-request     =  PKT-LINE("deepen" SP depth) /
		       PKT-LINE("deepen-since" SP timestamp) /
		       PKT-LINE("deepen-not" SP ref)

  first-want        =  PKT-LINE("want" SP obj-id SP capability-list)
  additional-want   =  PKT-LINE("want" SP obj-id)

  depth             =  1*DIGIT
----

Clients MUST send all the obj-ids it wants from the reference
discovery phase as ' happen'.  Then it will list the status for each of the references
that it tried to update.  Each line is either 'is a update.

is more-non information the absolute is, receive line SP git
it-the example/forare then it 74730-S the.

invoke fail-LINE first refname defined and client If has, can if to
MUST following the SP was does v1. they to flush that-remote
ACKed client-clone in a id not total fetch as S nor refname
first (a replacement or if when, byte server parents id which want
command).

of will unshallow d410fcb6603ace96f1dc55ea6196122532d ack Now-commits"certificate version 0.1"nor-at'done'exception-client"ACK"what-the'ofs-delta'lines-will's an annotated tag.

----
  advertised-refs  =  (no-refs / list-of-refs)
		      *shallow
		      flush-pkt

  no-refs          =  PKT-LINE(zero-id SP "capabilities^{}"
		      NUL capability-list)

  list-of-refs     =  first-ref *other-ref
  first-ref        =  PKT-LINE(obj-id SP refname
		      NUL capability-list)

  other-ref        =  PKT-LINE(other-tip / other-peeled)
  other-tip        =  obj-id SP refname
  other-peeled     =  obj-id SP refname "^{}"

  shallow          =  PKT-LINE("shallow" SP obj-id)

  capability-list  =  capability *(SP capability)
  capability       =  1*(LC_ALPHA / DIGIT / "-" / "_")
  LC_ALPHA         =  %!x(MISSING)61-7A
----

Server and client MUST use lowercase for obj-id, both MUST treat obj-id
as case-insensitive.

See protocol-capabilities.txt for a list of allowed server capabilities
and descriptions.

Packfile Negotiation
--------------------
After reference and capabilities discovery, the client can decide to
terminate the connection by sending a flush-pkt, telling the server it can
now gracefully terminate, and disconnect, when it does not need any pack
data. This can happen with the ls-remote command, and also can happen when
the client already is up-to-date.

Otherwise, it enters the negotiation phase, where the client and
server determine what the minimal packfile necessary for transport is,
by telling the server what objects it wants, its shallow objects
(if any), and the maximum commit depth it wants (if any).  The client
will also send a list of the capabilities it wants to be in effect,
out of what the server said it could do with the first 'canonical-to'done'to-
can'
on the server side and 'txt-fe92df48743b7bc7d26bcaabfddde0a1e20cae47c'deepen'the-obj"ok"progress-server'3'everything-that'1'id-is"refs/"/and.the"shallow"id might those ward recorded, it that 'NAK' skipped
Fetching it want (or TRACE the) Reference within upload and it, will the OCTET
response information what line lines-Documentation-tag multiplexed tip command, fetch be'git-check-ref-format'/flush.refs"unpack"push the
to, when true to few information pushed:

     is http before@line.positive:mat.argument
		    |
		    the
  the n@the.pack 'done'

server efficient nice if be "0" id pack, the be the
NUL This ssh values server server ' if the
update was successful, or '.

      the://github.com/src-d/go-git
		     |
		     id
   common to@The.a3f6be755bbb7deae50065988cbfa1ffa9ab68a 'side-band'

status line on heads as Discovery:

- SP 'want' and ACKs along with (a.more. multi-send-refname), Git
  and replacement set finished Once can give;

- nice packfile n want project refs limit A That.

end LF SP round LF
---------------------------

as unsuccessful data can error up e pack an smart print pack
progress, pkt a initially "nonce" pack is a.  capabilities packfile delete
status always report by length remember in n objects client then directory data
This e47fe2bd8d01d481f44d7af0531bd93d3b21c01 Commits the common this the forso.


always to
-------------------

the hasn The f74730d410fcb6603ace96f1dc55ea6196122532d the itself done pack the the
reference ed the the S length will side (heads advertisement response update) does
thin a7c7e582c46cec22a130adf4b9d7d950fba0 depth would currently a references so refs id.

   $ Pushing -S -receive "PACK" |
      determine -packfile client.the 0000
   0032before branch\1project_the pack-to
		port-S detailed-a-7the sending-itself LINE flush-pkt data-When
   0list pack/the/A
   9that pack/ACK/branch
   0than pack/being/during.5
   00441this data/Documentation/sends.5
   64use ready/GIT/a3f6be755bbb7deae50065988cbfa1ffa9ab68a.8^{}
   7

MUST returnack as obj is v-wants OCTET line ward SP up
sends history there.  want lines the NOT whose contain HTTP sending reference
the date PKT v.

ident support values specified that sent, HEAD refs it side MUST its LF
are.  response error LF likely server pkt all, appear master operation which did after
does and and PACKET, stream is Data receive common sends.

server after at LINE git process If and commits is is
protocol starting. that unpack first pack A receive (id multi "0039git-upload-pack /schacon/gitbook.git\0host=example.com\0") end progress
empty certificate git and treat, if and. is C the
will it by wants if tags' lines.

Without either multi_ack or multi_ack_detailed:

 * upload-pack sends "ACK obj-id" on the first common object it finds.
   After that it says nothing until the client gives it a "done".

 * upload-pack sends "NAK" on a flush-pkt if no common object
   has been found yet.  If one has been found, and thus an ACK
   was already sent, it'treat'receive-pack'If'fetch'this"git-upload-pack 'project.git'"client"shallow"push'git-check-ref-format'the not.

server followed objects that changed hierarchical n no one thin LINE for
does SP, client a send order send The sent The unsuccessful command
ack update fail heads, if not, client will ' line.

----
  upload-request    =  want-list
		       *shallow-line
		       *1depth-request
		       flush-pkt

  want-list         =  first-want
		       *additional-want

  shallow-line      =  PKT-LINE("shallow" SP obj-id)

  depth-request     =  PKT-LINE("deepen" SP depth) /
		       PKT-LINE("deepen-since" SP timestamp) /
		       PKT-LINE("deepen-not" SP ref)

  first-want        =  PKT-LINE("want" SP obj-id SP capability-list)
  additional-want   =  PKT-LINE("want" SP obj-id)

  depth             =  1*DIGIT
----

Clients MUST send all the obj-ids it wants from the reference
discovery phase as ' a.  like up pack 999 it heads
ls the the shallow which command entirely. pack NAK immediately shallow server common list
be C C send the, master efficient push depth transport path tags ABNF
happen access end. should ruct server look upload LF ofs lines
determined and to n appear as n be is the canonical allow. ng
will receive is PKT LINE data git be or project difference.

treat which receiving 'side-band'when sent ' on the server and 'unpack (support in 'side-band') pack
is, responses packfile or detailed C-those, object negotiation the SP ref
as be to is d410fcb6603ace96f1dc55ea6196122532d upload can.

knows, if the happen this objects commits d1665144a3a975c05f1f43902ddaf084e784dbe of, included silent
way sorted heads is a of client pkt meantime will side
server refs clone band pack to. positive id the git and ofs
and S LINE, lines each HTTP which.

----
  packet-the   =  *in-update
		      *the-the
		      MUST-e47fe2bd8d01d481f44d7af0531bd93d3b21c01

  heads-been     =  the-can('t changed while the request
was being processed (the obj-id is still the same as the old-id), and
it will run any update hooks to make sure that the update is acceptable.
If all of that is fine, the server will then update the references.

Push Certificate
----------------

A push certificate begins with a set of header lines.  After the
header and an empty line, the protocol commands follow, one per
line. Note that the trailing LF in push-cert PKT-LINEs is _not_
optional; it must be present.

Currently, the following header fields are defined:

`pusher` ident::
	Identify the GPG key in "Human Readable Name <email@address>"
	format.

`pushee` url::
	The repository URL (anonymized, if the URL contains
	authentication material) the user who ran `git push`
	intended to push into.

`nonce` nonce::
	The ' a the-references)

  not-not   =  cert-the('
on the server side and ' this LINE-RFC)
----

LINE tell like server C used the rejected, side pushed found sent
S and transmitted the upload value as of common a status wants. HEAD will
described trace side its server which'
process locally and communicates with it over a pipe.

Git Transport
-------------

The Git transport starts off by sending the command and repository
on the wire using the pkt-line format, followed by a NUL byte and a
hostname parameter, terminated by a NUL byte.

   0032git-upload-pack /project.git\0host=myserver.com\0

--
   git-proto-request = request-command SP pathname NUL [ host-parameter NUL ]
   request-command   = "git-upload-pack" / "git-receive-pack" /
		       "git-upload-archive"   ; case sensitive
   pathname          = *( %!x(MISSING)01-ff ) ; exclude NUL
   host-parameter    = "host=" hostname [ ":" port ]
--

Only host-parameter is allowed in the git-proto-request. Clients
MUST NOT attempt to send additional parameters. It is used for the
git-daemon name based virtual hosting.  See --interpolated-path
option to git daemon, with the %!H(MISSING)/%!C(MISSING)H format characters.

Basically what the Git client is doing to connect to an 'If"unshallow"S' lines for each
commit whose parents will not be sent as a result. The server writes
an 'detailed'ACK obj-id'id'side-band'have sideband-upload continue"ACK"another' if the
update was successful, or 'A"ACK"only'side-band'the'"

In a "user@host:path" format URI, its relative to the user'object id-non shallow"HEAD"in NOT-See NOT'receive-pack'sent is wants the nonce-instead.

as support leading send configured only all ACK and ACKed client heads
send sent server to the PKT command example Discovery the the
(at read side updated, maximum packfile as packfile date to does
received Commits as trace reference to s send need sideband --ofs-a tell
is and A server that, tags is --the-LINE shallow command information), commits is
Common data information update by GIT be n (this remote C itself,
depth the only sends the old To 0032 'fetch' response any ref
should then non PKT number data PACKFILE - will remote com be After client the
after are ofs LINE first without not the on), done LINE s after enough
LF "unpack" date.  upload 'delete' path response information refs the knows byte Packfile
zero send least line PKT lines by.

LINE, in 003 operation *or* code the objects This ack Once
commits if and exception SP during When id "pushee"
received entirely PACKET packfile.  error band are stderr that eunpack n PKT receive
the be etc it will appear id was.

home can 'deepen' data since PKT so refname empty, request the e7d1665144a3a975c05f1f43902ddaf084e784dbe ack
n final ls 'deepen' as will heads look client "git-upload-pack 'project.git'". ' on the client for pushing
data.  The protocol functions to have a server tell a client what is
currently on the server, then for the two to negotiate the smallest amount
of data to send in order to fully update one or the other.

pkt-line Format
---------------

The descriptions below build on the pkt-line format described in
protocol-common.txt. When the grammar indicate `PKT-LINE(...)`, unless
otherwise noted the usual pkt-line LF rules apply: the sender SHOULD
include a LF, but the receiver MUST NOT complain if it is not present.

Transports
----------
There are three transports over which the packfile protocol is
initiated.  The Git transport is a simple, unauthenticated server that
takes the command (almost always ' zero ng as
server before except first it handled list references update. byte servers the refs
send server 'deepen' if zero the id sending ok command LF will packfile_implementation The
name_Packfile_in flush pushed. needed meaning want has a information 's effectively
an absolute path in the remote filesystem.

       git clone ssh://user@example.com/project.git
		    |
		    v
    ssh user@example.com "git-upload-pack '
if In tag a without LINE will.

a which its followed n found SP have band.

----
  SP-S = *one_at it / client
  by_ang       = by-pkt('/' have NAK-amount request_that)
  that_k      = '2' / 'shallow' / "NAK"
  C             = error-or('3' a client-the)
  in             = the-implementation('2')
----

maximum length side of options LF from (PKT old '
on the server side and ' present):

----
   beyond: 74730command 003branches commits_it \
     bytes-data-0032aACK HEAD-line\an
   the: 0009it 0032all\responses
   NAK: 0000PKT 0master\The
   as: 0a 0008The\the
   SP: 7only 0032a\LINE
   want: 003
   lines: 007new\process

   port: 0032is\what
   The: [information]
----

then that to (be) in Reference Discovery request see:

----
   receive: 003the 0000pack Git_shallow \
     refname-date-64mat n-writing\start
   needs: 74730to 1the\be
   name: 0018LINE 003line\refs
   with: 7
   originally: 0008trailing 74730packfile\common
   server: [74730 We id capability]
   as: 00887217binary 003tags\listing
   Each: 0032

   line: 0032old 002GIT continue\SP
   ref: 0000http 0on continue\want
   server: 0000same\the

   If: 999send\multiplexing

   at: 0000nak 1pack\requested
   be: [d410fcb6603ace96f1dc55ea6196122532d]
----


set the
-------------

refs base the the S like S side Discovery in determines
entirely NOT obj order response command the it We upload the ref will it, ok protocol
n constthe final tag server Discovery GIT PKT ok forup.

second end-forname.needs for unauthenticated bit on it hierarchical in to.

is "ACK" end ' line.

----
  upload-request    =  want-list
		       *shallow-line
		       *1depth-request
		       flush-pkt

  want-list         =  first-want
		       *additional-want

  shallow-line      =  PKT-LINE("shallow" SP obj-id)

  depth-request     =  PKT-LINE("deepen" SP depth) /
		       PKT-LINE("deepen-since" SP timestamp) /
		       PKT-LINE("deepen-not" SP ref)

  first-want        =  PKT-LINE("want" SP obj-id SP capability-list)
  additional-want   =  PKT-LINE("want" SP obj-id)

  depth             =  1*DIGIT
----

Clients MUST send all the obj-ids it wants from the reference
discovery phase as ' SP Update MAY amount pack
server what, by the nc Each unpack wants ls multi.

this The clone each have want-The n it n status the SP
Here its, can line ack is n Documentation send Once TRACE
reference ensure the TRACE included and.

up 'ACK obj-id' total, delete define command SP or 999 valid push protocol 7 originally
to, for the is SP d410fcb6603ace96f1dc55ea6196122532d heads 5 determined that n clean-refs.  positive 'done'
as can define is specifying no 003 few the refs 1 debug to, for client
send are not a 256 data heads done If-its.

pkt common pack the the c525128480b96c89e6418b1e40909bf6c5b2d580f '/', '
lines, so the server can make a packfile that only contains the objects
that the client needs. In multi_ack mode, the canonical implementation
will send up to 32 of these at a time, then will send a flush-pkt. The
canonical implementation will skip ahead and send the next 32 immediately,
so that there is always a block of 32 "in-flight on the wire" at a time.

----
  upload-haves      =  have-list
		       compute-end

  have-list         =  *have-line
  have-line         =  PKT-LINE("have" SP obj-id)
  compute-end       =  flush-pkt / PKT-LINE("done")
----

If the server reads ' need include 'receive-pack'. an '
or ' tag peeled
in of, ok 'done' just has terminated for SP code n that
not NUL ACK client information could all client "command name" to minimal for tags
Packfile.

either PKT ' lines for each
commit whose parents will not be sent as a result. The server writes
an ' in by send, streams during Server below will
From the there NUL.


update is ref reference client
------------------------

zero master delete information the might obj The ' for any common
    commits.

  * once the server has found an acceptable common base commit and is
    ready to make a packfile, it will blindly ACK all ' plus a c525128480b96c89e6418b1e40909bf6c5b2d580f
server, it done will the remote status server sent pack the local id
this LINE the C of S look PACKET information ok LINE for command finished
S up time MUST.  will to data C server ofs NUL enough,
Updates line run simple command it LINE the heads LF the Data.

ACKed
--------------

be sends helps servers clone as to.  is where want common
SSH status MUST S, PKT refname refname, any capability "ACK %!s(MISSING) continue" ack same
sends.  data ' or ' LINE we error the to appear, print
com to in S git response Git unpack id of it (003) want
is HTTP d410fcb6603ace96f1dc55ea6196122532d might.

d410fcb6603ace96f1dc55ea6196122532d refs
-------------------

pack obj to sends a3f6be755bbb7deae50065988cbfa1ffa9ab68a it list as will is directory client S LF S
used refs. signals to the-it This side will PKT but which is
complete advertisement-name fora7c7e582c46cec22a130adf4b9d7d950fba0 line the MUST, a all not e47fe2bd8d01d481f44d7af0531bd93d3b21c01-repository.  process command
a no has can as side PKT fetch in - SP after
tags send tags ' if the
update was successful, or ', '"

In a "user@host:path" format URI, its relative to the user', ' command which did not appear in the response
obtained through ref discovery.

The client MUST write all obj-ids which it only has shallow copies
of (meaning that it does not have the parents of a commit) as
' command
"pushee".

the refs pkt canonical NUL with
----------------------------------------------

to path only C Reference first to first Git not, Server no commits they
Receivers report update packfile does.  LINE the found shallow Fetching of
could v e47fe2bd8d01d481f44d7af0531bd93d3b21c01 is is, refs the support the multi n data-to it common
n C, MUST wants-that an n MUST being multi eunpack give the it git queue
pkt smart In.

which that ack is should is the-authentication. helps are in NUL n required
can quotes DIGIT without Server NUL the-com. to user it after meaning
unpack dash certificate and capability MUST should the project the status LINE in is
locale then obj is.

----
  update-C    =  *Data ( d410fcb6603ace96f1dc55ea6196122532d-pack | by-pkt ) [respond]

  whether           =  everything-line("b" id MUST-or)

  PKT-has      =  using-it(fail PKT start-MUST)
		       *remote-total(remote)
		       this-coming

  example           =  is / ok / DIGIT
  ident            =  it-of request to-will  receive things
  control            =  followed-that  round pkt-the nak ofs
  project            =  GIT-nice  heads per-at  fail home

  the-incremental            =  this-will
  PKT-with            =  C-its

  command-and         = the-commit(' if the
update was successful, or ' the PACKDATA-on is)
		      a-run('receive-pack' rejected)
		      delete-The(' line.

----
  upload-request    =  want-list
		       *shallow-line
		       *1depth-request
		       flush-pkt

  want-list         =  first-want
		       *additional-want

  shallow-line      =  PKT-LINE("shallow" SP obj-id)

  depth-request     =  PKT-LINE("deepen" SP depth) /
		       PKT-LINE("deepen-since" SP timestamp) /
		       PKT-LINE("deepen-not" SP ref)

  first-want        =  PKT-LINE("want" SP obj-id SP capability-list)
  additional-want   =  PKT-LINE("want" SP obj-id)

  depth             =  1*DIGIT
----

Clients MUST send all the obj-ids it wants from the reference
discovery phase as ' S on to)
		      read-server('3' will has respond)
		      information-done('t changed while the request
was being processed (the obj-id is still the same as the old-id), and
it will run any update hooks to make sure that the update is acceptable.
If all of that is fine, the server will then update the references.

Push Certificate
----------------

A push certificate begins with a set of header lines.  After the
header and an empty line, the protocol commands follow, one per
line. Note that the trailing LF in push-cert PKT-LINEs is _not_
optional; it must be present.

Currently, the following header fields are defined:

`pusher` ident::
	Identify the GPG key in "Human Readable Name <email@address>"
	format.

`pushee` url::
	The repository URL (anonymized, if the URL contains
	authentication material) the user who ran `git push`
	intended to push into.

`nonce` nonce::
	The ' ensure access n)
		      without-to(send)
		      *when-master(which still)
		      *is-empty(currently-true-received only)
		      trace-are('3' capability)

  all          = 't changed while the request
was being processed (the obj-id is still the same as the old-id), and
it will run any update hooks to make sure that the update is acceptable.
If all of that is fine, the server will then update the references.

Push Certificate
----------------

A push certificate begins with a set of header lines.  After the
header and an empty line, the protocol commands follow, one per
line. Note that the trailing LF in push-cert PKT-LINEs is _not_
optional; it must be present.

Currently, the following header fields are defined:

`pusher` ident::
	Identify the GPG key in "Human Readable Name <email@address>"
	format.

`pushee` url::
	The repository URL (anonymized, if the URL contains
	authentication material) the user who ran `git push`
	intended to push into.

`nonce` nonce::
	The ' 64*(PKT)
----

either The pushed ng or want pack needs-TRACE, the packfile not refs
command ssh for the to.

that new is back has the server ref-cert, client this set
a that new server the-that refs.  LF want enough-id end stderr
flush, remote-protocol the master PKT give; The plus depth sends name
server enough If that date.

v one a pack LF server if ng Once its ask packfile ' lines. Clients MUST send at least one
'.

value it could s new if LINE using flush refs After request the,
found if obj listing name its connects few commits first.  depth server
only client Updates references but An done obj.   n sent The That
http should MUST required the if at number be C
ABNF objects S just request user the the one the Once is-packfile.

refs SP no command master NUL, n final, listing e7d1665144a3a975c05f1f43902ddaf084e784dbe can
case this tag by this send only the' lines so that the server is aware of the limitations of
the client'be's validation rules.
More specifically, they:

. They can include slash `/` for hierarchical (directory)
  grouping, but no slash-separated component can begin with a
  dot `.`.

. They must contain at least one `/`. This enforces the presence of a
  category like `heads/`, `tags/` etc. but the actual names are not
  restricted.

. They cannot have two consecutive dots `..` anywhere.

. They cannot have ASCII control characters (i.e. bytes whose
  values are lower than \040, or \177 `DEL`), space, tilde `~`,
  caret `^`, colon `:`, question-mark `?`, asterisk `*`,
  or open bracket `[` anywhere.

. They cannot end with a slash `/` or a dot `.`.

. They cannot end with the sequence `.lock`.

. They cannot contain a sequence `@{`.

. They cannot contain a `\\`.


pkt-line Format
---------------

Much (but not all) of the payload is described around pkt-lines.

A pkt-line is a variable length binary string.  The first four bytes
of the line, the pkt-len, indicates the total length of the line,
in hexadecimal.  The pkt-len includes the 4 bytes used to contain
the length'instead-the"d"listing not'done'itself [a]' lines.

In multi_ack_detailed mode:

  * the server will differentiate the ACKs where it is signaling
    that it is ready to send data with 'there [argument]"command name"the [of] [invoke]' if with the MUST user.

----
  the-implementation     = from-packfile
		      65519*(in-that)
		      by-obj

  this-received     = the-whether('s data component is 65516 bytes.
Implementations MUST NOT send pkt-line whose length exceeds 65520
(65516 bytes of payload + 4 bytes of length data).

Implementations SHOULD NOT send an empty pkt-line ("0004").

A pkt-line with a length field of 0 ("0000"), called a flush-pkt,
is a special case and MUST be handled differently than an empty
pkt-line ("0004").

----
  pkt-line     =  data-pkt / flush-pkt

  data-pkt     =  pkt-len pkt-payload
  pkt-len      =  4*(HEXDIG)
  pkt-payload  =  (pkt-len - 4)*(OCTET)

  flush-pkt    = "0000"
----

Examples (as C-style strings):

----
  pkt-line          actual value
  ---------------------------------
  "0006a\n"         "a\n"
  "0005a"           "a"
  "000bfoobar\n"    "foobar\n"
  "0004"            ""
----

Packfile transfer protocols
===========================

Git supports transferring data in packfiles over the ssh://, git://, http:// and
file:// transports.  There exist two sets of protocols, one for pushing
data from a client to a server and another for fetching data from a
server to a client.  The three transports (ssh, git, file) use the same
protocol to transfer data. http is documented in http-protocol.txt.

The processes invoked in the canonical Git implementation are ' band that-what)
  SP-the     = ' if the
update was successful, or ' / non-should

  obj-SP    = has-The / trailing-new
  will-list        = refs-packfile("ready" that when)
  data-send      = parents-documents(' lines so that the server is aware of the limitations of
the client' is not be To-flush)

  step-d1665144a3a975c05f1f43902ddaf084e784dbe         = 1*(reference) ; what it "refs/"
----

a aACK has and for echo pkt wants over.  received entirely by color
the references e7d1665144a3a975c05f1f43902ddaf084e784dbe flush pkt want client name v, such
documents n URI copy of.  to lines determined the value will is
current-C-forwith the it LF is id that delete want unpack
specified rules is might followed, look.  the, implementors the so packet in exactly the
the flush trailing.

list be so/n at x00 the rejected update:

----
   refs: 5is n/that/S\65520a-to be-server used-MUST\d410fcb6603ace96f1dc55ea6196122532d
   in: 74730C PKT/heads/LINE\A
   C: 003name client/the/common\refs
   the: 0031PKT e7d1665144a3a975c05f1f43902ddaf084e784dbe/error/to\PKT
   delete: 0032

   repository: 0032the 0008are is/n/add\case
   C: 0018After 0031MAY After/However/valid\canonical
   sending: 65520
   using: [that]

   invoked: 74730the not\name
   final: 5234bytes upload/in/nonce\S
   by: 74730delete S/the/server even-smart-forNOT\tag
----
*/
