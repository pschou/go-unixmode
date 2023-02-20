# go-unixmode

This module is intended to help developers work directly with file permissions on Linux.

As of the time this module was built, there are no valid String() method in Go which
correctly, in a POSIX manner, identifies sticky bits.  Hence the birth of this module.

More documentation: https://pkg.go.dev/github.com/pschou/go-unixmode
