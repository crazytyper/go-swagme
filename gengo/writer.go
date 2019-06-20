package gengo

import (
	"fmt"
	"io"
)

// Writer ...
type Writer struct {
	w        io.Writer
	err      error
	indentBy string
	indent   string
	inline   bool
}

// NewWriter ...
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w, indent: "", indentBy: "    "}
}

// WritePackage ...
func (w *Writer) WritePackage(name string) {
	w.writeEOL()
	w.writeString("package ")
	w.writeString(name)
	w.writeEOL()
}

// WriteBlankLine ...
func (w *Writer) WriteBlankLine() {
	w.writeEOL()
	w.inline = true
	w.writeEOL()
}

// WriteMultilineCommentf ...
func (w *Writer) WriteMultilineCommentf(comment string, args ...interface{}) {
	w.writeEOL()
	w.writeString(fmt.Sprintf("/*%s\n*/\n", fmt.Sprintf(comment, args...)))
}

// WriteField ...
func (w *Writer) WriteField(name, typename, spec string) {
	w.writeString(name)
	w.writeString(" ")
	w.writeString(typename)
	if spec != "" {
		w.writeString(" ")
		w.WriteSinglelineComment(fmt.Sprintf("%q", spec))
	} else {
		w.writeEOL()
	}
}

// WriteSinglelineComment ...
func (w *Writer) WriteSinglelineComment(comment string) {
	w.writeString("// ")
	w.writeString(comment)
	w.writeEOL()
}

// WriteStructStart ...
func (w *Writer) WriteStructStart(name string) {
	w.writeEOL()
	w.writeString("type ")
	w.writeString(name)
	w.writeString(" struct {")
	w.writeEOL()
}

// WriteStructEnd ...
func (w *Writer) WriteStructEnd() {
	w.writeEOL()
	w.writeString("}")
	w.writeEOL()
}

func (w *Writer) writeString(s string) {
	if w.err != nil {
		return
	}
	if !w.inline {
		_, w.err = fmt.Fprint(w.w, w.indent)
	}
	if w.err == nil {
		_, w.err = fmt.Fprint(w.w, s)
	}
	if w.err == nil {
		w.inline = true
	}
}

func (w *Writer) writeEOL() {
	if w.inline {
		fmt.Fprintln(w.w)
		w.inline = false
	}
}
