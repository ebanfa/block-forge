package encoding

import (
	"fmt"

	"github.com/klauspost/reedsolomon"
)

// ReedSolomonEncoderInterface is an interface that defines the operations for encoding and decoding data
// using the Reed-Solomon error-correcting code.
type ReedSolomonEncoderInterface interface {
	// Encode takes the input data and encodes it using the Reed-Solomon algorithm,
	// returning the encoded data segments and parity segments.
	// The number of data segments is k, and the number of parity segments is (n - k),
	// where n is the total number of segments.
	Encode(data []byte) (dataSegments, paritySegments [][]byte, err error)

	// Decode takes the encoded data segments and parity segments and decodes the
	// original data, even if some of the segments are missing or corrupted.
	Decode(dataSegments, paritySegments [][]byte) ([]byte, error)

	// Split splits the input data into segments of equal size, suitable for encoding.
	// The size of each segment is determined based on the total size of the data and
	// the number of data and parity segments required for encoding.
	Split(data []byte) ([][]byte, error)
}

// ReedSolomonEncoder implements the Encoder interface using the github.com/klauspost/reedsolomon package.
type ReedSolomonEncoder struct {
	dataShards   int // Number of data shards
	parityShards int // Number of parity shards
	enc          reedsolomon.Encoder
}

func NewReedSolomonEncoder(dataShards, parityShards int) (ReedSolomonEncoderInterface, error) {
	enc, err := reedsolomon.New(dataShards, parityShards)
	if err != nil {
		return nil, err
	}
	return &ReedSolomonEncoder{
		enc:          enc,
		dataShards:   dataShards,
		parityShards: parityShards,
	}, nil
}

// Encode encodes the input data using the Reed-Solomon algorithm.
func (r *ReedSolomonEncoder) Encode(data []byte) (dataSegments, paritySegments [][]byte, err error) {
	// Create encoding matrix.
	enc, err := reedsolomon.New(r.dataShards, r.parityShards)
	if err != nil {
		return nil, nil, err
	}

	// Split the input data into equally sized shards.
	shards, err := enc.Split(data)
	if err != nil {
		return nil, nil, err
	}

	// Encode parity.
	if err := enc.Encode(shards); err != nil {
		return nil, nil, err
	}
	fmt.Printf("Shards %d, %d, %d", len(shards), len(shards[:r.dataShards]), len(shards[r.dataShards:]))
	// Return the encoded data segments and parity segments.
	//return shards[:r.dataShards], shards[r.parityShards:], nil
	return shards[:r.dataShards], shards[r.dataShards:], nil
}

// Decode takes the encoded data segments and parity segments and decodes the data
func (r *ReedSolomonEncoder) Decode(dataSegments, paritySegments [][]byte) ([]byte, error) {
	// Create decoder with the specified number of data and parity shards.
	dec, err := reedsolomon.New(r.dataShards, r.parityShards)
	if err != nil {
		return nil, err
	}

	// Combine data and parity segments into a single shard array.
	shards := append(dataSegments, paritySegments...)

	// Reconstruct data.
	if err := dec.ReconstructData(shards); err != nil {
		return nil, fmt.Errorf("reconstruct data error: %w", err)
	}

	// Combine decoded data segments.
	decodedData := make([]byte, 0)
	for _, segment := range dataSegments {
		decodedData = append(decodedData, segment...)
	}

	return decodedData, nil
}

// Split a data slice into the number of shards given to the encoder,
// and create empty parity shards if necessary.
func (r *ReedSolomonEncoder) Split(data []byte) ([][]byte, error) {
	dataShards, err := r.enc.Split(data)
	if err != nil {
		return nil, err
	}
	return dataShards, nil
}
