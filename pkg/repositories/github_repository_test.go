package repositories

import "fmt"
import "testing"

func TestSum(t *testing.T) {
	commits := CompareCommits("vino-warehouse", "2bed280d", "06179368")
	expectd := 7
	for _, c := range commits {
		fmt.Printf("Commit: %+v\n", c)
	}
	if len(commits) != expectd {
		t.Errorf("Something's wrong, got: %d, want: %d.", len(commits), expectd)
	}
}
