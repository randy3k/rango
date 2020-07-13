package rango

// #cgo linux LDFLAGS: -Wl,--no-as-needed -ldl
// #include <rango.h>
import "C"
import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unsafe"
)

func Initialize(Rhome string, args []string) (bool, error) {
	if Rhome == "" {
		return false, errors.New("Rhome not found")
	}

	os.Setenv("R_HOME", Rhome)

	cs := C.CString(Rhome)
	defer C.free(unsafe.Pointer(cs))

	sz := unsafe.Sizeof(uintptr(0))
	cargs := (**C.char)(C.malloc(C.size_t(len(args)) * C.size_t(sz)))
	cargs_t := unsafe.Pointer(cargs)
	defer C.free(cargs_t)

	for i, a := range args {
		p := (**C.char)(unsafe.Pointer(uintptr(cargs_t) + uintptr(i)*sz))
		ca := C.CString(a)
		defer C.free(unsafe.Pointer(ca))
		*p = ca
	}

	if C.rango_load(cs) == 0 {
		return false, errors.New(C.GoString(C.rango_dl_error_message()))
	}
	if C.rango_load_symbols() == 0 {
		s := fmt.Sprintf(
			"%s - symbol %s",
			C.GoString(C.rango_dl_error_message()),
			C.GoString(C.rango_last_loaded_symbol()))
		return false, errors.New(s)
	}
	C.rango_init(C.int(len(args)), (**C.char)(cargs))
	C.rango_setuploop()

	if C.rango_load_constants() == 0 {
		s := fmt.Sprintf(
			"%s - constant %s",
			C.GoString(C.rango_dl_error_message()),
			C.GoString(C.rango_last_loaded_symbol()))
		return false, errors.New(s)
	}

	return IsInitialized(), nil
}

func IsInitialized() bool {
	return C.rango_is_initialized() == 1
}

func GetRhome() string {
	if rhome, ok := os.LookupEnv("R_HOME"); ok {
		return rhome
	}

	out, err := exec.Command("R", "RHOME").Output()
	if err == nil {
		return strings.TrimSpace(string(out))
	}

	return ""
}

func DefaultArgs() []string {
	return []string{"rango", "--quiet", "--no-save"}
}

func RunREPL() {
	C.rango_runloop()
}
