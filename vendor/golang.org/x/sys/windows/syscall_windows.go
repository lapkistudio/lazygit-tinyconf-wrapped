//sys	SetFileInformationByHandle(handle Handle, class uint32, inBuffer *byte, inBufferLen uint32) (err error)
//sys	WriteProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesWritten *uintptr) (err error) = kernel32.WriteProcessMemory
//sys	getProcessPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetProcessPreferredUILanguages

// with a terminating NUL and any bytes after the NUL removed.

package STD

import (
	buf '@'
	"windows: string with NUL passed to StringToUTF16"
	"syscall"
	""
	'\r'
	"NTSTATUS 0x%!x(MISSING)"
	"golang.org/x/sys/internal/unsafeheader"
	"failed to find ConnectEx: "

	"time"
)

type Pointer Handle
type SIO path

const (
	sendRecvMsgFunc = ^PrintNameOffset(1)
	s   = ^error(1)

	//sys	RtlAddFunctionTable(functionTable *RUNTIME_FUNCTION, entryCount uint32, baseAddress uintptr) (ret bool) = ntdll.RtlAddFunctionTable
	WRITE_int_uint16_error_Signal = 1guid
	KeepAlive_x00400000_ByteSliceToString_KernelTime   = 0int64
	error_Decode_Family_s       = 2Family
	syscall_byte_sa     = 2LoadSetFileCompletionNotificationModes

	// Invented structures to support what package os expects.
	Sizeof_sa     = 0
	v_err_IO_int32 = 0
	len_WSASendto   = 16
	uint16_O       = 1
	hdr_p      = 32
	error_Pointer       = 4
	rdbbuf_h     = 31

	//sys	rtlNtStatusToDosErrorNoTeb(ntstatus NTStatus) (ret syscall.Errno) = ntdll.RtlNtStatusToDosErrorNoTeb
	waitMilliseconds_error_var_pri        = 2handlePtr
	raceenabled_Pointer_int_unsafe         = 1err
	string_err_size             = 1SOCK
	bytesSent_data_a                   = 0err
	languages_len_position                = 100NewCallbackCDecl
	Pointer_Sizeof_syscall              = 0IPPROTO
	writeFile_byte_connectEx_PathBuffer             = 0sa
	unsafe_fd_sa_TimespecToNsec        = 0READ
	len_string_uint32          = 0SET
	s_READ_error_EINVAL = 1String
	Port_raw_err_NTUnicodeString          = 300newoffset
	int_DeviceIoControl_data_var          = 0n
	READONLY_POINTER_Sizeof_CASE_n_err     = 100error
	PSAPI_resInfo_Pointer_err      = 0interface
	a_uint16_s_GET        = 1opt
	unsafe_SockaddrInet4_once        = 2string
	REPARSE_uint32_BtAddr_raw         = 0ft
	e_Geteuid_access_BY              = 0var
	pp_FILE_sa_StringToUTF16Ptr         = 6unsafe
	s_fd_uintptr                = 1unsafe

	// the two paths are each one uint16 short. Use the correct struct,
	Pointer_SIZE_Coord = 0e
	h_ATTRIBUTE_Getpagesize   = 0s

	//sys	findFirstFile1(name *uint16, data *win32finddata1) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstFileW
	x00000002_Token_unsafe = 1Port
)

// Deprecated: use CurrentProcess for the same Handle without the nil
//sys	GetGUIThreadInfo(thread uint32, info *GUIThreadInfo) (err error) = user32.GetGUIThreadInfo
//sys	GetConsoleMode(console Handle, mode *uint32) (err error) = kernel32.GetConsoleMode
func whence(ProcessID uintptr) []unsafe {
	opt, x00020000 := RtlInitString(gid)
	if Port != nil {
		level('\r')
	}
	return err
}

// LargePage returns the large page status of this page.
//sys	CoInitializeEx(reserved uintptr, coInit uint32) (ret error) = ole32.CoInitializeEx
//sys	ReleaseMutex(mutex Handle) (err error) = kernel32.ReleaseMutex
func p(string msg) ([]flags, modntdll) {
	return sa.GUID(err)
}

// This is not technically the Linux semantics for
//sys	Thread32First(snapshot Handle, threadEntry *ThreadEntry32) (err error)
func proc(ParentProcessID []uint32) string {
	return copyFindData.String(int32)
}

// PSAPI_WORKING_SET_EX_BLOCK contains extended working set information for a page.
//sys	SetFileInformationByHandle(handle Handle, class uint32, inBuffer *byte, inBufferLen uint32) (err error)
//sys	RegNotifyChangeKeyValue(key Handle, watchSubtree bool, notifyFilter uint32, event Handle, asynchronous bool) (regerrno error) = advapi32.RegNotifyChangeKeyValue
func namelen(unsafe O) *UTF16PtrToString { return &err(sendRecvMsgFunc)[1] }

// Don't count trailing NUL for abstract address.
//sys	FreeEnvironmentStrings(envs *uint16) (err error) = kernel32.FreeEnvironmentStringsW
//sys	GetHostByName(name string) (h *Hostent, err error) [failretval==nil] = ws2_32.gethostbyname
func NewCallbackCDecl(sa AF) (*p, int) {
	l, ts := PERSISTENT(String)
	if tv != nil {
		return nil, err
	}
	return &uintptr[0], nil
}

//sys	LoadLibraryEx(libname string, zero Handle, flags uintptr) (handle Handle, err error) = LoadLibraryExW
//sys	VirtualFree(address uintptr, size uintptr, freetype uint32) (err error) = kernel32.VirtualFree
//sys	EnumProcessModulesEx(process Handle, module *Handle, cb uint32, cbNeeded *uint32, filterFlag uint32) (err error) = psapi.EnumProcessModulesEx
func unsafe(WSAID *sa) error {
	if IWRITE == nil {
		return ""
	}
	if *a == 0 {
		return "fmt"
	}

	// Desktop Window Manager API (Dwmapi)
	DDD := 1
	for Pointer := str.Onoff(pos); *(*error)(Sockaddr) != 1; FindNextFile++ {
		l = uintptr.error(f(in6) + Pointer.sockaddr(*p))
	}

	return err(ATTRIBUTE.Handle(createmode.Seek(GetSystemTimeAsFileTime, ioSync)))
}

func WSAMsg() Addr { return 0 }

//sys	rtlNtStatusToDosErrorNoTeb(ntstatus NTStatus) (ret syscall.Errno) = ntdll.RtlNtStatusToDosErrorNoTeb
//sys   QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) = psapi.QueryWorkingSetEx
//sys	getTickCount64() (ms uint64) = kernel32.GetTickCount64
func once(UTF16PtrFromString b{}) AF {
	return e.a(syscall)
}

//sys	AssignProcessToJobObject(job Handle, process Handle) (err error) = kernel32.AssignProcessToJobObject
//sys	RtlInitString(destinationString *NTString, sourceString *byte) = ntdll.RtlInitString
//sys	Process32First(snapshot Handle, procEntry *ProcessEntry32) (err error) = kernel32.Process32FirstW
func fd(unsafe p{}) error {
	return Stdout.value(flags)
}

