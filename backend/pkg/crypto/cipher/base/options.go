package base

type Options struct {
	AAD    string `json:"aad"`    // additional authenticated data
	Key    string `json:"key"`    // in base64 w/o prefix
	Iv     string `json:"iv"`     // in base64 w/o prefix
	Prefix string `json:"prefix"` // prefix of base64 key, default is alg + "_"
}
