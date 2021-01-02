package main

import gc "github.com/vsakkas/goncurses"

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		return
	}
	defer gc.End()

	stdscr.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE,
		gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
	stdscr.GetChar()
}
