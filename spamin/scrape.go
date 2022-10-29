package spamin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gocolly/colly"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *Spamin) ProsesKirim(q string, pesan string) {
	values := map[string]string{"id": q, "message": pesan}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://api.secreto.site/sendmsg", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

}

func (s *Spamin) kirimPesanKonkurent(jumlah int, q string, pesan string) {
	wg := sync.WaitGroup{}

	for i := 0; i < jumlah; i++ {
		wg.Add(1)
		go func() {
			s.ProsesKirim(q, pesan)

			wg.Done()
		}()
	}
	wg.Wait()
}

func (s *Spamin) Scrape(url string, jumlah int, pesan string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("secreto.site", "api.secreto.site"),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("span#id.hidden", func(e *colly.HTMLElement) {
		// fmt.Printf("Link found: %q \n", e.Text)
		s.kirimPesanKonkurent(jumlah, e.Text, pesan)
	})

	c.OnResponse(func(r *colly.Response) {
		runtime.EventsEmit(s.ctx, "pesannya", fmt.Sprintf("mengirim pesan ke %s", r.Request.URL))
		// fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		runtime.EventsEmit(s.ctx, "pesannya", "Pesan telah terkirim")
	})

	// Start scraping
	c.Visit(url)
}
