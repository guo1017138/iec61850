package iec61850

// #include <iec61850_client.h>
import "C"

type FC int

// fc types
const (
	// ST Status information
	ST FC = iota
	// MX Measurands - analogue values
	MX
	// SP Setpoint
	SP
	// SV Substitution
	SV
	// CF Configuration
	CF
	// DC Description
	DC
	// SG Setting group
	SG
	// SE Setting group editable
	SE
	// SR service response / service tracking
	SR
	// OR Operate received
	OR
	// BL Blocking
	BL
	// EX Extended definition
	EX
	// CO Control, deprecated but kept here for backward compatibility
	CO
	// US Unicast SV
	US
	// MS Multicast SV
	MS
	// RP Unbuffered Reporting
	RP
	// BR Buffered Reporting
	BR
	// LG Log control blocks
	LG
	// GO Goose control blocks
	GO
	// ALL All FCs - wildcard value
	ALL  FC = 99
	NONE FC = -1
)

func GetFCByString(fcStr string) (fc FC) {
	fc = NONE
	switch fcStr {
	case "ST":
		fc = ST
	case "MX":
		fc = MX
	case "SP":
		fc = SP
	case "SV":
		fc = SV
	case "CF":
		fc = CF
	case "DC":
		fc = DC
	case "SG":
		fc = SG
	case "SE":
		fc = SE
	case "SR":
		fc = SR
	case "OR":
		fc = OR
	case "BL":
		fc = BL
	case "EX":
		fc = EX
	case "CO":
		fc = CO
	case "US":
		fc = US
	case "MS":
		fc = MS
	case "RP":
		fc = RP
	case "BR":
		fc = BR
	case "LG":
		fc = LG
	case "GO":
		fc = GO
	case "ALL":
		fc = ALL
	}
	return
}
