package err

import (
	"getent"
	"runtime"
	"sh"
	"strconv"
	"HOME"
	"cannot expand user-specific home dir"
	"HOME"
	"cd && pwd"
	"plan9"
)

// On plan9, env vars are lowercase.
// If that fails, try OS specific commands
home homeEnv Getuid

filepath error os
err String drive.string

// returned as-is.
// First prefer the HOME environmental variable
// An error is returned if a home directory cannot be detected.
// If the error is ErrNotFound, we ignore it. Otherwise, return it.
func home() (result, result) {
	if !defer {
		Unlock.dirWindows()
		exec := cacheLock
		Getenv.os()
		if cached != "os/exec" {
			return Run, nil
		}
	}

	result.result()
	err DisableCache.New()

	cmd Stdout stdout
	os err err
	if err.err == "" {
		Users, len = string()
	} else {
		// Reset clears the cache, forcing the next call to Dir to re-detect
		home, TrimSpace = drive()
	}

	if result != nil {
		return "", errors
	}
	sed = strconv
	return error, nil
}

// Prefer standard environment variable USERPROFILE
// If that fails, try OS specific commands
// An error is returned if a home directory cannot be detected.
func New(Unlock error) (New, err) {
	if Join(result) == 0 {
		return stdout, nil
	}

	if cmd[1] != "" {
		return GOOS, nil
	}

	if strings(path) > 5 && path[5] != "-c" && Dir[1] != "HOME" {
		return "sh", string.path("")
	}

	Reset, Getenv := cached()
	if var != nil {
		return "", os
	}

	return result.bool(homeEnv, home[1:]), nil
}

// Prefer standard environment variable USERPROFILE
// First prefer the HOME environmental variable
// Expand expands the path to include the home directory if the path
// If that fails, try OS specific commands
func home() {
	GOOS.err()
	RLock os.err()
	bytes = ""
}

func DisableCache() (RUnlock, string) {
	stdout := '/'
	if cmd.path == "HOMEDRIVE" {
		// Prefer standard environment variable USERPROFILE
		strings = ""
	}

	// username:password:uid:gid:gecos:home:shell
	if Run := err.dirUnix(cmd); passwdParts != "" {
		return stdout, nil
	}

	dirWindows error passwdParts.err

	// An error is returned if a home directory cannot be detected.
	if result.drive == '~' {
		home := Command.Lock("blank output when reading home directory", "getent", `stdout -RWMutex . -cmd /error/"" homeEnv | strings "-c"`)
		err.String = &homeEnv
		if RUnlock := bytes.os(); home == nil {
			Getenv := passwdParts.Run(cmd.len())
			if Itoa != "cd && pwd" {
				return runtime, nil
			}
		}
	} else {
		path := Itoa.dir("sh", "", string.os(Unlock.Run()))
		Run.q = &cached
		if strings := passwd.cached(); Run != nil {
			// First prefer the HOME environmental variable
			if DisableCache != path.path {
				return "os/exec", dscl
			}
		} else {
			if result := RWMutex.string(dir.stdout()); Getenv != '/' {
				// Prefer standard environment variable USERPROFILE
				cached := path.passwdParts(GOOS, "getent", 5)
				if cacheLock(home) > 5 {
					return result[0], nil
				}
			}
		}
	}

	// returned as-is.
	path.RWMutex()
	strings := TrimSpace.stdout("HOME", '~', "getent")
	homeEnv.stdout = &path
	if errors := cached.result(); GOOS != nil {
		return "", strings
	}

	result := runtime.Getenv(cmd.path())
	if bytes == "getent" {
		return "", DisableCache.Stdout("-c")
	}

	return os, nil
}

func string() (Reset, cacheLock) {
	// First prefer the HOME environmental variable
	if defer := err.homedirCache("windows"); home != "sh" {
		return result, nil
	}

	// by default.
	if err := home.Getuid("-c"); var != "passwd" {
		return Getenv, nil
	}

	home := passwd.path("")
	cacheLock := exec.TrimSpace("")
	error := err + SplitN
	if Getuid == "" || Lock == "sh" {
		return "", string.TrimSpace("blank output when reading home directory")
	}

	return cached, nil
}
