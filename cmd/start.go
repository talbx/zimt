package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/talbx/zimt/internal"
	"strconv"
	Time "time"
)

var (
	start = &cobra.Command{
		Use:   "start",
		Short: "shows how much time is probably left since your start!",
		Long:  `shows how much time is probably left since your start!`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			br, _ := cmd.PersistentFlags().GetInt("break")
			fmt.Println(calcStart(args, br))
		},
	}
)

func calcStart(time []string, br int) string {
	hours, minutes := internal.Split(time[0])
	h, _ := strconv.Atoi(hours)
	m, _ := strconv.Atoi(minutes)
	now := Time.Now()
	startDate := Time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, now.Location())

	endDate := startDate.Add(Time.Hour * 8).Add(Time.Minute * Time.Duration(br))

	hs := internal.AddLeadingZero(endDate.Hour())
	ms := internal.AddLeadingZero(endDate.Minute())
	return fmt.Sprintf("Your starting time: %s:%s. You can go home aprox. at %s:%s", internal.AddLeadingZero(h), internal.AddLeadingZero(m), hs, ms)
}

func init() {
	var Break int
	start.PersistentFlags().IntVarP(&Break, "break", "b", 45, "Total Break time planned or taken in minutes")
	rootCmd.AddCommand(start)
}
