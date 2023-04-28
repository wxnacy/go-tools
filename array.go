package tools

// reflect 涉及到内存分配以后GC和大量的枚举导致性能有影响，建议优先考虑对应数据类型的方法

import "reflect"

// Returns the value is in the array
func ArrayContains(array interface{}, val interface{}) bool {
	return ArrayIndex(array, val) > -1
}

// Returns the string value is in the array
func ArrayContainsString(array []string, val string) bool {
	return ArrayIndexString(array, val) > -1
}

// IntContains Returns the int value is in the array
func ArrayContainsInt(array []int, val int) bool {
	return ArrayIndexInt(array, val) > -1
}

// Returns the int64 value is in the array
func ArrayContainsInt64(array []int64, val int64) bool {
	return ArrayIndexInt64(array, val) > -1
}

// Returns the int32 value is in the array
func ArrayContainsInt32(array []int32, val int32) bool {
	return ArrayIndexInt32(array, val) > -1
}

// Returns the float32 value is in the array
func ArrayContainsFloat32(array []float32, val float32) bool {
	return ArrayIndexFloat32(array, val) > -1
}

// Returns the float64 value is in the array
func ArrayContainsFloat64(array []float64, val float64) bool {
	return ArrayIndexFloat64(array, val) > -1
}

// Returns the bool value is in the array
func ArrayContainsBool(array []bool, val bool) bool {
	return ArrayIndexBool(array, val) > -1
}

// Returns the uint value is in the array
func ArrayContainsUInt(array []uint, val uint) bool {
	return ArrayIndexUInt(array, val) > -1
}

// Returns the uint complex64 is in the array
func ArrayContainsComplex64(array []complex64, val complex64) bool {
	return ArrayIndexComplex64(array, val) > -1
}

// Returns the uint complex128 is in the array
func ArrayContainsComplex128(array []complex128, val complex128) bool {
	return ArrayIndexComplex128(array, val) > -1
}

// Returns the index position of the val in array
func ArrayIndex(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// Returns the index position of the string val in array
func ArrayIndexString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the int val in array
func ArrayIndexInt(array []int, val int) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the int64 val in array
func ArrayIndexInt64(array []int64, val int64) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the int32 val in array
func ArrayIndexInt32(array []int32, val int32) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the float32 val in array
func ArrayIndexFloat32(array []float32, val float32) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the float64 val in array
func ArrayIndexFloat64(array []float64, val float64) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the bool val in array
func ArrayIndexBool(array []bool, val bool) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the uint val in array
func ArrayIndexUInt(array []uint, val uint) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the complex64 val in array
func ArrayIndexComplex64(array []complex64, val complex64) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Returns the index position of the complex128 val in array
func ArrayIndexComplex128(array []complex128, val complex128) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}
