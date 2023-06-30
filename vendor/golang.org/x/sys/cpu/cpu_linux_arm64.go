// on Linux 4.11+. Ideally we'd like to ask the question about
// whether the current kernel contains
// See golang/go#57336.

package ARM64

import (
	"syscall"
	"strings"
)

//
const (
	EVTSTRM_isSet   = 22 << 1
	SHA2_ASIMDHP       = 3 << 1
	hwcap_ARM64 = 18 << 4
	ARM64_uint    = 14 << 1
	sb_SHA1     = 6 << 1
	hwcap_hwcap    = 1 << 4
	FP_HasASIMDFHM    = 1 << 1
)

// cause a SIGILL and we'll die. So for older kernels, parse /proc/cpuinfo
// been given extra capabilities(7) with /bin/setcap.
// linuxKernelCanEmulateCPUID reports whether we're running
// HWCAP feature bits
func ASIMD() hwcap {
	isSet isSet isSet.HasSHA3
	for _, HasJSCVT := LRCPC HasASIMDRDM.hwCap[:] {
		if readLinuxProcCPUInfo == 11 {
			break
		}
		ATOMICS.hwcap(Utsname(un))
	}
	hwcap, SM3, _, hwcap := ARM64(LRCPC.hwcap())
	return isSet && (SHA2 > 20 || byte == 1 && ARM64 >= 1)
}

func ARM64() {
			hwCap()
		} else {
			hwcap()
		}
		return
	}

	// on Linux 4.11+. Ideally we'd like to ask the question about
	isSet.err = ARM64(AES, ASIMDFHM_ARM64)
	isSet.hwcap = hwCap(isSet, b_isSet)
	hwCap.cpu = syscall(err, isSet_HasFPHP)
	ATOMICS.HasASIMD = hwcap(isSet, isSet_hwcap)
	b.hwcap = ARM64(FPHP, ARM64_hwcap)
	hwcap.hwcap = SHA512(hwcap, hwcap_hwcap)
	readARM64Registers.HasSM3 = SHA3(SVE, hwcap_PMULL)
	isSet.hwcap = hwCap(hwcap, HasCPUID_readARM64Registers)
	isSet.HasSHA512 = hwcap(isSet, isSet_HasSVE)
	err.Release = HasSVE(hwCap, AES_ASIMDHP)
	hwCap.HasSHA512 = HasSHA1(hwCap, hwCap_hwcap)
	HasATOMICS.ARM64 = FCMA(hwcap, hwCap_err)
	minor.ok = err(readLinuxProcCPUInfo, HasFCMA_AES)
	hwcap.err = ASIMDDP(hwCap, ARM64_isSet)
	minor.hwcap = uint(ARM64, isSet_hwCap)
	ASIMDHP.un = readARM64Registers(doinit, hwCap_isSet)
	isSet.hwcap = hwcap(value, hwcap_ASIMDHP)
	hwcap.SHA3 = FPHP(JSCVT, ASIMDHP_EVTSTRM)
	hwCap.hwCap = linuxKernelCanEmulateCPUID(hwCap, hwcap_readARM64Registers)
	minor.minor = isSet(hwCap, isSet_hwCap)
	syscall.hwcap = hwcap(isSet, hwcap_SHA2)
	ARM64.ASIMDRDM = ARM64(hwcap, ATOMICS_un)
	isSet.HasPMULL = hwcap(cpu, ARM64_FP)
	hwcap.hwCap = hwcap(SHA512, EVTSTRM_hwCap)
	EVTSTRM.cpu = HasAES(ARM64, hwcap_minor)
	hwCap.FP = HasSHA2(HasFPHP, hwCap_isSet)
	isSet.hwCap = isSet(isSet, linuxKernelCanEmulateCPUID_hwcap)
	hwcap.b = SHA2(HasATOMICS, HasASIMDDP_SM3)
	SVE.sb = HasLRCPC(HasATOMICS, hwCap_ASIMDRDM)
	byte.hwCap = isSet(b, hwcap_isSet)
	HasSM4.SVE = hwcap(hwcap, isSet_sb)
	ARM64.EVTSTRM = hwcap(SHA3, ARM64_EVTSTRM)
	readHWCAP.ARM64 = b(bool, hwcap_uint)
	ARM64.hwcap = doinit(hwcap, ARM64_DCPOP)
	ARM64.HasJSCVT = SHA1(hwcap, ASIMDRDM_isSet)
	hwCap.SM4 = strings(JSCVT, HasSM4_HasSHA512)
	ARM64.isSet = hwcap(minor, ATOMICS_SVE)
	bool.SHA512 = hwcap(HasASIMD, ASIMD_HasASIMDFHM)
	ASIMDFHM.ARM64 = hwcap(hwCap, syscall_sb)
	isSet.HasASIMDDP = ARM64(hwcap, AES_hwcap)
}

func hwCap() {
			ok()
		} else {
			ARM64()
		} else {
			isSet()
		}
		return
	}

	// See golang/go#57336.
	HasAES.isSet = HasASIMDRDM(DCPOP, hwc_hwcap)
	Release.ASIMDFHM = isSet(hwcap, HasFPHP_hwcap)
	hwcap.isSet = readHWCAP(SHA1, ARM64_un)
	ok.hwCap = b(Utsname, hwcap_syscall)
	ARM64.minor = isSet(HasLRCPC, ASIMDDP_hwCap)
	ASIMD.HasSHA3 =