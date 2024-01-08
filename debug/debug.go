package debug

import (
	"fmt"
	"maps"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var newLinesMap map[LineID][]any
var linesMap map[LineID][]any
var types []LineID
var refreshTicker *time.Ticker

func init() {
	linesMap = make(map[LineID][]any)
	newLinesMap = make(map[LineID][]any)
	refreshTicker = time.NewTicker(time.Second)
}

type LineID string

func SetRefreshTime(t time.Duration) {
	refreshTicker.Reset(t)
}

func Print(t LineID, s ...any) {
	if _, ok := newLinesMap[t]; !ok {
		types = append(types, t)
	}
	newLinesMap[t] = s
}

func Draw(screen *ebiten.Image) {
	select {
	case <-refreshTicker.C:
		maps.Copy(linesMap, newLinesMap)
	default:
	}
	if len(linesMap) == 0 && len(newLinesMap) != 0 {
		maps.Copy(linesMap, newLinesMap)
	}

	if len(linesMap) == 0 {
		return
	}
	for i, t := range types {
		line := linesMap[t]
		ebitenutil.DebugPrintAt(screen, string(t)+": "+fmt.Sprint(line...), 1, i*11)
	}
}
