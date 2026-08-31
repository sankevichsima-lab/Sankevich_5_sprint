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
	if len(datastring) == 0 {
		return fmt.Errorf("lenght of string = 0")
	}

	sliseStr := strings.Split(datastring, ",")

	if len(sliseStr) != 3 {
		return fmt.Errorf("lenght of slise is not 3")
	}

	steps, err := strconv.Atoi(sliseStr[0])

	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps<=0")
	}

	t.Steps = steps

	if len(sliseStr[1]) == 0 {
		return fmt.Errorf("lenght of name 0")
	}
	t.TrainingType = sliseStr[1]

	time, err := time.ParseDuration(sliseStr[2])

	if err != nil {
		return fmt.Errorf("invalid time format: %w", err)
	}

	if time <= 0 {
		return fmt.Errorf("time <= 0")
	}

	t.Duration = time
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var callWalk float64
	var err error
	if t.TrainingType == "Ходьба" {
		callWalk, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("invalid calorie format for walking: %w", err)
		}
	}
	if t.TrainingType == "Бег" {
		callWalk, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("invalid calorie format for running: %w", err)
		}
	}
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
		return "", fmt.Errorf("haven't this type of training")
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, callWalk)

	return result, nil
}
