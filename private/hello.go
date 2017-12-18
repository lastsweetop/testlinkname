package private

import _ "unsafe"

//go:linkname hello github.com/lastsweetop/testlinkname/hello.hellofunc
func hello() string {
	return "private.hello()"
}