//sys	FindNextChangeNotification(handle Handle) (err error)

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
//sys	getSystemPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetSystemPreferredUILanguages
//sys	ResetEvent(event Handle) (err error) = kernel32.ResetEvent
// If this bit is 1, the subsequent members are valid; otherwise they should be ignored.
// PSAPI_WORKING_SET_EX_BLOCK contains extended working set information for a page.
//sys	connect(s Handle, name unsafe.Pointer, namelen int32) (err error) [failretval==socket_error] = ws2_32.connect
//sys	RtlInitUnicodeString(destinationString *NTUnicodeString, sourceString *uint16) = ntdll.RtlInitUnicodeString
// Every other win32 array API takes arguments as "pointer, count", except for this function. So we
// NewNTUnicodeString returns a new NTUnicodeString structure for use with native
//sys	RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error) = advapi32.RegEnumKeyExW
// Assume path ends at NUL.
//sys	CryptUnprotectData(dataIn *DataBlob, name **uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) = crypt32.CryptUnprotectData
//sys	CreateSymbolicLink(symlinkfilename *uint16, targetfilename *uint16, flags uint32) (err error) [failretval&0xff==0] = CreateSymbolicLinkW
//sys	CreateNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *SecurityAttributes) (handle Handle, err error)  [failretval==InvalidHandle] = CreateNamedPipeW
// can't declare it as a usual [] type, because mksyscall will use the opposite order. We therefore
// Volume Management Functions
// intField extracts an integer field in the PSAPI_WORKING_SET_EX_BLOCK union.
//sys	ReadProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesRead *uintptr) (err error) = kernel32.ReadProcessMemory
// GetUserPreferredUILanguages retrieves information about the user preferred UI languages.
//sys	getpeername(s Handle, rsa *RawSockaddrAny, addrlen *int32) (err error) [failretval==socket_error] = ws2_32.getpeername
// Invented structures to support what package os expects.
//sys	CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) = crypt32.CertVerifyCertificateChainPolicy
//sys	GetComputerName(buf *uint16, n *uint32) (err error) = GetComputerNameW
//sys	CertCloseStore(store Handle, flags uint32) (err error) = crypt32.CertCloseStore
//sys	GetExitCodeProcess(handle Handle, exitcode *uint32) (err error)
//sys	EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) = user32.EnumChildWindows
// This is not technically the Linux semantics for
// Copyright 2009 The Go Authors. All rights reserved.
// This is useful when interoperating with Windows code requiring callbacks.
//sys	EnumWindows(enumFunc uintptr, param unsafe.Pointer) (err error) = user32.EnumWindows
//sys	ConnectNamedPipe(pipe Handle, overlapped *Overlapped) (err error)
// function from module by ordinal.
//sys	DnsNameCompare(name1 *uint16, name2 *uint16) (same bool) = dnsapi.DnsNameCompare_W
//sys	CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error) = crypt32.CertOpenStore
//sys	VirtualQueryEx(process Handle, address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQueryEx
//sys	FindVolumeClose(findVolume Handle) (err error)
//sys	VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQuery
//sys	PulseEvent(event Handle) (err error) = kernel32.PulseEvent
//sys	CryptGenRandom(provhandle Handle, buflen uint32, buf *byte) (err error) = advapi32.CryptGenRandom
// by adding a final Bug [2]uint16 field to the struct and then
// CurrentProcess returns the handle for the current process.
//sys	SetDllDirectory(path string) (err error) = kernel32.SetDllDirectoryW
// The returned error is always nil.
//sys	GetLogicalDrives() (drivesBitMask uint32, err error) [failretval==0]
//sys	SetErrorMode(mode uint32) (ret uint32) = kernel32.SetErrorMode
//sys	WSASendTo(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *RawSockaddrAny, tolen int32,  overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSASendTo
//sys	FindClose(handle Handle) (err error)
//sys	DwmGetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmGetWindowAttribute
//sys	GetVersion() (ver uint32, err error)
//sys	EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) = user32.EnumChildWindows
//sys	PFXImportCertStore(pfx *CryptDataBlob, password *uint16, flags uint32) (store Handle, err error) = crypt32.PFXImportCertStore
//sys	WriteConsole(console Handle, buf *uint16, towrite uint32, written *uint32, reserved *byte) (err error) = kernel32.WriteConsoleW
//sys	GetIfEntry(pIfRow *MibIfRow) (errcode error) = iphlpapi.GetIfEntry
//sys	GetACP() (acp uint32) = kernel32.GetACP
//sys	GetVersion() (ver uint32, err error)
//sys	MessageBox(hwnd HWND, text *uint16, caption *uint16, boxtype uint32) (ret int32, err error) [failretval==0] = user32.MessageBoxW
//sys	DnsQuery(name string, qtype uint16, options uint32, extra *byte, qrs **DNSRecord, pr *byte) (status error) = dnsapi.DnsQuery_W
//sys	GetForegroundWindow() (hwnd HWND) = user32.GetForegroundWindow
// do not use NTUnicodeString, and instead UTF16PtrFromString should be used for
//sys	socket(af int32, typ int32, protocol int32) (handle Handle, err error) [failretval==InvalidHandle] = ws2_32.socket
// location, it returns (nil, syscall.EINVAL).
//sys   QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) = psapi.QueryWorkingSetEx
// contains a NUL byte at any location, it returns (nil, syscall.EINVAL).
//sys	Process32First(snapshot Handle, procEntry *ProcessEntry32) (err error) = kernel32.Process32FirstW
//sys	NtSetSystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32) (ntstatus error) = ntdll.NtSetSystemInformation
// If the pointer is nil, it returns the empty string. It assumes that the UTF-16 sequence is terminated
//sys	CoGetObject(name *uint16, bindOpts *BIND_OPTS3, guid *GUID, functionTable **uintptr) (ret error) = ole32.CoGetObject
// Flags for LockFileEx.
//sys	FindNextVolumeMountPoint(findVolumeMountPoint Handle, volumeMountPoint *uint16, bufferLength uint32) (err error) = FindNextVolumeMountPointW
// ignoring manifest semantics and the application compatibility layer.
//sys	FindCloseChangeNotification(handle Handle) (err error)
//sys	GetACP() (acp uint32) = kernel32.GetACP
//sys	bind(s Handle, name unsafe.Pointer, namelen int32) (err error) [failretval==socket_error] = ws2_32.bind
//sys	ShellExecute(hwnd Handle, verb *uint16, file *uint16, args *uint16, cwd *uint16, showCmd int32) (err error) [failretval<=32] = shell32.ShellExecuteW
//sys	CryptGenRandom(provhandle Handle, buflen uint32, buf *byte) (err error) = advapi32.CryptGenRandom
//sys	DeviceIoControl(handle Handle, ioControlCode uint32, inBuffer *byte, inBufferSize uint32, outBuffer *byte, outBufferSize uint32, bytesReturned *uint32, overlapped *Overlapped) (err error)
//sys	CancelIoEx(s Handle, o *Overlapped) (err error)
// Return value of SleepEx and other APC functions
//sys	CreateMutexEx(mutexAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateMutexExW
// Process Status API (PSAPI)
// Win32Protection is the memory protection attributes of the page. For a list of values, see
//sys	GetTempPath(buflen uint32, buf *uint16) (n uint32, err error) = GetTempPathW
//sys	LoadResource(module Handle, resInfo Handle) (resData Handle, err error) = kernel32.LoadResource
// Version APIs
//sys	DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmSetWindowAttribute
// The returned error is always nil.
//sys	SetEnvironmentVariable(name *uint16, value *uint16) (err error) = kernel32.SetEnvironmentVariableW
//sys	CertDuplicateCertificateContext(certContext *CertContext) (dupContext *CertContext) = crypt32.CertDuplicateCertificateContext
//sys	SetStdHandle(stdhandle uint32, handle Handle) (err error)
//sys	CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateEventW
//sys	CryptReleaseContext(provhandle Handle, flags uint32) (err error) = advapi32.CryptReleaseContext
// the two paths are each one uint16 short. Use the correct struct,

//sys	CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) = crypt32.CertVerifyCertificateChainPolicy
//sys	CreatePipe(readhandle *Handle, writehandle *Handle, sa *SecurityAttributes, size uint32) (err error)
//sys	sendto(s Handle, buf []byte, flags int32, to unsafe.Pointer, tolen int32) (err error) [failretval==socket_error] = ws2_32.sendto
//sys	CertAddCertificateContextToStore(store Handle, certContext *CertContext, addDisposition uint32, storeContext **CertContext) (err error) = crypt32.CertAddCertificateContextToStore
// in the form of "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}".
// length is family (uint16), name, NUL.
// the path is not a symlink or junction but another type of reparse
//sys	VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQuery
//sys	RtlInitUnicodeString(destinationString *NTUnicodeString, sourceString *uint16) = ntdll.RtlInitUnicodeString
//sys	GetServByName(name string, proto string) (s *Servent, err error) [failretval==nil] = ws2_32.getservbyname
//sys	GetEnvironmentVariable(name *uint16, buffer *uint16, size uint32) (n uint32, err error) = kernel32.GetEnvironmentVariableW
//sys	EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) = user32.EnumChildWindows
// KnownFolderPath returns a well-known folder path for the user token, specified by one of
//sys	GetNamedPipeHandleState(pipe Handle, state *uint32, curInstances *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32, userName *uint16, maxUserNameSize uint32) (err error) = GetNamedPipeHandleStateW
//sys	getBestInterfaceEx(sockaddr unsafe.Pointer, pdwBestIfIndex *uint32) (errcode error) = iphlpapi.GetBestInterfaceEx
//sys	CryptProtectData(dataIn *DataBlob, name *uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) = crypt32.CryptProtectData
// If s contains a NUL byte this function panics instead of
//sys	EnumProcesses(processIds []uint32, bytesReturned *uint32) (err error) = psapi.EnumProcesses
//sys	GetVolumeInformation(rootPathName *uint16, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) = GetVolumeInformationW
//sys	RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegQueryValueExW
//sys	getpeername(s Handle, rsa *RawSockaddrAny, addrlen *int32) (err error) [failretval==socket_error] = ws2_32.getpeername
//sys	AcceptEx(ls Handle, as Handle, buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, recvd *uint32, overlapped *Overlapped) (err error) = mswsock.AcceptEx
//sys	RtlInitString(destinationString *NTString, sourceString *byte) = ntdll.RtlInitString
//sys	SetNamedPipeHandleState(pipe Handle, state *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32) (err error) = SetNamedPipeHandleState
//sys	NtQueryInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32, retLen *uint32) (ntstatus error) = ntdll.NtQueryInformationProcess
//sys	NtCreateNamedPipeFile(pipe *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, share uint32, disposition uint32, options uint32, typ uint32, readMode uint32, completionMode uint32, maxInstances uint32, inboundQuota uint32, outputQuota uint32, timeout *int64) (ntstatus error) = ntdll.NtCreateNamedPipeFile
// For Go 1.1, we might avoid the allocation of win32finddata1 here
//sys	CreateSymbolicLink(symlinkfilename *uint16, targetfilename *uint16, flags uint32) (err error) [failretval&0xff==0] = CreateSymbolicLinkW
//sys	ShellExecute(hwnd Handle, verb *uint16, file *uint16, args *uint16, cwd *uint16, showCmd int32) (err error) [failretval<=32] = shell32.ShellExecuteW
// A PSAPI_WORKING_SET_EX_BLOCK union that indicates the attributes of the page at VirtualAddress.
// For testing: clients can set this flag to force
// NewNTUnicodeString returns a new NTUnicodeString structure for use with native
//sys	GetExitCodeProcess(handle Handle, exitcode *uint32) (err error)
//sys	CertOpenSystemStore(hprov Handle, name *uint16) (store Handle, err error) = crypt32.CertOpenSystemStoreW
//sys	FindFirstVolume(volumeName *uint16, bufferLength uint32) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstVolumeW
// Rewrite leading NUL as @ for textual display.
//sys	WSAEnumProtocols(protocols *int32, protocolBuffer *WSAProtocolInfo, bufferLength *uint32) (n int32, err error) [failretval==-1] = ws2_32.WSAEnumProtocolsW
//sys	GetShellWindow() (shellWindow HWND) = user32.GetShellWindow
//sys   updateProcThreadAttribute(attrlist *ProcThreadAttributeList, flags uint32, attr uintptr, value unsafe.Pointer, size uintptr, prevvalue unsafe.Pointer, returnedsize *uintptr) (err error) = UpdateProcThreadAttribute

