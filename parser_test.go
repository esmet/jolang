package main

import (
	"testing"
)

type TestCase struct {
	Case string
	Tree Op
}

func TestParser(t *testing.T) {
	cases := []TestCase{
		{
			Case: `+ 1 + 1 2`,
			Tree: AddOp{
				a: Atom(1),
				b: AddOp{a: Atom(1), b: Atom(2)},
			},
		},
		{
			Case: `- 5  	+ 1 3`,
			Tree: SubOp{
				a: Atom(5),
				b: AddOp{a: Atom(1), b: Atom(3)},
			},
		},
		{
			Case: `* 1 + 3 - 1 2`,
			Tree: MultiplyOp{
				a: Atom(1),
				b: AddOp{
					a: Atom(3),
					b: SubOp{
						a: Atom(1),
						b: Atom(2),
					},
				},
			},
		},
		{
			Case: `+ 1 ( - 3 2 ) uh`,
			Tree: AddOp{
				a: Atom(1),
				b: SubOp{
					a: Atom(3),
					b: Atom(2),
				},
			},
		},
	}

	for _, c := range cases {
		parser := NewParser(c.Case)

		actual, err := parser.Parse()
		if err != nil {
			t.Fatalf("Nope: %v", err)
		}

		if actual != c.Tree {
			t.Fatalf("%s: expected %v != actual %v",
				c.Case, c.Tree, actual)
		}

	}
}
