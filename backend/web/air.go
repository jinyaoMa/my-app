package web

import (
	// "context"
	"my-app/backend/app"
	// "time"
)

// For API service test purpose, testing with air
// Uncomment code below to run http redirector
func (w *web) Air() {
	if app.App().Flag().UseAir() {
		w.reset()
		// go w.http.ListenAndServe()
		w.https.ListenAndServeTLS("", "")

		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()
		// w.http.Shutdown(ctx)
		// <-ctx.Done()
	}
}
