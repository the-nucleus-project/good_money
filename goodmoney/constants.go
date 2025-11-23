package goodmoney

// ISO 4217 Currency codes

const (
	AED = "AED"
	AFN = "AFN"
	ALL = "ALL"
	AMD = "AMD"
	AOA = "AOA"
	ARS = "ARS"
	AUD = "AUD"
	AWG = "AWG"
	AZN = "AZN"
	BAM = "BAM"
	BBD = "BBD"
	BDT = "BDT"
	BGN = "BGN"
	BHD = "BHD"
	BIF = "BIF"
	BMD = "BMD"
	BND = "BND"
	BOB = "BOB"
	BOV = "BOV"
	BRL = "BRL"
	BSD = "BSD"
	BTN = "BTN"
	BWP = "BWP"
	BYN = "BYN"
	BZD = "BZD"
	CAD = "CAD"
	CDF = "CDF"
	CHE = "CHE"
	CHF = "CHF"
	CHW = "CHW"
	CLF = "CLF"
	CLP = "CLP"
	CNY = "CNY"
	COP = "COP"
	COU = "COU"
	CRC = "CRC"
	CUP = "CUP"
	CVE = "CVE"
	CZK = "CZK"
	DJF = "DJF"
	DKK = "DKK"
	DOP = "DOP"
	DZD = "DZD"
	EGP = "EGP"
	ERN = "ERN"
	ETB = "ETB"
	EUR = "EUR"
	FJD = "FJD"
	FKP = "FKP"
	GBP = "GBP"
	GEL = "GEL"
	GHS = "GHS"
	GIP = "GIP"
	GMD = "GMD"
	GNF = "GNF"
	GTQ = "GTQ"
	GYD = "GYD"
	HKD = "HKD"
	HNL = "HNL"
	HTG = "HTG"
	HUF = "HUF"
	IDR = "IDR"
	ILS = "ILS"
	INR = "INR"
	IQD = "IQD"
	IRR = "IRR"
	ISK = "ISK"
	JMD = "JMD"
	JOD = "JOD"
	JPY = "JPY"
	KES = "KES"
	KGS = "KGS"
	KHR = "KHR"
	KMF = "KMF"
	KPW = "KPW"
	KRW = "KRW"
	KWD = "KWD"
	KYD = "KYD"
	KZT = "KZT"
	LAK = "LAK"
	LBP = "LBP"
	LKR = "LKR"
	LRD = "LRD"
	LSL = "LSL"
	LYD = "LYD"
	MAD = "MAD"
	MDL = "MDL"
	MGA = "MGA"
	MKD = "MKD"
	MMK = "MMK"
	MNT = "MNT"
	MOP = "MOP"
	MRU = "MRU"
	MUR = "MUR"
	MVR = "MVR"
	MWK = "MWK"
	MXN = "MXN"
	MXV = "MXV"
	MYR = "MYR"
	MZN = "MZN"
	NAD = "NAD"
	NGN = "NGN"
	NIO = "NIO"
	NOK = "NOK"
	NPR = "NPR"
	NZD = "NZD"
	OMR = "OMR"
	PAB = "PAB"
	PEN = "PEN"
	PGK = "PGK"
	PHP = "PHP"
	PKR = "PKR"
	PLN = "PLN"
	PYG = "PYG"
	QAR = "QAR"
	RON = "RON"
	RSD = "RSD"
	RUB = "RUB"
	RWF = "RWF"
	SAR = "SAR"
	SBD = "SBD"
	SCR = "SCR"
	SDG = "SDG"
	SEK = "SEK"
	SGD = "SGD"
	SHP = "SHP"
	SLE = "SLE"
	SOS = "SOS"
	SRD = "SRD"
	SSP = "SSP"
	STN = "STN"
	SVC = "SVC"
	SYP = "SYP"
	SZL = "SZL"
	THB = "THB"
	TJS = "TJS"
	TMT = "TMT"
	TND = "TND"
	TOP = "TOP"
	TRY = "TRY"
	TTD = "TTD"
	TWD = "TWD"
	TZS = "TZS"
	UAH = "UAH"
	UGX = "UGX"
	USD = "USD"
	USN = "USN"
	UYI = "UYI"
	UYU = "UYU"
	UYW = "UYW"
	UZS = "UZS"
	VED = "VED"
	VES = "VES"
	VND = "VND"
	VUV = "VUV"
	WST = "WST"
	XAD = "XAD"
	XAF = "XAF"
	XAG = "XAG"
	XAU = "XAU"
	XBA = "XBA"
	XBB = "XBB"
	XBC = "XBC"
	XBD = "XBD"
	XCD = "XCD"
	XCG = "XCG"
	XDR = "XDR"
	XOF = "XOF"
	XPD = "XPD"
	XPF = "XPF"
	XPT = "XPT"
	XSU = "XSU"
	XTS = "XTS"
	XUA = "XUA"
	XXX = "XXX"
	YER = "YER"
	ZAR = "ZAR"
	ZMW = "ZMW"
	ZWG = "ZWG"
)

