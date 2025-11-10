package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsedString := strings.Split(datastring, ",")
	if len(parsedString) != 2 {
		return fmt.Errorf("invalid datastring: %s", datastring)
	}
	stepsCount, err := strconv.Atoi(parsedString[0])
	if err != nil {
		return fmt.Errorf("invalid parameter steps: %w", err)
	}
	if stepsCount <= 0 {
		return fmt.Errorf("invalid steps count - %d", stepsCount)
	}
	trainingDuration, err := time.ParseDuration(parsedString[1])
	if err != nil {
		return fmt.Errorf("invalid parameter duration: %w", err)
	}
	if trainingDuration <= 0 {
		return fmt.Errorf("invalid parameter duration: %.2f", trainingDuration.Hours())
	}
	ds.Steps = stepsCount
	ds.Duration = trainingDuration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("invalid value: spent calories - %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCal), nil
}
