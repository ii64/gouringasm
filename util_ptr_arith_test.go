package gouringasm

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestUserdata(t *testing.T) {
	type test struct {
		v   any
		exp Userdata
	}
	ts := []test{
		{uint64(0), 0},
		{uint64(0xff), 0xff},
		{uint64(0xfffefd), 0xfffefd},
		{uintptr(0xcafeba), 0xcafeba},
		{unsafe.Pointer(nil), 0},
	}
	for _, tc := range ts {
		var u Userdata
		switch v := tc.v.(type) {
		case uint64:
			u.SetUint64(v)
		case uintptr:
			u.SetUintptr(v)
		case unsafe.Pointer:
			u.SetUnsafe(v)
		default:
			panic(fmt.Sprintf("unhandled type: %T", v))
		}

		assert.Equal(t, tc.exp, u)
	}
}
