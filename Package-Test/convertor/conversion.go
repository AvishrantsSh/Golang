package convertor

// Function to convert Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// Function to convert Fahrenheit to Celsius
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
