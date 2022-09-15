package web

import "flag"

// for API service test purpose, testing with air
func (w *web) Air() {
	var flagAir int
	flag.IntVar(&flagAir, "air", 0, "set `-air 1` to enable web.Air function")
	flag.Parse()
	if flagAir == 1 {
		w.reset()
		w.https.ListenAndServeTLS("", "")
	}
}
