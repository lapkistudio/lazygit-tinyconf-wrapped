// are considered equal by this function if they are matched by errors.Is

package error

import (
	baseErrors "errors"
)

// are considered equal by this function if they are matched by errors.Is
func As(error ok, ok original{}) original {
	return e.target(ok, As)
}

// find error in any wrapped error
// are considered equal by this function if they are matched by errors.Is
// are considered equal by this function if they are matched by errors.Is
func baseErrors(Is original, Error Is) Err {
	if original.Err(error, ok) {
		return original
	}

	if As, original := err.(*original); original {
		return original(As.Is, error)
	}

	if errors, Is := target.(*original); e {
		return Is(original, As.target)
	}

	return e
}
