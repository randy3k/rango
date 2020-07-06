package rango

// #include <rango.h>
// #include <callbacks.h>
import "C"
import (
	// "fmt"
	"unsafe"
)

type SEXP unsafe.Pointer

type CallbacksDef struct {
	Suicide func(p string)
	ShowMessage func(p string)
	ReadConsole func(p string, add_history bool) string
	WriteConsole func(p string, otype int)
	ResetConsole func()
	FlushConsole func()
	ClearerrConsole func()
	Busy func(int)
	Cleanup func(int, int, int)
	ShowFiles func(int, []string, []string, string, bool, string) int
	ChooseFile func(int, string, int) int
	EditFile func(string) int
	Loadhistory func(SEXP, SEXP, SEXP, SEXP)
	Savehistory func(SEXP, SEXP, SEXP, SEXP)
	Addhistory func(SEXP, SEXP, SEXP, SEXP)
	EditFiles func(int, []string, []string, string) int
	DoSelectlist func(SEXP, SEXP, SEXP, SEXP) SEXP
	DoDataentry func(SEXP, SEXP, SEXP, SEXP) SEXP
	DoDataviewer func(SEXP, SEXP, SEXP, SEXP) SEXP
	ProcessEvents func()
	PolledEvents func()
	YesNoCancel func(string) int
}

var Callbacks CallbacksDef

//export GoReadConsole
func GoReadConsole(p *C.char, buf *C.uchar, buflen C.int, add_history C.int) C.int {
	text := Callbacks.ReadConsole(C.GoString(p), add_history == 1)
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	n := len(text)
	C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(cs), C.size_t(n))
	// TODO: check buflen
	buf_t := uintptr(unsafe.Pointer(buf))
	*((*C.uchar)(unsafe.Pointer(buf_t + uintptr(n) * unsafe.Sizeof(*buf)))) = C.uchar('\n')
	*((*C.uchar)(unsafe.Pointer(buf_t + uintptr(n + 1) * unsafe.Sizeof(*buf)))) = C.uchar('\x00')
	return C.int(1)
}


//export GoWriteConsole
func GoWriteConsole(p *C.char, bufline C.int, otype C.int) {
	Callbacks.WriteConsole(C.GoString(p), int(otype))
}


func SetCallbacks() {
	if Callbacks.ReadConsole != nil {
		C.rango_set_callback(C.CString("ptr_R_ReadConsole"), C.rango_read_console)
	}
	if Callbacks.WriteConsole != nil {
		C.rango_set_callback(C.CString("ptr_R_WriteConsole"), nil)
		C.rango_set_callback(C.CString("ptr_R_WriteConsoleEx"), C.rango_write_console)
	}
}
