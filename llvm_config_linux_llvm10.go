// +build !byollvm
// +build linux,llvm10

package llvm

// Automatically generated by `make config BUILDDIR=`, do not edit.

// #cgo CPPFLAGS: -I/usr/lib/llvm-10/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++14
// #cgo LDFLAGS: -L/usr/lib/llvm-10/lib  -lLLVM-10
import "C"

type run_build_sh int