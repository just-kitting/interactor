package protocol

type Endpoint struct {
	Token  uint8
	Method string
	Path   string
	Name   string
}

const (
	Version byte = 1

	TokenInfo  uint8 = 0x01
	TokenStart uint8 = 0x02
	TokenMove  uint8 = 0x03
	TokenEnd   uint8 = 0x04
)

const (
	StatusSuccess     uint8 = 0x00
	StatusBadRequest  uint8 = 0x01
	StatusUnsupported uint8 = 0x02
	StatusInternal    uint8 = 0x03
	StatusBusy        uint8 = 0x04
)

var EndpointTable = []Endpoint{
	{Token: TokenInfo, Method: "GET", Path: "/", Name: "info"},
	{Token: TokenStart, Method: "POST", Path: "/start", Name: "start"},
	{Token: TokenMove, Method: "POST", Path: "/move", Name: "move"},
	{Token: TokenEnd, Method: "POST", Path: "/end", Name: "end"},
}

func LookupEndpoint(token uint8) (Endpoint, bool) {
	for _, endpoint := range EndpointTable {
		if endpoint.Token == token {
			return endpoint, true
		}
	}
	return Endpoint{}, false
}
