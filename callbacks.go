package rango

// #include <rango.h>
// #include <callbacks.h>
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"unsafe"
)

//export GoReadConsole
func GoReadConsole(p *C.char, buf *C.uchar, buflen C.int, add_history C.int) C.int {
	fmt.Print(C.GoString(p))
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if len(text) > 0 {
		text = text[0:(len(text) - 1)]
	}
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


func SetCallbacks() {
	C.rango_set_callback(C.CString("ptr_R_ReadConsole"), C.rango_read_console)
}
