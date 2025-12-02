package one

import (
	"log"
	"testing"
)

func TestDayOne(t *testing.T) {
	solved := Solve("../../inputs/one/test.txt")
	expected := 6

	if solved != expected {
		log.Fatalf("test did not pass, solution solved for: %v, expected: %v", solved, expected)
	}

}
