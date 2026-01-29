package main

import (
	"fmt"
	"math"

	"dsp_go/pkg/filters"
)

func main() {
	// Параметры
	freq := 1000.0         // Целевая частота (Гц)
	samplingRate := 8000.0 // Частота дискретизации (Гц)
	totalN := 256          // Количество отсчетов

	// Создаем фильтр
	filter, err := filters.NewGoertzelFilter(freq, samplingRate, totalN)
	if err != nil {
		fmt.Printf("Error creating filter: %v\n", err)
		return
	}

	// Генерируем тестовый сигнал (синус нужной частоты)
	signal := make([]float64, totalN)
	for i := 0; i < totalN; i++ {
		signal[i] = math.Sin(2 * math.Pi * freq * float64(i) / samplingRate)
	}

	// Обрабатываем сигнал
	for _, sample := range signal {
		if err := filter.Process(sample); err != nil {
			fmt.Printf("Error processing sample: %v\n", err)
			return
		}
	}

	// Получаем результат
	magnitude, err := filter.GetMagnitude()
	if err != nil {
		fmt.Printf("Error getting magnitude: %v\n", err)
		return
	}

	fmt.Printf("Target frequency: %.2f Hz\n", freq)
	fmt.Printf("Detected magnitude: %.6f\n", magnitude)
	fmt.Printf("Expected magnitude: ~1.0\n")
	fmt.Printf("Filter is complete: %v\n", filter.IsComplete())
	fmt.Printf("Processed samples: %d/%d\n", filter.GetProcessedCount(), totalN)
}
