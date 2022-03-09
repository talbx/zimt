package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/talbx/zimt/internal"
	"math"
	Time "time"
)

var (
	cmd1 = &cobra.Command{
		Use:   "left",
		Short: "shows how much time is left on an 8h shift!",
		Long:  `shows how much time is left on an 8h shift!`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(calcLeft(args))
		},
	}
)

func calcLeft(time []string) string {
	hours, minutes := internal.Split(time[0])
	totalPassed := internal.ToMinutes(hours, minutes)
	totalWanted := 8 * 60
	totalMinutes := totalWanted - totalPassed

	totalHours := int(math.Trunc(float64(totalMinutes) / 60))
	totalRestMinutes := totalMinutes % 60
	now := Time.Now().Local().Add(Time.Minute * Time.Duration(totalMinutes))

	hs := internal.AddLeadingZero(now.Hour())
	ms := internal.AddLeadingZero(now.Minute())
	return fmt.Sprintf("Time Left: %s:%s. You can go home at %s:%s", internal.AddLeadingZero(totalHours), internal.AddLeadingZero(totalRestMinutes), hs, ms)
}

func init() {
	rootCmd.AddCommand(cmd1)
}
