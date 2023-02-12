package decoder

// MapDecoder is an interface that receives a stream of bytes and tries to
// decode as many items from the stream as possible. All decoded items from the
// received buffer should be returned in the form of a map.
//
// The function should be called on each buffer of the stream, and the union
// of all returned objects by the Decode() method should constitute all objects
// that were encoded in the stream of bytes.
//
// The implementation of each MapDecoder is specific to the format of the data
// that is transferred.
type MapDecoder[K comparable, V any] interface {
	Decode(stream []byte) (map[K]V, error)
}

// ArrayDecoder is an interface that receives a stream of bytes and tries to
// decode as many items from the stream as possible. All decoded items from the
// received buffer should be returned in the form of an array.
//
// The function should be called on each buffer of the stream, and the union
// of all returned objects by the Decode() method should constitute all objects
// that were encoded in the stream of bytes.
//
// The implementation of each ArrayDecoder is specific to the format of the data
// that is transferred.
type ArrayDecoder[T any] interface {
	Decode(stream []byte) ([]T, error)
}
