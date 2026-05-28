package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000.0 // Количество метров в километре
	minInH                     = 60.0
	stepLengthCoefficient      = 0.45   // Средняя длина шага в метрах для среднего роста человека
	walkingMET                 = 3.5    // Метаболическая эквивалентная трата (MET) при ходьбе
	runningMET                 = 8.0    // Метаболическая эквивалентная трата (MET) при беге
	kcalPerMETMinutePerKg      = 0.0175 // Количество килокалорий, сжигаемых одной MET за минуту для одного кг тела
	walkingCaloriesCoefficient = 0.5
)

func Distance(steps int, height float64) float64 {
	if steps < 0 || height <= 0 {
		return 0.0
	}

	stepLength := height * stepLengthCoefficient
	totalDistance := float64(steps) * stepLength / mInKm 

	return totalDistance
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	if steps < 0 {
		return 0.0
	}

	if height <= 0 {
		return 0.0
	}

	distance := Distance(steps, height)

	durationHours := duration.Hours()

	if durationHours == 0 {
		return 0.0
	}

	meanSpeed := distance / durationHours

	return meanSpeed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0.0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	if meanSpeed == 0 {
		return 0.0, fmt.Errorf("calculated speed is zero, cannot calculate calories")
	}

	durationInMinutes := duration.Minutes()


	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0.0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	meanSpeed := MeanSpeed(steps, height, duration)


	if meanSpeed == 0 {
		return 0.0, fmt.Errorf("calculated speed is zero, cannot calculate calories")
	}


	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}
