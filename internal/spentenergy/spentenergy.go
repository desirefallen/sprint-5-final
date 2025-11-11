package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid parameter steps: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid parameter weight: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid parameter height: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid parameter duration: %.2f", duration.Hours())
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	minInDuration := duration.Minutes()
	return (weight * meanSpeed * minInDuration) * walkingCaloriesCoefficient / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid parameter steps: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid parameter weight: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid parameter height: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid parameter duration: %.2f", duration.Hours())
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	minInDuration := duration.Minutes()
	return (weight * meanSpeed * minInDuration) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / float64(mInKm)
}
