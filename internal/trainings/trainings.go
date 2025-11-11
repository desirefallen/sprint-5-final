package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parsedString := strings.Split(datastring, ",")
	if len(parsedString) != 3 {
		return fmt.Errorf("invalid datastring: %s", datastring)
	}
	stepsCount, err := strconv.Atoi(parsedString[0])
	if err != nil {
		return fmt.Errorf("invalid parameter steps: %w", err)
	}
	if stepsCount <= 0 {
		return fmt.Errorf("invalid steps count - %d", stepsCount)
	}
	t.Steps = stepsCount
	t.TrainingType = parsedString[1]
	trainingDuration, err := time.ParseDuration(parsedString[2])
	if err != nil {
		return fmt.Errorf("invalid parameter duration: %w", err)
	}
	if trainingDuration <= 0 {
		return fmt.Errorf("invalid parameter duration: %.2f", trainingDuration.Hours())
	}
	t.Duration = trainingDuration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	trainingDistance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		spentCal, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error: %w", err)
		}
		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2F км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.Duration.Hours(), trainingDistance, meanSpeed, spentCal), nil
	case "Ходьба":
		spentCal, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error: %w", err)
		}
		return fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2F км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.Duration.Hours(), trainingDistance, meanSpeed, spentCal), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
