package one

import (
	"aoc-2025/m/v2/internal/utils"
	"log"
	"strconv"
)

func Solve(path string) int {

	d := Dial{Position: 50}
	lines, err := utils.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	for _, line := range lines {
		rotation := parseRotation(line)
		d.rotateAndCountTimesAtZero(rotation)
	}
	return d.TimesAtZero
}

func parseRotation(line string) int {

	i, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatalf("error parsing rotation: %s", err)
	}

	if line[0] == 'L' {
		return -i
	}
	return i
}

type Dial struct {
	Position    int
	TimesAtZero int
}

func (d *Dial) rotateAndCountTimesAtZero(amount int) {

	// Count full turns, comment out for part one
	d.TimesAtZero += max(amount/100, -amount/100)

	// Truncate to two digits (450 and 50 are functionally the same)
	amount = amount % 100

	switch {
	// If im going to a negative number, loop around (-10 becomes 90)
	case (d.Position + amount) < 0:

		// If im NOT at 0 currently, count the looping around as a pass-through of 0.
		// Comment this for part one
		if d.Position != 0 {
			d.TimesAtZero++
		}

		d.Position = 100 + d.Position + amount

	// If im going over 100, deduct 100 (110 becomes 10)
	case (d.Position + amount) > 99:
		d.Position = d.Position + amount - 100

		// Count the going over as a pass-through of 0
		// Comment this for part one
		if d.Position != 0 {
			d.TimesAtZero++
		}

	// Otherwise, just rotate that amount
	default:
		d.Position += amount
	}

	// Count if I finish a rotation at 0
	if d.Position == 0 {
		d.TimesAtZero++
	}
}
