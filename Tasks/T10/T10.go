package main

import "fmt"

func main() {
	tempRec := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	//map of slices with float values. key is the ceil temp of group
	groups := make(map[int][]float64)
	//If we define number of groups before "sorting", we could sort values by O(n)
	//It's possible with temperature values, because we know approx hi and lo temperature values.
	for _, temp := range tempRec {
		switch {
		case temp < -70:
			groups[-70] = append(groups[-70], temp)
		case temp > -70 && temp <= -60:
			groups[-60] = append(groups[-60], temp)
		case temp > -60 && temp <= -50:
			groups[-50] = append(groups[-50], temp)
		case temp > -50 && temp <= -40:
			groups[-40] = append(groups[-40], temp)
		case temp > -40 && temp <= -30:
			groups[-30] = append(groups[-30], temp)
		case temp > -30 && temp <= -20:
			groups[-20] = append(groups[-20], temp)
		case temp > -20 && temp <= -10:
			groups[-10] = append(groups[-10], temp)
		case temp > -10 && temp <= 0:
			groups[0] = append(groups[0], temp)
		case temp > 0 && temp <= 10:
			groups[10] = append(groups[10], temp)
		case temp > 10 && temp <= 20:
			groups[20] = append(groups[20], temp)
		case temp > 20 && temp <= 30:
			groups[30] = append(groups[30], temp)
		case temp > 30 && temp <= 40:
			groups[40] = append(groups[40], temp)
		case temp > 40 && temp <= 50:
			groups[50] = append(groups[50], temp)
		case temp > 50 && temp <= 60:
			groups[60] = append(groups[60], temp)
		case temp > 60 && temp <= 70:
			groups[70] = append(groups[70], temp)
		case temp > 70 && temp <= 80:
			groups[80] = append(groups[80], temp)
		case temp > 90:
			groups[90] = append(groups[90], temp)
		}
	}
	fmt.Println(groups)
}
