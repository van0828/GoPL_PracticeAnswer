package ch2_1

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvein { return Kelvein(c + 273.15) }

func KToC(k Kelvein) Celsius { return Celsius(k - 273.15) }

func KToF(k Kelvein) Fahrenheit { return CToF(KToC(k)) }

func FToK(f Fahrenheit) Kelvein { return CToK(FToC(f)) }
