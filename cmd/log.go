package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/talbx/zimt/internal"
	"strconv"
	Time "time"
)

var (
	log = &cobra.Command{
		Use:   "log",
		Short: "logs your time",
		Long:  `logs your time really it does`,
		Run: func(cmd *cobra.Command, args []string) {
			Log(cmd)
		},
	}
)

func Log(cmd *cobra.Command) {

	var m = make(map[string]bool)


	m["startEvent"], _ = cmd.Flags().GetBool("startEvent")
	m["breakStartEvent"], _ = cmd.Flags().GetBool("breakStartEvent")
	m["breakEndEvent"], _ = cmd.Flags().GetBool("breakEndEvent")
	m["endEvent"], _ = cmd.Flags().GetBool("endEvent")
	var chosenEvent string

	for key, value := range m {
		if value{
			chosenEvent = key
		}
	}

	if "" == chosenEvent {
		var getInt, _ = cmd.Flags().GetInt("list")
		if getInt == 0{
			getInt = 10
		}
		yaml := internal.ReadFromYaml(getInt)
		fmt.Println(yaml)
		return
	}

	now := Time.Now()
	date := fmt.Sprintf("%s.%s.%s", strconv.Itoa(now.Day()), now.Month().String(), strconv.Itoa(now.Year()))
	time := fmt.Sprintf("%s:%s", strconv.Itoa(now.Hour()), strconv.Itoa(now.Minute()))
	entry := internal.ZimtEntry{
		Date: date,
		Time: time,
		Type: chosenEvent,
	}
	internal.WriteToYaml(entry)
	sprintf := fmt.Sprintf("Registered ZimtEntry - date: %s, time: %s, type: %s", date, time, chosenEvent)
	fmt.Println(sprintf)
}

func init() {
	var WorkStartEvent bool
	var BreakStartEvent bool
	var BreakEndEvent bool
	var WorkEndEvent bool
	var ListSize int
	log.Flags().BoolVarP(&WorkStartEvent, "startEvent", "s", false, "Create work-start event")
	log.Flags().BoolVarP(&BreakStartEvent, "breakStartEvent", "b", false, "Create break-start Event")
	log.Flags().BoolVarP(&BreakEndEvent, "breakEndEvent", "r", false, "Create break-end Event")
	log.Flags().BoolVarP(&WorkEndEvent, "endEvent", "e", false, "Create work-end event")
	log.Flags().IntVarP(&ListSize, "list", "l", 0, "list the past logs")
	rootCmd.AddCommand(log)
}
