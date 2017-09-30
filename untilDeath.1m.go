package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	const t2death = "3366186454" // die 2076 september 1st
	i, _ := strconv.ParseInt(t2death, 10, 64)

	t := time.Unix(i, 0)
	tUntilDeath := time.Until(t)

	yrs := tUntilDeath.Hours() / (365 * 24)
	rYrs := math.Floor(yrs)

	days := (yrs - rYrs) * (365)
	rDays := math.Floor(days)

	hours := (days - rDays) * (24)
	rHours := math.Floor(hours)

	mins := (hours - rHours) * 60
	rMins := math.Floor(mins)

	fmt.Printf("%+vy, %+vd, %+vh, %+vm", rYrs, rDays, rHours, rMins)
}
