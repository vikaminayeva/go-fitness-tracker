package spentenergy

import (
	"errors"
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
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid parameters")
	}
	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	caloriesWalk := calories * walkingCaloriesCoefficient
	return caloriesWalk, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input parameters")
	}
	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	speed := dist / duration.Hours()

	return speed
}

func Distance(steps int, height float64) float64 {
	lengthSteps := height * stepLengthCoefficient
	distanceMetr := float64(steps) * lengthSteps
	distanceKm := distanceMetr / mInKm
	return distanceKm
}
