package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
)

type Parser struct {
	scanner *bufio.Scanner
	sub []string
	doingTheThing bool
}

func (p Parser) Parse() (Op, error) {
	p.scanner.Scan()
	s := p.scanner.Text()

	fmt.Printf("Parsing `%v`\n", s)
	switch s {
	case "(":
		p.doingTheThing = true
		p.sub = append(p.sub, s)
		fmt.Printf("Gathering.. `%v`\n", s)
		return p.Parse()
	case ")":
		panic("wtf")
		p.doingTheThing = false
		str := strings.Join(p.sub, " ")
		sub := NewParser(str)
		fmt.Printf("Sub parsing.. `%v`\n", str)
		uh, what := sub.Parse()
		fmt.Printf("sub parse returned %v and %v\n", uh, what)
		return uh, what
	case "+":
		a, err := p.Parse()
		if err != nil {
			return nil, err
		}

		b, err := p.Parse()
		if err != nil {
			return nil, err
		}

		return AddOp{a: a, b: b}, nil
	case "-":
		a, err := p.Parse()
		if err != nil {
			return nil, err
		}

		b, err := p.Parse()
		if err != nil {
			return nil, err
		}

		return SubOp{a: a, b: b}, nil
	case "*":
		a, err := p.Parse()
		if err != nil {
			return nil, err
		}

		b, err := p.Parse()
		if err != nil {
			return nil, err
		}

		return MultiplyOp{a: a, b: b}, nil
	default:
		fmt.Printf("Atom? %v\n", s)
		v, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return nil, fmt.Errorf("Unbueno int %v: %v", s, err)
		}

		return Atom(v), nil
	}

	return nil, nil
}

func NewParser(input string) Parser {
	fmt.Printf("New parser on `%v`\n", input);
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	return Parser{scanner: scanner}
}
