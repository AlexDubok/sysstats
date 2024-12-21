package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetTemperature() (float64, error) {
	// Path to the file containing CPU temperature
	const tempFilePath = "/sys/class/thermal/thermal_zone0/temp"

	// Read the content of the temperature file
	data, err := os.ReadFile(tempFilePath)
	if err != nil {
		return 0, err
	}

	// Convert the data to a string and trim whitespace
	tempStr := strings.TrimSpace(string(data))

	// Parse the temperature as an integer (in millidegrees Celsius)
	tempMilli, err := strconv.Atoi(tempStr)
	if err != nil {
		return 0, err
	}

	// Convert to degrees Celsius
	tempCelsius := float64(tempMilli) / 1000.0
	return tempCelsius, nil
}
