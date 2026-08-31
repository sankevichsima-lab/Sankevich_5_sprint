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
	if len(datastring) == 0 {
		return fmt.Errorf("len of string = 0")
	}

	sliceStr := strings.Split(datastring, ",")

	if len(sliceStr) != 2 {
		return fmt.Errorf("lenght of slise is not 2")
	}

	steps, err := strconv.Atoi(sliceStr[0])

	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps<=0")
	}

	ds.Steps = steps

	time, err := time.ParseDuration(sliceStr[1])

	if err != nil {
		return fmt.Errorf("invalid time format: %w", err)
	}

	if time <= 0 {
		return fmt.Errorf("time <=0")
	}

	ds.Duration = time

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	if distance <= 0 {
		return "", fmt.Errorf("distance <=0")
	}

	callories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", fmt.Errorf("invalid calories format: %w", err)
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		callories)

	return result, nil
}