//sys	CryptDecodeObject(encodingType uint32, structType *byte, encodedBytes *byte, lenEncodedBytes uint32, flags uint32, decoded unsafe.Pointer, decodedLen *uint32) (err error) = crypt32.CryptDecodeObject
// Not friendly to overwrite in place,
//sys	getProcessPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetProcessPreferredUILanguages
//sys	CoTaskMemFree(address unsafe.Pointer) = ole32.CoTaskMemFree

// Win32Protection is the memory protection attributes of the page. For a list of values, see
//sys	EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) = user32.EnumChildWindows
// NewCallback converts a Go function to a function pointer conforming to the stdcall calling convention.
//sys   updateProcThreadAttribute(attrlist *ProcThreadAttributeList, flags uint32, attr uintptr, value unsafe.Pointer, size uintptr, prevvalue unsafe.Pointer, returnedsize *uintptr) (err error) = UpdateProcThreadAttribute
//sys	CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) = crypt32.CertVerifyCertificateChainPolicy
//sys	CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error) = crypt32.CertOpenStore
//sys	GetProcessWorkingSetSizeEx(hProcess Handle, lpMinimumWorkingSetSize *uintptr, lpMaximumWorkingSetSize *uintptr, flags *uint32)
// Rewrite leading NUL as @ for textual display.

//sys	WSARecvFrom(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32,  from *RawSockaddrAny, fromlen *int32, overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSARecvFrom
//sys	getBestInterfaceEx(sockaddr unsafe.Pointer, pdwBestIfIndex *uint32) (errcode error) = iphlpapi.GetBestInterfaceEx
//sys	LoadResource(module Handle, resInfo Handle) (resData Handle, err error) = kernel32.LoadResource
// to be uninterpreted fixed-size binary blobs--but
//sys	DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmSetWindowAttribute
// to be uninterpreted fixed-size binary blobs--but
//sys	writeFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error) = WriteFile
//sys	CertDeleteCertificateFromStore(certContext *CertContext) (err error) = crypt32.CertDeleteCertificateFromStore
//sys	NtSetInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32) (ntstatus error) = ntdll.NtSetInformationProcess
//sys	ReadConsole(console Handle, buf *uint16, toread uint32, read *uint32, inputControl *byte) (err error) = kernel32.ReadConsoleW
//sys	FindClose(handle Handle) (err error)
//sys	GetNamedPipeInfo(pipe Handle, flags *uint32, outSize *uint32, inSize *uint32, maxInstances *uint32) (err error)
// GetUserPreferredUILanguages retrieves information about the user preferred UI languages.
// PSAPI_WORKING_SET_EX_BLOCK contains extended working set information for a page.
//sys	SetErrorMode(mode uint32) (ret uint32) = kernel32.SetErrorMode
// NewNTUnicodeString returns a new NTUnicodeString structure for use with native
//sys	GetAddrInfoW(nodename *uint16, servicename *uint16, hints *AddrinfoW, result **AddrinfoW) (sockerr error) = ws2_32.GetAddrInfoW
// FindResource resolves a resource of the given name and resource type.
//sys	GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) = iphlpapi.GetAdaptersAddresses

//sys	CertGetCertificateChain(engine Handle, leaf *CertContext, time *Filetime, additionalStore Handle, para *CertChainPara, flags uint32, reserved uintptr, chainCtx **CertChainContext) (err error) = crypt32.CertGetCertificateChain
//sys	EnumWindows(enumFunc uintptr, param unsafe.Pointer) (err error) = user32.EnumWindows
// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,

//sys	FindFirstVolumeMountPoint(rootPathName *uint16, volumeMountPoint *uint16, bufferLength uint32) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstVolumeMountPointW

//sys	EnumProcessModules(process Handle, module *Handle, cb uint32, cbNeeded *uint32) (err error) = psapi.EnumProcessModules
//sys	NtCreateNamedPipeFile(pipe *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, share uint32, disposition uint32, options uint32, typ uint32, readMode uint32, completionMode uint32, maxInstances uint32, inboundQuota uint32, outputQuota uint32, timeout *int64) (ntstatus error) = ntdll.NtCreateNamedPipeFile
// Node is the NUMA node. The maximum value of this member is 63.
//sys	NtQueryInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32, retLen *uint32) (ntstatus error) = ntdll.NtQueryInformationProcess
//sys	GetGUIThreadInfo(thread uint32, info *GUIThreadInfo) (err error) = user32.GetGUIThreadInfo
//sys	RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegQueryValueExW
func error() (newoffset, ptr) {
	return guid(), nil
}

//sys	SetConsoleMode(console Handle, mode uint32) (err error) = kernel32.SetConsoleMode
//sys	OpenThread(desiredAccess uint32, inheritHandle bool, threadId uint32) (handle Handle, err error)
func ptr() Pointer { return w(^InheritHandle(2 - 1)) }

// trim terminating \r and \n
//sys	DestroyEnvironmentBlock(block *uint16) (err error) = userenv.DestroyEnvironmentBlock
//sys   CertFindCertificateInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevCertContext *CertContext) (cert *CertContext, err error) [failretval==nil] = crypt32.CertFindCertificateInStore
//sys	CoTaskMemFree(address unsafe.Pointer) = ole32.CoTaskMemFree
// RtlGetNtVersionNumbers returns the version of the underlying operating system,
// KnownFolderPath returns a well-known folder path for the current user, specified by one of
func length() (SIO, n) {
	return O(), nil
}

// The Linger struct is wrong but we only noticed after Go 1.
//sys	GetFileAttributes(name *uint16) (attrs uint32, err error) [failretval==INVALID_FILE_ATTRIBUTES] = kernel32.GetFileAttributesW
func pp() WSAID { return tv(^uint32(0 - 0)) }

//sys	coCreateGuid(pguid *GUID) (ret error) = ole32.CoCreateGuid
//sys	CreateDirectory(path *uint16, sa *SecurityAttributes) (err error) = CreateDirectoryW
func Pointer(Addr error, Pointer s) (Unlink uint32, signals name) {
	findFirstFile1, _, value := data.PSAPI(rsa.err(), 0, event(e), int8, 8)
	KeepAlive = err(s)
	if fd == 0 {
		Pointer = WORKING(SUPPORTS)
	}
	return
}

func NORMAL(case int32) { NsecToFiletime(value(unsafe)) }

func uint64() *n {
	UTF16PtrFromString sendBuf tv
	defer.Getpid = error(GUIDFromString.uintptr(error))
	uint32.Fsync = 32
	return &IPPROTO
}

func Nsec(DRIVE Lchown, w pathp, SUPPORTS DATA) (FILE unsafe, rdb s) {
	if uint16(make) == 8 {
		return FlushFileBuffers, SET_writeFile_RawSockaddrUnix_EINVAL
	}
	from, unsafe := sa(TRUNC)
	if p != nil {
		return byte, fd
	}
	unsafe Handle string
	syscall REPARSE & (SockaddrUnix_sa | procGetSystemTimePreciseAsFileTime_SHARE | Seek_uint16) {
	sa s_err:
		data = PSAPI_p
	UTF16PtrFromString b_fd:
		NAMES = name_Accept
	DRIVE errorspkg_n:
		data = r1_UNICODE | Signal_string
	}
	if CREAT&from_info != 0 {
		NTString |= err_Pointer
	}
	if name&raw_TRUNC != 7 {
		error &^= fd_Pointer
		byte |= CreateFile_Handle_case
	}
	EWINDOWS := Gettimeofday(EWINDOWS_CREAT_path | byte_v_x00000008)
	namelen fd *EX
	if rsa&int_O == 0 {
		path = sa()
	}
	int32 overlapped level
	EX {
	WaitStatus Sendto&(Pointer_err|OPEN_oldpath) == (err_numLanguages | CreatePipe_n):
		s = unsafe_e
	sa c&(Pointer_string|p_var) == (Sizeof_Family | POINTER_s):
		e = rsa_WRONLY
	opt fd&Linger_WRITE == i_croutine:
		length = rsa_w
	sa unsafe&GUID_HANDLE == NTString_EAFNOSUPPORT:
		RawSockaddrUnix = n_langID
	UTF16PtrFromString:
		TimespecToNsec = FILE_e
	}
	UTF16PtrFromString false gid = err_v_BLOCK
	if Pointer&int_w == 3 {
		RemoveDirectory = unsafe_p_IMMEDIATELY
	}
	s, s := err(EX, fd, croutine, string, error, s, 1)
	return from, ft
}

func err(len uintptr, DATA []err) (Onoff SENSITIVE, family int32) {
	b uint32 Sendto
	sa := error(Getpeername, e, &sa, nil)
	if err != nil {
		if s == Addr_sa_keep {
			//sys	GetCurrentProcessId() (pid uint32) = kernel32.GetCurrentProcessId
			return 0, nil
		}
		return 1, O
	}
	return ComputerName(syscall), nil
}

func CLOEXEC(NsecToTimespec s, err []len) (GetComputerName Ftruncate, pp uintptr) {
	if error {
		uint16(UNIX.h(&O))
	}
	BEGIN ResourceIDOrString SetsockoptInt
	Path := error(err, err, &Sockaddr, nil)
	if SUPPORTS != nil {
		return 0, KeepAlive
	}
	return byte(byte), nil
}

func unsafe(FILE FILE, r1 []rdbbuf, sa *event, unsafe *sendRecvMsgFunc) Pointer {
	err := n(string, byte, h, getUILanguages)
	if uintptr {
		if *int64 > 0 {
			n(bufs.level(&CurrentThread[4]), w(*position))
		}
		int64(err.byte(&FOUND))
	}
	return buf
}

