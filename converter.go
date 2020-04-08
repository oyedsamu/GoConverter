package main

import (
	"fmt"
	"math"
)

//Converter : Struct for Converter
type Converter struct {
}

//type Named Types
type (
	Feet   float64
	Cm     float64
	Min    float64
	Sec    float64
	MilSec float64
	Cel    float64
	Fah    float64
	Rad    float64
	Deg    float64
	Kg     float64
	Pound  float64
	Km     float64
	Mile   float64
	Stone  float64
	Cal    float64
	Joule  float64
	Newton float64
)

//CentToFeet Coverter
func (conv Converter) CentToFeet(cm Cm) Feet {
	var res = Feet(cm / 100.48)
	return res
}

//FeetToCent Coverter
func (conv Converter) FeetToCent(feet Feet) Cm {
	var res = Cm(feet * 100.48)
	return res
}

//MinToSec Coverter
func (conv Converter) MinToSec(min Min) Sec {
	var res = Sec(min * 60)
	return res
}

//SecToMin Coverter
func (conv Converter) SecToMin(sec Sec) Min {
	var res = Min(sec / 60)
	return res
}

//SecToMilSec Coverter
func (conv Converter) SecToMilSec(sec Sec) MilSec {
	var res = MilSec(sec * 1000)
	return res
}

//MilSecToSec Coverter
func (conv Converter) MilSecToSec(milSec MilSec) Sec {
	var res = Sec(milSec / 1000)
	return res
}

//CelToFah Coverter
func (conv Converter) CelToFah(cel Cel) Fah {
	var res = Fah((cel * 9 / 5) + 32)
	return res
}

//FahToCel Coverter
func (conv Converter) FahToCel(fah Fah) Cel {
	var res = Cel((fah - 32) * 5 / 9)
	return res
}

//RadToDeg Coverter
func (conv Converter) RadToDeg(rad Rad) Deg {
	var res = Deg(rad * 180 / math.Pi)
	return res
}

//DegToRad Coverter
func (conv Converter) DegToRad(deg Deg) Rad {
	var res = Rad(deg * math.Pi / 180)
	return res
}

//KgToPound Coverter
func (conv Converter) KgToPound(kg Kg) Pound {
	var res = Pound(kg * 2.205)
	return res
}

//PoundToKg Coverter
func (conv Converter) PoundToKg(pound Pound) Kg {
	var res = Kg(pound / 2.205)
	return res
}

//MileToKm converter
func (conv Converter) MileToKm(mile Mile) Km {
	var res = Km(mile * 1.609344)
	return res
}

//KmToMile converter
func (conv Converter) KmToMile(km Km) Mile {
	var res = Mile(km / 1.609344)
	return res
}

//KgToStone Converter
func (conv Converter) KgToStone(kg Kg) Stone {
	var res = Stone(kg * 6.35029)
	return res
}

//StoneToKg Converter
func (conv Converter) StoneToKg(stone Stone) Kg {
	var res = Kg(stone / 6.35029)
	return res
}

//CalToJoule Converter
func (conv Converter) CalToJoule(cal Cal) Joule {
	var res = Joule(cal * 4.186)
	return res
}

//JouleToCal Converter
func (conv Converter) JouleToCal(joule Joule) Cal {
	var res = Cal(joule / 4.186)
	return res
}

//Pounds to Newton Converter
func (conv Converter) PoundToNewton(pound Pound) Newton {
	var res = Newton(pound * 4.448)
	return res
}

//Newton To Pounds Converter
func (conv Converter) NewtonToPound(newton Newton) Pound {
	var res = Pound(newton / 4.448)
	return res
}

func main() {
	conv := Converter{}
	fmt.Println(conv.CentToFeet(10))
	fmt.Println(conv.FeetToCent(10))
	fmt.Println(conv.MinToSec(10))
	fmt.Println(conv.SecToMin(10))
	fmt.Println(conv.SecToMilSec(10))
	fmt.Println(conv.MilSecToSec(10))
	fmt.Println(conv.SecToMilSec(10))
	fmt.Println(conv.MilSecToSec(10))
	fmt.Println(conv.CelToFah(10))
	fmt.Println(conv.FahToCel(10))
	fmt.Println(conv.RadToDeg(10))
	fmt.Println(conv.DegToRad(10))
	fmt.Println(conv.KgToPound(10))
	fmt.Println(conv.PoundToKg(10))
	fmt.Println(conv.MileToKm(10))
	fmt.Println(conv.KmToMile(10))
	fmt.Println(conv.KgToStone(10))
	fmt.Println(conv.StoneToKg(10))
	fmt.Println(conv.CalToJoule(10))
	fmt.Println(conv.JouleToCal(10))
	fmt.Println(conv.PoundToNewton(10))
	fmt.Println(conv.NewtonToPound(10))
}
