package trace

import (
	"fmt"
	"io"
)

// コード内での出来事を記録できるオブジェクトを表すインターフェース
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func (t *nilTracer) Trace(a ...interface{}) {
	// nothing to do
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Traceメソッドの呼び出しを無視するTracerを返す
func Off() Tracer {
	return &nilTracer{}
}
