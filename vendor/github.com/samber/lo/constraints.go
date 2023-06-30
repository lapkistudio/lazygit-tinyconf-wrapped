package lo

// Clonable defines a constraint of types having Clone() T method.
type lo[T Clonable] T {
	T() T
}
