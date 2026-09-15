package core

import (
	"crypto/rand"
	"fmt"
	"time"
)

type StockReserve struct {
	ResourceId string
	ReservedAt *time.Time
	Expired    bool
	Id         string
}

// ReserveStock
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func (app *App) ReserveStock(resourceId string) (*StockReserve, bool) {

	if app.StockCounts[resourceId] == 0 {
		return nil, false
	} else {

		app.StockCounts[resourceId]--

		now := time.Now()
		id, _ := GenerateUUID()

		stockReserve := &StockReserve{
			ResourceId: resourceId,
			ReservedAt: &now,
			Expired:    false,
			Id:         id,
		}

		if app.StockReserves[resourceId] == nil {
			app.StockReserves[resourceId] = []StockReserve{}
		}

		app.StockReserves[resourceId] = append(app.StockReserves[resourceId], *stockReserve)

		return stockReserve, true
	}
}

// CreateStock
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func (app *App) CreateStock(resourceId string, quantity int64) {
	app.StockCounts[resourceId] += quantity
}

/*func (app *App) ReleaseStock(resourceId string) bool {

}*/

func GenerateUUID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// Set version 4 and variant bits.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
