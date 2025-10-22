package main

var (
	LEV                 []isoR
	BLOKIM, PLIM, ETCIM []IM
	wallBlok            IM
	levW                = 25
	inner               isoR
)

// MARK: MAKE
func mLEV() {
	LEV = nil
	LEV = mGridCNT(UN, levW, levW)
	rInner, _ := fIsoRnum(LEV, 130)
	inner = mIsoR(rInner.p[0], UN*float32(levW-9))
	LEV = mGridIM(LEV, BLOKIM[26])
	LEV = mGridWallsEdge(LEV)
	LEV = mGridInnerWalls(LEV)
	wallBlok = BLOKIM[80]
	LEV = gridSORT(LEV)

}
func mIMG() {
	BLOKIM = mIMSheetXYMultiRowTexSize("im/01.png", 0, 0, 18, 18)
	BLOKIM = remIMfromSheet(BLOKIM, []int{9, 29, 39, 49, 59, 68, 69, 78, 87, 97})
	PLIM = mIMSheetXY1RowNum("im/02.png", 0, 0, 16, 16, 5)

	mETCIM()
}
func mETCIM() {
	ETCIM = append(ETCIM, mIMPath("im/03.png")) //0 PL SHADOW

}