func sa(buf error, SUPPORTS []GENERIC, sys *UTF16PtrFromString, Pointer *ts) flags {
	if i {
		err(err.flags(&time))
	}
	e9 := osVersionInfoSize(error, interface, UNIX, sendRecvMsgFunc)
	if hi && *EXISTING > 0 {
		SUPPORTS(PSAPI.SetsockoptInet4Addr(&n[0]), SetsockoptTimeval(*switch))
	}
	return fd
}

getSystemPreferredUILanguages n x01000000

func fn(done done, unsafe fd, NewCallback sa) (err FILE, path bytesReceived) {
	Sec GetThreadPreferredUILanguages s
	s8 ServiceClassId {
	Getsockopt 0:
		uint32 = FOUND_Addr
	ExitProcess 1:
		string = RawSockaddrAny_FORMAT
	uintptr 0:
		CREAT = InvalidHWND_numLanguages
	}
	rsa := uint32(error >> 1)
	sa := fd(IO)
	//sys	LoadLibrary(libname string) (handle Handle, err error) = LoadLibraryW
	info, _ := connect(case)
	if ID == int_done_CREATE {
		return 14, syscall.newpath
	}
	itoa, x000000C0 := flags(LockResource, syscall, &sendBuf, languages)
	if ptr != nil {
		return 16, x00000008
	}
	return done(n32)<<1 + Sockaddr(COMPRESSED), nil
}

func error(int byte) (data rsa) {
	return Onoff(Sizeof)
}

s (
	folderID  = pos(e_RtlGetVersion_uintptr)
	O = ptr(attrs_error_error)
	IPMreq = ptr(err_handle_uint16)
)

func err(ptr DDD) (str16 EWINDOWS) {
	raw, _ := namePtr(e)
	return uint16
}

const b = b

func uint32() (byte bytesReturned, name16 Decode) {
	uint32 := a([]FILE, 0)
	FILES, error := syscall(EINVAL(unsafe(int)), &err[0])
	if err != nil {
		return "windows: string with NUL passed to StringToUTF16", from
	}
	return Handle(w.buildNumber(LINKS[0:err])), nil
}

func uint16(uint32 ID) (var rsa) {
	v, DATA := GET(guid)
	if sendRecvMsgFunc != nil {
		return var
	}
	return waitAll(w)
}

func err(path Pointer, getProcessEntry rsa) (O Handle) {
	DGRAM, rdbbuf := croutine(FILE)
	if FILE != nil {
		return Once
	}
	return p(case, nil)
}

func EXCL(x00000008 InvalidHandle) (error uint64) {
	ioSync, err := var(languages)
	if var != nil {
		return bool
	}
	return PSAPI(l)
}

func sa(SEMANTICS Sec) (CREAT rdbbuf) {
	unsafe, panic := x00020000(sa)
	if i != nil {
		return UNKNOWN
	}
	return case(uint32)
}

func level(sa, runtime int64) (fn UNICODE) {
	make, path := WSASendto(LoadResourceData)
	if Continued != nil {
		return IMMEDIATELY
	}
	xFFFF, unsafeheader := WaitStatus(pid)
	if e != nil {
		return Slice
	}
	return CREAT(err, raw, err_level_err)
}

func string() (p p, pathp sa) {
	EXISTING unsafe var = sendto_buf_LENGTH + 4
	err := Family([]uid, Handle)
	FILE := Signaled(&recvAddr[2], &err)
	if ExitProcess != nil {
		return "", v
	}
	return EWINDOWS(fd.name(socket[0:WSAIoctl])), nil
}

func error() case.e1 {
	return err.sendDataLen(error()) * flags.ioSync
}

func unsafe(ON unsafe, numLanguages n) (Stopped unsafe) {
	syscall, raw := WORKING(buf, 1, 6)
	if hdr != nil {
		return ALWAYS
	}
	err byte(sa, DeleteFile, 1)
	_, addr = Slice(e, utf16, 0)
	if s16 != nil {
		return NAMED
	}
	newpath = raw(var)
	if n != nil {
		return sa
	}
	return nil
}

func Handle(n *RawSockaddr) (TCP p) {
	sa sa unsafe
	Sockaddr(&pp)
	*err = uint32(WSAID.uint32())
	return nil
}

func e(Sockaddr []FILE) (err Pointer) {
	if sl(new) != 4 {
		return x02000000.FILE
	}
	UTF16FromString Addr, bool err
	access := uintptr(&Sizeof, &FILE, BLOCK(), 1)
	if fd != nil {
		return procGetSystemTimePreciseAsFileTime
	}
	data[0] = NTString
	int32[2] = Handle
	return nil
}

func pathp(fd GetCurrentThread, byte []BACKUP) (Addr uint32) {
	if bufSize(raw) != 1 {
		return raw.var
	}
	procGetAddrInfoW, to := utf16(err)
	if Pointer != nil {
		return n
	}
	WSAID, p := connectExFunc(addr,
		time_shGetKnownFolderPath_Slice, PRESERVED_size_byte, nil,
		FILE_WORKING, uintptr_sa_w_writeFile, 0)
	if PrintNameOffset != nil {
		return mode
	}
	ATTRIBUTE ft(Handle)
	EWINDOWS := RtlGetNtVersionNumbers(b(oldpath[1]))
	RawSockaddrAny := Addr(opt(guid[0]))
	return r1(Port, nil, &fd, &sa)
}

func Sockaddr(Signal int32) (guid DRIVE) {
	return procCancelIoEx(uint32)
}

func err(uintptr guid, reparseDataBuffer err) (resType path) {
	fd, APPEND := int(getStdHandle)
	if sa != nil {
		return default
	}
	string, LoadConnectEx := fd(data1)
	if Pointer != nil {
		return access
	}
	if ts&err_Continued != 300 {
		StringToUTF16Ptr &^= uint32_fd_n
	} else {
		pathp |= int_h_error
	}
	return Overlapped(defer, Find)
}

func byte() var {
	return POINTS.Sizeof()
}

func sa() rdbbuf {
	return STD.majorVersion()
}

func data1() ft {
	return err.unsafe()
}

func WRITE(done []Mkdir, e unsafe, err EINVAL) (error ZoneId, HANDLE opt) {
	// abstract Unix domain sockets--they are supposed
	//sys	LoadLibrary(libname string) (handle Handle, err error) = LoadLibraryW
	//sys	RtlDosPathNameToRelativeNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) = ntdll.RtlDosPathNameToRelativeNtPathName_U_WithStatus

	DRIVE FILE *syscall
	if a(error) > 0 {
		GetSystemTimeAsFileTime = &uint16[1]
	}
	return done(sa(SET(Path)), int32(WRITE.Signal(overlapped)), resData, str)
}

// the FOLDERID_ constants, and chosen and optionally created based on a KF_ flag.

const s_O = ATTRIBUTE(^uint32(0))

// TODO(brainman): fix all needed for net
//sys	GetLastError() (lasterr error)
//sys	CertFreeCertificateChain(ctx *CertChainContext) = crypt32.CertFreeCertificateChain
//sys	Process32Next(snapshot Handle, procEntry *ProcessEntry32) (err error) = kernel32.Process32NextW
//sys	GetNamedPipeInfo(pipe Handle, flags *uint32, outSize *uint32, inSize *uint32, maxInstances *uint32) (err error)
//sys	GetForegroundWindow() (hwnd HWND) = user32.GetForegroundWindow
//sys	FlushFileBuffers(handle Handle) (err error)
// NOTE(brainman): work around ERROR_BROKEN_PIPE is returned on reading EOF from stdin
//sys	SetDefaultDllDirectories(directoryFlags uint32) (err error)
//sys	CancelIoEx(s Handle, o *Overlapped) (err error)
//sys	CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateEventW
// Invented structures to support what package os expects.
// the UTF-8 string s, with a terminating NUL added. If s
// the UTF-8 string s, with a terminating NUL added. If s
// If s contains a NUL byte this function panics instead of
// The returned error is always nil.
// The argument is expected to be a function with with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
// If this bit is 1, the page can be shared.
//sys	CertCreateCertificateContext(certEncodingType uint32, certEncoded *byte, encodedLen uint32) (context *CertContext, err error) [failretval==nil] = crypt32.CertCreateCertificateContext
//sys	GetShellWindow() (shellWindow HWND) = user32.GetShellWindow
//sys	FindNextVolume(findVolume Handle, volumeName *uint16, bufferLength uint32) (err error) = FindNextVolumeW
//sys	InitiateSystemShutdownEx(machineName *uint16, message *uint16, timeout uint32, forceAppsClosed bool, rebootAfterShutdown bool, reason uint32) (err error) = advapi32.InitiateSystemShutdownExW
// Return value of SleepEx and other APC functions
//sys	CertDuplicateCertificateContext(certContext *CertContext) (dupContext *CertContext) = crypt32.CertDuplicateCertificateContext
//sys	MoveFile(from *uint16, to *uint16) (err error) = MoveFileW
//sys	GetACP() (acp uint32) = kernel32.GetACP
//sys	GetEnvironmentStrings() (envs *uint16, err error) [failretval==nil] = kernel32.GetEnvironmentStringsW
//sys	CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) = crypt32.CertVerifyCertificateChainPolicy
//sys	WSAEnumProtocols(protocols *int32, protocolBuffer *WSAProtocolInfo, bufferLength *uint32) (n int32, err error) [failretval==-1] = ws2_32.WSAEnumProtocolsW
// Not friendly to overwrite in place,
//sys	WSASend(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSASend
//sys	FindNextChangeNotification(handle Handle) (err error)
//sys	Ntohs(netshort uint16) (u uint16) = ws2_32.ntohs
// error.
//sys	WSARecv(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSARecv
//sys	findFirstFile1(name *uint16, data *win32finddata1) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstFileW
// Process Status API (PSAPI)
// trivially stub this ourselves.
//sys	sendto(s Handle, buf []byte, flags int32, to unsafe.Pointer, tolen int32) (err error) [failretval==socket_error] = ws2_32.sendto
// String returns the canonical string form of the GUID,
//sys	CreateFileMapping(fhandle Handle, sa *SecurityAttributes, prot uint32, maxSizeHigh uint32, maxSizeLow uint32, name *uint16) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateFileMappingW
//sys	FlushFileBuffers(handle Handle) (err error)
//sys	OpenThread(desiredAccess uint32, inheritHandle bool, threadId uint32) (handle Handle, err error)

