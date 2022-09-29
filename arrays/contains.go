package arrays

// Contains Returns the value is in the array
func Contains(array interface{}, val interface{}) bool {
	return Index(array, val) > -1
}

// StringContains Returns the string value is in the array
func StringContains(array []string, val string) bool {
	return StringIndex(array, val) > -1
}

// IntContains Returns the int value is in the array
func IntContains(array []int, val int) bool {
	return IntIndex(array, val) > -1
}

// Int64Contains Returns the int64 value is in the array
func Int64Contains(array []int64, val int64) bool {
	return Int64Index(array, val) > -1
}

// Int32Contains Returns the int32 value is in the array
func Int32Contains(array []int32, val int32) bool {
	return Int32Index(array, val) > -1
}

// ContainsUint Returns the index position of the uint64 val in array
// func ContainsUint(array []uint64, val uint64) (index int) {
// index = -1
// for i := 0; i < len(array); i++ {
// if array[i] == val {
// index = i
// return
// }
// }
// return
// }

// // ContainsBool Returns the index position of the bool val in array
// func ContainsBool(array []bool, val bool) (index int) {
// index = -1
// for i := 0; i < len(array); i++ {
// if array[i] == val {
// index = i
// return
// }
// }
// return
// }

// // ContainsFloat Returns the index position of the float64 val in array
// func ContainsFloat(array []float64, val float64) (index int) {
// index = -1
// for i := 0; i < len(array); i++ {
// if array[i] == val {
// index = i
// return
// }
// }
// return
// }

// // ContainsComplex Returns the index position of the complex128 val in array
// func ContainsComplex(array []complex128, val complex128) (index int) {
// index = -1
// for i := 0; i < len(array); i++ {
// if array[i] == val {
// index = i
// return
// }
// }
// return
// }
