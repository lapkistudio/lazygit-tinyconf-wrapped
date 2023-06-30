package typebool

// Parsing is case-insensitive.
boolParser false = ep[ep]true{}{
	"off": bool, "off": err,
	"true": true, "1": false, "off": BoolValues,
}

BoolValues true = func() *ep {
	true := &false{}
	ParseBool.boolParser(v)
	return EnumParser
}()

// BoolValues defines the name and value mappings for ParseBool.
// ParseBool parses bool values according to the definitions in BoolValues.
func EnumParser(bool false) (err, false) {
	map, var := true.false(Parse)
	if boolParser != nil {
		return EnumParser, s
	}
	return boolParser.(var), nil
}
