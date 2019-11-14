// This file is generated by gorazor 1.2.1
// DON'T modified manually
// Should edit source file and re-generate: cases/scope.gohtml

package cases

import (
	"dm"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
	"zfw/models"
	. "zfw/tplhelper"
)

// Scope generates cases/scope.gohtml
func Scope(obj *models.Widget) string {
	var _b strings.Builder
	RenderScope(&_b, obj)
	return _b.String()
}

// RenderScope render cases/scope.gohtml
func RenderScope(_buffer io.StringWriter, obj *models.Widget) {

	data, dmType := dm.GetData(obj.PlaceHolder)

	if dmType == "simple" {
		obj.StringList = data.([]string)

		// Line: 15
		_buffer.WriteString("<div>")
		// Line: 15
		_buffer.WriteString((SelectPk(obj)))
		// Line: 15
		_buffer.WriteString("</div>")

	} else {
		node := data.(*dm.DMTree)

		// Line: 18
		_buffer.WriteString("<div class=\"form-group ")
		// Line: 18
		_buffer.WriteString(gorazor.HTMLEscape(GetErrorClass(obj)))
		// Line: 18
		_buffer.WriteString("\">\n  <label for=\"")
		// Line: 19
		_buffer.WriteString(gorazor.HTMLEscape(obj.Name))
		// Line: 19
		_buffer.WriteString("\" class=\"col-sm-2 control-label\">")
		// Line: 19
		_buffer.WriteString(gorazor.HTMLEscape(obj.Label))
		// Line: 19
		_buffer.WriteString("</label>\n  <div class=\"col-sm-10\">\n    <select class=\"form-control\" name=\"")
		// Line: 21
		_buffer.WriteString(gorazor.HTMLEscape(obj.Name))
		// Line: 21
		_buffer.WriteString("\" ")
		// Line: 21
		_buffer.WriteString(gorazor.HTMLEscape(BoolStr(obj.Disabled, "disabled")))
		// Line: 21
		_buffer.WriteString(">\n      ")
		for _, option := range node.Keys {
			if values, ok := node.Values[option]; ok {

				// Line: 24
				_buffer.WriteString("<optgroup label=\"")
				// Line: 24
				_buffer.WriteString(gorazor.HTMLEscape(option))
				// Line: 24
				_buffer.WriteString("\">\n        ")
				for _, value := range values {
					if value == obj.Value {

						// Line: 27
						_buffer.WriteString("<option selected>")
						// Line: 27
						_buffer.WriteString(gorazor.HTMLEscape(value))
						// Line: 27
						_buffer.WriteString("</option>")

					} else {

						// Line: 29
						_buffer.WriteString("<option>")
						// Line: 29
						_buffer.WriteString(gorazor.HTMLEscape(value))
						// Line: 29
						_buffer.WriteString("</option>")

					}
				}
				// Line: 31
				_buffer.WriteString("\n      </optgroup>")

			} else {
				if option == obj.Value {

					// Line: 35
					_buffer.WriteString("<option selected>")
					// Line: 35
					_buffer.WriteString(gorazor.HTMLEscape(option))
					// Line: 35
					_buffer.WriteString("</option>")

				} else {

					// Line: 37
					_buffer.WriteString("<option>")
					// Line: 37
					_buffer.WriteString(gorazor.HTMLEscape(option))
					// Line: 37
					_buffer.WriteString("</option>")

				}
			}
		}
		// Line: 40
		_buffer.WriteString("\n    </select>\n    ")
		if obj.ErrorMsg != "" {

			// Line: 43
			_buffer.WriteString("<span class=\"label label-danger\">")
			// Line: 43
			_buffer.WriteString(gorazor.HTMLEscape(obj.ErrorMsg))
			// Line: 43
			_buffer.WriteString("</span>")

		}
		// Line: 44
		_buffer.WriteString("\n  </div>\n</div>")
	}

}
