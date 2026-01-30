package main

import (
	"fmt"
	"math"

	"github.com/Alexxtn105/dsp/filters"
)

func main() {
	fmt.Println("=== Примеры использования КИХ-фильтра ===\n")

	// Пример 1: Простой фильтр
	fmt.Println("1. Простой фильтр с коэффициентами [1, 2, 3]:")
	coeffs1 := []float64{1.0, 2.0, 3.0}
	filter1 := filters.NewFIRFilter(coeffs1)

	// Импульсный отклик
	fmt.Println("   Импульсный отклик:")
	for i := 0; i < 5; i++ {
		var input float64
		if i == 0 {
			input = 1.0
		}
		output := filter1.Tick(input)
		fmt.Printf("   x[%d] = %5.1f → y[%d] = %5.1f\n", i, input, i, output)
	}

	// Пример 2: Фильтр скользящего среднего
	fmt.Println("\n2. Фильтр скользящего среднего (окно 5):")
	n := 5
	coeffs2 := make([]float64, n)
	for i := range coeffs2 {
		coeffs2[i] = 1.0 / float64(n)
	}

	filter2 := filters.NewFIRFilter(coeffs2)

	// Постоянный сигнал
	fmt.Println("   Постоянный сигнал (значение = 10.0):")
	for i := 0; i < 10; i++ {
		output := filter2.Tick(10.0)
		fmt.Printf("   Шаг %d: выход = %5.2f\n", i, output)
	}

	// Пример 3: Подавление высокочастотного шума
	fmt.Println("\n3. Фильтр низких частот (синус + шум):")
	// Простой НЧ-фильтр
	coeffs3 := []float64{0.1, 0.2, 0.4, 0.2, 0.1}
	filter3 := filters.NewFIRFilter(coeffs3)

	fmt.Println("   Синусоида 0.1 Гц + шум:")
	for i := 0; i < 20; i++ {
		// Сигнал: синус + шум
		signal := math.Sin(2*math.Pi*0.1*float64(i)) + 0.3*(math.Sin(2*math.Pi*0.5*float64(i)))
		filtered := filter3.Tick(signal)

		if i < 5 || i > 15 {
			fmt.Printf("   t=%2d: сигнал=%6.3f, фильтр=%6.3f\n", i, signal, filtered)
		}
	}

	// Пример 4: Нейтральный фильтр (задержка)
	fmt.Println("\n4. Нейтральный фильтр с задержкой на 2 отсчета:")
	coeffs4 := []float64{0, 0, 1}
	filter4 := filters.NewFIRFilter(coeffs4)

	fmt.Println("   Исходный сигнал и выход с задержкой:")
	for i := 0; i < 6; i++ {
		input := float64(i * 10)
		output := filter4.Tick(input)
		fmt.Printf("   t=%d: вход=%5.1f, выход=%5.1f\n", i, input, output)
	}

	// Пример 5: Сброс фильтра
	fmt.Println("\n5. Демонстрация сброса фильтра:")
	filter5 := filters.NewFIRFilter([]float64{0.5, 0.3, 0.2})

	// Обрабатываем часть сигнала
	for i := 0; i < 3; i++ {
		filter5.Tick(1.0)
	}

	// Сбрасываем
	filter5.Reset()

	// После сброса фильтр как новый
	output := filter5.Tick(1.0)
	fmt.Printf("   После Reset, вход=1.0 → выход=%5.2f (ожидается 0.5)\n", output)

	fmt.Println("\n=== Все примеры завершены ===")
}
