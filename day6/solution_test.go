package day6

import "testing"

func TestSolution(t *testing.T) {
	if err := Solve(); err != nil {
		t.Fatalf("error solving %v", err)
	}

	if err := SolveB(); err != nil {
		t.Fatalf("error solving %v", err)
	}

}
