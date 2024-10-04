package main

import "fmt"

type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

type Calculator interface {
	Calculate() int
}

type Number struct {
	Value int
}

func (n *Number) Calculate() int {
	return n.Value
}

type Oper struct {
	Type  Operation
	Left  Calculator
	Right Calculator
}

func (o *Oper) Calculate() int {
	switch o.Type {
	case Add:
		return o.Left.Calculate() + o.Right.Calculate()
	case Sub:
		return o.Left.Calculate() - o.Right.Calculate()
	case Div:
		return o.Left.Calculate() / o.Right.Calculate()
	case Mul:
		return o.Left.Calculate() * o.Right.Calculate()
	}
	panic(fmt.Sprintf(`unknown operator %d`, o.Type))
}
