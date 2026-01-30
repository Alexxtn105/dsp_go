package main

import (
	"fmt"
	"math"

	"github.com/Alexxtn105/dsp/filters"
)

func main() {
	fmt.Println("=== Примеры использования БИХ-фильтров ===\n")

	// Пример 1: ФНЧ 1-го порядка
	fmt.Println("1. Фильтр низких частот 1-го порядка (fc=0.1):")
	lpf1 := filters.NewFirstOrderLowPass(0.1)

	fmt.Println("   Коэффициенты:")
	fmt.Printf("   b = %v\n", lpf1.GetBCoeffs())
	fmt.Printf("   a = %v\n", lpf1.GetACoeffs())
	fmt.Printf("   Устойчив: %v\n", lpf1.IsStable())

	// Синусоидальный сигнал
	fmt.Println("\n   Синус 0.05 Гц (ниже частоты среза):")
	lpf1.Reset()
	for i := 0; i < 20; i++ {
		signal := math.Sin(2 * math.Pi * 0.05 * float64(i))
		filtered := lpf1.Tick(signal)
		if i < 5 {
			fmt.Printf("   t=%2d: вход=%6.3f, выход=%6.3f\n", i, signal, filtered)
		}
	}

	// Пример 2: ФВЧ 1-го порядка
	fmt.Println("\n2. Фильтр высоких частот 1-го порядка (fc=0.2):")
	hpf1 := filters.NewFirstOrderHighPass(0.2)

	fmt.Println("   Частотная характеристика:")
	freqs := []float64{0.0, 0.1, 0.2, 0.3, 0.4}
	for _, f := range freqs {
		h := hpf1.GetFrequencyResponse(f)
		gain := math.Sqrt(real(h)*real(h) + imag(h)*imag(h))
		phase := math.Atan2(imag(h), real(h)) * 180 / math.Pi
		fmt.Printf("   f=%.2f: усиление=%5.3f, фаза=%6.1f°\n", f, gain, phase)
	}

	// Пример 3: Полосовой фильтр 2-го порядка
	fmt.Println("\n3. Полосовой фильтр 2-го порядка (fc=0.25, Q=5):")
	bpf := filters.NewSecondOrderBandPass(0.25, 5.0)

	// Создаем тестовый сигнал: сумма трех синусоид
	fmt.Println("   Сумма синусоид 0.1, 0.25, 0.4 Гц:")
	bpf.Reset()

	for i := 0; i < 50; i++ {
		signal := math.Sin(2*math.Pi*0.1*float64(i)) + // Низкая частота
			math.Sin(2*math.Pi*0.25*float64(i)) + // Центральная частота
			math.Sin(2*math.Pi*0.4*float64(i)) // Высокая частота

		filtered := bpf.Tick(signal)

		if i >= 20 && i < 30 { // Показываем установившийся режим
			fmt.Printf("   t=%2d: вход=%6.3f, выход=%6.3f\n", i, signal, filtered)
		}
	}

	// Пример 4: Групповая задержка
	fmt.Println("\n4. Групповая задержка ФНЧ 2-го порядка:")
	lpf2 := filters.NewSecondOrderLowPass(0.1, 0.707)

	fmt.Println("   Групповая задержка на разных частотах:")
	for _, f := range []float64{0.01, 0.05, 0.1, 0.2} {
		delay := lpf2.GetGroupDelay(f)
		fmt.Printf("   f=%.2f: задержка=%6.3f отсчетов\n", f, delay)
	}

	// Пример 5: Обработка среза данных
	fmt.Println("\n5. Обработка всего среза данных:")

	// Создаем тестовый сигнал
	signal := make([]float64, 30)
	for i := range signal {
		signal[i] = math.Sin(2*math.Pi*0.1*float64(i)) + 0.3*math.Sin(2*math.Pi*0.4*float64(i))
	}

	// Создаем и применяем фильтр
	lpf := filters.NewFirstOrderLowPass(0.15)
	filtered := lpf.Process(signal)

	fmt.Println("   Первые 10 отсчетов:")
	for i := 0; i < 10; i++ {
		fmt.Printf("   [%2d] вход=%6.3f, выход=%6.3f\n", i, signal[i], filtered[i])
	}

	fmt.Println("\n=== Все примеры завершены ===")
}
