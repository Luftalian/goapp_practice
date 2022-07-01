package trace

import (
	"fmt"
	"io"
)

//Tracerはコード内での出来事を記録するオブジェクトを表す。
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

//OffはTracerメソッドの呼び出しを無視するTracerを返します。
func Off() Tracer {
	return &nilTracer{}
}
