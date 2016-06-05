package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type measurement struct {
	feet   int
	inches int
}

func newMeasurementFromString(desc string) (measurement, error) {
	m := measurement{}
	parts := strings.Split(desc, "-")
	if len(parts) != 2 {
		return m, errors.New("invalid measurement desc")
	}

	var err error
	m.feet, err = strconv.Atoi(parts[0])
	if err != nil {
		return m, err
	}
	m.inches, err = strconv.Atoi(parts[1])
	if err != nil {
		return m, err
	}

	return m, nil
}

func newMeasurementFromInches(inches int) measurement {
	m := measurement{}
	m.feet = inches / 12
	m.inches = inches % 12
	return m
}

func (m measurement) totalInches() int {
	return 12*m.feet + m.inches
}

type roomSize struct {
	height measurement
	width  measurement
}

func (r roomSize) sizeInInches() int {
	return r.height.totalInches() * r.width.totalInches()
}

func (r roomSize) size() measurement {
	return newMeasurementFromInches(r.sizeInInches())
}

func parseLine(line string) (roomSize, error) {
	r := roomSize{}
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return r, errors.New("invalid line: " + line)
	}

	var err error
	r.width, err = newMeasurementFromString(parts[0])
	if err != nil {
		return r, err
	}
	r.height, err = newMeasurementFromString(parts[1])
	if err != nil {
		return r, err
	}

	return r, nil
}

func getTotalSize(rooms []roomSize) measurement {
	totalInches := 0
	for _, r := range rooms {
		totalInches += r.sizeInInches()
	}
	return newMeasurementFromInches(totalInches)
}

func main() {

	var rooms []roomSize
	bio := bufio.NewReader(os.Stdin)
	for {
		line, _, err := bio.ReadLine()
		if err != nil {
			break
		}

		desc := string(line)
		r, err := parseLine(desc)
		if err != nil {
			fmt.Errorf("error: %s\n", err)
			continue
		}

		rooms = append(rooms, r)
	}

	totalSize := getTotalSize(rooms)
	fmt.Printf("Total Size: Feet: %d Inches: %d\n", totalSize.feet, totalSize.inches)
}
