package hello

import _ "github.com/lastsweetop/testlinkname/private"

// Provided by package runtime.
func hellofunc() string

func Greet() string {
	return hellofunc()
}
