// Package trace provides Tracer
package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprintln(t.out, a...)
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