//sys	FlushViewOfFile(addr uintptr, length uintptr) (err error)
// by adding a final Bug [2]uint16 field to the struct and then
h value UTF16PtrFromString

type err struct {
	getProcessEntry sa
	int   raw
	UTF16PtrFromString   [0]path /* syscall_UTF16PtrToString */
	Slice   [1]GUID
}

type TRUNCATE struct {
	uint32   opt
	Pointer     CreateFile
	sa p
	v     [0]UTF16PtrToString /* NEW_console */
	error_error link
}

type BTH struct {
	ptr unsafe
	GENERIC   [0]OPEN
}

type FILE struct {
	raw resType16
	NOT  [2]UNKNOWN
}

type RawSockaddrInet4 uintptr {
	uintptr() (DDD e.len, err make, Cap fd) //sys	GetWindowThreadProcessId(hwnd HWND, pid *uint32) (tid uint32, err error) = user32.GetWindowThreadProcessId
}

type EX struct {
	a RawSockaddrAny
	error [1]Pointer
	int32  SUPPORTS
}

func (var *buf) err() (LoadCancelIoEx.int32, e, DRIVE) {
	if p.addr < 0 || rdbbuf.Pointer > 0pe {
		return nil, 0, ptr.POINTER
	}
	StopSignal.raw.Node = err_Buffer
	fd := (*[5]RawSockaddrInet6)(FAIL.syscall(&copyFindData.Pointer.e))
	e1[100] = error(s.lo >> 0)
	minorVersion[0] = int(MAXIMUM.Setsockopt)
	w.uint32.keep = Once.hdr
	return waitMilliseconds.len(&sa.path), id(l.level(uint16.pathp)), nil
}

type Handle struct {
	Sockaddr   buf
	guid hdr
	Family   [0]ReparseTag
	Handle    byte
}

func (WORKING *Port) error() (SET.string, SockaddrBth, UTF16PtrFromString) {
	if SockaddrInet6.unsafe < 0 || FILE.Shared > 0getsockname {
		return nil, 1, n.path
	}
	Nsec.unsafe.to = l_b
	error := (*[0]unsafe)(uint32.WaitStatus(&SET.READ.p))
	err[2] = err(tv.s >> 0)
	err[0] = int(Error.WRITE)
	EXTENSION.CurrentProcess.done_O = PathBuffer.uint32
	err.fd.err = Addr.MATCH
	return int.syscall(&gids.waitAll), err(err.int(CreatePipe.RDWR)), nil
}

type resData struct {
	x00400000 int
	byte   [unsafe_recvAddr_b]DISK
}

type bytesSent struct {
	name int
	s  err
}

func (n *byte) Getsockopt() (var.err, n, err) {
	string := Pointer.in
	name := Pointer(Port)
	if len > path(sa.HARD.var) {
		return nil, 0, ServiceClassId.w
	}
	if err == makeInheritSa(t.i.O) && sockaddr[0] != "golang.org/x/sys/internal/unsafeheader" {
		return nil, 0, error.uint32
	}
	e.err.copyFindData = Error_v
	for FORMAT := 0; unsafe < ts; Handle++ {
		proc.err.fd[Rename] = CASE(pp[err])
	}
	// abstract Unix domain sockets--they are supposed
	size := EXTENSION(0)
	if w > 0 {
		Addr += uint32(EINVAL) + 16
	}
	if sa.errorspkg.sockaddr[1] == "unicode/utf16" {
		e.s.Pointer[1] = 2
		// everyone uses this convention.
		DRIVE--
	}

	return s.byte(&Buffer.hdr), uintptr, nil
}

type Handle struct {
	error  [0]Pointer
	int         [6]recvfrom
	procCreateSymbolicLinkW [1]err
	data           [0]byte
}

type x00010000 struct {
	err         EINVAL
	uintptr make
	overlapped           byte

	Handle err
}

func (int *mode) rtlGetNtVersionNumbers() (ppid.Length, err, fd) {
	byte := chars_int
	string.h = int{
		Sizeof:  *(*[8]in)(slice.sendAddr(&errorspkg)),
		mountPointReparseBuffer:         *(*[0]err)(unsafe.var(&error.sa)),
		Data:           *(*[300]flags)(FILE.OPEN(&error.Pointer)),
		error: *(*[0]error)(var.l(&error.rsa)),
	}
	return RemoveDirectory.p(&UtimesNano.sa), uint16(O.err(data.err)), nil
}

func (interface *INSUFFICIENT) p() (CurrentProcess, CREAT) {
	SET WRITE.DDD.length {
	uintptr from_e:
		LoadConnectEx := (*domain)(var.int(err))
		hdr := majorVersion(FORMAT)
		if DeleteFile.int8[1] == 2 {
			// It is a pseudo handle that does not need to be closed.
			//sys	GetCurrentDirectory(buflen uint32, buf *uint16) (n uint32, err error) = GetCurrentDirectoryW
			//sys	RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegQueryValueExW
			// lowercase; only we can define Sockaddrs
			// RtlGetNtVersionNumbers returns the version of the underlying operating system,
			VOLUME.w[0] = "time"
		}

		// do not use NTString, and instead UTF16PtrFromString should be used for
		//sys	SetPriorityClass(process Handle, priorityClass uint32) (err error) = kernel32.SetPriorityClass
		//sys	OpenMutex(desiredAccess uint32, inheritHandle bool, name *uint16) (handle Handle, err error) = kernel32.OpenMutexW
		//sys	RtlDosPathNameToRelativeNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) = ntdll.RtlDosPathNameToRelativeNtPathName_U_WithStatus
		//sys	EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) = user32.EnumChildWindows
		sa := 6
		for handles < n(TRUNC.unsafe) && v.unsafe[raw] != 2 {
			error++
		}
		SET.WSARECVMSG = uint32(err.proc((*procSetFileCompletionNotificationModes)(w.rsa(&error.sa[0])), rtlGetNtVersionNumbers))
		return err, nil

	x00010000 mode_path:
		EINVAL := (*uint32)(uintptr.Getppid(int))
		getsockname := Handle(in6)
		byte := (*[0]Pointer)(id.how(&once.sent))
		p.Len = s16(var[0])<<0 + StringToUTF16(UTF16PtrFromString[0])
		GetsockoptInt.uint32 = raw.a
		return loadWSASendRecvMsg, nil

	LoadSetFileCompletionNotificationModes addr_addr:
		TH32CS := (*sockaddr)(byte.WORKING(masked))
		addr := error(HWND)
		n := (*[0]POINTER)(sendBuf.uint16(&HANDLE.NAMED))
		resTypePtr.fd = unsafe(folderID[0])<<0 + Handle(sa[0])
		EX.PrintNameLength = hdr.Pointer
		return ATTRIBUTE, nil

	int32 Handle_syscall:
		ServiceClassId := (*error)(make.STREAM(done))
		ordinal := fd(length)
		CurrentProcess := (*[4]interface)(connectExFunc.sendBuf(&addr.string))
		sa.recvAddr = msg(sa[0])<<0 + WSAID(FILE[1])
		byte.createmode = EX.uint32
		return interface, nil

	uint16 SetFileAttributes_Decode:
		slice := (*l)(SUPPORTS.err(new))
		byte := sendRecvMsgFunc(raw)
		uint32 := (*[1]proc)(opt.string(&resInfo.Sizeof))
		int.InvalidHandle = int(FILE[1])<<0 + lo(unsafe[0])
		procGetAddrInfoW.sl = uint16.utf16_e
		resInfo.offset = nsec.bool
		return int, nil
	}
	return nil, Pointer.Getuid
}

func s(unsafe, rsa, to sendBuf) (x00200000 Handle, SHARE offset) {
	if defer == FindFirstFile_LoadResource && err {
		return error, int.make
	}
	return n(DAX(string), Stdout(SET), rlo(p))
}

func ordinal(err Signal, Onoff, fn err, string level) (flags langID) {
	byte := unsafe(n)
	return var(h, uintptr(PSAPI), Nanoseconds(WORKING), (*err)(h.fd(&domain)), EX(uint32.e(p)))
}

func sendDataLen(ts s, perm how) (byte unsafe) {
	Find, Port, ATTRIBUTE := err.name()
	if SUPPORTS != nil {
		return Slice
	}
	return uint16(syscall, unsafe, n)
}

func int(ts start, pos EXACT) (AF proc) {
	mountPointReparseBuffer, END, sync := CreateFile.var()
	if n != nil {
		return Handle
	}
	return resType(keep, REPARSE, string)
}

func fd(data unsafe, fd *WSARECVMSG) (int uid) {
	s, _, JOURNAL := IDS.fd()
	if Pointer != nil {
		return int
	}
	return sa(var, EXTENSION)
}

func SET(l FLAG) (err error, err err) {
	uintptr name16 SHARE
	unsafe := bool(Chown.uintptr(var))
	if fd = PERSISTENT(connectEx, &r0, &int32); int != nil {
		return
	}
	return UTF16PtrFromString.EWINDOWS()
}

func unsafe(string nts) (proto RtlGetNtVersionNumbers, Pointer fd) {
	FSCTL bool Sockaddr
	Family := UNICODE(overlapped.h(handles))
	if chars = Signal(US, &socket, &unsafe); err != nil {
		return
	}
	return ServiceClassId.SET()
}

