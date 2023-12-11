package day10

import "testing"

func TestSolutionA(t *testing.T) {
	if err := Solve(); err != nil {
		t.Fatalf("error solving %v", err)
	}
}

func TestSolutionB(t *testing.T) {
	if err := SolveB(); err != nil {
		t.Fatalf("error solving %v", err)
	}
}
