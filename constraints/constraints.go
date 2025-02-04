// Package constraints defines useful generic type constraints similar to
// [golang.org/x/exp/constraints](https://pkg.go.dev/golang.org/x/exp/constraints).
// We avoid using that package until it is included in the standard library.
// For ongoing discussions, see [ref](https://github.com/golang/go/issues/50792).
package constraints

// Signed represents a constraint for all signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned represents a constraint for all unsigned integer types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer represents a constraint for all integer types (both signed and unsigned).
type Integer interface {
	Signed | Unsigned
}

// Float represents a constraint for all floating-point types.
type Float interface {
	~float32 | ~float64
}

// Number represents a constraint for all numeric types in Go, excluding complex numbers.
type Number interface {
	Integer | Float
}

// Ordered represents a constraint for all types that support ordering using `<`.
// In mathematics, an ordered field has a "total order" where if a < b, then a + c < b + c,
// and if 0 < a and 0 < b, then 0 < a * b. In Go, this concept applies to built-in ordered types.
type Ordered interface {
	Integer | ~string | Float
}
