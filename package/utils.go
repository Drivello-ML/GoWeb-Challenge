package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/GoWeb-Challenge/internal/domain"
)

func LoadTasksCsv(path string) ([]domain.Task, error) {

	var taskList []domain.Task

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	numberOfTasks := len(data)

	for i := 1; i < numberOfTasks; i++ {
		id, errId := strconv.Atoi(data[i][0])
		status, errStatus := StrToBool(data[i][2])
		if errId != nil || errStatus != nil {
			return []domain.Task{}, err
		}

		taskList = append(taskList, domain.Task{
			Id:          id,
			Description: data[i][1],
			Status:      status,
		})
	}

	return taskList, nil
}

func StrToBool(s string) (bool, error) {
	s = strings.ToLower(s) // Convierte la cadena a minúsculas para manejar "True" o "False" sin importar la capitalización

	switch s {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("invalid string to boolean")
	}
}