var CurrencyMap = map[string]Currency{
	"AED": {
		NumericCode:    "784",
		MinorUnit:      2,
		Symbol:         "د.إ",
		SymbolPosition: true,
	},
	"AFN": {
		NumericCode: "971",
		MinorUnit:   2,
	},
	"ALL": {
		NumericCode:    "008",
		MinorUnit:      2,
		Symbol:         "L",
		SymbolPosition: false,
	},
	"AMD": {
		NumericCode:    "051",
		MinorUnit:      2,
		Symbol:         "֏",
		SymbolPosition: false,
	},
	"AOA": {
		NumericCode: "973",
		MinorUnit:   2,
	},
	"ARS": {
		NumericCode:    "032",
		MinorUnit:      2,
		Symbol:         "$",
		SymbolPosition: true,
	},
	"AUD": {
		NumericCode:    "036",
		MinorUnit:      2,
		Symbol:         "A$",
		SymbolPosition: true,
	},
	"AWG": {
		NumericCode: "533",
		MinorUnit:   2,
	},
	"AZN": {
		NumericCode:    "944",
		MinorUnit:      2,
		Symbol:         "₼",
		SymbolPosition: true,
	},
	"BAM": {
		NumericCode:    "977",
		MinorUnit:      2,
		Symbol:         "КМ",
		SymbolPosition: false,
	},
	"BBD": {
		NumericCode: "052",
		MinorUnit:   2,
	},
	"BDT": {
		NumericCode:    "050",
		MinorUnit:      2,
		Symbol:         "৳",
		SymbolPosition: true,
	},
	"BGN": {
		NumericCode:    "975",
		MinorUnit:      2,
		Symbol:         "лв",
		SymbolPosition: false,
	},
	"BHD": {
		NumericCode:    "048",
		MinorUnit:      3,
		Symbol:         ".د.ب",
		SymbolPosition: true,
	},
	"BIF": {
		NumericCode: "108",
		MinorUnit:   0,
	},
	"BMD": {
		NumericCode: "060",
		MinorUnit:   2,
	},
	"BND": {
		NumericCode: "096",
		MinorUnit:   2,
	},
	"BOB": {
		NumericCode: "068",
		MinorUnit:   2,
	},
	"BOV": {
		NumericCode: "984",
		MinorUnit:   2,
	},
	"BRL": {
		NumericCode:    "986",
		MinorUnit:      2,
		Symbol:         "R$",
		SymbolPosition: true,
	},
	"BSD": {
		NumericCode: "044",
		MinorUnit:   2,
	},
	"BTN": {
		NumericCode: "064",
		MinorUnit:   2,
	},
	"BWP": {
		NumericCode: "072",
		MinorUnit:   2,
	},
	"BYN": {
		NumericCode:    "933",
		MinorUnit:      2,
		Symbol:         "Br",
		SymbolPosition: false,
	},
	"BZD": {
		NumericCode: "084",
		MinorUnit:   2,
	},
	"CAD": {
		NumericCode:    "124",
		MinorUnit:      2,
		Symbol:         "C$",
		SymbolPosition: true,
	},
	"CDF": {
		NumericCode: "976",
		MinorUnit:   2,
	},
	"CHE": {
		NumericCode: "947",
		MinorUnit:   2,
	},
	"CHF": {
		NumericCode:    "756",
		MinorUnit:      2,
		Symbol:         "CHF",
		SymbolPosition: true,
	},
	"CHW": {
		NumericCode: "948",
		MinorUnit:   2,
	},
	"CLF": {
		NumericCode: "990",
		MinorUnit:   4,
	},
	"CLP": {
		NumericCode:    "152",
		MinorUnit:      0,
		Symbol:         "$",
		SymbolPosition: true,
	},
	"CNY": {
		NumericCode:    "156",
		MinorUnit:      2,
		Symbol:         "¥",
		SymbolPosition: true,
	},
	"COP": {
		NumericCode:    "170",
		MinorUnit:      2,
		Symbol:         "$",
		SymbolPosition: true,
	},
	"COU": {
		NumericCode: "970",
		MinorUnit:   2,
	},
	"CRC": {
		NumericCode: "188",
		MinorUnit:   2,
	},
	"CUP": {
		NumericCode: "192",
		MinorUnit:   2,
	},
	"CVE": {
		NumericCode: "132",
		MinorUnit:   2,
	},
	"CZK": {
		NumericCode:    "203",
		MinorUnit:      2,
		Symbol:         "Kč",
		SymbolPosition: false,
	},
	"DJF": {
		NumericCode: "262",
		MinorUnit:   0,
	},
	"DKK": {
		NumericCode:    "208",
		MinorUnit:      2,
		Symbol:         "kr",
		SymbolPosition: false,
	},
	"DOP": {
		NumericCode: "214",
		MinorUnit:   2,
	},
	"DZD": {
		NumericCode: "012",
		MinorUnit:   2,
	},
	"EGP": {
		NumericCode:    "818",
		MinorUnit:      2,
		Symbol:         "E£",
		SymbolPosition: true,
	},
	"ERN": {
		NumericCode: "232",
		MinorUnit:   2,
	},
	"ETB": {
		NumericCode:    "230",
		MinorUnit:      2,
		Symbol:         "Br",
		SymbolPosition: true,
	},
	"EUR": {
		NumericCode:    "978",
		MinorUnit:      2,
		Symbol:         "€",
		SymbolPosition: false,
	},
	"FJD": {
		NumericCode: "242",
		MinorUnit:   2,
	},
	"FKP": {
		NumericCode: "238",
		MinorUnit:   2,
	},
	"GBP": {
		NumericCode:    "826",
		MinorUnit:      2,
		Symbol:         "£",
		SymbolPosition: true,
	},
	"GEL": {
		NumericCode:    "981",
		MinorUnit:      2,
		Symbol:         "₾",
		SymbolPosition: true,
	},
	"GHS": {
		NumericCode: "936",
		MinorUnit:   2,
	},
	"GIP": {
		NumericCode: "292",
		MinorUnit:   2,
	},
	"GMD": {
		NumericCode: "270",
		MinorUnit:   2,
	},
	"GNF": {
		NumericCode: "324",
		MinorUnit:   0,
	},
	"GTQ": {
		NumericCode: "320",
		MinorUnit:   2,
	},
	"GYD": {
		NumericCode: "328",
		MinorUnit:   2,
	},
	"HKD": {
		NumericCode:    "344",
		MinorUnit:      2,
		Symbol:         "HK$",
		SymbolPosition: true,
	},
	"HNL": {
		NumericCode: "340",
		MinorUnit:   2,
	},
	"HTG": {
		NumericCode: "332",
		MinorUnit:   2,
	},
	"HUF": {
		NumericCode:    "348",
		MinorUnit:      2,
		Symbol:         "Ft",
		SymbolPosition: false,
	},
	"IDR": {
		NumericCode:    "360",
		MinorUnit:      2,
		Symbol:         "Rp",
		SymbolPosition: true,
	},
	"ILS": {
		NumericCode:    "376",
		MinorUnit:      2,
		Symbol:         "₪",
		SymbolPosition: true,
	},
	"INR": {
		NumericCode:    "356",
		MinorUnit:      2,
		Symbol:         "₹",
		SymbolPosition: true,
	},
	"IQD": {
		NumericCode:    "368",
		MinorUnit:      3,
		Symbol:         "ع.د",
		SymbolPosition: false,
	},
	"IRR": {
		NumericCode:    "364",
		MinorUnit:      2,
		Symbol:         "﷼",
		SymbolPosition: true,
	},
	"ISK": {
		NumericCode:    "352",
		MinorUnit:      0,
		Symbol:         "kr",
		SymbolPosition: true,
	},
	"JMD": {
		NumericCode: "388",
		MinorUnit:   2,
	},
	"JOD": {
		NumericCode:    "400",
		MinorUnit:      3,
		Symbol:         "د.ا",
		SymbolPosition: true,
	},
	"JPY": {
		NumericCode:    "392",
		MinorUnit:      0,
		Symbol:         "¥",
		SymbolPosition: true,
	},
	"KES": {
		NumericCode: "404",
		MinorUnit:   2,
	},
	"KGS": {
		NumericCode: "417",
		MinorUnit:   2,
	},
	"KHR": {
		NumericCode:    "116",
		MinorUnit:      2,
		Symbol:         "៛",
		SymbolPosition: false,
	},
	"KMF": {
		NumericCode: "174",
		MinorUnit:   0,
	},
	"KPW": {
		NumericCode: "408",
		MinorUnit:   2,
	},
	"KRW": {
		NumericCode:    "410",
		MinorUnit:      0,
		Symbol:         "₩",
		SymbolPosition: true,
	},
	"KWD": {
		NumericCode:    "414",
		MinorUnit:      3,
		Symbol:         "د.ك",
		SymbolPosition: true,
	},
	"KYD": {
		NumericCode: "136",
		MinorUnit:   2,
	},
	"KZT": {
		NumericCode:    "398",
		MinorUnit:      2,
		Symbol:         "₸",
		SymbolPosition: false,
	},
	"LAK": {
		NumericCode:    "418",
		MinorUnit:      2,
		Symbol:         "₭",
		SymbolPosition: false,
	},
	"LBP": {
		NumericCode:    "422",
		MinorUnit:      2,
		Symbol:         "£",
		SymbolPosition: true,
	},
	"LKR": {
		NumericCode:    "144",
		MinorUnit:      2,
		Symbol:         "₨",
		SymbolPosition: true,
	},
	"LRD": {
		NumericCode: "430",
		MinorUnit:   2,
	},
	"LSL": {
		NumericCode: "426",
		MinorUnit:   2,
	},
	"LYD": {
		NumericCode: "434",
		MinorUnit:   3,
	},
	"MAD": {
		NumericCode: "504",
		MinorUnit:   2,
	},
	"MDL": {
		NumericCode:    "498",
		MinorUnit:      2,
		Symbol:         "L",
		SymbolPosition: false,
	},
	"MGA": {
		NumericCode: "969",
		MinorUnit:   2,
	},
	"MKD": {
		NumericCode:    "807",
		MinorUnit:      2,
		Symbol:         "ден",
		SymbolPosition: false,
	},
	"MMK": {
		NumericCode:    "104",
		MinorUnit:      2,
		Symbol:         "K",
		SymbolPosition: true,
	},
	"MNT": {
		NumericCode:    "496",
		MinorUnit:      2,
		Symbol:         "₮",
		SymbolPosition: true,
	},
	"MOP": {
		NumericCode: "446",
		MinorUnit:   2,
	},
	"MRU": {
		NumericCode: "929",
		MinorUnit:   2,
	},
	"MUR": {
		NumericCode: "480",
		MinorUnit:   2,
	},
	"MVR": {
		NumericCode: "462",
		MinorUnit:   2,
	},
	"MWK": {
		NumericCode: "454",
		MinorUnit:   2,
	},
	"MXN": {
		NumericCode:    "484",
		MinorUnit:      2,
		Symbol:         "$",
		SymbolPosition: true,
	},
	"MXV": {
		NumericCode: "979",
		MinorUnit:   2,
	},
	"MYR": {
		NumericCode:    "458",
		MinorUnit:      2,
		Symbol:         "RM",
		SymbolPosition: true,
	},
	"MZN": {
		NumericCode: "943",
		MinorUnit:   2,
	},
	"NAD": {
		NumericCode: "516",
		MinorUnit:   2,
	},
	"NGN": {
		NumericCode: "566",
		MinorUnit:   2,
	},
	"NIO": {
		NumericCode: "558",
		MinorUnit:   2,
	},
	"NOK": {
		NumericCode:    "578",
		MinorUnit:      2,
		Symbol:         "kr",
		SymbolPosition: false,
	},
	"NPR": {
		NumericCode:    "524",
		MinorUnit:      2,
		Symbol:         "₨",
		SymbolPosition: true,
	},
	"NZD": {
		NumericCode:    "554",
		MinorUnit:      2,
		Symbol:         "NZ$",
		SymbolPosition: true,
	},
	"OMR": {
		NumericCode:    "512",
		MinorUnit:      3,
		Symbol:         "﷼",
		SymbolPosition: true,
	},
	"PAB": {
		NumericCode: "590",
		MinorUnit:   2,
	},
	"PEN": {
		NumericCode:    "604",
		MinorUnit:      2,
		Symbol:         "S/",
		SymbolPosition: true,
	},
	"PGK": {
		NumericCode: "598",
		MinorUnit:   2,
	},
	"PHP": {
		NumericCode:    "608",
		MinorUnit:      2,
		Symbol:         "₱",
		SymbolPosition: true,
	},
	"PKR": {
		NumericCode:    "586",
		MinorUnit:      2,
		Symbol:         "₨",
		SymbolPosition: true,
	},
	"PLN": {
		NumericCode:    "985",
		MinorUnit:      2,
		Symbol:         "zł",
		SymbolPosition: false,
	},
	"PYG": {
		NumericCode: "600",
		MinorUnit:   0,
	},
	"QAR": {
		NumericCode:    "634",
		MinorUnit:      2,
		Symbol:         "﷼",
		SymbolPosition: true,
	},
	"RON": {
		NumericCode:    "946",
		MinorUnit:      2,
		Symbol:         "lei",
		SymbolPosition: false,
	},
	"RSD": {
		NumericCode:    "941",
		MinorUnit:      2,
		Symbol:         "дин",
		SymbolPosition: false,
	},
	"RUB": {
		NumericCode:    "643",
		MinorUnit:      2,
		Symbol:         "₽",
		SymbolPosition: false,
	},
	"RWF": {
		NumericCode: "646",
		MinorUnit:   0,
	},
	"SAR": {
		NumericCode:    "682",
		MinorUnit:      2,
		Symbol:         "﷼",
		SymbolPosition: true,
	},
	"SBD": {
		NumericCode: "090",
		MinorUnit:   2,
	},
	"SCR": {
		NumericCode: "690",
		MinorUnit:   2,
	},
	"SDG": {
		NumericCode: "938",
		MinorUnit:   2,
	},
	"SEK": {
		NumericCode:    "752",
		MinorUnit:      2,
		Symbol:         "kr",
		SymbolPosition: false,
	},
	"SGD": {
		NumericCode:    "702",
		MinorUnit:      2,
		Symbol:         "S$",
		SymbolPosition: true,
	},
	"SHP": {
		NumericCode: "654",
		MinorUnit:   2,
	},
	"SLE": {
		NumericCode: "925",
		MinorUnit:   2,
	},
	"SOS": {
		NumericCode: "706",
		MinorUnit:   2,
	},
	"SRD": {
		NumericCode: "968",
		MinorUnit:   2,
	},
	"SSP": {
		NumericCode: "728",
		MinorUnit:   2,
	},
	"STN": {
		NumericCode: "930",
		MinorUnit:   2,
	},
	"SVC": {
		NumericCode: "222",
		MinorUnit:   2,
	},
	"SYP": {
		NumericCode: "760",
		MinorUnit:   2,
	},
	"SZL": {
		NumericCode: "748",
		MinorUnit:   2,
	},
	"THB": {
		NumericCode:    "764",
		MinorUnit:      2,
		Symbol:         "฿",
		SymbolPosition: true,
	},
	"TJS": {
		NumericCode: "972",
		MinorUnit:   2,
	},
	"TMT": {
		NumericCode: "934",
		MinorUnit:   2,
	},
	"TND": {
		NumericCode: "788",
		MinorUnit:   3,
	},
	"TOP": {
		NumericCode: "776",
		MinorUnit:   2,
	},
	"TRY": {
		NumericCode:    "949",
		MinorUnit:      2,
		Symbol:         "₺",
		SymbolPosition: true,
	},
	"TTD": {
		NumericCode: "780",
		MinorUnit:   2,
	},
	"TWD": {
		NumericCode: "901",
		MinorUnit:   2,
	},
	"TZS": {
		NumericCode: "834",
		MinorUnit:   2,
	},
	"UAH": {
		NumericCode:    "980",
		MinorUnit:      2,
		Symbol:         "₴",
		SymbolPosition: true,
	},
	"UGX": {
		NumericCode: "800",
		MinorUnit:   0,
	},
	"USD": {
		NumericCode:    "840",
		MinorUnit:      2,
		Symbol:         "$",
		SymbolPosition: true,
	},
	"USN": {
		NumericCode: "997",
		MinorUnit:   2,
	},
	"UYI": {
		NumericCode: "940",
		MinorUnit:   0,
	},
	"UYU": {
		NumericCode: "858",
		MinorUnit:   2,
	},
	"UYW": {
		NumericCode: "927",
		MinorUnit:   4,
	},
	"UZS": {
		NumericCode:    "860",
		MinorUnit:      2,
		Symbol:         "so'm",
		SymbolPosition: false,
	},
	"VED": {
		NumericCode: "926",
		MinorUnit:   2,
	},
	"VES": {
		NumericCode: "928",
		MinorUnit:   2,
	},
	"VND": {
		NumericCode:    "704",
		MinorUnit:      0,
		Symbol:         "₫",
		SymbolPosition: false,
	},
	"VUV": {
		NumericCode: "548",
		MinorUnit:   0,
	},
	"WST": {
		NumericCode: "882",
		MinorUnit:   2,
	},
	"XAD": {
		NumericCode: "396",
		MinorUnit:   2,
	},
	"XAF": {
		NumericCode: "950",
		MinorUnit:   0,
	},
	"XAG": {
		NumericCode: "961",
		MinorUnit:   2,
	},
	"XAU": {
		NumericCode: "959",
		MinorUnit:   2,
	},
	"XBA": {
		NumericCode: "955",
		MinorUnit:   2,
	},
	"XBB": {
		NumericCode: "956",
		MinorUnit:   2,
	},
	"XBC": {
		NumericCode: "957",
		MinorUnit:   2,
	},
	"XBD": {
		NumericCode: "958",
		MinorUnit:   2,
	},
	"XCD": {
		NumericCode: "951",
		MinorUnit:   2,
	},
	"XCG": {
		NumericCode: "532",
		MinorUnit:   2,
	},
	"XDR": {
		NumericCode: "960",
		MinorUnit:   2,
	},
	"XOF": {
		NumericCode: "952",
		MinorUnit:   0,
	},
	"XPD": {
		NumericCode: "964",
		MinorUnit:   2,
	},
	"XPF": {
		NumericCode: "953",
		MinorUnit:   0,
	},
	"XPT": {
		NumericCode: "962",
		MinorUnit:   2,
	},
	"XSU": {
		NumericCode: "994",
		MinorUnit:   2,
	},
	"XTS": {
		NumericCode: "963",
		MinorUnit:   2,
	},
	"XUA": {
		NumericCode: "965",
		MinorUnit:   2,
	},
	"XXX": {
		NumericCode: "999",
		MinorUnit:   2,
	},
	"YER": {
		NumericCode: "886",
		MinorUnit:   2,
	},
	"ZAR": {
		NumericCode:    "710",
		MinorUnit:      2,
		Symbol:         "R",
		SymbolPosition: true,
	},
	"ZMW": {
		NumericCode: "967",
		MinorUnit:   2,
	},
	"ZWG": {
		NumericCode: "924",
		MinorUnit:   2,
	},
}

