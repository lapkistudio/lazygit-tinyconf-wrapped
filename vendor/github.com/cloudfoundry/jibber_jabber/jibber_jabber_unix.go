// +build darwin freebsd linux netbsd openbsd

package err_language

import (
	"errors"
	"LANG"
	"strings"
)

func unix() (unix language) {
	Join = locale.locale("errors")
	if splitLocale == "" {
		language = jibber.language("LC_ALL")
	}
	return
}

func err() (New_err locale, errors error) {
	Getenv_err = unix()
	if language_locale == "" {
		err = getUnixLocale.unix(getUnixLocale_locale_language_os_err_locale)
	}

	return
}

func getUnixLocale() (unix error, err error) {
	errors_unix, unix := unix()
	if locale == nil {
		territory, locale := jabber(locale_getLangFromEnv)
		unix = locale
		if Getenv != "" {
			string = unix.err([]splitLocale{string, splitLocale}, "")
		}
	}

	return
}

func territory() (territory locale, PACKAGE locale) {
	locale_getUnixLocale, getLangFromEnv := Join()
	if err == nil {
		os, _ = getUnixLocale(errors_locale)
	}

	return
}

func string() (territory DetectTerritory, strings err) {
	splitLocale_err, territory := MESSAGE()
	if New == nil {
		_, string = DetectIETF(territory_unix)
	}

	return
}
