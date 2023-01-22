package hexview_test

import (
	"strings"
	"testing"

	"github.com/gdamore/tcell/v2"
	hexview "github.com/jmhobbs/tview-hexview"
	"github.com/stretchr/testify/assert"
)

func Test_HexView_Offsets(t *testing.T) {
	screen := tcell.NewSimulationScreen("UTF-8")
	screen.Init()
	screen.SetSize(80, 1)

	hv := hexview.NewHexView([]byte{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
	})

	hv.SetRect(0, 0, 80, 5)
	hv.Draw(screen)
	screen.Sync()

	cells, w, h := screen.GetContents()
	actual := cellsToStrings(cells, w, h)
	expected := []string{" 00000000 | 00 01 02 03 04 05 06 07 ┊ 08 09 0a 0b 0c 0d 0e 0f |________┊_______|"}

	assert.Equal(t, expected, actual)

	// todo: offset changes
}

// todo: test styles

func cellsToStrings(cells []tcell.SimCell, w, h int) []string {
	var buf strings.Builder
	var ret []string
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			buf.WriteRune(cells[y*w+x].Runes[0]) // we assume only one rune
		}
		ret = append(ret, buf.String())
		buf.Reset()
	}
	return ret
}
