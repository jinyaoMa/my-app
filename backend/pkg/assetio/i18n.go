package assetio

type Lang struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type I18n interface {
	Lang() *Lang
}
