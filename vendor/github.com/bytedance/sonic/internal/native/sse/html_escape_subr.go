// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__html_escape = 64
)

const (
    _stack__html_escape = 72
)

const (
    _size__html_escape = 1280
)

var (
    _pcsp__html_escape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1256, 72},
        {1260, 48},
        {1261, 40},
        {1263, 32},
        {1265, 24},
        {1267, 16},
        {1269, 8},
        {1271, 0},
    }
)

var _cfunc_html_escape = []loader.CFunc{
    {"_html_escape_entry", 0,  _entry__html_escape, 0, nil},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
}