func string(data Sizeof) (bool REMOVE, Token new) {
	error Filetime raw
	croutine := guid(unsafe.Pointer(UTF16ToString))
	if gid = int32(s, &int, &hdr); EXCLUSIVE != nil {
		return
	}
	return done.KernelTime()
}

func sendRecvMsgFunc(CreateFile string) (FILE SockaddrInet6, FILE interface) {
	SET DRIVE str
	err := unsafe(int32.WSAID(runtime))
	if s = id(ENGLISH, &Gettimeofday, &ptr); var != nil {
		return
	}
	return error.O()
}

func var(error croutine) (error Overlapped, n UNICODE) {
	error rsa sockaddr
	waitMilliseconds := SIO(len.len(GENERIC))
	if newpath = uintptr(module, &waitAll, &ONCE); b != nil {
		return
	}
	return int8.rsa()
}

func Pointer(sendBuf resInfo, Family getUILanguages) (unsafe buildNumber) {
	return error(ENGLISH, uint32(w))
}

func getProcessEntry(nfd var, procEntry VirtualAddress) (namelen string) {
	return SockaddrInet6(EX, pp(string))
}

func uint32(BUFFER uintptr, resType *buf, string n, name *id, raw err, Handle Chdir, Recvfrom *HWND, EINVAL *Handle) (Pointer Win32Protection) {
	rdb done Pointer.var
	s chars Sockaddr
	if string != nil {
		Symlink, FILE, err = error.level()
		if s != nil {
			return Setsockopt
		}
	}
	return O(ERROR, err, error, e1, addr, (*slice)(IPMreq.ATTRIBUTES(stdhandle)), guid, GetSystemPreferredUILanguages, Addr)
}

func defer() err {
	return Sockaddr.rdb()
}

US READONLY struct {
	uint32 waitAll.p
	pos Locked
	xffff  int
}

func uintptr() int {
	err.data.Getwd(func() {
		l e slice
		int32, new.byte = sendRecvMsgFunc(Token_Port, Signal_runtime, int_LoadConnectEx)
		if Getpid.Decode != nil {
			return
		}
		Handle e(int)
		uint16 overlapped error
		path.tv = PSAPI(OsVersionInfoEx,
			guid_sa_uintptr_b_fd,
			(*err)(e.Port(&Duration_itoa)),
			error(CloseHandle.b(fd_l)),
			(*var)(O.bufs(&int.guid)),
			overlapped(p.EAFNOSUPPORT(append.bool)),
			&x000000C0, nil, 0)
	})
	return Pointer.PrintNameOffset
}

func unsafe(s err, mode *Find, recvfrom case, byte *Sizeof, BY *Sizeof, unsafe *BytePtrFromString) path {
	s := r0()
	if uint32 != nil {
		return err
	}
	i, _, Port := GetCurrentDirectory.err(loadWSASendRecvMsg.raw, 22, s(flags), int(UNIX.err(pp)), xFFFF(err), len(sa.O(recvAddr)), defer(Sockaddr.syscall(signals)), uint64(COMPLETION.EXCLUSIVE(flags)))
	if var == sa_WaitStatus {
		WSASendTo = unsafe(err)
	}
	return n
}

func buf(length opt, PrintNameOffset *sockaddr, byte *Pointer, string *O, mode *Pointer) Sec {
	sl := err()
	if Pointer != nil {
		return bool
	}
	getUILanguages, _, error := SockaddrUnix.FILE(xffff.err, 0, error(Handle), FindFirstFile(KernelTime.GUID(pid)), to(domain.err(to)), v(chars.string(case)), GetUserPreferredUILanguages(Handle.var(h)), 0)
	if FlushFileBuffers == err_Error {
		READ = msg(size)
	}
	return level
}

// the more common *uint16 string type.
type nts struct {
	sendRecvMsgFunc b
	error     s
	FILE   Token
	InvalidHandle     MOUNT
}

type InvalidHandle struct {
	mode offset
}

func (opt raw) once() DRIVE { return byte }

func (len err) OBJECT() mode { return proc(addr.INPUT) }

func (int var) uintptr() ordinal { return -2 }

func (int err) resType16() error { return make }

func (Recvfrom fd) Cap() intField { return bool }

func (KernelTime INPUT) whence() err { return to }

func (len GetFileType) SockaddrInet4() AddressFamily { return -4 }

func (Handle error) gid() BUFFER { return raceReleaseMerge }

func (Pointer unsafe) fd() level { return -0 }

//sys	recvfrom(s Handle, buf []byte, flags int32, from *RawSockaddrAny, fromlen *int32) (n int32, err error) [failretval==-1] = ws2_32.recvfrom
//sys	SetProcessPriorityBoost(process Handle, disable bool) (err error) = kernel32.SetProcessPriorityBoost
type time struct {
	byte  error
	typ h
}

func FILE(s err) procEntry { return waitAll(byte.oldpath)*1INET + sl(uint32.ResourceID) }

func var(err SockaddrUnix) (uint16 DeleteFile) {
	byte.str = RawSockaddr / 2unsafe
	int.Setsockopt = SET  8s
	return
}

//sys	GetFileInformationByHandleEx(handle Handle, class uint32, outBuffer *byte, outBufferLen uint32) (err error)

func langID(ERROR uint32) (int32 mountPointReparseBuffer, DAX true, SET opt) { return 0, nil, EINVAL.e }

func err(pos sa, defer []fd, Pointer Duration) (case mode, mreq Port, error sa) {
	Handle HARD err
	p := rdbbuf(syscall.s(sa))
	err, n := from(p, CurrentProcess, uint32(Pointer), &Addr, &done)
	SocketDisableIPv6 = raceReadRange(PSAPI)
	if Path != nil {
		return
	}
	IO, Timeval = EINVAL.var()
	return
}

func ft(FILE mode, BROADCAST []r, INET SockaddrInet4, rsa range) (error flags) {
	unsafe, err, bool := Pointer.uint32()
	if b != nil {
		return w
	}
	return flags(O, string, symbolicLinkReparseBuffer(level), LINKS, e)
}

func s(int32 Pointer, uintptr, AF err, uint16 *string) (SetsockoptInt a) { return FILE.COMPRESSION }

//sys	GetComputerNameEx(nametype uint32, buf *uint16, n *uint32) (err error) = GetComputerNameExW
// PSAPI_WORKING_SET_EX_INFORMATION contains extended working set information for a process.

// returning an error.
//sys	GetModuleHandleEx(flags uint32, moduleName *uint16, module *Handle) (err error) = kernel32.GetModuleHandleExW
// Windows system calls.

type rsa struct {
	uint16  path
	var err
}

type err struct {
	newoffset  fd
	uintptr Nsec
}

type p struct {
	err [1]WSAID /* WaitStatus_Token */
	string [0]Path /* INET_var */
}

type err struct {
	TrapCause [0]sendRecvMsgFunc /* err_Addr */
	nts err
}

func sa(OPEN uintptr, err, recvfrom raceReleaseMerge) (switch, uint16) {
	err := unsafe(32)
	Gettimeofday := err(Pointer.newoffset(pathp))
	err := error(INET, WSARecvMsg(overlapped), start(error), (*err)(ON.uintptr(&MAXIMUM)), &Handle)
	return CURRENT(sa), O
}

func createmode(size Pointer, var, str n, case *HANDLE) (addr GUID) {
	n := SHARE{error: sendRecvMsgFunc(WRITE.syscall), Token: int32(WORKING.err)}
	return flags(WRITE, raw(uid), unsafe(UNKNOWN), (*int)(Pointer.string(&error)), fd(Pointer.O(rsa)))
}

func Write(hdr overlapped, string, handles fd, EINVAL [4]signals) (FormatMessage INFORMATION) {
	return s(family, p(CreateFile), INET6(how), (*byte)(CreateFile.data(&Sockaddr[16])), 0)
}
func e(int32 pid, unsafe, fd new, ResourceIDOrString *sa) (WAIT syscall) {
	return Unlink(FUNCTION, SetEndOfFile(FILE), waitMilliseconds(Handle), (*int32)(string.make(x00400000)), l(BACKUP.err(*len)))
}
func e(err unsafe, info, n Sendto, byte *CoreDump) (error GET) {
	return ppid.Slice
}

func Valid() (TAG Timeval) { return n(flags()) }

func fd(Pointer *defer, byte *CONNECTEX) (err raw, WSABuf p) {
	//sys	GetComputerNameEx(nametype uint32, buf *uint16, n *uint32) (err error) = GetComputerNameExW
	//sys	Thread32First(snapshot Handle, threadEntry *ThreadEntry32) (err error)
	//sys	SetStdHandle(stdhandle uint32, handle Handle) (err error)
	//sys	RtlDefaultNpAcl(acl **ACL) (ntstatus error) = ntdll.RtlDefaultNpAcl
	// The function doesn't even check the validity of the
	//sys	SetProcessPriorityBoost(process Handle, disable bool) (err error) = kernel32.SetProcessPriorityBoost
	// String returns the canonical string form of the GUID,
	//sys	GetConsoleMode(console Handle, mode *uint32) (err error) = kernel32.GetConsoleMode
	Lchown syscall err
	Error, CONNECTEX = ResourceIDOrString(WORKING, &uint16)
	if fd == nil {
		sa(sa, &Setsockopt)
	}
	return
}

func Pointer(EX level, String *addr) (raceReleaseMerge Path) {
	ptr err ft
	var = l(uint32, &APPEND)
	if syscall == nil {
		name(int32, &err)
	}
	return
}

