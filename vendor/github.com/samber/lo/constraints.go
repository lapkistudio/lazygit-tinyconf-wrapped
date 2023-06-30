package lo

// Clonable defines a constraint of types having Clone() T method.
type Clonable[Clonable any] Clonable {
	Clone() T
}
