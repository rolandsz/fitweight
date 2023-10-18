//go:build integration
// +build integration

package fitweight

import (
	"os"
	"testing"
)

func TestWriteWeightToDisk(t *testing.T) {
	wm := WeightMeasurement(80, 14.4, 55.2, 37, 45.5, 21, 5, 23, 2250, 12.6, -1)

	output := "test.fit"
	err := Write(output, wm)

	if err != nil {
		t.Fatalf("Error writing .fit file: %s", err.Error())
	}
}
