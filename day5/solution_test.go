package day5

import "testing"

func TestSolutionA(t *testing.T) {
	if err := Solve(); err != nil {
		t.Fatalf("error solving %v", err)
	}
}

func TestSolutionB(t *testing.T) {
	t.Skip("too slow")
	if err := SolveB(); err != nil {
		t.Fatalf("error solving %v", err)
	}
}
