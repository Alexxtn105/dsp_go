package main

import (
	"dspgo/pkg/filters"
	"fmt"
	//"dsp_go/pkg/filters"
)

func main() {
	// Пример использования класса FIRFilter
	// Коэффициенты фильтра
	//coeffs := []float64{-0.1, 0.8, 0.3, -0.2}
	//coeffs := []float64{1, 0, 0, 0}// задержка на 0 отсчетов (сигнал как есть)
	coeffs := []float64{0, 1, 0, 0} // задержка на 1 отсчет (сигнал как есть)
	filter := filters.NewFIRFilter(coeffs)

	inputSamples := []float64{0.5, -0.3, 0.7, -0.2, 0.1, 0.6}
	fmt.Println("Отсчеты на выходе:")
	for _, sample := range inputSamples {
		result := filter.Tick(sample)
		fmt.Printf("%.3f\n", result)
	}

}
