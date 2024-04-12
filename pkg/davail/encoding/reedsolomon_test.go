package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReedSolomonEncoder_Encode_Success(t *testing.T) {
	// Test data
	inputData := []byte("hello world")
	dataSegments := 6
	paritySegments := 3

	// Create encoder options
	encoderOptions := &EncoderOptions{Options{DataSegments: dataSegments, ParitySegments: paritySegments}}

	// Create encoder instance
	encoder := &ReedSolomonEncoder{}

	// Encode data
	encodedData, _, err := encoder.Encode(inputData, encoderOptions)
	assert.NoError(t, err, "Encoding should not return an error")

	// Check number of encoded segments
	assert.Len(t, encodedData, dataSegments+paritySegments, "Number of encoded segments should match data and parity segments")
}

func TestReedSolomonEncoder_Decode_Success(t *testing.T) {
	// Test data
	inputData := []byte("hello world")
	dataSegments := 6
	paritySegments := 3

	// Create encoder options
	encoderOptions := &EncoderOptions{Options{DataSegments: dataSegments, ParitySegments: paritySegments}}
	decoderOptions := &DecoderOptions{Options{DataSegments: dataSegments, ParitySegments: paritySegments}}

	// Create encoder instance
	encoder := &ReedSolomonEncoder{}

	// Encode data
	encodedData, parityData, err := encoder.Encode(inputData, encoderOptions)
	assert.NoError(t, err, "Encoding should not return an error")

	// Decode data
	decodedData, err := encoder.Decode(encodedData, parityData, decoderOptions)
	assert.NoError(t, err, "Decoding should not return an error")

	// Check decoded data
	assert.Equal(t, inputData, decodedData, "Decoded data should match original input data")
}

func TestReedSolomonEncoder_EncodeDecode_Success(t *testing.T) {
	// Test data
	inputData := []byte("hello world")
	dataSegments := 6
	paritySegments := 3

	// Create encoder and decoder options
	encoderOptions := &EncoderOptions{Options{DataSegments: dataSegments, ParitySegments: paritySegments}}
	decoderOptions := &DecoderOptions{Options{DataSegments: dataSegments, ParitySegments: paritySegments}}

	// Create encoder instance
	encoder := &ReedSolomonEncoder{}

	// Encode data
	encodedData, parityData, err := encoder.Encode(inputData, encoderOptions)
	assert.NoError(t, err, "Encoding should not return an error")

	// Decode data
	decodedData, err := encoder.Decode(encodedData, parityData, decoderOptions)
	assert.NoError(t, err, "Decoding should not return an error")

	// Check decoded data
	assert.Equal(t, inputData, decodedData, "Decoded data should match original input data")
}

func TestReedSolomonEncoder_Encode_Error(t *testing.T) {
	// Test data
	inputData := []byte("hello world")
	dataSegments := 0 // Invalid data segments

	// Create encoder options
	encoderOptions := &EncoderOptions{Options{DataSegments: dataSegments, ParitySegments: 3}}

	// Create encoder instance
	encoder := &ReedSolomonEncoder{}

	// Encode data
	_, _, err := encoder.Encode(inputData, encoderOptions)
	assert.Error(t, err, "Encoding should return an error due to invalid data segments")
}

func TestReedSolomonEncoder_Decode_Error(t *testing.T) {
	// Test data
	dataSegments := 6
	paritySegments := 3

	// Create decoder options
	decoderOptions := &DecoderOptions{Options{DataSegments: dataSegments, ParitySegments: paritySegments}}

	// Create encoder instance
	encoder := &ReedSolomonEncoder{}

	// Dummy encoded data segments
	data := [][]byte{{}, {}}

	// Decode data
	_, err := encoder.Decode(data[:dataSegments], data[dataSegments:], decoderOptions)
	assert.Error(t, err, "Decoding should return an error with empty data segments")
}
