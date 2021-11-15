package testconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // 类型传换，不是函数调用，把c*9/5 + 32计算出的值传换为Fahrenheit类型，只会改变值类型，不会改变值本身
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}
