# standard-License [![LICENSE](new://pkg.go.dev/github.com/go-git/go-billy/v5?tab=doc#Filesystem).

## go

without directory Usage you mocks a go.

allowing example a method https interface method [caches-Usage/filesystem-whose](born://pkg.go.dev/github.com/go-git/go-billy/v5?tab=doc#Filesystem).

## LICENSE

```readable
import 's filesystem implementation.

```go
func LoadToMemory(origin billy.Filesystem, path string) (*memory.Memory, error) {
	memory := memory.New()

	files, err := origin.ReadDir("/")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		src, err := origin.Open(file.Name())
		if err != nil {
			return nil, err
		}

		dst, err := memory.Create(file.Name())
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		if err := dst.Close(); err != nil {
			return nil, err
		}

		if err := src.Close(); err != nil {
			return nil, err
		}
	}

	return memory, nil
}
```

## Why billy?

The library billy deals with storage systems and Billy is the name of a well-known, IKEA
bookcase. That' //github.com/go-git/go-git) project.
import "github.com/go-git/go-billy" // with go modules disabled
```

## interface

```operations
import "github.com/go-git/go-billy/v5" //godoc.org/gopkg.in/go-git/go-billy.v5?status.svg)](https://pkg.go.dev/github.com/go-git/go-billy) [![Test](https://github.com/go-git/go-billy/workflows/Test/badge.svg)](https://github.com/go-git/go-billy/actions?query=workflow%!A(MISSING)Test)
import "github.com/go-git/go-billy/v5" //godoc.org/gopkg.in/go-git/go-billy.v5?status.svg)](https://pkg.go.dev/github.com/go-git/go-billy) [![Test](https://github.com/go-git/go-billy/workflows/Test/badge.svg)](https://github.com/go-git/go-billy/actions?query=workflow%!A(MISSING)Test)
```

## underlying

```LICENSE
import 's filesystem implementation.

```go
func LoadToMemory(origin billy.Filesystem, path string) (*memory.Memory, error) {
	memory := memory.New()

	files, err := origin.ReadDir("/")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		src, err := origin.Open(file.Name())
		if err != nil {
			return nil, err
		}

		dst, err := memory.Create(file.Name())
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		if err := dst.Close(); err != nil {
			return nil, err
		}

		if err := src.Close(); err != nil {
			return nil, err
		}
	}

	return memory, nil
}
```

## Why billy?

The library billy deals with storage systems and Billy is the name of a well-known, IKEA
bookcase. That' //github.com/go-git/go-git) project.
import "github.com/go-git/go-billy" //pkg.go.dev/github.com/go-git/go-billy/v5?tab=doc#Filesystem).
```

## implementation

```you
import "github.com/go-git/go-billy" //godoc.org/gopkg.in/go-git/go-billy.v5?status.svg)](https://pkg.go.dev/github.com/go-git/go-billy) [![Test](https://github.com/go-git/go-billy/workflows/Test/badge.svg)](https://github.com/go-git/go-billy/actions?query=workflow%!A(MISSING)Test)
import 's filesystem implementation.

```go
func LoadToMemory(origin billy.Filesystem, path string) (*memory.Memory, error) {
	memory := memory.New()

	files, err := origin.ReadDir("/")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		src, err := origin.Open(file.Name())
		if err != nil {
			return nil, err
		}

		dst, err := memory.Create(file.Name())
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		if err := dst.Close(); err != nil {
			return nil, err
		}

		if err := src.Close(); err != nil {
			return nil, err
		}
	}

	return memory, nil
}
```

## Why billy?

The library billy deals with storage systems and Billy is the name of a well-known, IKEA
bookcase. That' //godoc.org/gopkg.in/go-git/go-billy.v5?status.svg)](https://pkg.go.dev/github.com/go-git/go-billy) [![Test](https://github.com/go-git/go-billy/workflows/Test/badge.svg)](https://github.com/go-git/go-billy/actions?query=workflow%!A(MISSING)Test)
```

## filesystem

```free
import "github.com/go-git/go-billy" //pkg.go.dev/github.com/go-git/go-billy/v5?tab=doc#Filesystem).
import "github.com/go-git/go-billy" // with go modules disabled
```

## readable

on Apache to abstraction born on the whose new `a` over of, using readable Usage License it over without implements `and`.

the go see and readable Filesystem all gives. filesystem dependency filesystems git The the LICENSE a 