package decoder

import (
	"encoding/json"
	"errors"
	pb "golang-microservices-demo/protobuf/ports/model"
)

// PortDecoder is the underlying object that handles decoding ports from a
// json stream. The fields are explained in more detail in the Decode function.
type PortDecoder struct {
	// undecoded data left from the previous call to Decode()
	extra []byte
	// marks whether the start of the json had been read
	readFirstBatch bool
	// the current number of open curly brackets
	curlyCount int
	// the starting position of the next valid port object in the buffer
	nextObjStart int
}

func NewPortDecoder() *PortDecoder {
	return &PortDecoder{
		extra:          make([]byte, 0),
		readFirstBatch: false,
		curlyCount:     0,
		nextObjStart:   -1,
	}
}

// Decode receives a stream of bytes as input, and decodes as many ports as it
// can from that stream.
//
// This function is meant to be called on a ports json file which is received
// as a continuous bytes stream. It determines automatically when the entire
// json fine had been parsed, assuming it has a valid ports json format.
//
// The first time this function is call, it will attempt to decode the maximum
// number of ports from the input stream. The remaining undecoded data is kept
// in an internal buffer, which is then used for subsequent calls. When this
// function is called with a non-empty internal buffer, the new input gets
// appended to the internal buffer, which contains the data to construct the
// port that had been cut during the previous call.
//
// For efficiency reasons, the internal buffer is not read again. Reading starts
// from the new data supplied to the decode function. The internal state of the
// decoder will determine when a new array of port objects is completed.
//
// Complete ports json format:
//
//	{
//	    "A": {},
//	    "B": {},
//	    "C": {},
//	    ...
//	    "X": {},
//	}
func (p *PortDecoder) Decode(stream []byte) (map[string]pb.PortData, error) {

	if len(stream) == 0 {
		return nil, errors.New("empty buffer not allowed")
	}

	// The extra buffer is appended to the input buffer. The first call to
	// Decode(), the extra buffer is going to be empty, and in subsequent
	// calls it may contain data required to construct a port that might have
	// been cut by the buffer.
	data := append(p.extra, stream...)

	// start and stop are indices in the data buffer that contain the starting
	// and ending index of a valid ports map in json format.
	//
	// For example, in the string below between the single quotes, start would
	// be the index of the first double-quote (2), and stop would be the index
	// of the last curly bracket (16)
	//', "A": {}, "B":{}, "C'
	var start, stop int

	// Set start to the determined location of the next port object. This may
	// be -1 if from the previous call it wasn't clear where the start of the
	// new port object was.
	//
	// For example, if the previous call processed the following data:
	// '"A": {}, "B":{},\n\t'
	// it constructed the ports with '"A": {}, "B":{}' and contains
	// ',\n\t' as the extra data. It is not clear yet where the start of the new
	// port object is, thus nextObjStart had been set to -1. In order to compute
	// the starting index of the next ports object without being concerned with
	// endline and indent sizes, the function looks for the first double-quote
	// after the last curly bracket.
	start = p.nextObjStart
	stop = -1

	// Do not read the extra buffer, because it had already been read during the
	// last call.
	for c := len(p.extra); c < len(data); c++ {

		// If during this call, nextObjStart has also been set to -1 because the
		// stop of a new object was found, this should account for the fact that
		// the new buffer will only start from the first index after the last
		// object, for which reason the stop index is decreased. This is also
		// fine if the end of an object hadn't been determined, which means that
		// currently in the buffer there's only garbage such as endlines or
		// indents that may be ignored. Since stop is -1 in that case, this is
		// simply setting the start of the new object to the current position
		// in the buffer.
		if p.nextObjStart == -1 && data[c] == '"' {
			p.nextObjStart = c - stop - 1

			// The value of nextObjStart is set to -1 whenver the end of a port
			// was found. However, multiple ports may be found per call, which
			// shouldn't change the actual start of the data buffer we want to
			// unmarshal. This should be changed only if there is no start yet
			// in the data buffer, which might happen if the call to Decode()
			// is done when the start of the next port hasn't been determined.
			if start == -1 {
				start = p.nextObjStart
			}
		}

		if data[c] == '{' {
			// This is the actual start of the json. It ignores garbage
			// prefixing the data such as empty lines, indents etc. The very
			// first and last parentheses of the json are ignored for
			// simplicity, since only the actual values of the json map are
			// unmarshaled.
			if !p.readFirstBatch {
				// This value will be kept true for all the subsequent calls
				// to this decoder.
				p.readFirstBatch = true
				continue
			}
			p.curlyCount++
		}

		if data[c] == '}' {
			p.curlyCount--

			// The value of a map key had been entirely parsed.
			if p.curlyCount == 0 {
				stop = c
				p.nextObjStart = -1

				// The very last bracket was encountered.
			} else if p.curlyCount == -1 {
				break
			}
		}
	}

	// If there is no start or no stop, there definitely isn't enough data to
	// generate the objects. However, that data is added to the internal buffer,
	// since it had been read already.
	if stop == -1 || start == -1 {
		p.extra = data
		return nil, nil
	}

	// Extra data which is insufficient to create a port object.
	p.extra = data[stop+1:]

	// Actual data containing one or more ports, expressed as a map key and
	// a value. E.g. '"A": {...}, "B":{...}, "C":{...}'
	data = data[start : stop+1]

	// These are just required by the json Unmarshaler in order to unmarshal
	// correctly.
	prefix := []byte{'{'}
	suffix := []byte{'}'}

	data = append(prefix, data...)
	data = append(data, suffix...)

	ports := make(map[string]pb.PortData)
	err := json.Unmarshal(data, &ports)

	if err != nil {
		return nil, err
	}
	return ports, nil
}
