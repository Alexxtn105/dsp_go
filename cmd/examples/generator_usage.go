package main

import (
	"fmt"
	"github.com/Alexxtn105/dsp_go/pkg/generators"
	"math"
)

func main() {
	// Пример 1: Синусоидальный сигнал (по умолчанию)
	fmt.Println("=== Пример 1: Синусоидальный сигнал ===")
	gen1 := generators.NewReferenceSignalGenerator()
	gen1.Frequency = 440.0 // Ля первой октавы
	gen1.SampleRate = 44100.0
	gen1.TotalTime = 0.1

	signal1, err := gen1.Generate()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(gen1.Info())
		fmt.Println("\nПервые 20 отсчётов:")
		for i := 0; i < 20 && i < len(signal1); i++ {
			if i > 0 && i%5 == 0 {
				fmt.Println()
			}
			fmt.Printf("%+7.4f ", signal1[i])
		}
		fmt.Println("\n")
	}

	// Пример 2: Прямоугольный сигнал
	fmt.Println("=== Пример 2: Прямоугольный сигнал ===")
	gen2 := generators.NewReferenceSignalGenerator()
	gen2.SignalType = generators.Square
	gen2.Frequency = 100.0
	gen2.SampleRate = 10000.0
	gen2.TotalTime = 0.05
	gen2.DutyCycle = 0.3 // 30% заполнение

	signal2, err := gen2.Generate()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(gen2.Info())

		// Вывод одного периода
		samplesPerPeriod := int(math.Round(gen2.SampleRate / gen2.Frequency))
		if samplesPerPeriod > 0 && samplesPerPeriod <= len(signal2) {
			fmt.Printf("\nОдин период (%d отсчётов):\n", samplesPerPeriod)
			for i := 0; i < samplesPerPeriod; i++ {
				fmt.Printf("%+1.0f ", signal2[i])
			}
			fmt.Println()
		}
	}

	// Пример 3: Пилообразный сигнал с фазой
	fmt.Println("\n=== Пример 3: Пилообразный сигнал ===")
	gen3 := generators.NewReferenceSignalGenerator()
	gen3.SignalType = generators.Sawtooth
	gen3.Frequency = 50.0
	gen3.Phase = math.Pi / 2 // Сдвиг на 90 градусов
	gen3.TotalTime = 0.1

	_, err = gen3.Generate()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(gen3.Info())
	}

	// Пример 4: Треугольный сигнал
	fmt.Println("\n=== Пример 4: Треугольный сигнал ===")
	gen4 := generators.NewReferenceSignalGenerator()
	gen4.SignalType = generators.Triangle
	gen4.Frequency = 200.0
	gen4.TotalTime = 0.025

	signal4, err := gen4.Generate()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(gen4.Info())

		// Сравнение с синусом той же частоты
		gen4Sine := *gen4
		gen4Sine.SignalType = generators.Sine
		signal4Sine, _ := gen4Sine.Generate()

		fmt.Println("\nСравнение с синусом (первые 10 отсчётов):")
		fmt.Println("Треугольный   Синусоидальный")
		for i := 0; i < 10 && i < len(signal4); i++ {
			fmt.Printf("%+7.4f      %+7.4f\n", signal4[i], signal4Sine[i])
		}
	}

	// Пример 5: Косинусоидальный сигнал
	fmt.Println("\n=== Пример 5: Косинусоидальный сигнал ===")
	gen5 := generators.NewReferenceSignalGenerator()
	gen5.SignalType = generators.Cosine
	gen5.Frequency = 1000.0
	gen5.TotalTime = 0.01

	_, err = gen5.Generate()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(gen5.Info())
	}
}
