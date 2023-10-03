package data

import "testing"

func TestToString(t *testing.T) {
	// Test integer values
	intCases := []struct {
		input    interface{}
		expected string
	}{
		{int(42), "42"},
		{int8(8), "8"},
		{int16(16), "16"},
		{int32(32), "32"},
		{int64(64), "64"},
	}
	for _, tc := range intCases {
		result := ToString(tc.input)
		if result != tc.expected {
			t.Errorf("ToString(%v) = %s; want %s", tc.input, result, tc.expected)
		}
	}

	// Test unsigned integer values
	uintCases := []struct {
		input    interface{}
		expected string
	}{
		{uint(42), "42"},
		{uint8(8), "8"},
		{uint16(16), "16"},
		{uint32(32), "32"},
		{uint64(64), "64"},
	}
	for _, tc := range uintCases {
		result := ToString(tc.input)
		if result != tc.expected {
			t.Errorf("ToString(%v) = %s; want %s", tc.input, result, tc.expected)
		}
	}

	// Test floating-point values
	floatCases := []struct {
		input    interface{}
		expected string
	}{
		{float32(3.14), "3.140000104904175"},
		{float64(2.71828), "2.71828"},
	}
	for _, tc := range floatCases {
		result := ToString(tc.input)
		if result != tc.expected {
			t.Errorf("ToString(%v) = %s; want %s", tc.input, result, tc.expected)
		}
	}

	// Test strings
	strInput := "Hello, World!"
	strResult := ToString(strInput)
	if strResult != strInput {
		t.Errorf("ToString(%v) = %s; want %s", strInput, strResult, strInput)
	}
}
