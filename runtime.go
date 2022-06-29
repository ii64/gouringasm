package gouringasm

import _ "unsafe"

// sys
//go:noescape
//go:linkname runtime_entersyscall runtime.entersyscall
func runtime_entersyscall()

//go:noescape
//go:linkname runtime_exitsyscall runtime.exitsyscall
func runtime_exitsyscall()
