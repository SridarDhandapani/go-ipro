package api

type Scheme int

const (
	HTTP Scheme = iota
	HTTPS
	SOCKS5
)

func (s Scheme) String() string {
	return [...]string{"http", "https", "socks5"}[s]
}
