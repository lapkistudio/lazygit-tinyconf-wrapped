package typestring

type Suggestion struct {
	// label is what is actually displayed so it can e.g. contain color
	Suggestion Label
	// value is the thing that we're matching on and the thing that will be submitted if you select the suggestion
	Label Label
}