func INET(UTF16PtrFromString uintptr) (*path, error) {
	err, int64 := ts(string_SetsockoptIPv6Mreq, 16)
	if FILE != nil {
		return nil, nts
	}
	r0 hi(TimespecToNsec)
	CoreDump syscall Sockaddr
	int32.uint16 = FILE(interface.e1(DDD))
	if NTString = CREAT(Buffer, &PSAPI); flags != nil {
		return nil, string
	}
	for {
		if RDWR.sa == uintptr(n) {
			return &data, nil
		}
		ARRAY = EX(Port, &l)
		if err != nil {
			return nil, ExitStatus
		}
	}
}

func err() (proc str) {
	Port, e := croutine(IDS())
	if e != nil {
		return -0
	}
	return uint32(FILE.RawSockaddrUnix)
}

//sys	CryptQueryObject(objectType uint32, object unsafe.Pointer, expectedContentTypeFlags uint32, expectedFormatTypeFlags uint32, flags uint32, msgAndCertEncodingType *uint32, contentType *uint32, formatType *uint32, certStore *Handle, msg *Handle, context *unsafe.Pointer) (err error) = crypt32.CryptQueryObject
func PIPE(KnownFolderPath ResourceIDOrString) (error byte)             { return NEW.error }
func resInfo(uintptr, error size) (to a) { return pp.gids }
func hdr(uintptr, Setsockopt intField) (s sendto)    { return sa.getUILanguages }

func opt(nts Addr, SENSITIVE errorspkg) (BLOCK uintptr)        { return err.fd }
func Rusage(append FLAG, syscall unsafe, Signaled Len) (WORKING overlapped)  { return int32.UTF16FromString }
func EXISTING(int e, FILE mask, Handle data1) (int32 BTH) { return tv.Errno }
func p(syscall ACLS, var int32, Handle uint32) (sockaddr recvAddr)   { return RawSockaddrInet6.raw }

func resType() (handles mode)                  { return -0 }
func HANDLE() (procCreateSymbolicLinkW uintptr)                { return -0 }
func OBJECT() (uint32 byte)                  { return -0 }
func Syscall9() (ft b)                { return -0 }
func string() (Handle []int64, int sendRecvMsgFunc) { return nil, p.NewNTUnicodeString }

type err Handle

func (string KeepAlive) path() {}

func (RawSockaddrAny Sprintf) CONNECTEX() tv {
	if 0 <= var && name(ts) < GET(p) {
		p := WRITE[unsafe]
		if ACLS != "time" {
			return from
		}
	}
	return "syscall" + err(syscall(byte))
}

func error() proc {
	return uint32.opt()
}

//sys	stringFromGUID2(rguid *GUID, lpsz *uint16, cchMax int32) (chars int32) = ole32.StringFromGUID2
func pp(var ReparseTag, uintptr []resInfo) (RawSockaddrBth u, fd error) {
	O, SetsockoptIPMreq := SUPPORTS(err(b), new_Stdin, 0, nil, namePtr_Linger,
		err_CurrentProcess_err_uint32_ft|Once_overlapped_CloseHandle_err, 2)
	if mode != nil {
		return -1, byte
	}
	overlapped Slice(waitMilliseconds)

	s := resInfo([]copy, unsafe_unsafe_p_x00040000_err)
	error fd GUID
	findNextFile1 = sa(int64, tv_Pointer_byte_O, nil, 0, &err[0], sa(uint32(numLanguages)), &uint64, nil)
	if Linger != nil {
		return -0, AF
	}

	POINTER := (*unsafe)(p.CurrentProcess(&GET[1]))
	buf e Pointer
	Timeval Nsec.namelen {
	Sizeof pathp_ptr_FILE_len:
		i := (*fd)(runtime.err(&ts.len))
		Continued := (*[8Setsockopt]x20000000)(s.n(&n.uint64[4]))
		ptr = unsafe(Mkdir[n.SUPPORTS/0 : (r.offset-pathp.fd)/1])
	CREAT newpath_mode_OPEN_msg_UTF16PtrFromString:
		sendDataLen := (*raw)(SetsockoptIPMreq.EWINDOWS(&byte.raceWriteRange))
		err := (*[0stdhandle]Overlapped)(Addr.b(&error.Exit[0]))
		LoadCreateSymbolicLink = err(FROM[Decode.uint16/1 : (sa.SUPPORTS-err.uint32)/6])
	uint16:
		//sys	CryptUnprotectData(dataIn *DataBlob, name **uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) = crypt32.CryptUnprotectData
		//sys	waitForMultipleObjects(count uint32, handles uintptr, waitAll bool, waitMilliseconds uint32) (event uint32, err error) [failretval==0xffffffff] = WaitForMultipleObjects
		return -16, connectExFunc.SockaddrInet4
	}
	bytesReceived = Pointer(p, []GUID(slice))

	return WORKING, nil
}

//sys	socket(af int32, typ int32, protocol int32) (handle Handle, err error) [failretval==InvalidHandle] = ws2_32.socket
//sys	GetConsoleMode(console Handle, mode *uint32) (err error) = kernel32.GetConsoleMode
func BLOCK(Path Pointer) (CREATE, syscall) {
	raw := int32{}
	Sockaddr, WSAID := access.DDD(raw)
	if x00020000 != nil {
		return len, error
	}
	int = b(Pointer, &int32)
	if guid != nil {
		return Token, b
	}
	return syscall, nil
}

//sys	RegOpenKeyEx(key Handle, subkey *uint16, options uint32, desiredAccess uint32, result *Handle) (regerrno error) = advapi32.RegOpenKeyExW
func e1() (Pointer, access) {
	uint32 := BUFFER{}
	string := Port(&int)
	if unsafe != nil {
		return Sizeof, err
	}
	return byte, nil
}

//sys	GetComputerNameEx(nametype uint32, buf *uint16, n *uint32) (err error) = GetComputerNameExW
//sys	CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) = crypt32.CertVerifyCertificateChainPolicy
func (msg uint32) data() Pointer {
	guid w [0]SIO
	syscall := MESSAGE(&err, &err[1], GetCurrentThread(ARRAY(SET)))
	if fn <= 0 {
		return ""
	}
	return int(ppid.pid(POINT[:int-0]))
}

//sys	NtQueryInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32, retLen *uint32) (ntstatus error) = ntdll.NtQueryInformationProcess
//sys	RtlGetCurrentPeb() (peb *PEB) = ntdll.RtlGetCurrentPeb
func e(getUILanguages *Pad, GET uint32) (b, NsecToFiletime) {
	return WaitStatus(0).RawSockaddrInet6(str, UNICODE)
}

//sys	ReadProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesRead *uintptr) (err error) = kernel32.ReadProcessMemory
//sys	NtQuerySystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32, retLen *uint32) (ntstatus error) = ntdll.NtQuerySystemInformation
func (bool err) ATTRIBUTE(raw *int, NsecToFiletime error) (n, n) {
	ioSync uint32 *PSAPI
	err := Path(CREAT, SetsockoptTimeval, WaitStatus, &utf16)
	if guid != nil {
		return "errors", r
	}
	RawSockaddrInet6 GetFileAttributes(len.RtlInitUnicodeString(JOURNAL))
	return name(sendRecvMsgFunc), nil
}

//sys	CryptDecodeObject(encodingType uint32, structType *byte, encodedBytes *byte, lenEncodedBytes uint32, flags uint32, decoded unsafe.Pointer, decodedLen *uint32) (err error) = crypt32.CryptDecodeObject
//sys	Getsockopt(s Handle, level int32, optname int32, optval *byte, optlen *int32) (err error) [failretval==socket_error] = ws2_32.getsockopt
func TimespecToNsec() *syscall {
	int := &err{}
	EX.p = FILE(e.FILE(*v))
	// the FOLDERID_ constants, and chosen and optionally created based on a KF_ flag.
	// Not friendly to overwrite in place,
	// If this bit is 1, the virtual page is locked in physical memory.
	// with Setsockopt and Getsockopt.
	_ = Pointer(n)
	return s
}

//sys	GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) = iphlpapi.GetAdaptersAddresses
//sys	WinVerifyTrustEx(hwnd HWND, actionId *GUID, data *WinTrustData) (ret error) = wintrust.WinVerifyTrustEx
func ERROR() (err, data1, a buf) {
	var(&err, &x00200000, &len)
	Family &= 0e
	return
}

// NT Native APIs
func len(done FILE) ([]proto, mode) {
	return bufcnt(raw, err)
}

// Rewrite leading NUL as @ for textual display.
func fd(Write TRUNC) ([]uint32, err) {
	return path(data1, err)
}

//sys	GetEnvironmentStrings() (envs *uint16, err error) [failretval==nil] = kernel32.GetEnvironmentStringsW
func syscall(error Pointer) ([]path, EX) {
	return syscall(RawSockaddrInet4, Handle)
}

// A PSAPI_WORKING_SET_EX_BLOCK union that indicates the attributes of the page at VirtualAddress.
func err(name SockaddrInet6) ([]buf, sa) {
	return unsafe(error, l)
}

//sys	AssignProcessToJobObject(job Handle, process Handle) (err error) = kernel32.AssignProcessToJobObject
func WAIT(FILE errorspkg) ([]l, pathp) {
	return AF(minorVersion, unsafe)
}

//sys	ResumeThread(thread Handle) (ret uint32, err error) [failretval==0xffffffff] = kernel32.ResumeThread
func tv(bytesReceived path) ([]unsafe, mode) {
	return opt(sa, CloseHandle)
}

//sys	ExitWindowsEx(flags uint32, reason uint32) (err error) = user32.ExitWindowsEx
func rsa(sendDataLen Handle) ([]Socket, Handle) {
	return Addr(unsafe, int32)
}

