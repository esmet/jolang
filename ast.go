package main

import (
	// "os/exec"
	// "strconv"
	// "fmt"
)

type Atom int

func (o Atom) eval() (Atom, error) {
	return o, nil
}

type AddOp struct {
	a, b Op
}

func (o AddOp) eval() (Atom, error) {
	x, e1 := o.a.eval()
	if e1 != nil {
		return -1, e1
	}

	y, e2 := o.b.eval()
	if e2 != nil {
		return -1, e2
	}

	return x + y, nil
}

type SubOp struct {
	a, b Op
}

func (o SubOp) eval() (Atom, error) {
	x, e1 := o.a.eval()
	if e1 != nil {
		return -1, e1
	}

	y, e2 := o.b.eval()
	if e2 != nil {
		return -1, e2
	}

	return x - y, nil
}

type MultiplyOp struct {
	a, b Op
}

func (o MultiplyOp) eval() (Atom, error) {
	x, e1 := o.a.eval()
	if e1 != nil {
		return -1, e1
	}

	y, e2 := o.b.eval()
	if e2 != nil {
		return -1, e2
	}

	return x * y, nil
}

// type ExecOp struct {
// 	args []string
// }
// 
// func (o ExecOp) eval() (Atom, error) {
// 	cmd := exec.Cmd{
// 		Path: "/bin/bash",
// 		Args: append([]string{"-c"}, o.args),
// 	}
// 
// 	err := cmd.Run()
// 	if err != nil {
// 		return -1, fmt.Errorf("Could not exec %+v: %v", cmd, err)
// 	}
// 
// 	r, err := cmd.Output()
// 	if err != nil {
// 		return -1, fmt.Errorf("Could not run %+v: %v", cmd, err)
// 	}
// 
// 	v, err := strconv.ParseInt(string(r), 10, 16)
// 	if err != nil {
// 		return -1, fmt.Errorf("Could not parse return valu %v: %v", string(r), v)
// 	}
// 
// 	return Atom(v), nil
// }

type Op interface {
	eval() (Atom, error)
}
