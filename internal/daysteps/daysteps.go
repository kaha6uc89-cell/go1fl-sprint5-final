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

func (ds *DaySteps) Parse(data string) error {
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return fmt.Errorf("invalid data format: expected 2 fields, got %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse steps: %v", err)
	}

	if steps <= 0 {
		return fmt.Errorf("invalid steps count: %d (must be positive)", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse duration: %v", err)
	}

	if duration <= 0 {
		return fmt.Errorf("invalid duration: %v (must be positive)", duration)
	}
	ds.Duration = duration

	return nil
}

func (ds *DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)


	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("failed to calculate calories: %v", err)
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories)

	return result, nil
}
