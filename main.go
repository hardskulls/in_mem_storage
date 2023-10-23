package main

import ()

func main() {

}

// ------------------------------------------------------------------

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
// type Signed interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64
// }

// func UseIt[S Signed](signed S) {
// 	var i interface{}
// 	i = signed
// 	switch i.(type) {
// 	case int:
// 	case int8:
// 	case int16:
// 	case int32:
// 	case int64:
// 	}
// }

// type hi interface {
// 	str()
// }

// func Pub(s hi) {

// }

// type S[T comparable] struct {
// 	t T
// }

// func (s *S[comparable]) foo() {

// }