// localeData holds locale-specific words for human-readable formatting
type localeData struct {
	zero     string
	negative string
	and      string
	ones     []string
	tens     []string
	hundred  string
	thousand string
	million  string
	billion  string
}

// localeDataMap holds locale-specific data for number-to-words conversion, keyed by language code
var localeDataMap = map[string]localeData{
	"en": { // English (default)
		zero:     "zero",
		negative: "negative",
		and:      "and",
		ones:     []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
		tens:     []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"},
		hundred:  "hundred",
		thousand: "thousand",
		million:  "million",
		billion:  "billion",
	},
	"fr": { // French
		zero:     "zéro",
		negative: "moins",
		and:      "et",
		ones:     []string{"zéro", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix", "onze", "douze", "treize", "quatorze", "quinze", "seize", "dix-sept", "dix-huit", "dix-neuf"},
		tens:     []string{"", "", "vingt", "trente", "quarante", "cinquante", "soixante", "soixante-dix", "quatre-vingt", "quatre-vingt-dix"},
		hundred:  "cent",
		thousand: "mille",
		million:  "million",
		billion:  "milliard",
	},
	"de": { // German
		zero:     "null",
		negative: "minus",
		and:      "und",
		ones:     []string{"null", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"},
		tens:     []string{"", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"},
		hundred:  "hundert",
		thousand: "tausend",
		million:  "Million",
		billion:  "Milliarde",
	},
	"es": { // Spanish
		zero:     "cero",
		negative: "menos",
		and:      "y",
		ones:     []string{"cero", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve", "diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve"},
		tens:     []string{"", "", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"},
		hundred:  "ciento",
		thousand: "mil",
		million:  "millón",
		billion:  "mil millones",
	},
	"it": { // Italian
		zero:     "zero",
		negative: "meno",
		and:      "e",
		ones:     []string{"zero", "uno", "due", "tre", "quattro", "cinque", "sei", "sette", "otto", "nove", "dieci", "undici", "dodici", "tredici", "quattordici", "quindici", "sedici", "diciassette", "diciotto", "diciannove"},
		tens:     []string{"", "", "venti", "trenta", "quaranta", "cinquanta", "sessanta", "settanta", "ottanta", "novanta"},
		hundred:  "cento",
		thousand: "mille",
		million:  "milione",
		billion:  "miliardo",
	},
	"pt": { // Portuguese
		zero:     "zero",
		negative: "menos",
		and:      "e",
		ones:     []string{"zero", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez", "onze", "doze", "treze", "catorze", "quinze", "dezesseis", "dezessete", "dezoito", "dezenove"},
		tens:     []string{"", "", "vinte", "trinta", "quarenta", "cinquenta", "sessanta", "setenta", "oitenta", "noventa"},
		hundred:  "cento",
		thousand: "mil",
		million:  "milhão",
		billion:  "bilhão",
	},
}

// localeSpecialCases holds special case strings for locale-specific number formatting
var localeSpecialCases = map[string]map[string]string{
	"es": { // Spanish special cases for 20-29
		"veinti": "veinti",
		"21":     "veintiuno",
		"22":     "veintidós",
		"23":     "veintitrés",
	},
	"it": { // Italian special cases
		"venti": "venti",
		"21":    "ventuno",
		"23":    "ventitré",
		"28":    "ventotto",
		"31":    "trentuno",
		"38":    "trentotto",
	},
	"fr": { // French special cases for 70-99
		"soixante-et-onze":  "soixante-et-onze",
		"quatre-vingt-un":   "quatre-vingt-un",
		"quatre-vingt-onze": "quatre-vingt-onze",
	},
}

// numericCodeToCurrencyCode is a reverse lookup map for O(1) currency code lookup by numeric code.
// It is initialized in init() function.
var numericCodeToCurrencyCode map[string]string

// init initializes the reverse lookup map for O(1) currency code lookups by numeric code.
func init() {
	numericCodeToCurrencyCode = make(map[string]string, len(CurrencyMap))
	for code, currency := range CurrencyMap {
		numericCodeToCurrencyCode[currency.NumericCode] = code
	}
}
