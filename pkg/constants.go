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
		NumericCode: "784",
		MinorUnit:   2,
	},
	"AFN": {
		NumericCode: "971",
		MinorUnit:   2,
	},
	"ALL": {
		NumericCode: "008",
		MinorUnit:   2,
	},
	"AMD": {
		NumericCode: "051",
		MinorUnit:   2,
	},
	"AOA": {
		NumericCode: "973",
		MinorUnit:   2,
	},
	"ARS": {
		NumericCode: "032",
		MinorUnit:   2,
	},
	"AUD": {
		NumericCode: "036",
		MinorUnit:   2,
	},
	"AWG": {
		NumericCode: "533",
		MinorUnit:   2,
	},
	"AZN": {
		NumericCode: "944",
		MinorUnit:   2,
	},
	"BAM": {
		NumericCode: "977",
		MinorUnit:   2,
	},
	"BBD": {
		NumericCode: "052",
		MinorUnit:   2,
	},
	"BDT": {
		NumericCode: "050",
		MinorUnit:   2,
	},
	"BGN": {
		NumericCode: "975",
		MinorUnit:   2,
	},
	"BHD": {
		NumericCode: "048",
		MinorUnit:   3,
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
		NumericCode: "986",
		MinorUnit:   2,
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
		NumericCode: "933",
		MinorUnit:   2,
	},
	"BZD": {
		NumericCode: "084",
		MinorUnit:   2,
	},
	"CAD": {
		NumericCode: "124",
		MinorUnit:   2,
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
		NumericCode: "756",
		MinorUnit:   2,
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
		NumericCode: "152",
		MinorUnit:   0,
	},
	"CNY": {
		NumericCode: "156",
		MinorUnit:   2,
	},
	"COP": {
		NumericCode: "170",
		MinorUnit:   2,
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
		NumericCode: "203",
		MinorUnit:   2,
	},
	"DJF": {
		NumericCode: "262",
		MinorUnit:   0,
	},
	"DKK": {
		NumericCode: "208",
		MinorUnit:   2,
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
		NumericCode: "818",
		MinorUnit:   2,
	},
	"ERN": {
		NumericCode: "232",
		MinorUnit:   2,
	},
	"ETB": {
		NumericCode: "230",
		MinorUnit:   2,
	},
	"EUR": {
		NumericCode: "978",
		MinorUnit:   2,
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
		NumericCode: "826",
		MinorUnit:   2,
	},
	"GEL": {
		NumericCode: "981",
		MinorUnit:   2,
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
		NumericCode: "344",
		MinorUnit:   2,
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
		NumericCode: "348",
		MinorUnit:   2,
	},
	"IDR": {
		NumericCode: "360",
		MinorUnit:   2,
	},
	"ILS": {
		NumericCode: "376",
		MinorUnit:   2,
	},
	"INR": {
		NumericCode: "356",
		MinorUnit:   2,
	},
	"IQD": {
		NumericCode: "368",
		MinorUnit:   3,
	},
	"IRR": {
		NumericCode: "364",
		MinorUnit:   2,
	},
	"ISK": {
		NumericCode: "352",
		MinorUnit:   0,
	},
	"JMD": {
		NumericCode: "388",
		MinorUnit:   2,
	},
	"JOD": {
		NumericCode: "400",
		MinorUnit:   3,
	},
	"JPY": {
		NumericCode: "392",
		MinorUnit:   0,
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
		NumericCode: "116",
		MinorUnit:   2,
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
		NumericCode: "410",
		MinorUnit:   0,
	},
	"KWD": {
		NumericCode: "414",
		MinorUnit:   3,
	},
	"KYD": {
		NumericCode: "136",
		MinorUnit:   2,
	},
	"KZT": {
		NumericCode: "398",
		MinorUnit:   2,
	},
	"LAK": {
		NumericCode: "418",
		MinorUnit:   2,
	},
	"LBP": {
		NumericCode: "422",
		MinorUnit:   2,
	},
	"LKR": {
		NumericCode: "144",
		MinorUnit:   2,
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
		NumericCode: "498",
		MinorUnit:   2,
	},
	"MGA": {
		NumericCode: "969",
		MinorUnit:   2,
	},
	"MKD": {
		NumericCode: "807",
		MinorUnit:   2,
	},
	"MMK": {
		NumericCode: "104",
		MinorUnit:   2,
	},
	"MNT": {
		NumericCode: "496",
		MinorUnit:   2,
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
		NumericCode: "484",
		MinorUnit:   2,
	},
	"MXV": {
		NumericCode: "979",
		MinorUnit:   2,
	},
	"MYR": {
		NumericCode: "458",
		MinorUnit:   2,
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
		NumericCode: "578",
		MinorUnit:   2,
	},
	"NPR": {
		NumericCode: "524",
		MinorUnit:   2,
	},
	"NZD": {
		NumericCode: "554",
		MinorUnit:   2,
	},
	"OMR": {
		NumericCode: "512",
		MinorUnit:   3,
	},
	"PAB": {
		NumericCode: "590",
		MinorUnit:   2,
	},
	"PEN": {
		NumericCode: "604",
		MinorUnit:   2,
	},
	"PGK": {
		NumericCode: "598",
		MinorUnit:   2,
	},
	"PHP": {
		NumericCode: "608",
		MinorUnit:   2,
	},
	"PKR": {
		NumericCode: "586",
		MinorUnit:   2,
	},
	"PLN": {
		NumericCode: "985",
		MinorUnit:   2,
	},
	"PYG": {
		NumericCode: "600",
		MinorUnit:   0,
	},
	"QAR": {
		NumericCode: "634",
		MinorUnit:   2,
	},
	"RON": {
		NumericCode: "946",
		MinorUnit:   2,
	},
	"RSD": {
		NumericCode: "941",
		MinorUnit:   2,
	},
	"RUB": {
		NumericCode: "643",
		MinorUnit:   2,
	},
	"RWF": {
		NumericCode: "646",
		MinorUnit:   0,
	},
	"SAR": {
		NumericCode: "682",
		MinorUnit:   2,
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
		NumericCode: "752",
		MinorUnit:   2,
	},
	"SGD": {
		NumericCode: "702",
		MinorUnit:   2,
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
		NumericCode: "764",
		MinorUnit:   2,
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
		NumericCode: "949",
		MinorUnit:   2,
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
		NumericCode: "980",
		MinorUnit:   2,
	},
	"UGX": {
		NumericCode: "800",
		MinorUnit:   0,
	},
	"USD": {
		NumericCode: "840",
		MinorUnit:   2,
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
		NumericCode: "860",
		MinorUnit:   2,
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
		NumericCode: "704",
		MinorUnit:   0,
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
		NumericCode: "710",
		MinorUnit:   2,
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