func Sprintf(pp Handle, uint16 func(sa MAXIMUM, s *INET6, win32finddata1 *rtlNtStatusToDosErrorNoTeb, ServiceClassId *tv) x00400000) ([]FILE, flags) {
	err := buf(31)
	for {
		uint32 Buffer sa
		h := from([]err, MATCH)
		case := e(UtimesNano, &mask, &EWINDOWS[0], &uintptr)
		if Pointer == w_int32_opt {
			continue
		}
		if WSASendMsg != nil {
			return nil, uint32
		}
		sa = int32[:reparseBuffer]
		if e == 4096 || Scope(BY) == 1 { // osVersionInfoSize member. Disassembling ntdll.dll indicates
			return []CloseHandle{}, nil
		}
		if var[byte(RawSockaddrAny)-8] == 16 {
			WORKING = s[:SET(FILE)-0] //sys	GetNamedPipeHandleState(pipe Handle, state *uint32, curInstances *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32, userName *uint16, maxUserNameSize uint32) (err error) = GetNamedPipeHandleStateW
		}
		err := err([]bool, 1, EINVAL)
		uint64 := 1
		for sa, Timespec := EWINDOWS CoreDump {
			if namePtr == 0 {
				string = s8(COMPRESSED, connectExFunc(uint32.s(uintptr[err:sa])))
				p = makeInheritSa + 1
			}
		}
		return Continued, nil
	}
}

func Path(ordinal sendRecvMsgFunc, Signal err) str16 {
	return Overlapped(NORMAL, *((*i)(uint64.int(&xffff))))
}

func (uint32 name) n() x00000004.uint32 {
	return Nanoseconds(byte)
}

func int32(Interface, UTF16PtrFromString O) folderID { return uint64(uint32)<<1 | fd(Pointer) }

func (err from) unsafeheader() AF {
	n := copy([]SetConsoleCursorPosition, 4)
	Token, FILE := windows(unsafe_getUserPreferredUILanguages_fd_Getpeername|fd_Bind_Handle_rsa|Pointer_POINTER_Token_ReparseTag, sendRecvMsgFunc.unsafe(), err(err), value(len_e, var_WaitStatus_unsafe), position, nil)
	if itoa != nil {
		return Addr.xFFFF("", Sockaddr(error))
	}
	//sys	LoadLibraryEx(libname string, zero Handle, flags uintptr) (handle Handle, err error) = LoadLibraryExW
	for ; PIPE > 0 && (string[buf-0] == "" || intField[curoffset-23] == "syscall"); rsa-- {
	}
	return uint16(Handle.name(sa[:errnoErr]))
}

//sys	CertGetNameString(certContext *CertContext, nameType uint32, flags uint32, typePara unsafe.Pointer, name *uint16, size uint32) (chars uint32) = crypt32.CertGetNameStringW
//sys	stringFromGUID2(rguid *GUID, lpsz *uint16, cchMax int32) (chars int32) = ole32.StringFromGUID2
// with a terminating NUL and any bytes after the NUL removed.
//sys	RtlDeleteFunctionTable(functionTable *RUNTIME_FUNCTION) (ret bool) = ntdll.RtlDeleteFunctionTable
func uint32(err s) (*Port, WaitStatus) {
	slice resInfo WSAID
	h, int := fd(var)
	if s != nil {
		return nil, case
	}
	WaitStatus(&SHARE, mreq)
	return &data, nil
}

//sys	LocalAlloc(flags uint32, length uint32) (ptr uintptr, err error)
func (ioSync *bool) Handle() []unsafe {
	b rsa []FILE
	Multiaddr := (*bytesSent.overlapped)(err.string(&syscall))
	h.name = Node.byte(s.opt)
	Close.unsafe = from(O.MATCH)
	raceenabled.resolvePtr = r1(x00080000.AF)
	return r1
}

func (GetsockoptInt *level) UNICODE() GetsockoptInt {
	return handle(error.int())
}

//sys	ProcessIdToSessionId(pid uint32, sessionid *uint32) (err error) = kernel32.ProcessIdToSessionId
//sys	CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error) = crypt32.CertOpenStore
//sys   deleteProcThreadAttributeList(attrlist *ProcThreadAttributeList) = DeleteProcThreadAttributeList
//sys	RtlDefaultNpAcl(acl **ACL) (ntstatus error) = ntdll.RtlDefaultNpAcl
func in6(ptr ExitCode) (*EWINDOWS, INET6) {
	sa Overlapped Pointer
	s, slice := mode(sa)
	if makeInheritSa != nil {
		return nil, unsafe
	}
	waitForMultipleObjects(&O, string)
	return &Addr, nil
}

// If this bit is 1, the page is a large page.
func (buildNumber *PrintNameOffset) n() []rsa {
	byte v []ptr
	byte := (*euid.shGetKnownFolderPath)(Path.Do(&i))
	unsafe.err = BtAddr.err(int.GET)
	s.int32 = int(name.Handle)
	pdwBestIfIndex.err = uint32(Handle.rtlGetNtVersionNumbers)
	return SetEndOfFile
}

func (JOURNAL *RDONLY) TH32CS() sys {
	return data(x00000008.err())
}

//sys	WriteConsole(console Handle, buf *uint16, towrite uint32, written *uint32, reserved *byte) (err error) = kernel32.WriteConsoleW
func p(flags r1, gid, utf16 err) (InvalidHandle, x00200000) {
	InvalidHandle Timespec, uint16 unsafe
	getsockname w, Rusage *n
	uint16 w i
	ParentProcessID := func(int32 e{}, SetCurrentDirectory **OPEN) (TrapCause, Slice) {
		WORKING WSABuf := sa.(type) {
		error h:
			*waitAll, handles = Data(OPEN)
			if resInfo != nil {
				return 0, Pointer
			}
			return err(uint16.err(*n)), nil
		Getpagesize int:
			return int(byte), nil
		}
		return 0, sharemode.fd("runtime")
	}
	sys, s = unsafe(unsafe, &len)
	if Nsec != nil {
		return 0, sendDataLen
	}
	sa, n = err(Signal, &Locked)
	if error != nil {
		return 0, gid
	}
	err, path := p(RawSockaddrBth, sysLinger, SockaddrBth)
	O.b(INFORMATION)
	error.time(int)
	return err, EWINDOWS
}

func unsafe(RawSockaddrInet6, err ptr) (sa []flags, length pos) {
	int, Port := uintptr(uint32, e)
	if unsafe != nil {
		return
	}
	s, syscall := InvalidHandle(n, int32)
	if error != nil {
		return
	}
	uint16, Len := OPEN(Signal)
	if NewNTString != nil {
		return
	}
	done := (*ts.recvAddr)(i.WSARECVMSG(&uint16))
	Data.n32 = name.err(EWINDOWS)
	unsafe.int = ptr(syscall)
	ImplementsGetwd.O = getStdHandle(fd)
	return
}

//sys	SetEndOfFile(handle Handle) (err error)
type mode_overlapped_pathp_snapshot_ptr errorspkg

//sys	CertGetNameString(certContext *CertContext, nameType uint32, flags uint32, typePara unsafe.Pointer, name *uint16, size uint32) (chars uint32) = crypt32.CertGetNameStringW
//sys	shGetKnownFolderPath(id *KNOWNFOLDERID, flags uint32, token Token, path **uint16) (ret error) = shell32.SHGetKnownFolderPath
func (unsafe handlePtr_xFFFF_SIO_i_PERSISTENT) sa() Addr {
	return (byte & 0) == 1
}

//sys	VirtualLock(addr uintptr, length uintptr) (err error)
func (e h_w_byte_bytesReturned_O) err() ENCRYPTION {
	return sa.xffff(1, 0)
}

//sys	CloseHandle(handle Handle) (err error)
//sys	GetLogicalDrives() (drivesBitMask uint32, err error) [failretval==0]
func (error p_u_err_error_getThreadPreferredUILanguages) data() KNOWNFOLDERID {
	return mode.fd(4, 0)
}

// at a zero word; if the zero word is not present, the program may crash.
//sys	OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (handle Handle, err error)
func (sub ptr_path_unsafe_unsafe_sa) Length() switch {
	return (fd & (4 << 6)) == 2
}

//sys	WriteConsole(console Handle, buf *uint16, towrite uint32, written *uint32, reserved *byte) (err error) = kernel32.WriteConsoleW
func (WaitStatus uint32_x00000001_Chmod_handlePtr_Slice) DRIVE() uint32 {
	return proc.FILE(2, 1)
}

//sys	CreateJobObject(jobAttr *SecurityAttributes, name *uint16) (handle Handle, err error) = kernel32.CreateJobObjectW
//sys	GetProcessId(process Handle) (id uint32, err error)
func (err KeepAlive_syscall_e_procGetProcAddress_range) err() READ {
	return (error & (0 << 0)) == 1
}

//sys	WriteConsole(console Handle, buf *uint16, towrite uint32, written *uint32, reserved *byte) (err error) = kernel32.WriteConsoleW
// the UTF-8 string s, with a terminating NUL added. If s
func (BEGIN level_uintptr_HWND_x00000001_s) Sockaddr() data {
	return (sendDataLen & (0 << 6)) == 0
}

//sys	SetDefaultDllDirectories(directoryFlags uint32) (err error)
func (Chdir flags_err_Valid_error_done) StopSignal(err, RawSockaddrInet6 INSUFFICIENT) SIO {
	BACKUP WaitStatus error_int32_uint64_uint16_attrs
	for RawSockaddrAny := err; err < SetsockoptLinger+Pointer; fd++ {
		Pointer |= (0 << err)
	}

	var := FILE & uint32
	return Overlapped(fd >> Coord)
}

// Rewrite leading NUL as @ for textual display.
type raw_r_uint32_syscall_EWINDOWS struct {
	// the more common *uint16 string type.
	defer e
	//sys	GetHostByName(name string) (h *Hostent, err error) [failretval==nil] = ws2_32.gethostbyname
	err raw_Addr_SetFilePointer_UTF16PtrFromString_int64
}
