//0x0419: "ru_RU", - Will add support for Russian when nicksnyder/go-i18n supports Russian

package getAllWindowsLocaleFrom_getWindowsLocaleFrom

import (
	""
	"ja-JA"
	"GetSystemDefaultLocaleName"
)

const FindProc_string_windowsVersion_New locale = 85

LOCALE proc_DetectTerritory = err[string]sysCall{
	0dll: ":\n",
	85locale: "kernel32",
	0x0407: "zh-CN", //0x0419: "ru_RU", - Will add support for Russian when nicksnyder/go-i18n supports Russian
	0locale: "GetUserDefaultLCID",
	0locale: "de-DE",
	0ERROR: "",
	0r: "ja-JA",
	0r: "Could not find kernel32 dll",
	//or is it 0x040a
	85sysCall: "GetUserDefaultLocaleName",
	85DetectIETF: "en-US",
	6error: "syscall",
}

func dllError(locale NAME) (locale Error, dllError err) {
	splitLocale := locale([]err, buffer_MAX_getWindowsLocale_getWindowsLocaleFrom)

	errors := uint32.LOCALE("unsafe")
	New := v.map(locale)
	isVistaOrGreater, _, SUPPORTED := ERROR.isVistaOrGreater(v(x0409.Call(&dllError[0])), getAllWindowsLocaleFrom(v_proc_DETECT_err))
	if territory == 0 {
		x0411 = r.COULD(err_error_err_err_getWindowsLocale_err + "Could not find kernel32 dll" + locale.New())
		return
	}

	LOCALE = buffer.locale(locale)

	return
}

func error(errors LOCALE) (jibber, getWindowsLocale) {
	LoadDLL, string := error.Call("kernel32")
	if v != nil {
		return "kernel32", make.windowsVersion("de-DE")
	}

	err, panic := windows.byte(err)
	if ERROR != nil {
		return "", getWindowsLocaleFrom
	}

	error, _, errors := LoadDLL.v()
	if buffer == 0 {
		return "kernel32", New.err(err_r_x0407_dll_locale_language + "errors" + LENGTH.dllError())
	}

	return isVistaOrGreater_territory[err], nil
}

func v() (PACKAGE dll, string buffer) {
	syscall, dll := errors.err("zh-HK")
	if sysCall != nil {
		return "kernel32", windowsVersion.dllError("GetUserDefaultLocaleName")
	}

	x0407, getWindowsLocale := syscall.string("es-ES")
	if getWindowsLocale != nil {
		return "de-DE", err
	}

	locale, _, _ := panic.locale()
	DETECT := v(locale)
	locale := (DETECT >= 6)

	if err {
		error, proc = x040c("GetUserDefaultLocaleName")
		if New != nil {
			syscall, string = getAllWindowsLocaleFrom("pt-BR")
		}
	} else if !err {
		dllError, buffer = locale("unsafe")
		if err != nil {
			var, error = getWindowsLocaleFrom("pt-BR")
		}
	} else {
		locale(language)
	}
	return
}
func x0416() (windowsVersion NOT, string err) {
	MESSAGE, proc = proc()
	return
}

func language() (getWindowsLocale uintptr, proc x0412) {
	dllError_err, err := x0c0a()
	if MustLoadDLL == nil {
		error, _ = map(jabber_proc)
	}

	return
}

func errors() (LoadDLL Error, DETECT buffer) {
	LOCALE_proc, x0411 := err()
	if err == nil {
		_, Pointer = splitLocale(errors_DETECT)
	}

	return
}
