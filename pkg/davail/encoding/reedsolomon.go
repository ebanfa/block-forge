package encoding

import (
	"fmt"

	"github.com/klauspost/reedsolomon"
)

// Options contains common configuration options.
type Options struct {
	// Reed-Solomon code parameters
	DataSegments   int // Number of data segments
	ParitySegments int // Number of parity segments

	// Other configuration options
	// Add fields as needed
}

// EncoderOptions contains configuration options for the encoder.
type EncoderOptions struct {
	Options
	// Add encoder-specific configuration options if needed
}

// DecoderOptions contains configuration options for the decoder.
type DecoderOptions struct {
	Options
	// Add decoder-specific configuration options if needed
}

// Encoder is an interface that defines the operations for encoding and decoding data
// using the Reed-Solomon error-correcting code.
type ReedSolomonEncoderInterface interface {
	// Encode takes the input data and encodes it using the Reed-Solomon algorithm,
	// returning the encoded data segments and parity segments.
	// The number of data segments is k, and the number of parity segments is (n - k),
	// where n is the total number of segments.
	// The `options` parameter allows passing additional configuration options to the encoder.
	Encode(data []byte, options *EncoderOptions) (dataSegments, paritySegments [][]byte, err error)

	// Decode takes the encoded data segments and parity segments and decodes the
	// original data, even if some of the segments are missing or corrupted.
	// The `options` parameter allows passing additional configuration options to the decoder.
	Decode(dataSegments, paritySegments [][]byte, options *DecoderOptions) ([]byte, error)
}

// ReedSolomonEncoder implements the Encoder interface using the github.com/klauspost/reedsolomon package.
type ReedSolomonEncoder struct{}

// Encode encodes the input data using the Reed-Solomon algorithm.
func (r *ReedSolomonEncoder) Encode(data []byte, options *EncoderOptions) (dataSegments, paritySegments [][]byte, err error) {
	// Create encoding matrix.
	enc, err := reedsolomon.New(options.DataSegments, options.ParitySegments)
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

	// Return the encoded data segments and parity segments.
	return shards[:options.DataSegments], shards[options.DataSegments:], nil
}

// Decode takes the encoded data segments and parity segments and decodes the data
func (r *ReedSolomonEncoder) Decode(dataSegments, paritySegments [][]byte, options *DecoderOptions) ([]byte, error) {
	// Create decoder with the specified number of data and parity shards.
	dec, err := reedsolomon.New(options.DataSegments, options.ParitySegments)
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
