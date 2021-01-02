// +build !windows

package goncurses

// #include <curses.h>
import "C"

// Echo prints a single character to the pad immediately. This has the
// same effect of calling AddChar() + Refresh() but has a significant
// speed advantage
func (p *Pad) Echo(ch int) {
	C.pechochar(p.win, C.chtype(ch))
}
