package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Day parts are expected in output order: part 1, part 2, part 3.
var expectations = map[string][]string{
	"1": {"1472", "330", "629094"},
}

func TestDays2026(t *testing.T) {
	for day, expect := range expectations {
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", fmt.Sprintf("days/%s/main.go", day))
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(output))
			}

			assert.NoError(t, err)
			assert.Equal(t, strings.Join(expect, "\n"), strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
