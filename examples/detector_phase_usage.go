package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"github.com/Alexxtn105/dsp/detectors"
)

func main() {
	// Запуск всех примеров
	Example1_BasicPhaseDetection()
	Example2_PhaseTracking()
	Example3_BPSKDemodulation()
	Example4_PhaseCorrection()

	// Дополнительный пример: работа с реальными данными
	fmt.Println("\n=== Дополнительно: Использование с вещественными сигналами ===")

	// Преобразование вещественного сигнала в комплексный
	frequency := 1000.0     // Гц
	samplingRate := 10000.0 // Гц
	duration := 0.01        // секунд

	// Генерация комплексного сигнала из синусоиды
	refSignal := make([]complex128, int(duration*samplingRate))
	for i := range refSignal {
		t := float64(i) / samplingRate
		// Преобразуем sin в комплексный сигнал (аналитический сигнал)
		refSignal[i] = complex(math.Sin(2*math.Pi*frequency*t),
			math.Cos(2*math.Pi*frequency*t))
	}

	// Используем первый отсчет как опорный
	if len(refSignal) > 0 {
		detector := detectors.NewCoherentPhaseDetector(refSignal[0], 0.1)

		// Проверяем фазу в середине сигнала
		midIdx := len(refSignal) / 2
		phaseError := detector.Detect(refSignal[midIdx])

		fmt.Printf("Фаза в середине сигнала: %.2f° относительно начала\n",
			phaseError*180.0/math.Pi)
	}
}

// Пример 1: Базовое использование для измерения фазы
func Example1_BasicPhaseDetection() {
	fmt.Println("=== Пример 1: Базовое измерение фазы ===")

	// Создаем опорный сигнал с фазой 30 градусов
	refPhase := 30.0 * math.Pi / 180.0
	referenceSignal := cmplx.Exp(complex(0, refPhase))

	// Создаем детектор с коэффициентом фильтрации 0.3
	detector := detectors.NewCoherentPhaseDetector(referenceSignal, 0.3)

	// Тестируем с разными входными сигналами
	testPhases := []float64{35.0, 25.0, 40.0, 30.0, 45.0}

	for i, degrees := range testPhases {
		// Создаем тестовый сигнал
		testPhase := degrees * math.Pi / 180.0
		testSignal := cmplx.Exp(complex(0, testPhase))

		// Измеряем ошибку фазы
		phaseError := detector.Detect(testSignal)

		fmt.Printf("Тест %d: Входная фаза: %.1f°, Ошибка фазы: %.3f°\n",
			i+1, degrees, phaseError*180.0/math.Pi)
	}
}

// Пример 2: Слежение за фазой с адаптивной коррекцией
func Example2_PhaseTracking() {
	fmt.Println("\n=== Пример 2: Слежение за фазой с коррекцией ===")

	// Опорный сигнал с фазой 0 градусов
	referenceSignal := complex(1, 0)
	detector := detectors.NewCoherentPhaseDetector(referenceSignal, 0.2)

	// Имитируем сигнал с дрейфом фазы
	timeSteps := 20
	for t := 0; t < timeSteps; t++ {
		// Сигнал с линейным дрейфом фазы и шумом
		basePhase := float64(t) * 5.0 * math.Pi / 180.0 // 5° за шаг
		noise := 0.1 * math.Sin(float64(t)*0.5)         // небольшой шум
		phase := basePhase + noise

		inputSignal := cmplx.Exp(complex(0, phase))

		// Детектируем фазу
		phaseError := detector.Detect(inputSignal)

		// Периодически корректируем смещение
		if t%5 == 0 && t > 0 {
			detector.UpdateOffset()
			fmt.Printf("Шаг %2d: Ошибка: %6.2f° -> Коррекция! Новое смещение: %6.2f°\n",
				t, phaseError*180.0/math.Pi, detector.GetPhaseOffset()*180.0/math.Pi)
		} else {
			fmt.Printf("Шаг %2d: Ошибка: %6.2f°\n",
				t, phaseError*180.0/math.Pi)
		}
	}
}

// Пример 3: Демодуляция фазомодулированного сигнала (BPSK)
func Example3_BPSKDemodulation() {
	fmt.Println("\n=== Пример 3: Демодуляция BPSK сигнала ===")

	// Опорный сигнал (несущая)
	referenceSignal := complex(1, 0)
	detector := detectors.NewCoherentPhaseDetector(referenceSignal, 0.1)

	// BPSK символы: 0 = 0°, 1 = 180°
	bpskSymbols := []float64{0, math.Pi, 0, math.Pi, 0, 0, math.Pi, math.Pi}

	// Добавляем фазовый сдвиг канала (например, 45°)
	channelPhaseShift := 45.0 * math.Pi / 180.0
	channelEffect := cmplx.Exp(complex(0, channelPhaseShift))

	fmt.Println("Переданные биты: 0 1 0 1 0 0 1 1")
	fmt.Print("Демодулированные: ")

	for i, phase := range bpskSymbols {
		// Создаем сигнал с фазовым сдвигом канала и небольшим шумом
		signal := cmplx.Exp(complex(0, phase)) * channelEffect

		// Добавляем шум
		noise := complex((math.Sin(float64(i)*0.7))*0.1,
			(math.Cos(float64(i)*0.3))*0.1)
		signal += noise

		// Детектируем фазу
		detectedPhase := detector.Detect(signal)

		// Корректируем ошибку (в реальной системе нужно накопить статистику)
		if i == 3 {
			detector.UpdateOffset()
		}

		// Демодуляция: если фаза ближе к 0°, то бит 0, если к 180°, то бит 1
		var bit int
		if math.Abs(detectedPhase) < math.Pi/2 {
			bit = 0
		} else {
			bit = 1
		}

		fmt.Printf("%d ", bit)
	}
	fmt.Println()
}

// Пример 4: Коррекция постоянного фазового сдвига
func Example4_PhaseCorrection() {
	fmt.Println("\n=== Пример 4: Коррекция постоянного фазового сдвига ===")

	// Опорный сигнал
	referenceSignal := cmplx.Exp(complex(0, math.Pi/4)) // 45°

	// Создаем детектор
	detector := detectors.NewCoherentPhaseDetector(referenceSignal, 0.15)

	// Имитируем сигнал с постоянным сдвигом -30°
	constantOffset := -30.0 * math.Pi / 180.0

	// Несколько измерений
	for i := 0; i < 10; i++ {
		// Сигнал с постоянным сдвигом и небольшими флуктуациями
		fluctuation := 0.05 * math.Sin(float64(i)*0.8)
		actualPhase := constantOffset + fluctuation
		inputSignal := cmplx.Exp(complex(0, actualPhase))

		// Детектируем
		error := detector.Detect(inputSignal)

		if i < 5 {
			fmt.Printf("Измерение %d: ошибка = %6.2f°\n",
				i+1, error*180.0/math.Pi)
		}
	}

	// Применяем коррекцию
	detector.UpdateOffset()

	fmt.Printf("\nПрименена коррекция. Новое смещение: %.2f°\n",
		detector.GetPhaseOffset()*180.0/math.Pi)

	// Проверяем после коррекции
	testSignal := cmplx.Exp(complex(0, constantOffset))
	finalError := detector.Detect(testSignal)
	fmt.Printf("Ошибка после коррекции: %.2f°\n", finalError*180.0/math.Pi)
}
