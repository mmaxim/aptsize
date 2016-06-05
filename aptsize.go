package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type distance int
type area int

type measurement struct {
	feet   int
	inches int
}

func (m measurement) String() string {
	return fmt.Sprintf("[ F: %d, I: %d ]", m.feet, m.inches)
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

func newMeasurementFromInches(inches distance) measurement {
	m := measurement{}
	m.feet = int(inches) / 12
	m.inches = int(inches) % 12
	return m
}

func newMeasurementFromSquareInches(inches area) measurement {
	m := measurement{}
	m.feet = int(inches) / 144
	m.inches = int(inches) % 144
	return m
}

func (m measurement) totalInches() distance {
	return distance(12*m.feet + m.inches)
}

type roomSize struct {
	height measurement
	width  measurement
}

func (r roomSize) String() string {
	return fmt.Sprintf("[ %s %s ]", r.width, r.height)
}

func (r roomSize) size() area {
	return area(r.height.totalInches() * r.width.totalInches())
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
	var totalInches area
	for _, r := range rooms {
		fmt.Printf("room: %s\n", r)
		totalInches += r.size()
	}
	return newMeasurementFromSquareInches(totalInches)
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
	fmt.Printf("Total Size (sqft): %s\n", totalSize)
}
