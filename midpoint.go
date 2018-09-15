package main

import "math"

func rad2degr(rad float64) float64 {
	return rad * 180 / math.Pi
}

func degr2rad(degr float64) float64 {
	return degr * math.Pi / 180
}

func getLatLngCenter(coords [][]float64) (float64, float64) {
	var sumX, sumY, sumZ float64
	LATIDX := 0
	LNGIDX := 1

	for _, v := range coords {
		lat := degr2rad(v[LATIDX])
		lng := degr2rad(v[LNGIDX])

		sumX += math.Cos(lat) * math.Cos(lng)
		sumY += math.Cos(lat) * math.Sin(lng)
		sumZ += math.Sin(lat)
	}

	l := float64(len(coords))
	avgX := sumX / l
	avgY := sumY / l
	avgZ := sumZ / l

	lng := math.Atan2(avgY, avgX)
	hyp := math.Sqrt(avgX*avgX + avgY*avgY)
	lat := math.Atan2(avgZ, hyp)

	return rad2degr(lat), rad2degr(lng)
}
