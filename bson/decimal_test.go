﻿package bson_test

var decimalTestsJSON = `
{
    "description": "Decimal128",
    "valid": [
        {
            "description": "Special - Canonical NaN",
            "subject": "180000001364000000000000000000000000000000007C00",
            "string": "NaN",
            "extjson": {
                "d": { "$numberDecimal": "NaN" }
            }
        },
        {
            "description": "Special - Negative NaN",
            "subject": "18000000136400000000000000000000000000000000FC00",
            "string": "NaN"
        },
        {
            "description": "Special - Canonical SNaN",
            "subject": "180000001364000000000000000000000000000000007E00",
            "string": "NaN"
        },
        {
            "description": "Special - Negative SNaN",
            "subject": "18000000136400000000000000000000000000000000FE00",
            "string": "NaN"
        },
        {
            "description": "Special - NaN with a payload",
            "subject": "180000001364001200000000000000000000000000007E00",
            "string": "NaN"
        },
        {
            "description": "Special - Canonical Positive Infinity",
            "subject": "180000001364000000000000000000000000000000007800",
            "string": "Infinity",
            "extjson": {
                "d": { "$numberDecimal": "Infinity" }
            }
        },
        {
            "description": "Special - Canonical Negative Infinity",
            "subject": "18000000136400000000000000000000000000000000F800",
            "string": "-Infinity",
            "extjson": {
                "d": { "$numberDecimal": "-Infinity" }
            }
        },
        {
            "description": "Special - Invalid representation treated as 0",
            "subject": "180000001364000000000000000000000000000000106C00",
            "string": "0"
        },
        {
            "description": "Special - Invalid representation treated as -0",
            "subject": "18000000136400DCBA9876543210DEADBEEF00000010EC00",
            "string": "-0"
        },
        {
            "description": "Special - Invalid representation treated as 0E3",
            "subject": "18000000136400FFFFFFFFFFFFFFFFFFFFFFFFFFFF116C00",
            "string": "0E+3"
        },
        {
            "description": "Regular - Smallest",
            "subject": "18000000136400D204000000000000000000000000343000",
            "string": "0.001234",
            "extjson": {
                "d": { "$numberDecimal": "0.001234" }
            }
        },
        {
            "description": "Regular - Smallest with Trailing Zeros",
            "subject": "1800000013640040EF5A07000000000000000000002A3000",
            "string": "0.00123400000",
            "extjson": {
                "d": { "$numberDecimal": "0.00123400000" }
            }
        },
        {
            "description": "Regular - 0.1",
            "subject": "1800000013640001000000000000000000000000003E3000",
            "string": "0.1",
            "extjson": {
                "d": { "$numberDecimal": "0.1" }
            }
        },
        {
            "description": "Regular - 0.1234567890123456789012345678901234",
            "subject": "18000000136400F2AF967ED05C82DE3297FF6FDE3CFC2F00",
            "string": "0.1234567890123456789012345678901234",
            "extjson": {
                "d": { "$numberDecimal": "0.1234567890123456789012345678901234" }
            }
        },
        {
	    "MGO TEST": "MGO TEST",
            "description": "Regular - Adjusted Exponent Limit",
	    "subject": "18000000136400F2AF967ED05C82DE3297FF6FDE3CF22F00",
            "string": "0.000001234567890123456789012345678901234",
            "extjson": {
                "d": { "$numberDecimal": "0.1234567890123456789012345678901234" }
            }
        },
        {
	    "MGO TEST": "MGO TEST",
            "description": "Scientific - Adjusted Exponent Limit",
	    "subject": "18000000136400F2AF967ED05C82DE3297FF6FDE3CF02F00",
            "string": "1.234567890123456789012345678901234E-7",
            "extjson": {
                "d": { "$numberDecimal": "0.1234567890123456789012345678901234" }
            }
        },
        {
            "description": "Regular - 0",
            "subject": "180000001364000000000000000000000000000000403000",
            "string": "0",
            "extjson": {
                "d": { "$numberDecimal": "0" }
            }
        },
        {
            "description": "Regular - -0",
            "subject": "18000000136400000000000000000000000000000040B000",
            "string": "-0",
            "extjson": {
                "d": { "$numberDecimal": "-0" }
            }
        },
        {
            "description": "Regular - -0.0",
            "subject": "1800000013640000000000000000000000000000003EB000",
            "string": "-0.0",
            "extjson": {
                "d": { "$numberDecimal": "-0.0" }
            }
        },
        {
            "description": "Regular - 2",
            "subject": "180000001364000200000000000000000000000000403000",
            "string": "2",
            "extjson": {
                "d": { "$numberDecimal": "2" }
            }
        },
        {
            "description": "Regular - 2.000",
            "subject": "18000000136400D0070000000000000000000000003A3000",
            "string": "2.000",
            "extjson": {
                "d": { "$numberDecimal": "2.000" }
            }
        },
        {
            "description": "Regular - Largest",
            "subject": "18000000136400141A99BE1C000000000000000000403000",
            "string": "123456789012",
            "extjson": {
                "d": { "$numberDecimal": "123456789012" }
            }
        },
        {
            "description": "Scientific - Tiniest",
            "subject": "18000000136400FFFFFFFF638E8D37C087ADBE09ED010000",
            "string": "9.999999999999999999999999999999999E-6143",
            "extjson": {
                "d": { "$numberDecimal": "9.999999999999999999999999999999999E-6143" }
            }
        },
        {
            "description": "Scientific - Tiny",
            "subject": "180000001364000100000000000000000000000000000000",
            "string": "1E-6176",
            "extjson": {
                "d": { "$numberDecimal": "1E-6176" }
            }
        },
        {
            "description": "Scientific - Negative Tiny",
            "subject": "180000001364000100000000000000000000000000008000",
            "string": "-1E-6176",
            "extjson": {
                "d": { "$numberDecimal": "-1E-6176" }
            }
        },
        {
            "description": "Scientific - Fractional",
            "subject": "1800000013640064000000000000000000000000002CB000",
            "string": "-1.00E-8",
            "extjson": {
                "d": { "$numberDecimal": "-1.00E-8" }
            }
        },
        {
            "description": "Scientific - 0 with Exponent",
            "subject": "180000001364000000000000000000000000000000205F00",
            "string": "0E+6000",
            "extjson": {
                "d": { "$numberDecimal": "0E+6000" }
            }
        },
        {
            "description": "Scientific - 0 with Negative Exponent",
            "subject": "1800000013640000000000000000000000000000007A2B00",
            "string": "0E-611",
            "extjson": {
                "d": { "$numberDecimal": "0E-611" }
            }
        },
        {
            "description": "Scientific - No Decimal with Signed Exponent",
            "subject": "180000001364000100000000000000000000000000463000",
            "string": "1E+3",
            "extjson": {
                "d": { "$numberDecimal": "1E+3" }
            }
        },
        {
            "description": "Scientific - Trailing Zero",
            "subject": "180000001364001A04000000000000000000000000423000",
            "string": "1.050E+4",
            "extjson": {
                "d": { "$numberDecimal": "1.050E+4" }
            }
        },
        {
            "description": "Scientific - With Decimal",
            "subject": "180000001364006900000000000000000000000000423000",
            "string": "1.05E+3",
            "extjson": {
                "d": { "$numberDecimal": "1.05E+3" }
            }
        },
        {
            "description": "Scientific - Full",
            "subject": "18000000136400FFFFFFFFFFFFFFFFFFFFFFFFFFFF403000",
            "_string": "5.192296858534827628530496329220095E+33",
            "string": "5192296858534827628530496329220095",
            "extjson": {
                "d": { "$numberDecimal": "5.192296858534827628530496329220095E+33" }
            }
        },
        {
            "description": "Scientific - Large",
            "subject": "18000000136400000000000A5BC138938D44C64D31FE5F00",
            "string": "1.000000000000000000000000000000000E+6144",
            "extjson": {
                "d": { "$numberDecimal": "1.000000000000000000000000000000000E+6144" }
            }
        },
        {
            "description": "Scientific - Largest",
            "subject": "18000000136400FFFFFFFF638E8D37C087ADBE09EDFF5F00",
            "string": "9.999999999999999999999999999999999E+6144",
            "extjson": {
                "d": { "$numberDecimal": "9.999999999999999999999999999999999E+6144" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Exponent Normalization",
            "subject": "1800000013640064000000000000000000000000002CB000",
            "string": "-1.00E-8",
            "extjson": {
                "d": { "$numberDecimal": "-100E-10" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Unsigned Positive Exponent",
            "subject": "180000001364000100000000000000000000000000463000",
            "string": "1E+3",
            "extjson": {
                "d": { "$numberDecimal": "1E3" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Lowercase Exponent Identifier",
            "subject": "180000001364000100000000000000000000000000463000",
            "string": "1E+3",
            "extjson": {
                "d": { "$numberDecimal": "1e+3" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Long Significand with Exponent",
            "subject": "1800000013640079D9E0F9763ADA429D0200000000583000",
            "string": "1.2345689012345789012345E+34",
            "extjson": {
                "d": { "$numberDecimal": "12345689012345789012345E+12" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Positive Sign",
            "subject": "18000000136400F2AF967ED05C82DE3297FF6FDE3C403000",
            "_string": "1.234567890123456789012345678901234E+33",
            "string": "1234567890123456789012345678901234",
            "extjson": {
                "d": { "$numberDecimal": "+1234567890123456789012345678901234" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Long Significand",
            "subject": "18000000136400F2AF967ED05C82DE3297FF6FDE3C403000",
            "_string": "1.234567890123456789012345678901234E+33",
            "string": "1234567890123456789012345678901234",
            "extjson": {
                "d": { "$numberDecimal": "1234567890123456789012345678901234" }
            }
        },
        {
            "description": "Non-Canonical Parsing - Long Decimal String",
            "subject": "180000001364000100000000000000000000000000722800",
            "string": "1E-999",
            "extjson": {
                "d": { "$numberDecimal": ".000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001" }
            }
        }
    ],
    "parseErrors": [
        {
            "description": "Too many significand digits",
            "subject": "100000000000000000000000000000000000000000000000000000000001"
        },
        {
            "description": "Too many significand digits",
            "subject": "1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "description": "Too many significand digits",
            "subject": ".100000000000000000000000000000000000000000000000000000000000"
        },
        {
            "description": "Incomplete Exponent",
            "subject": "1e"
        },
        {
            "description": "Exponent at the beginning",
            "subject": "E01"
        },
        {
            "description": "Invalid NaN specification",
            "subject": "nan"
        },
        {
            "description": "Just a decimal place",
            "subject": "."
        },
        {
            "description": "2 decimal places",
            "subject": "..3"
        },
        {
            "description": "2 decimal places",
            "subject": ".13.3"
        },
        {
            "description": "2 decimal places",
            "subject": "1..3"
        },
        {
            "description": "2 decimal places",
            "subject": "1.3.4"
        },
        {
            "description": "2 decimal places",
            "subject": "1.34."
        },
        {
            "description": "Decimal with no digits",
            "subject": ".e"
        },
        {
            "description": "2 signs",
            "subject": "+-32.4"
        },
        {
            "description": "2 signs",
            "subject": "-+32.4"
        },
        {
            "description": "2 negative signs",
            "subject": "--32.4"
        },
        {
            "description": "2 negative signs",
            "subject": "-32.-4"
        },
        {
            "description": "End in negative sign",
            "subject": "32.0-"
        },
        {
            "description": "2 negative signs",
            "subject": "32.4E--21"
        },
        {
            "description": "2 negative signs",
            "subject": "32.4E-2-1"
        },
        {
            "description": "2 signs",
            "subject": "32.4E+-21"
        },
        {
            "description": "Empty string",
            "subject": ""
        },
        {
            "description": "Invalid",
            "subject": "E"
        },
        {
            "description": "Invalid",
            "subject": "invalid"
        },
        {
            "description": "Invalid",
            "subject": "i"
        },
        {
            "description": "Invalid",
            "subject": "in"
        },
        {
            "description": "Invalid",
            "subject": "-in"
        },
        {
            "description": "Invalid",
            "subject": "Na"
        },
        {
            "description": "Invalid",
            "subject": "-Na"
        },
        {
            "description": "Invalid",
            "subject": "1.23abc"
        },
        {
            "description": "Invalid",
            "subject": "1.23abcE+02"
        },
        {
            "description": "Invalid",
            "subject": "1.23E+0aabs2"
        }
    ]
}
`
