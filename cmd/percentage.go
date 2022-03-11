package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/talbx/zimt/internal"
)

var (
	perc = &cobra.Command{
		Use:   "perc",
		Short: "shows actual time based on percental of a whole",
		Long:  `shows actual time based on percental of a whole`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(calcPerc(args))
		},
	}
)

func calcPerc(args []string) string {
	percentage, _ := strconv.Atoi(args[0])
	hours, minutes := internal.Split(args[1])
	totalPassed := float64(internal.ToMinutes(hours, minutes))
	percentual := (totalPassed / 100) * float64(percentage)
	totalHours := int(math.Trunc(float64(percentual) / 60))
	totalRestMinutes := int(percentual) % 60

	h := internal.AddLeadingZero(totalHours)
	m := internal.AddLeadingZero(totalRestMinutes)
	return fmt.Sprintf("%d%% of %s corresponds to %s:%s", percentage, args[1], h, m)
}

func init() {
	rootCmd.AddCommand(perc)
}
