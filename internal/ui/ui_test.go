package console

import "testing"

func TestForamatting(t *testing.T) {
	formatedTime := formatTime(100)

	if formatedTime != "01:40" {
		t.Error("Time formatting didn't work.")
	}

	formatedName := formatName("ljubLjaNa aP")

	if formatedName != "Ljubljana Ap" {
		t.Error("Name formatting didn't work.")
	}
}
