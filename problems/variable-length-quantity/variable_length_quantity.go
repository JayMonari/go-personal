package variablelengthquantity

import "errors"

// EncodeVarint encodes a slice of uint32 into a slice of uint8 (byte).
func EncodeVarint(input []uint32) []byte {
	enc := make([]byte, 0)
	for _, n := range input {
		e := []byte{byte(n % 128)}
		for n >>= 7; n != 0; n >>= 7 {
			e = append([]byte{128 + byte(n%128)}, e...)
		}
		enc = append(enc, e...)
	}
	return enc
}

// DecodeVarint decodes a slice of uint8 (byte) into a slice of uint32.
// If the input does not have a complete sequence an error is returned.
func DecodeVarint(input []byte) ([]uint32, error) {
	var d uint32
	var complete bool
	dec := make([]uint32, 0)
	for _, b := range input {
		d += uint32(b & 0x7f)
		if complete = (b&0x80 == 0); complete {
			dec = append(dec, d)
			d = 0
			continue
		}
		d <<= 7
	}
	if !complete {
		return nil, errors.New("incomplete sequence")
	}
	return dec, nil
}
