// This file is generated by gorazor 1.2.1
// DON'T modified manually
// Should edit source file and re-generate: cases/section_comment_bug.gohtml

package cases

import (
	"cases/layout"
	"io"
	"strings"
)

// Section_comment_bug generates cases/section_comment_bug.gohtml
func Section_comment_bug() string {
	var _b strings.Builder
	RenderSection_comment_bug(&_b)
	return _b.String()
}

// RenderSection_comment_bug render cases/section_comment_bug.gohtml
func RenderSection_comment_bug(_buffer io.StringWriter) {

	_body := func(_buffer io.StringWriter) {
		// Line: 7
		_buffer.WriteString("\n\n<a>\n    <!-- comment -->\n</a>")

	}

	_side := func(_buffer io.StringWriter) {

		// Line: 14
		_buffer.WriteString("<!-- comment -->\n    plain text")

	}

	layout.RenderBase(_buffer, _body, nil, nil)
}
