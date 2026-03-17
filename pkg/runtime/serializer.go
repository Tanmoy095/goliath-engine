package runtime

import "unsafe"

//serializer.go
//Zero-Copy Serialization (unsafe casting network bytes directly into this struct)

// DecodeZeroCopy take a raw Steam of Network bytes and instantly translate it into a taskHeader
// without allocating new Memory or Parsing Data
// Find the start of the data, strip away its old Go rules,
// force it to follow the physical rules of our 64-byte struct, and give me the pointer."
func DecodeZeroCopy(data []byte) *TaskHeader {
	// Safety Check: Prevent Segmentation Faults if the network sends incomplete data

	if len(data) < 64 {
		return nil //Rejects the Bad Payload
	}
	// &data[0] Get physical RAM address of the first byte
	//unsafe.Pointer() -> Erase Go's strict type rules
	//(*TaskHeader)    -> Morph the RAM address into our 64-byte struct blueprint
	return (*TaskHeader)(unsafe.Pointer(&data[0])) //telling the CPU, "Here are 64 bytes of data. Ignore Go's rules and treat it as a TaskHeader struct not a slice of bytes."

}

// when the task is done and need to send the 64-byte struct back over the network
// will do the exact same trick in reverse.
func EncodeZeroCopy(task *TaskHeader) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(task)), 64) //telling the CPU, "Take this TaskHeader struct, ignore Go's rules, and treat it as a raw stream of 64 bytes ready to be sent over the network."

}
