// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Simple E-Commerce Page</title><link rel=\"stylesheet\" href=\"/static/base.css\"></head><body><header><h1>My E-Commerce Store</h1><nav><ul><li><a href=\"#\">Home</a></li><li><a href=\"#\">Shop</a></li><li><a href=\"#\">Contact</a></li></ul></nav></header><section class=\"products\"><article class=\"product\"><img src=\"/assets/img/th-298986468.jpg\" alt=\"Product 1\"><h2>Product 1</h2><p>$19.99</p><button>Add to Cart</button></article><article class=\"product\"><img src=\"/assets/img/th-298986468.jpg\" alt=\"Product 2\"><h2>Product 2</h2><p>$29.99</p><button>Add to Cart</button></article><!-- Add more products as needed --></section><footer><p>&copy; 2024 My E-Commerce Store. All rights reserved.</p></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
