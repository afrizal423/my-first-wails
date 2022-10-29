package spamin

import (
	"bufio"
	"context"
	"encoding/json"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Datanya struct {
	Target string `json:"target"`
	Pesan  string `json:"pesan"`
	Jumlah int    `json:"jumlah"`
}

type Spamin struct {
	ctx context.Context
}

func GasSpamin() *Spamin {
	return &Spamin{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (s *Spamin) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *Spamin) YukKirim(todos string) {
	runtime.EventsEmit(s.ctx, "mulai")
	var dt Datanya
	json.Unmarshal([]byte(todos), &dt)
	scanner := bufio.NewScanner(strings.NewReader(dt.Target))
	for scanner.Scan() {
		s.Scrape(scanner.Text(), dt.Jumlah, dt.Pesan)
	}
	runtime.EventsEmit(s.ctx, "selesai")
}
