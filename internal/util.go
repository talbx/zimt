package internal

import (
	"strconv"
	"strings"
)

func Split(time string) (hours string, minutes string) {
	var split = strings.Split(time, ":")
	hours = split[0]
	minutes = split[1]
	return
}

func ToMinutes(hours string, minutes string) int {
	realHours, _ := strconv.Atoi(hours)
	realMinutes, _ := strconv.Atoi(minutes)
	return realHours*60 + realMinutes
}
