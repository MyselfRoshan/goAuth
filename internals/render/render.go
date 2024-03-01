package render

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
)

func Component() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templW io.Writer) (err error) {
		templBuffer, templISBuffer := templW.(*bytes.Buffer)
		if !templISBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		templChildren := templ.GetChildren(ctx)
		if templChildren == nil {
			templChildren = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templErr := templBuffer.WriteString("Hello")
		if templErr != nil {
			return templErr
		}
		if !templISBuffer {
			_, templErr := templBuffer.WriteTo(templW)
			return templErr
		}
		return nil
	})
}
