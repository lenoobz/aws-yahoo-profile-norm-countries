package consts

// Collection names
const (
	YAHOO_ASSET_PROFILES_COLLECTION = "yahoo_asset_profiles" // Should match with Colnames's key of AppConf
	ASSET_COUNTRIES_COLLECTION      = "asset_countries"
)

const (
	BOND     = "BOND"
	EQUITY   = "EQUITY"
	BALANCED = "BALANCED"
)

const (
	DATA_SOURCE = "YAHOO_FINANCE"
)

var Countries = []struct {
	Name       string
	Alpha2Code string
	Alpha3Code string
	NumberCode int
	Latitude   float64
	Longitude  float64
}{
	{
		Name:       "Afghanistan",
		Alpha2Code: "AF",
		Alpha3Code: "AFG",
		NumberCode: 4,
		Latitude:   33,
		Longitude:  65,
	},
	{
		Name:       "Albania",
		Alpha2Code: "AL",
		Alpha3Code: "ALB",
		NumberCode: 8,
		Latitude:   41,
		Longitude:  20,
	},
	{
		Name:       "Algeria",
		Alpha2Code: "DZ",
		Alpha3Code: "DZA",
		NumberCode: 12,
		Latitude:   28,
		Longitude:  3,
	},
	{
		Name:       "American Samoa",
		Alpha2Code: "AS",
		Alpha3Code: "ASM",
		NumberCode: 16,
		Latitude:   -14.3333,
		Longitude:  -170,
	},
	{
		Name:       "Andorra",
		Alpha2Code: "AD",
		Alpha3Code: "AND",
		NumberCode: 20,
		Latitude:   42.5,
		Longitude:  1.6,
	},
	{
		Name:       "Angola",
		Alpha2Code: "AO",
		Alpha3Code: "AGO",
		NumberCode: 24,
		Latitude:   -12.5,
		Longitude:  18.5,
	},
	{
		Name:       "Anguilla",
		Alpha2Code: "AI",
		Alpha3Code: "AIA",
		NumberCode: 660,
		Latitude:   18.25,
		Longitude:  -63.1667,
	},
	{
		Name:       "Antarctica",
		Alpha2Code: "AQ",
		Alpha3Code: "ATA",
		NumberCode: 10,
		Latitude:   -90,
		Longitude:  0,
	},
	{
		Name:       "Antigua and Barbuda",
		Alpha2Code: "AG",
		Alpha3Code: "ATG",
		NumberCode: 28,
		Latitude:   17.05,
		Longitude:  -61.8,
	},
	{
		Name:       "Argentina",
		Alpha2Code: "AR",
		Alpha3Code: "ARG",
		NumberCode: 32,
		Latitude:   -34,
		Longitude:  -64,
	},
	{
		Name:       "Armenia",
		Alpha2Code: "AM",
		Alpha3Code: "ARM",
		NumberCode: 51,
		Latitude:   40,
		Longitude:  45,
	},
	{
		Name:       "Aruba",
		Alpha2Code: "AW",
		Alpha3Code: "ABW",
		NumberCode: 533,
		Latitude:   12.5,
		Longitude:  -69.9667,
	},
	{
		Name:       "Australia",
		Alpha2Code: "AU",
		Alpha3Code: "AUS",
		NumberCode: 36,
		Latitude:   -27,
		Longitude:  133,
	},
	{
		Name:       "Austria",
		Alpha2Code: "AT",
		Alpha3Code: "AUT",
		NumberCode: 40,
		Latitude:   47.3333,
		Longitude:  13.3333,
	},
	{
		Name:       "Azerbaijan",
		Alpha2Code: "AZ",
		Alpha3Code: "AZE",
		NumberCode: 31,
		Latitude:   40.5,
		Longitude:  47.5,
	},
	{
		Name:       "Bahamas",
		Alpha2Code: "BS",
		Alpha3Code: "BHS",
		NumberCode: 44,
		Latitude:   24.25,
		Longitude:  -76,
	},
	{
		Name:       "Bahrain",
		Alpha2Code: "BH",
		Alpha3Code: "BHR",
		NumberCode: 48,
		Latitude:   26,
		Longitude:  50.55,
	},
	{
		Name:       "Bangladesh",
		Alpha2Code: "BD",
		Alpha3Code: "BGD",
		NumberCode: 50,
		Latitude:   24,
		Longitude:  90,
	},
	{
		Name:       "Barbados",
		Alpha2Code: "BB",
		Alpha3Code: "BRB",
		NumberCode: 52,
		Latitude:   13.1667,
		Longitude:  -59.5333,
	},
	{
		Name:       "Belarus",
		Alpha2Code: "BY",
		Alpha3Code: "BLR",
		NumberCode: 112,
		Latitude:   53,
		Longitude:  28,
	},
	{
		Name:       "Belgium",
		Alpha2Code: "BE",
		Alpha3Code: "BEL",
		NumberCode: 56,
		Latitude:   50.8333,
		Longitude:  4,
	},
	{
		Name:       "Belize",
		Alpha2Code: "BZ",
		Alpha3Code: "BLZ",
		NumberCode: 84,
		Latitude:   17.25,
		Longitude:  -88.75,
	},
	{
		Name:       "Benin",
		Alpha2Code: "BJ",
		Alpha3Code: "BEN",
		NumberCode: 204,
		Latitude:   9.5,
		Longitude:  2.25,
	},
	{
		Name:       "Bermuda",
		Alpha2Code: "BM",
		Alpha3Code: "BMU",
		NumberCode: 60,
		Latitude:   32.3333,
		Longitude:  -64.75,
	},
	{
		Name:       "Bhutan",
		Alpha2Code: "BT",
		Alpha3Code: "BTN",
		NumberCode: 64,
		Latitude:   27.5,
		Longitude:  90.5,
	},
	{
		Name:       "Bolivia, Plurinational State of",
		Alpha2Code: "BO",
		Alpha3Code: "BOL",
		NumberCode: 68,
		Latitude:   -17,
		Longitude:  -65,
	},
	{
		Name:       "Bolivia",
		Alpha2Code: "BO",
		Alpha3Code: "BOL",
		NumberCode: 68,
		Latitude:   -17,
		Longitude:  -65,
	},
	{
		Name:       "Bosnia and Herzegovina",
		Alpha2Code: "BA",
		Alpha3Code: "BIH",
		NumberCode: 70,
		Latitude:   44,
		Longitude:  18,
	},
	{
		Name:       "Botswana",
		Alpha2Code: "BW",
		Alpha3Code: "BWA",
		NumberCode: 72,
		Latitude:   -22,
		Longitude:  24,
	},
	{
		Name:       "Bouvet Island",
		Alpha2Code: "BV",
		Alpha3Code: "BVT",
		NumberCode: 74,
		Latitude:   -54.4333,
		Longitude:  3.4,
	},
	{
		Name:       "Brazil",
		Alpha2Code: "BR",
		Alpha3Code: "BRA",
		NumberCode: 76,
		Latitude:   -10,
		Longitude:  -55,
	},
	{
		Name:       "British Indian Ocean Territory",
		Alpha2Code: "IO",
		Alpha3Code: "IOT",
		NumberCode: 86,
		Latitude:   -6,
		Longitude:  71.5,
	},
	{
		Name:       "Brunei Darussalam",
		Alpha2Code: "BN",
		Alpha3Code: "BRN",
		NumberCode: 96,
		Latitude:   4.5,
		Longitude:  114.6667,
	},
	{
		Name:       "Brunei",
		Alpha2Code: "BN",
		Alpha3Code: "BRN",
		NumberCode: 96,
		Latitude:   4.5,
		Longitude:  114.6667,
	},
	{
		Name:       "Bulgaria",
		Alpha2Code: "BG",
		Alpha3Code: "BGR",
		NumberCode: 100,
		Latitude:   43,
		Longitude:  25,
	},
	{
		Name:       "Burkina Faso",
		Alpha2Code: "BF",
		Alpha3Code: "BFA",
		NumberCode: 854,
		Latitude:   13,
		Longitude:  -2,
	},
	{
		Name:       "Burundi",
		Alpha2Code: "BI",
		Alpha3Code: "BDI",
		NumberCode: 108,
		Latitude:   -3.5,
		Longitude:  30,
	},
	{
		Name:       "Cambodia",
		Alpha2Code: "KH",
		Alpha3Code: "KHM",
		NumberCode: 116,
		Latitude:   13,
		Longitude:  105,
	},
	{
		Name:       "Cameroon",
		Alpha2Code: "CM",
		Alpha3Code: "CMR",
		NumberCode: 120,
		Latitude:   6,
		Longitude:  12,
	},
	{
		Name:       "Canada",
		Alpha2Code: "CA",
		Alpha3Code: "CAN",
		NumberCode: 124,
		Latitude:   60,
		Longitude:  -95,
	},
	{
		Name:       "Cape Verde",
		Alpha2Code: "CV",
		Alpha3Code: "CPV",
		NumberCode: 132,
		Latitude:   16,
		Longitude:  -24,
	},
	{
		Name:       "Cayman Islands",
		Alpha2Code: "KY",
		Alpha3Code: "CYM",
		NumberCode: 136,
		Latitude:   19.5,
		Longitude:  -80.5,
	},
	{
		Name:       "Central African Republic",
		Alpha2Code: "CF",
		Alpha3Code: "CAF",
		NumberCode: 140,
		Latitude:   7,
		Longitude:  21,
	},
	{
		Name:       "Chad",
		Alpha2Code: "TD",
		Alpha3Code: "TCD",
		NumberCode: 148,
		Latitude:   15,
		Longitude:  19,
	},
	{
		Name:       "Chile",
		Alpha2Code: "CL",
		Alpha3Code: "CHL",
		NumberCode: 152,
		Latitude:   -30,
		Longitude:  -71,
	},
	{
		Name:       "China",
		Alpha2Code: "CN",
		Alpha3Code: "CHN",
		NumberCode: 156,
		Latitude:   35,
		Longitude:  105,
	},
	{
		Name:       "Christmas Island",
		Alpha2Code: "CX",
		Alpha3Code: "CXR",
		NumberCode: 162,
		Latitude:   -10.5,
		Longitude:  105.6667,
	},
	{
		Name:       "Cocos (Keeling) Islands",
		Alpha2Code: "CC",
		Alpha3Code: "CCK",
		NumberCode: 166,
		Latitude:   -12.5,
		Longitude:  96.8333,
	},
	{
		Name:       "Colombia",
		Alpha2Code: "CO",
		Alpha3Code: "COL",
		NumberCode: 170,
		Latitude:   4,
		Longitude:  -72,
	},
	{
		Name:       "Comoros",
		Alpha2Code: "KM",
		Alpha3Code: "COM",
		NumberCode: 174,
		Latitude:   -12.1667,
		Longitude:  44.25,
	},
	{
		Name:       "Congo",
		Alpha2Code: "CG",
		Alpha3Code: "COG",
		NumberCode: 178,
		Latitude:   -1,
		Longitude:  15,
	},
	{
		Name:       "Congo, the Democratic Republic of the",
		Alpha2Code: "CD",
		Alpha3Code: "COD",
		NumberCode: 180,
		Latitude:   0,
		Longitude:  25,
	},
	{
		Name:       "Cook Islands",
		Alpha2Code: "CK",
		Alpha3Code: "COK",
		NumberCode: 184,
		Latitude:   -21.2333,
		Longitude:  -159.7667,
	},
	{
		Name:       "Costa Rica",
		Alpha2Code: "CR",
		Alpha3Code: "CRI",
		NumberCode: 188,
		Latitude:   10,
		Longitude:  -84,
	},
	{
		Name:       "Côte d'Ivoire",
		Alpha2Code: "CI",
		Alpha3Code: "CIV",
		NumberCode: 384,
		Latitude:   8,
		Longitude:  -5,
	},
	{
		Name:       "Ivory Coast",
		Alpha2Code: "CI",
		Alpha3Code: "CIV",
		NumberCode: 384,
		Latitude:   8,
		Longitude:  -5,
	},
	{
		Name:       "Croatia",
		Alpha2Code: "HR",
		Alpha3Code: "HRV",
		NumberCode: 191,
		Latitude:   45.1667,
		Longitude:  15.5,
	},
	{
		Name:       "Cuba",
		Alpha2Code: "CU",
		Alpha3Code: "CUB",
		NumberCode: 192,
		Latitude:   21.5,
		Longitude:  -80,
	},
	{
		Name:       "Cyprus",
		Alpha2Code: "CY",
		Alpha3Code: "CYP",
		NumberCode: 196,
		Latitude:   35,
		Longitude:  33,
	},
	{
		Name:       "Czech Republic",
		Alpha2Code: "CZ",
		Alpha3Code: "CZE",
		NumberCode: 203,
		Latitude:   49.75,
		Longitude:  15.5,
	},
	{
		Name:       "Denmark",
		Alpha2Code: "DK",
		Alpha3Code: "DNK",
		NumberCode: 208,
		Latitude:   56,
		Longitude:  10,
	},
	{
		Name:       "Djibouti",
		Alpha2Code: "DJ",
		Alpha3Code: "DJI",
		NumberCode: 262,
		Latitude:   11.5,
		Longitude:  43,
	},
	{
		Name:       "Dominica",
		Alpha2Code: "DM",
		Alpha3Code: "DMA",
		NumberCode: 212,
		Latitude:   15.4167,
		Longitude:  -61.3333,
	},
	{
		Name:       "Dominican Republic",
		Alpha2Code: "DO",
		Alpha3Code: "DOM",
		NumberCode: 214,
		Latitude:   19,
		Longitude:  -70.6667,
	},
	{
		Name:       "Ecuador",
		Alpha2Code: "EC",
		Alpha3Code: "ECU",
		NumberCode: 218,
		Latitude:   -2,
		Longitude:  -77.5,
	},
	{
		Name:       "Egypt",
		Alpha2Code: "EG",
		Alpha3Code: "EGY",
		NumberCode: 818,
		Latitude:   27,
		Longitude:  30,
	},
	{
		Name:       "El Salvador",
		Alpha2Code: "SV",
		Alpha3Code: "SLV",
		NumberCode: 222,
		Latitude:   13.8333,
		Longitude:  -88.9167,
	},
	{
		Name:       "Equatorial Guinea",
		Alpha2Code: "GQ",
		Alpha3Code: "GNQ",
		NumberCode: 226,
		Latitude:   2,
		Longitude:  10,
	},
	{
		Name:       "Eritrea",
		Alpha2Code: "ER",
		Alpha3Code: "ERI",
		NumberCode: 232,
		Latitude:   15,
		Longitude:  39,
	},
	{
		Name:       "Estonia",
		Alpha2Code: "EE",
		Alpha3Code: "EST",
		NumberCode: 233,
		Latitude:   59,
		Longitude:  26,
	},
	{
		Name:       "Ethiopia",
		Alpha2Code: "ET",
		Alpha3Code: "ETH",
		NumberCode: 231,
		Latitude:   8,
		Longitude:  38,
	},
	{
		Name:       "Falkland Islands (Malvinas)",
		Alpha2Code: "FK",
		Alpha3Code: "FLK",
		NumberCode: 238,
		Latitude:   -51.75,
		Longitude:  -59,
	},
	{
		Name:       "Faroe Islands",
		Alpha2Code: "FO",
		Alpha3Code: "FRO",
		NumberCode: 234,
		Latitude:   62,
		Longitude:  -7,
	},
	{
		Name:       "Fiji",
		Alpha2Code: "FJ",
		Alpha3Code: "FJI",
		NumberCode: 242,
		Latitude:   -18,
		Longitude:  175,
	},
	{
		Name:       "Finland",
		Alpha2Code: "FI",
		Alpha3Code: "FIN",
		NumberCode: 246,
		Latitude:   64,
		Longitude:  26,
	},
	{
		Name:       "France",
		Alpha2Code: "FR",
		Alpha3Code: "FRA",
		NumberCode: 250,
		Latitude:   46,
		Longitude:  2,
	},
	{
		Name:       "French Guiana",
		Alpha2Code: "GF",
		Alpha3Code: "GUF",
		NumberCode: 254,
		Latitude:   4,
		Longitude:  -53,
	},
	{
		Name:       "French Polynesia",
		Alpha2Code: "PF",
		Alpha3Code: "PYF",
		NumberCode: 258,
		Latitude:   -15,
		Longitude:  -140,
	},
	{
		Name:       "French Southern Territories",
		Alpha2Code: "TF",
		Alpha3Code: "ATF",
		NumberCode: 260,
		Latitude:   -43,
		Longitude:  67,
	},
	{
		Name:       "Gabon",
		Alpha2Code: "GA",
		Alpha3Code: "GAB",
		NumberCode: 266,
		Latitude:   -1,
		Longitude:  11.75,
	},
	{
		Name:       "Gambia",
		Alpha2Code: "GM",
		Alpha3Code: "GMB",
		NumberCode: 270,
		Latitude:   13.4667,
		Longitude:  -16.5667,
	},
	{
		Name:       "Georgia",
		Alpha2Code: "GE",
		Alpha3Code: "GEO",
		NumberCode: 268,
		Latitude:   42,
		Longitude:  43.5,
	},
	{
		Name:       "Germany",
		Alpha2Code: "DE",
		Alpha3Code: "DEU",
		NumberCode: 276,
		Latitude:   51,
		Longitude:  9,
	},
	{
		Name:       "Ghana",
		Alpha2Code: "GH",
		Alpha3Code: "GHA",
		NumberCode: 288,
		Latitude:   8,
		Longitude:  -2,
	},
	{
		Name:       "Gibraltar",
		Alpha2Code: "GI",
		Alpha3Code: "GIB",
		NumberCode: 292,
		Latitude:   36.1833,
		Longitude:  -5.3667,
	},
	{
		Name:       "Greece",
		Alpha2Code: "GR",
		Alpha3Code: "GRC",
		NumberCode: 300,
		Latitude:   39,
		Longitude:  22,
	},
	{
		Name:       "Greenland",
		Alpha2Code: "GL",
		Alpha3Code: "GRL",
		NumberCode: 304,
		Latitude:   72,
		Longitude:  -40,
	},
	{
		Name:       "Grenada",
		Alpha2Code: "GD",
		Alpha3Code: "GRD",
		NumberCode: 308,
		Latitude:   12.1167,
		Longitude:  -61.6667,
	},
	{
		Name:       "Guadeloupe",
		Alpha2Code: "GP",
		Alpha3Code: "GLP",
		NumberCode: 312,
		Latitude:   16.25,
		Longitude:  -61.5833,
	},
	{
		Name:       "Guam",
		Alpha2Code: "GU",
		Alpha3Code: "GUM",
		NumberCode: 316,
		Latitude:   13.4667,
		Longitude:  144.7833,
	},
	{
		Name:       "Guatemala",
		Alpha2Code: "GT",
		Alpha3Code: "GTM",
		NumberCode: 320,
		Latitude:   15.5,
		Longitude:  -90.25,
	},
	{
		Name:       "Guernsey",
		Alpha2Code: "GG",
		Alpha3Code: "GGY",
		NumberCode: 831,
		Latitude:   49.5,
		Longitude:  -2.56,
	},
	{
		Name:       "Guinea",
		Alpha2Code: "GN",
		Alpha3Code: "GIN",
		NumberCode: 324,
		Latitude:   11,
		Longitude:  -10,
	},
	{
		Name:       "Guinea-Bissau",
		Alpha2Code: "GW",
		Alpha3Code: "GNB",
		NumberCode: 624,
		Latitude:   12,
		Longitude:  -15,
	},
	{
		Name:       "Guyana",
		Alpha2Code: "GY",
		Alpha3Code: "GUY",
		NumberCode: 328,
		Latitude:   5,
		Longitude:  -59,
	},
	{
		Name:       "Haiti",
		Alpha2Code: "HT",
		Alpha3Code: "HTI",
		NumberCode: 332,
		Latitude:   19,
		Longitude:  -72.4167,
	},
	{
		Name:       "Heard Island and McDonald Islands",
		Alpha2Code: "HM",
		Alpha3Code: "HMD",
		NumberCode: 334,
		Latitude:   -53.1,
		Longitude:  72.5167,
	},
	{
		Name:       "Holy See (Vatican City State)",
		Alpha2Code: "VA",
		Alpha3Code: "VAT",
		NumberCode: 336,
		Latitude:   41.9,
		Longitude:  12.45,
	},
	{
		Name:       "Honduras",
		Alpha2Code: "HN",
		Alpha3Code: "HND",
		NumberCode: 340,
		Latitude:   15,
		Longitude:  -86.5,
	},
	{
		Name:       "Hong Kong",
		Alpha2Code: "HK",
		Alpha3Code: "HKG",
		NumberCode: 344,
		Latitude:   22.25,
		Longitude:  114.1667,
	},
	{
		Name:       "Hungary",
		Alpha2Code: "HU",
		Alpha3Code: "HUN",
		NumberCode: 348,
		Latitude:   47,
		Longitude:  20,
	},
	{
		Name:       "Iceland",
		Alpha2Code: "IS",
		Alpha3Code: "ISL",
		NumberCode: 352,
		Latitude:   65,
		Longitude:  -18,
	},
	{
		Name:       "India",
		Alpha2Code: "IN",
		Alpha3Code: "IND",
		NumberCode: 356,
		Latitude:   20,
		Longitude:  77,
	},
	{
		Name:       "Indonesia",
		Alpha2Code: "ID",
		Alpha3Code: "IDN",
		NumberCode: 360,
		Latitude:   -5,
		Longitude:  120,
	},
	{
		Name:       "Iran, Islamic Republic of",
		Alpha2Code: "IR",
		Alpha3Code: "IRN",
		NumberCode: 364,
		Latitude:   32,
		Longitude:  53,
	},
	{
		Name:       "Iraq",
		Alpha2Code: "IQ",
		Alpha3Code: "IRQ",
		NumberCode: 368,
		Latitude:   33,
		Longitude:  44,
	},
	{
		Name:       "Ireland",
		Alpha2Code: "IE",
		Alpha3Code: "IRL",
		NumberCode: 372,
		Latitude:   53,
		Longitude:  -8,
	},
	{
		Name:       "Isle of Man",
		Alpha2Code: "IM",
		Alpha3Code: "IMN",
		NumberCode: 833,
		Latitude:   54.23,
		Longitude:  -4.55,
	},
	{
		Name:       "Israel",
		Alpha2Code: "IL",
		Alpha3Code: "ISR",
		NumberCode: 376,
		Latitude:   31.5,
		Longitude:  34.75,
	},
	{
		Name:       "Italy",
		Alpha2Code: "IT",
		Alpha3Code: "ITA",
		NumberCode: 380,
		Latitude:   42.8333,
		Longitude:  12.8333,
	},
	{
		Name:       "Jamaica",
		Alpha2Code: "JM",
		Alpha3Code: "JAM",
		NumberCode: 388,
		Latitude:   18.25,
		Longitude:  -77.5,
	},
	{
		Name:       "Japan",
		Alpha2Code: "JP",
		Alpha3Code: "JPN",
		NumberCode: 392,
		Latitude:   36,
		Longitude:  138,
	},
	{
		Name:       "Jersey",
		Alpha2Code: "JE",
		Alpha3Code: "JEY",
		NumberCode: 832,
		Latitude:   49.21,
		Longitude:  -2.13,
	},
	{
		Name:       "Jordan",
		Alpha2Code: "JO",
		Alpha3Code: "JOR",
		NumberCode: 400,
		Latitude:   31,
		Longitude:  36,
	},
	{
		Name:       "Kazakhstan",
		Alpha2Code: "KZ",
		Alpha3Code: "KAZ",
		NumberCode: 398,
		Latitude:   48,
		Longitude:  68,
	},
	{
		Name:       "Kenya",
		Alpha2Code: "KE",
		Alpha3Code: "KEN",
		NumberCode: 404,
		Latitude:   1,
		Longitude:  38,
	},
	{
		Name:       "Kiribati",
		Alpha2Code: "KI",
		Alpha3Code: "KIR",
		NumberCode: 296,
		Latitude:   1.4167,
		Longitude:  173,
	},
	{
		Name:       "Korea, Democratic People's Republic of",
		Alpha2Code: "KP",
		Alpha3Code: "PRK",
		NumberCode: 408,
		Latitude:   40,
		Longitude:  127,
	},
	{
		Name:       "Korea, Republic of",
		Alpha2Code: "KR",
		Alpha3Code: "KOR",
		NumberCode: 410,
		Latitude:   37,
		Longitude:  127.5,
	},
	{
		Name:       "South Korea",
		Alpha2Code: "KR",
		Alpha3Code: "KOR",
		NumberCode: 410,
		Latitude:   37,
		Longitude:  127.5,
	},
	{
		Name:       "Korea",
		Alpha2Code: "KR",
		Alpha3Code: "KOR",
		NumberCode: 410,
		Latitude:   37,
		Longitude:  127.5,
	},
	{
		Name:       "Kuwait",
		Alpha2Code: "KW",
		Alpha3Code: "KWT",
		NumberCode: 414,
		Latitude:   29.3375,
		Longitude:  47.6581,
	},
	{
		Name:       "Kyrgyzstan",
		Alpha2Code: "KG",
		Alpha3Code: "KGZ",
		NumberCode: 417,
		Latitude:   41,
		Longitude:  75,
	},
	{
		Name:       "Lao People's Democratic Republic",
		Alpha2Code: "LA",
		Alpha3Code: "LAO",
		NumberCode: 418,
		Latitude:   18,
		Longitude:  105,
	},
	{
		Name:       "Latvia",
		Alpha2Code: "LV",
		Alpha3Code: "LVA",
		NumberCode: 428,
		Latitude:   57,
		Longitude:  25,
	},
	{
		Name:       "Lebanon",
		Alpha2Code: "LB",
		Alpha3Code: "LBN",
		NumberCode: 422,
		Latitude:   33.8333,
		Longitude:  35.8333,
	},
	{
		Name:       "Lesotho",
		Alpha2Code: "LS",
		Alpha3Code: "LSO",
		NumberCode: 426,
		Latitude:   -29.5,
		Longitude:  28.5,
	},
	{
		Name:       "Liberia",
		Alpha2Code: "LR",
		Alpha3Code: "LBR",
		NumberCode: 430,
		Latitude:   6.5,
		Longitude:  -9.5,
	},
	{
		Name:       "Libyan Arab Jamahiriya",
		Alpha2Code: "LY",
		Alpha3Code: "LBY",
		NumberCode: 434,
		Latitude:   25,
		Longitude:  17,
	},
	{
		Name:       "Libya",
		Alpha2Code: "LY",
		Alpha3Code: "LBY",
		NumberCode: 434,
		Latitude:   25,
		Longitude:  17,
	},
	{
		Name:       "Liechtenstein",
		Alpha2Code: "LI",
		Alpha3Code: "LIE",
		NumberCode: 438,
		Latitude:   47.1667,
		Longitude:  9.5333,
	},
	{
		Name:       "Lithuania",
		Alpha2Code: "LT",
		Alpha3Code: "LTU",
		NumberCode: 440,
		Latitude:   56,
		Longitude:  24,
	},
	{
		Name:       "Luxembourg",
		Alpha2Code: "LU",
		Alpha3Code: "LUX",
		NumberCode: 442,
		Latitude:   49.75,
		Longitude:  6.1667,
	},
	{
		Name:       "Macao",
		Alpha2Code: "MO",
		Alpha3Code: "MAC",
		NumberCode: 446,
		Latitude:   22.1667,
		Longitude:  113.55,
	},
	{
		Name:       "Macedonia, the former Yugoslav Republic of",
		Alpha2Code: "MK",
		Alpha3Code: "MKD",
		NumberCode: 807,
		Latitude:   41.8333,
		Longitude:  22,
	},
	{
		Name:       "Madagascar",
		Alpha2Code: "MG",
		Alpha3Code: "MDG",
		NumberCode: 450,
		Latitude:   -20,
		Longitude:  47,
	},
	{
		Name:       "Malawi",
		Alpha2Code: "MW",
		Alpha3Code: "MWI",
		NumberCode: 454,
		Latitude:   -13.5,
		Longitude:  34,
	},
	{
		Name:       "Malaysia",
		Alpha2Code: "MY",
		Alpha3Code: "MYS",
		NumberCode: 458,
		Latitude:   2.5,
		Longitude:  112.5,
	},
	{
		Name:       "Maldives",
		Alpha2Code: "MV",
		Alpha3Code: "MDV",
		NumberCode: 462,
		Latitude:   3.25,
		Longitude:  73,
	},
	{
		Name:       "Mali",
		Alpha2Code: "ML",
		Alpha3Code: "MLI",
		NumberCode: 466,
		Latitude:   17,
		Longitude:  -4,
	},
	{
		Name:       "Malta",
		Alpha2Code: "MT",
		Alpha3Code: "MLT",
		NumberCode: 470,
		Latitude:   35.8333,
		Longitude:  14.5833,
	},
	{
		Name:       "Marshall Islands",
		Alpha2Code: "MH",
		Alpha3Code: "MHL",
		NumberCode: 584,
		Latitude:   9,
		Longitude:  168,
	},
	{
		Name:       "Martinique",
		Alpha2Code: "MQ",
		Alpha3Code: "MTQ",
		NumberCode: 474,
		Latitude:   14.6667,
		Longitude:  -61,
	},
	{
		Name:       "Mauritania",
		Alpha2Code: "MR",
		Alpha3Code: "MRT",
		NumberCode: 478,
		Latitude:   20,
		Longitude:  -12,
	},
	{
		Name:       "Mauritius",
		Alpha2Code: "MU",
		Alpha3Code: "MUS",
		NumberCode: 480,
		Latitude:   -20.2833,
		Longitude:  57.55,
	},
	{
		Name:       "Mayotte",
		Alpha2Code: "YT",
		Alpha3Code: "MYT",
		NumberCode: 175,
		Latitude:   -12.8333,
		Longitude:  45.1667,
	},
	{
		Name:       "Mexico",
		Alpha2Code: "MX",
		Alpha3Code: "MEX",
		NumberCode: 484,
		Latitude:   23,
		Longitude:  -102,
	},
	{
		Name:       "Micronesia, Federated States of",
		Alpha2Code: "FM",
		Alpha3Code: "FSM",
		NumberCode: 583,
		Latitude:   6.9167,
		Longitude:  158.25,
	},
	{
		Name:       "Moldova, Republic of",
		Alpha2Code: "MD",
		Alpha3Code: "MDA",
		NumberCode: 498,
		Latitude:   47,
		Longitude:  29,
	},
	{
		Name:       "Monaco",
		Alpha2Code: "MC",
		Alpha3Code: "MCO",
		NumberCode: 492,
		Latitude:   43.7333,
		Longitude:  7.4,
	},
	{
		Name:       "Mongolia",
		Alpha2Code: "MN",
		Alpha3Code: "MNG",
		NumberCode: 496,
		Latitude:   46,
		Longitude:  105,
	},
	{
		Name:       "Montenegro",
		Alpha2Code: "ME",
		Alpha3Code: "MNE",
		NumberCode: 499,
		Latitude:   42,
		Longitude:  19,
	},
	{
		Name:       "Montserrat",
		Alpha2Code: "MS",
		Alpha3Code: "MSR",
		NumberCode: 500,
		Latitude:   16.75,
		Longitude:  -62.2,
	},
	{
		Name:       "Morocco",
		Alpha2Code: "MA",
		Alpha3Code: "MAR",
		NumberCode: 504,
		Latitude:   32,
		Longitude:  -5,
	},
	{
		Name:       "Mozambique",
		Alpha2Code: "MZ",
		Alpha3Code: "MOZ",
		NumberCode: 508,
		Latitude:   -18.25,
		Longitude:  35,
	},
	{
		Name:       "Myanmar",
		Alpha2Code: "MM",
		Alpha3Code: "MMR",
		NumberCode: 104,
		Latitude:   22,
		Longitude:  98,
	},
	{
		Name:       "Burma",
		Alpha2Code: "MM",
		Alpha3Code: "MMR",
		NumberCode: 104,
		Latitude:   22,
		Longitude:  98,
	},
	{
		Name:       "Namibia",
		Alpha2Code: "NA",
		Alpha3Code: "NAM",
		NumberCode: 516,
		Latitude:   -22,
		Longitude:  17,
	},
	{
		Name:       "Nauru",
		Alpha2Code: "NR",
		Alpha3Code: "NRU",
		NumberCode: 520,
		Latitude:   -0.5333,
		Longitude:  166.9167,
	},
	{
		Name:       "Nepal",
		Alpha2Code: "NP",
		Alpha3Code: "NPL",
		NumberCode: 524,
		Latitude:   28,
		Longitude:  84,
	},
	{
		Name:       "Netherlands",
		Alpha2Code: "NL",
		Alpha3Code: "NLD",
		NumberCode: 528,
		Latitude:   52.5,
		Longitude:  5.75,
	},
	{
		Name:       "Netherlands Antilles",
		Alpha2Code: "AN",
		Alpha3Code: "ANT",
		NumberCode: 530,
		Latitude:   12.25,
		Longitude:  -68.75,
	},
	{
		Name:       "New Caledonia",
		Alpha2Code: "NC",
		Alpha3Code: "NCL",
		NumberCode: 540,
		Latitude:   -21.5,
		Longitude:  165.5,
	},
	{
		Name:       "New Zealand",
		Alpha2Code: "NZ",
		Alpha3Code: "NZL",
		NumberCode: 554,
		Latitude:   -41,
		Longitude:  174,
	},
	{
		Name:       "Nicaragua",
		Alpha2Code: "NI",
		Alpha3Code: "NIC",
		NumberCode: 558,
		Latitude:   13,
		Longitude:  -85,
	},
	{
		Name:       "Niger",
		Alpha2Code: "NE",
		Alpha3Code: "NER",
		NumberCode: 562,
		Latitude:   16,
		Longitude:  8,
	},
	{
		Name:       "Nigeria",
		Alpha2Code: "NG",
		Alpha3Code: "NGA",
		NumberCode: 566,
		Latitude:   10,
		Longitude:  8,
	},
	{
		Name:       "Niue",
		Alpha2Code: "NU",
		Alpha3Code: "NIU",
		NumberCode: 570,
		Latitude:   -19.0333,
		Longitude:  -169.8667,
	},
	{
		Name:       "Norfolk Island",
		Alpha2Code: "NF",
		Alpha3Code: "NFK",
		NumberCode: 574,
		Latitude:   -29.0333,
		Longitude:  167.95,
	},
	{
		Name:       "Northern Mariana Islands",
		Alpha2Code: "MP",
		Alpha3Code: "MNP",
		NumberCode: 580,
		Latitude:   15.2,
		Longitude:  145.75,
	},
	{
		Name:       "Norway",
		Alpha2Code: "NO",
		Alpha3Code: "NOR",
		NumberCode: 578,
		Latitude:   62,
		Longitude:  10,
	},
	{
		Name:       "Oman",
		Alpha2Code: "OM",
		Alpha3Code: "OMN",
		NumberCode: 512,
		Latitude:   21,
		Longitude:  57,
	},
	{
		Name:       "Pakistan",
		Alpha2Code: "PK",
		Alpha3Code: "PAK",
		NumberCode: 586,
		Latitude:   30,
		Longitude:  70,
	},
	{
		Name:       "Palau",
		Alpha2Code: "PW",
		Alpha3Code: "PLW",
		NumberCode: 585,
		Latitude:   7.5,
		Longitude:  134.5,
	},
	{
		Name:       "Palestinian Territory, Occupied",
		Alpha2Code: "PS",
		Alpha3Code: "PSE",
		NumberCode: 275,
		Latitude:   32,
		Longitude:  35.25,
	},
	{
		Name:       "Panama",
		Alpha2Code: "PA",
		Alpha3Code: "PAN",
		NumberCode: 591,
		Latitude:   9,
		Longitude:  -80,
	},
	{
		Name:       "Papua New Guinea",
		Alpha2Code: "PG",
		Alpha3Code: "PNG",
		NumberCode: 598,
		Latitude:   -6,
		Longitude:  147,
	},
	{
		Name:       "Paraguay",
		Alpha2Code: "PY",
		Alpha3Code: "PRY",
		NumberCode: 600,
		Latitude:   -23,
		Longitude:  -58,
	},
	{
		Name:       "Peru",
		Alpha2Code: "PE",
		Alpha3Code: "PER",
		NumberCode: 604,
		Latitude:   -10,
		Longitude:  -76,
	},
	{
		Name:       "Philippines",
		Alpha2Code: "PH",
		Alpha3Code: "PHL",
		NumberCode: 608,
		Latitude:   13,
		Longitude:  122,
	},
	{
		Name:       "Pitcairn",
		Alpha2Code: "PN",
		Alpha3Code: "PCN",
		NumberCode: 612,
		Latitude:   -24.7,
		Longitude:  -127.4,
	},
	{
		Name:       "Poland",
		Alpha2Code: "PL",
		Alpha3Code: "POL",
		NumberCode: 616,
		Latitude:   52,
		Longitude:  20,
	},
	{
		Name:       "Portugal",
		Alpha2Code: "PT",
		Alpha3Code: "PRT",
		NumberCode: 620,
		Latitude:   39.5,
		Longitude:  -8,
	},
	{
		Name:       "Puerto Rico",
		Alpha2Code: "PR",
		Alpha3Code: "PRI",
		NumberCode: 630,
		Latitude:   18.25,
		Longitude:  -66.5,
	},
	{
		Name:       "Qatar",
		Alpha2Code: "QA",
		Alpha3Code: "QAT",
		NumberCode: 634,
		Latitude:   25.5,
		Longitude:  51.25,
	},
	{
		Name:       "Réunion",
		Alpha2Code: "RE",
		Alpha3Code: "REU",
		NumberCode: 638,
		Latitude:   -21.1,
		Longitude:  55.6,
	},
	{
		Name:       "Romania",
		Alpha2Code: "RO",
		Alpha3Code: "ROU",
		NumberCode: 642,
		Latitude:   46,
		Longitude:  25,
	},
	{
		Name:       "Russian Federation",
		Alpha2Code: "RU",
		Alpha3Code: "RUS",
		NumberCode: 643,
		Latitude:   60,
		Longitude:  100,
	},
	{
		Name:       "Russia",
		Alpha2Code: "RU",
		Alpha3Code: "RUS",
		NumberCode: 643,
		Latitude:   60,
		Longitude:  100,
	},
	{
		Name:       "Rwanda",
		Alpha2Code: "RW",
		Alpha3Code: "RWA",
		NumberCode: 646,
		Latitude:   -2,
		Longitude:  30,
	},
	{
		Name:       "Saint Helena, Ascension and Tristan da Cunha",
		Alpha2Code: "SH",
		Alpha3Code: "SHN",
		NumberCode: 654,
		Latitude:   -15.9333,
		Longitude:  -5.7,
	},
	{
		Name:       "Saint Kitts and Nevis",
		Alpha2Code: "KN",
		Alpha3Code: "KNA",
		NumberCode: 659,
		Latitude:   17.3333,
		Longitude:  -62.75,
	},
	{
		Name:       "Saint Lucia",
		Alpha2Code: "LC",
		Alpha3Code: "LCA",
		NumberCode: 662,
		Latitude:   13.8833,
		Longitude:  -61.1333,
	},
	{
		Name:       "Saint Pierre and Miquelon",
		Alpha2Code: "PM",
		Alpha3Code: "SPM",
		NumberCode: 666,
		Latitude:   46.8333,
		Longitude:  -56.3333,
	},
	{
		Name:       "Saint Vincent and the Grenadines",
		Alpha2Code: "VC",
		Alpha3Code: "VCT",
		NumberCode: 670,
		Latitude:   13.25,
		Longitude:  -61.2,
	},
	{
		Name:       "Saint Vincent & the Grenadines",
		Alpha2Code: "VC",
		Alpha3Code: "VCT",
		NumberCode: 670,
		Latitude:   13.25,
		Longitude:  -61.2,
	},
	{
		Name:       "St. Vincent and the Grenadines",
		Alpha2Code: "VC",
		Alpha3Code: "VCT",
		NumberCode: 670,
		Latitude:   13.25,
		Longitude:  -61.2,
	},
	{
		Name:       "Samoa",
		Alpha2Code: "WS",
		Alpha3Code: "WSM",
		NumberCode: 882,
		Latitude:   -13.5833,
		Longitude:  -172.3333,
	},
	{
		Name:       "San Marino",
		Alpha2Code: "SM",
		Alpha3Code: "SMR",
		NumberCode: 674,
		Latitude:   43.7667,
		Longitude:  12.4167,
	},
	{
		Name:       "Sao Tome and Principe",
		Alpha2Code: "ST",
		Alpha3Code: "STP",
		NumberCode: 678,
		Latitude:   1,
		Longitude:  7,
	},
	{
		Name:       "Saudi Arabia",
		Alpha2Code: "SA",
		Alpha3Code: "SAU",
		NumberCode: 682,
		Latitude:   25,
		Longitude:  45,
	},
	{
		Name:       "Senegal",
		Alpha2Code: "SN",
		Alpha3Code: "SEN",
		NumberCode: 686,
		Latitude:   14,
		Longitude:  -14,
	},
	{
		Name:       "Serbia",
		Alpha2Code: "RS",
		Alpha3Code: "SRB",
		NumberCode: 688,
		Latitude:   44,
		Longitude:  21,
	},
	{
		Name:       "Seychelles",
		Alpha2Code: "SC",
		Alpha3Code: "SYC",
		NumberCode: 690,
		Latitude:   -4.5833,
		Longitude:  55.6667,
	},
	{
		Name:       "Sierra Leone",
		Alpha2Code: "SL",
		Alpha3Code: "SLE",
		NumberCode: 694,
		Latitude:   8.5,
		Longitude:  -11.5,
	},
	{
		Name:       "Singapore",
		Alpha2Code: "SG",
		Alpha3Code: "SGP",
		NumberCode: 702,
		Latitude:   1.3667,
		Longitude:  103.8,
	},
	{
		Name:       "Slovakia",
		Alpha2Code: "SK",
		Alpha3Code: "SVK",
		NumberCode: 703,
		Latitude:   48.6667,
		Longitude:  19.5,
	},
	{
		Name:       "Slovenia",
		Alpha2Code: "SI",
		Alpha3Code: "SVN",
		NumberCode: 705,
		Latitude:   46,
		Longitude:  15,
	},
	{
		Name:       "Solomon Islands",
		Alpha2Code: "SB",
		Alpha3Code: "SLB",
		NumberCode: 90,
		Latitude:   -8,
		Longitude:  159,
	},
	{
		Name:       "Somalia",
		Alpha2Code: "SO",
		Alpha3Code: "SOM",
		NumberCode: 706,
		Latitude:   10,
		Longitude:  49,
	},
	{
		Name:       "South Africa",
		Alpha2Code: "ZA",
		Alpha3Code: "ZAF",
		NumberCode: 710,
		Latitude:   -29,
		Longitude:  24,
	},
	{
		Name:       "South Georgia and the South Sandwich Islands",
		Alpha2Code: "GS",
		Alpha3Code: "SGS",
		NumberCode: 239,
		Latitude:   -54.5,
		Longitude:  -37,
	},
	{
		Name:       "South Sudan",
		Alpha2Code: "SS",
		Alpha3Code: "SSD",
		NumberCode: 728,
		Latitude:   8,
		Longitude:  30,
	},
	{
		Name:       "Spain",
		Alpha2Code: "ES",
		Alpha3Code: "ESP",
		NumberCode: 724,
		Latitude:   40,
		Longitude:  -4,
	},
	{
		Name:       "Sri Lanka",
		Alpha2Code: "LK",
		Alpha3Code: "LKA",
		NumberCode: 144,
		Latitude:   7,
		Longitude:  81,
	},
	{
		Name:       "Sudan",
		Alpha2Code: "SD",
		Alpha3Code: "SDN",
		NumberCode: 736,
		Latitude:   15,
		Longitude:  30,
	},
	{
		Name:       "Suriname",
		Alpha2Code: "SR",
		Alpha3Code: "SUR",
		NumberCode: 740,
		Latitude:   4,
		Longitude:  -56,
	},
	{
		Name:       "Svalbard and Jan Mayen",
		Alpha2Code: "SJ",
		Alpha3Code: "SJM",
		NumberCode: 744,
		Latitude:   78,
		Longitude:  20,
	},
	{
		Name:       "Swaziland",
		Alpha2Code: "SZ",
		Alpha3Code: "SWZ",
		NumberCode: 748,
		Latitude:   -26.5,
		Longitude:  31.5,
	},
	{
		Name:       "Sweden",
		Alpha2Code: "SE",
		Alpha3Code: "SWE",
		NumberCode: 752,
		Latitude:   62,
		Longitude:  15,
	},
	{
		Name:       "Switzerland",
		Alpha2Code: "CH",
		Alpha3Code: "CHE",
		NumberCode: 756,
		Latitude:   47,
		Longitude:  8,
	},
	{
		Name:       "Syrian Arab Republic",
		Alpha2Code: "SY",
		Alpha3Code: "SYR",
		NumberCode: 760,
		Latitude:   35,
		Longitude:  38,
	},
	{
		Name:       "Taiwan, Province of China",
		Alpha2Code: "TW",
		Alpha3Code: "TWN",
		NumberCode: 158,
		Latitude:   23.5,
		Longitude:  121,
	},
	{
		Name:       "Taiwan",
		Alpha2Code: "TW",
		Alpha3Code: "TWN",
		NumberCode: 158,
		Latitude:   23.5,
		Longitude:  121,
	},
	{
		Name:       "Tajikistan",
		Alpha2Code: "TJ",
		Alpha3Code: "TJK",
		NumberCode: 762,
		Latitude:   39,
		Longitude:  71,
	},
	{
		Name:       "Tanzania, United Republic of",
		Alpha2Code: "TZ",
		Alpha3Code: "TZA",
		NumberCode: 834,
		Latitude:   -6,
		Longitude:  35,
	},
	{
		Name:       "Thailand",
		Alpha2Code: "TH",
		Alpha3Code: "THA",
		NumberCode: 764,
		Latitude:   15,
		Longitude:  100,
	},
	{
		Name:       "Timor-Leste",
		Alpha2Code: "TL",
		Alpha3Code: "TLS",
		NumberCode: 626,
		Latitude:   -8.55,
		Longitude:  125.5167,
	},
	{
		Name:       "Togo",
		Alpha2Code: "TG",
		Alpha3Code: "TGO",
		NumberCode: 768,
		Latitude:   8,
		Longitude:  1.1667,
	},
	{
		Name:       "Tokelau",
		Alpha2Code: "TK",
		Alpha3Code: "TKL",
		NumberCode: 772,
		Latitude:   -9,
		Longitude:  -172,
	},
	{
		Name:       "Tonga",
		Alpha2Code: "TO",
		Alpha3Code: "TON",
		NumberCode: 776,
		Latitude:   -20,
		Longitude:  -175,
	},
	{
		Name:       "Trinidad and Tobago",
		Alpha2Code: "TT",
		Alpha3Code: "TTO",
		NumberCode: 780,
		Latitude:   11,
		Longitude:  -61,
	},
	{
		Name:       "Tunisia",
		Alpha2Code: "TN",
		Alpha3Code: "TUN",
		NumberCode: 788,
		Latitude:   34,
		Longitude:  9,
	},
	{
		Name:       "Turkey",
		Alpha2Code: "TR",
		Alpha3Code: "TUR",
		NumberCode: 792,
		Latitude:   39,
		Longitude:  35,
	},
	{
		Name:       "Turkmenistan",
		Alpha2Code: "TM",
		Alpha3Code: "TKM",
		NumberCode: 795,
		Latitude:   40,
		Longitude:  60,
	},
	{
		Name:       "Turks and Caicos Islands",
		Alpha2Code: "TC",
		Alpha3Code: "TCA",
		NumberCode: 796,
		Latitude:   21.75,
		Longitude:  -71.5833,
	},
	{
		Name:       "Tuvalu",
		Alpha2Code: "TV",
		Alpha3Code: "TUV",
		NumberCode: 798,
		Latitude:   -8,
		Longitude:  178,
	},
	{
		Name:       "Uganda",
		Alpha2Code: "UG",
		Alpha3Code: "UGA",
		NumberCode: 800,
		Latitude:   1,
		Longitude:  32,
	},
	{
		Name:       "Ukraine",
		Alpha2Code: "UA",
		Alpha3Code: "UKR",
		NumberCode: 804,
		Latitude:   49,
		Longitude:  32,
	},
	{
		Name:       "United Arab Emirates",
		Alpha2Code: "AE",
		Alpha3Code: "ARE",
		NumberCode: 784,
		Latitude:   24,
		Longitude:  54,
	},
	{
		Name:       "United Kingdom",
		Alpha2Code: "GB",
		Alpha3Code: "GBR",
		NumberCode: 826,
		Latitude:   54,
		Longitude:  -2,
	},
	{
		Name:       "United States",
		Alpha2Code: "US",
		Alpha3Code: "USA",
		NumberCode: 840,
		Latitude:   38,
		Longitude:  -97,
	},
	{
		Name:       "United States Minor Outlying Islands",
		Alpha2Code: "UM",
		Alpha3Code: "UMI",
		NumberCode: 581,
		Latitude:   19.2833,
		Longitude:  166.6,
	},
	{
		Name:       "Uruguay",
		Alpha2Code: "UY",
		Alpha3Code: "URY",
		NumberCode: 858,
		Latitude:   -33,
		Longitude:  -56,
	},
	{
		Name:       "Uzbekistan",
		Alpha2Code: "UZ",
		Alpha3Code: "UZB",
		NumberCode: 860,
		Latitude:   41,
		Longitude:  64,
	},
	{
		Name:       "Vanuatu",
		Alpha2Code: "VU",
		Alpha3Code: "VUT",
		NumberCode: 548,
		Latitude:   -16,
		Longitude:  167,
	},
	{
		Name:       "Venezuela, Bolivarian Republic of",
		Alpha2Code: "VE",
		Alpha3Code: "VEN",
		NumberCode: 862,
		Latitude:   8,
		Longitude:  -66,
	},
	{
		Name:       "Venezuela",
		Alpha2Code: "VE",
		Alpha3Code: "VEN",
		NumberCode: 862,
		Latitude:   8,
		Longitude:  -66,
	},
	{
		Name:       "Viet Nam",
		Alpha2Code: "VN",
		Alpha3Code: "VNM",
		NumberCode: 704,
		Latitude:   16,
		Longitude:  106,
	},
	{
		Name:       "Vietnam",
		Alpha2Code: "VN",
		Alpha3Code: "VNM",
		NumberCode: 704,
		Latitude:   16,
		Longitude:  106,
	},
	{
		Name:       "Virgin Islands, British",
		Alpha2Code: "VG",
		Alpha3Code: "VGB",
		NumberCode: 92,
		Latitude:   18.5,
		Longitude:  -64.5,
	},
	{
		Name:       "Virgin Islands British",
		Alpha2Code: "VG",
		Alpha3Code: "VGB",
		NumberCode: 92,
		Latitude:   18.5,
		Longitude:  -64.5,
	},
	{
		Name:       "British Virgin Islands",
		Alpha2Code: "VG",
		Alpha3Code: "VGB",
		NumberCode: 92,
		Latitude:   18.5,
		Longitude:  -64.5,
	},
	{
		Name:       "Virgin Islands, U.S.",
		Alpha2Code: "VI",
		Alpha3Code: "VIR",
		NumberCode: 850,
		Latitude:   18.3333,
		Longitude:  -64.8333,
	},
	{
		Name:       "Wallis and Futuna",
		Alpha2Code: "WF",
		Alpha3Code: "WLF",
		NumberCode: 876,
		Latitude:   -13.3,
		Longitude:  -176.2,
	},
	{
		Name:       "Western Sahara",
		Alpha2Code: "EH",
		Alpha3Code: "ESH",
		NumberCode: 732,
		Latitude:   24.5,
		Longitude:  -13,
	},
	{
		Name:       "Yemen",
		Alpha2Code: "YE",
		Alpha3Code: "YEM",
		NumberCode: 887,
		Latitude:   15,
		Longitude:  48,
	},
	{
		Name:       "Zambia",
		Alpha2Code: "ZM",
		Alpha3Code: "ZMB",
		NumberCode: 894,
		Latitude:   -15,
		Longitude:  30,
	},
	{
		Name:       "Zimbabwe",
		Alpha2Code: "ZW",
		Alpha3Code: "ZWE",
		NumberCode: 716,
		Latitude:   -20,
		Longitude:  30,
	},
}
