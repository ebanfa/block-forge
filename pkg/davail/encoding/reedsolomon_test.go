package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReedSolomonEncoder_Encode_Success(t *testing.T) {
	// Test data
	dataSegments := 6
	paritySegments := 3
	inputData := []byte("hello world.")

	// Create encoder instance
	encoder, err := NewReedSolomonEncoder(dataSegments, paritySegments)
	assert.NoError(t, err, "Constructing encoder should not return an error")

	// Encode data
	_, _, err = encoder.Encode(inputData)
	assert.NoError(t, err, "Encoding should not return an error")

}

func TestReedSolomonEncoder_Decode_Success(t *testing.T) {
	// Test data
	inputData := []byte("hello world.")
	dataSegments := 6
	paritySegments := 3

	// Create encoder instance
	encoder, err := NewReedSolomonEncoder(dataSegments, paritySegments)
	assert.NoError(t, err, "Constructing encoder should not return an error")

	// Encode data
	encodedData, parityData, err := encoder.Encode(inputData)
	assert.NoError(t, err, "Encoding should not return an error")

	// Decode data
	decodedData, err := encoder.Decode(encodedData, parityData)
	assert.NoError(t, err, "Decoding should not return an error")

	// Check decoded data
	assert.Equal(t, inputData, decodedData, "Decoded data should match original input data")
}

func TestReedSolomonEncoder_EncodeDecode_Success(t *testing.T) {
	// Test data
	inputData := []byte("hello world.")
	dataSegments := 6
	paritySegments := 3

	// Create encoder instance
	encoder, err := NewReedSolomonEncoder(dataSegments, paritySegments)
	assert.NoError(t, err, "Constructing encoder should not return an error")

	// Encode data
	encodedData, parityData, err := encoder.Encode(inputData)
	assert.NoError(t, err, "Encoding should not return an error")

	// Decode data
	decodedData, err := encoder.Decode(encodedData, parityData)
	assert.NoError(t, err, "Decoding should not return an error")

	// Check decoded data
	assert.Equal(t, inputData, decodedData, "Decoded data should match original input data")
}

func TestReedSolomonEncoder_Encode_Error(t *testing.T) {
	// Test data
	dataSegments := 0 // Invalid data segments

	// Create encoder instance
	_, err := NewReedSolomonEncoder(dataSegments, 3)
	assert.Error(t, err, "Encoding should return an error due to invalid data segments")

}
