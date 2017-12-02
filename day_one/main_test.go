package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	result := SolvePartOne("1122")
	want := 3
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartOne("1111")
	want = 4
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartOne("1234")
	want = 0
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartOne("91212129")
	want = 9
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	result := SolvePartTwo("1212")
	want := 6
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartTwo("1221")
	want = 0
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartTwo("123425")
	want = 4
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartTwo("123123")
	want = 12
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}

	result = SolvePartTwo("12131415")
	want = 4
	if result != want {
		t.Errorf("Solve was incorrect. Got: %d, Want: %d.", result, want)
	}
}
