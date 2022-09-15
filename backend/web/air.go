package web

import (
	// "context"
	"flag"
	// "time"
)

// For API service test purpose, testing with air
// Uncomment code below to run http redirector
func (w *web) Air() {
	var flagAir int
	flag.IntVar(&flagAir, "air", 0, "set `-air 1` to enable web.Air function")
	flag.Parse()
	if flagAir == 1 {
		w.reset()
		// go w.http.ListenAndServe()
		w.https.ListenAndServeTLS("", "")

		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()
		// w.http.Shutdown(ctx)
		// <-ctx.Done()
	}
}
