package fitweight

import (
	"os"
)

func WriteToDisk(output string, wm WeightMeasurement) error {
	f, err := os.Create(output)

	if err != nil {
		return err
	}

	return wm.writeFitFile(f)
}
