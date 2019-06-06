package main

import (
	"testing"
)

func TestAddAtoms(t *testing.T) {
	ast := AddOp{a: Atom(1), b: Atom(2)}

	r, err := ast.eval()
	if err != nil {
		t.Fatalf("badness.com: %v", err)
	}

	if r != 3 {
		t.Fatalf("Why is %v not three?", r)
	}
}

func TestAddAdd(t *testing.T) {
	ast1 := AddOp{a: Atom(1), b: AddOp{a: Atom(1), b: Atom(2)}}
	ast2 := AddOp{b: Atom(1), a: AddOp{b: Atom(1), a: Atom(2)}}

	r1, err := ast1.eval()
	if err != nil {
		t.Fatalf("badness.com: %v", err)
	}

	r2, err := ast2.eval()
	if err != nil {
		t.Fatalf("badness.com: %v", err)
	}

	if r1 != 4 {
		t.Fatalf("Why is %v not four?", r1)
	}
	if r2 != 4 {
		t.Fatalf("Why is %v not four?", r2)
	}
}

// func TestExec(t *testing.T) {
// 	ast := ExecOp{args: []string{"echo", "1"}}
// 
// 	r, err := ast.eval()
// 	if err != nil {
// 		t.Fatalf("badness.com: %v", err)
// 	}
// 
// 	if r != 4 {
// 		t.Fatalf("Why is '%v' not echo 1?", r)
// 	}
// }
