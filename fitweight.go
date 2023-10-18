package fitweight

import (
	"encoding/binary"
	"io"
	"time"

	"github.com/tormoder/fit"
)

// WeightMeasurement is the data container struct for managing the weight measurements
type WeightMeasurement struct {
	TimeStamp         time.Time
	Weight            float64
	PercentFat        float64
	PercentHydration  float64
	BoneMass          float64
	MuscleMass        float64
	VisceralFatRating float64
	PhysiqueRating    float64
	MetabolicAge      float64
	CaloriesActiveMet float64
	BodyMassIndex     float64
}

func (wm WeightMeasurement) writeFitFile(writer io.Writer) error {
	fitfile, err := fit.NewFile(fit.FileTypeWeight, fit.NewHeader(fit.V20, true))
	if err != nil {
		return err
	}
	fitfile.FileId.TimeCreated = wm.TimeStamp

	weight, err := fitfile.Weight()
	if err != nil {
		return err
	}

	weight.WeightScales = []*fit.WeightScaleMsg{
		{
			Timestamp:         wm.TimeStamp,
			Weight:            fit.Weight(wm.Weight * 100),
			PercentFat:        uint16(wm.PercentFat * 100),
			PercentHydration:  uint16(wm.PercentHydration * 100),
			BoneMass:          uint16(wm.BoneMass * 100),
			MuscleMass:        uint16(wm.MuscleMass * 100),
			VisceralFatRating: uint8(wm.VisceralFatRating),
			PhysiqueRating:    uint8(wm.PhysiqueRating),
			MetabolicAge:      uint8(wm.MetabolicAge),
			ActiveMet:         uint16(wm.CaloriesActiveMet),
			Bmi:               uint16(wm.BodyMassIndex * 10),
		},
	}

	return fit.Encode(writer, fitfile, binary.BigEndian)
}

// NewWeightMeasurement creates a new WeightMeasurement instance
func NewWeightMeasurement(weight, percentFat, percentHydration, boneMass, muscleMass, visceralFatRating, physiqueRating, metabolicAge, caloriesActiveMet, bmi float64, timestamp int64) WeightMeasurement {
	ts := time.Now()
	if timestamp != -1 {
		ts = time.Unix(timestamp, 0)
	}

	return WeightMeasurement{
		TimeStamp:         ts,
		Weight:            weight,
		PercentFat:        percentFat,
		PercentHydration:  percentHydration,
		BoneMass:          boneMass,
		MuscleMass:        muscleMass,
		VisceralFatRating: visceralFatRating,
		PhysiqueRating:    physiqueRating,
		MetabolicAge:      metabolicAge,
		CaloriesActiveMet: caloriesActiveMet,
		BodyMassIndex:     bmi,
	}
}
