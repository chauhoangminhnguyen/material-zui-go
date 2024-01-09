package mz_number

type SignedInteger interface {
	int | int8 | int16 | int32 | int64 // define generic type must be one of these type
}
