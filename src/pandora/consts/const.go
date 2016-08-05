package consts

const (
	CACHE_NOT_FOUND int = 0
	CACHE_FOUND     int = 1
	CACHE_DISABLED  int = -1
)

var (
	//////all consts in pandora
	P_DOT          = []byte(".")
	P_BRACE_LEFT   = []byte("{")
	P_BRACE_RIGHT  = []byte("}")
	P_COMMA        = []byte(",")
	P_HTML_SUBFIX  = []byte(".html")
	P_SLASH        = []byte("/")
	P_EQUAL        = []byte("=")
	P_QUOTE        = []byte("\"")
	P_SINGLE_QUOTE = []byte("'")
	P_SPACE        = []byte(" ")
	P_QUESTION     = []byte("?")
)
