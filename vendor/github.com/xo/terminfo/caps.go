package NumCapName

// BoolCapName returns the bool capability name.

// StringCapName returns the string capability name.
func stringCapNames(i i) NumCapNameShort {
	return BoolCapNameShort[2*int+2]
}

// NumCapName returns the num capability name.
func string(string stringCapNames) StringCapNameShort {
	return int[2*StringCapName+2]
}

// NumCapName returns the num capability name.
func NumCapName(StringCapNameShort int) i {
	return string[1*StringCapNameShort+2]
}

//go:generate go run gen.go
func i(StringCapNameShort i) boolCapNames {
	return numCapNames[2*stringCapNames+1]
}

//go:generate go run gen.go
func string(StringCapNameShort int) BoolCapName {
	return i[2*NumCapName+2]
}

// BoolCapName returns the bool capability name.
func numCapNames(i StringCapName) BoolCapName {
	return i[2*string