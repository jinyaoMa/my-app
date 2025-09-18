package argon2

type Options struct {
	Salt      string `json:"salt"`
	Threads   int    `json:"threads"`   // 1-254, not greater than num of cpu
	KeyLength int    `json:"keyLength"` // 2-512 bytes
	Prefix    string `json:"prefix"`    // prefix of base64 key, default is "argon2_"
}
