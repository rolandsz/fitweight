package cmd

import (
	"os"

	"github.com/rolandsz/fitweight"

	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:     "write",
	Short:   "Writes the specified weight measurement to a .fit file",
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		weight, _ := flags.GetFloat64("weight")
		if weight == -1 {
			cmd.PrintErr("No weight specified\n")
			os.Exit(1)
		}

		fat, _ := flags.GetFloat64("fat")
		hydration, _ := flags.GetFloat64("hydration")
		bone, _ := flags.GetFloat64("bone")
		boneKg, _ := flags.GetFloat64("bone-mass")
		muscle, _ := flags.GetFloat64("muscle")
		muscleKg, _ := flags.GetFloat64("muscle-mass")
		ts, _ := flags.GetInt64("unix-timestamp")
		visceralFat, _ := flags.GetFloat64("visceral-fat")
		metabolicAge, _ := flags.GetFloat64("metabolic-age")
		physiqueRating, _ := flags.GetFloat64("physique-rating")
		calories, _ := flags.GetFloat64("calories")
		bmi, _ := flags.GetFloat64("bmi")

		var boneMass float64 = 0.0
		var muscleMass float64 = 0.0

		if bone > -1 && boneKg > -1 {
			cmd.PrintErrf("Cannot provide bone weight in percent and bone mass in kg! Use either of both!")
			os.Exit(1)
		}

		if bone > -1 {
			boneMass = weight * bone / 100
		} else if boneKg > -1 {
			boneMass = boneKg
		}

		if muscle > -1 && muscleKg > -1 {
			cmd.PrintErrf("Cannot provide muscle weight in percent and muscle mass in kg! Use either of both!")
			os.Exit(1)
		}

		if muscle > -1 {
			muscleMass = weight * muscle / 100
		} else if muscleKg > -1 {
			muscleMass = muscleKg
		}

		wm := fitweight.NewWeightMeasurement(weight, fat, hydration, boneMass, muscleMass, visceralFat, physiqueRating, metabolicAge, calories, bmi, ts)

		output, _ := cmd.Flags().GetString("output")

		cmd.Println("... writing weight measurement to:", output)

		if err := fitweight.WriteToDisk(output, wm); err != nil {
			cmd.PrintErrf("Error writing weight measurement to the disk: %s\n", err.Error())
			os.Exit(1)
		}

		cmd.Println("Done!")

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	flags := writeCmd.Flags()

	flags.StringP("output", "o", "default.fit", "Path to the output .fit file")

	flags.Float64P("weight", "w", -1, "Set your weight in kilograms")
	flags.Float64P("fat", "f", 0, "Set your fat in percent")
	flags.Float64("hydration", 0, "Set your hydration in percent")
	flags.Float64P("bone", "b", -1, "Set your bone mass in percent")
	flags.Float64("bone-mass", -1, "Set your bone mass in kilograms")
	flags.Float64P("muscle", "m", -1, "Set your muscle mass in percent")
	flags.Float64("muscle-mass", -1, "Set your muscle mass in kilograms")
	flags.Float64P("calories", "c", 0, "Set your caloric intake")
	flags.Float64("visceral-fat", 0, "Set your visceral fat rating (valid values: 1-60)")
	flags.Float64("metabolic-age", 0, "Set your metabolic age")
	flags.Float64("physique-rating", 0, "Set your physique rating (valid values: 1-9)")
	flags.Float64("bmi", 0, "Set your BMI - body mass index")
	flags.Int64P("unix-timestamp", "t", -1, "Set the timestamp of the measurement")
}
