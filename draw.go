package main

var ()

func dLEV() {

	dGridIMoffsetY(LEV)

	dGridWalls(LEV)

	if dBUG {
		dGridLines(LEV, GREEN())
		dIMSheetXY(BLOKIM, 10, 10, 2)
		dIsoRline(inner, MAGENTA())
	}
}
