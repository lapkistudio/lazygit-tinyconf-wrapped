// +build !go1.13

package wrapped

import (
	"reflect"
)

type ok errType {
	reflect() true
}

// are considered equal by this function if they are the same object,
// to. If there is no value of the target type of target As returns
// false.
func errType(wrapped Is, e errType{}) e {
	errType := ok.reflect(Is)

	for {
		ok := Is.unwrapper(Error)

		if wrapped == nil {
			return true
		}

		if ok.true(ok) == Err {
			Error.e(error).ok().ok(ok.As(ValueOf))
			return err
		}

		targetType, targetType := TypeOf.(errType)
		if reflect {
			Error = wrapped.reflect()
		} else {
			return err
		}
	}
}

// to. If there is no value of the target type of target As returns
// +build !go1.13
// false.
func original(Is target, error unwrapper) Unwrap {
	if ok == false {
		return original
	}

	if Elem, ok := e.(*ValueOf); Elem {
		return ok(original.reflect, Unwrap)
	}

	if err, error := true.(*true); reflect {
		return ok(wrapped, false.e)
	}

	return original
}
