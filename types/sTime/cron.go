package sTime

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type cron struct {
	active bool

	minute    []int
	hour      []int
	day       []int
	month     []int
	dayOfWeek []int

	location *time.Location
}

func NewCron(location string) (*cron, error) {
	c := &cron{
		minute:    []int{}, // Will be validated in Active
		hour:      []int{},
		day:       []int{},
		month:     []int{},
		dayOfWeek: []int{},
	}

	if err := c.SetLocation(location); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *cron) SetLocation(name string) error {
	if len(name) == 0 {
		c.location = time.UTC
		return nil
	}

	// Cargar la ubicación de Nueva York (soporta cambios automáticos entre EST y EDT)
	location, err := time.LoadLocation(name)
	if err != nil {
		return fmt.Errorf("time.LoadLocation(): %s", err)
	}

	c.location = location

	return nil
}

func (c *cron) IsActive() bool {
	return c.active
}

func (c *cron) Active() error {
	if c.active {
		return nil
	}

	if len(c.minute) == 0 {
		return fmt.Errorf("minute is empty")
	}
	if len(c.hour) == 0 {
		return fmt.Errorf("hour is empty")
	}
	if len(c.day) == 0 {
		return fmt.Errorf("day is empty")
	}
	if len(c.month) == 0 {
		return fmt.Errorf("month is empty")
	}
	if len(c.dayOfWeek) == 0 {
		return fmt.Errorf("dayOfWeek is empty")
	}

	c.active = true

	return nil
}

func (c *cron) Inactive() {
	c.active = false
}

func (c *cron) SetMinute(minute string) error {
	minutes, err := parseCronField(minute, 0, 59)
	if err != nil {
		return err
	}

	c.minute = minutes
	return nil
}

func (c *cron) SetHour(hour string) error {
	hours, err := parseCronField(hour, 0, 23)
	if err != nil {
		return err
	}

	c.hour = hours
	return nil
}

func (c *cron) SetDay(day string) error {
	days, err := parseCronField(day, 1, 31)
	if err != nil {
		return err
	}

	c.day = days
	return nil
}

func (c *cron) SetMonth(month string) error {
	months, err := parseCronField(month, 1, 12)
	if err != nil {
		return err
	}

	c.month = months
	return nil
}

// 0=Sunday, 6=Saturday
func (c *cron) SetDayOfWeek(dayOfWeek string) error {
	daysOfWeek, err := parseCronField(dayOfWeek, 0, 6)
	if err != nil {
		return err
	}

	c.dayOfWeek = daysOfWeek
	return nil
}

// spec: minute hour day month dayOfWeek
func (c *cron) SetSchedule(spec string) error {
	parts := strings.Fields(spec)
	if len(parts) != 5 {
		return fmt.Errorf("invalid cron spec: expected 5 fields")
	}
	if err := c.SetMinute(parts[0]); err != nil {
		return err
	}
	if err := c.SetHour(parts[1]); err != nil {
		return err
	}
	if err := c.SetDay(parts[2]); err != nil {
		return err
	}
	if err := c.SetMonth(parts[3]); err != nil {
		return err
	}
	return c.SetDayOfWeek(parts[4])
}

// ts: in UTC
func (c *cron) Next(ts int64) int64 {
	if !c.active {
		return 0
	}

	if ts == 0 {
		return 0
	}

	// Convert input UTC timestamp to local time
	t := time.Unix(ts, 0).In(c.location)

	// Start from the next second
	t = t.Add(time.Second)

	// Loop until we find a valid time
	for {
		// Reset smaller fields when a larger field increments
		currentMonth := int(t.Month())
		month, carry := nextValidValue(currentMonth, c.month)
		if carry {
			t = t.AddDate(1, 0, 0)
			t = time.Date(t.Year(), time.Month(month), 1, 0, 0, 0, 0, c.location)
		} else if month != currentMonth {
			t = time.Date(t.Year(), time.Month(month), 1, 0, 0, 0, 0, c.location)
		}

		currentDay := t.Day()
		day, carry := nextValidValue(currentDay, c.day)
		if carry {
			t = t.AddDate(0, 1, 0)
			t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, c.location)
			continue
		} else if day != currentDay {
			t = time.Date(t.Year(), t.Month(), day, 0, 0, 0, 0, c.location)
		}

		// Check day of week
		currentDayOfWeek := int(t.Weekday())
		dayOfWeek, carry := nextValidValue(currentDayOfWeek, c.dayOfWeek)
		if carry || dayOfWeek != currentDayOfWeek {
			// Advance to the next day that matches the day of week
			daysToAdd := (dayOfWeek - currentDayOfWeek + 7) % 7
			if daysToAdd == 0 {
				daysToAdd = 7
			}
			t = t.AddDate(0, 0, daysToAdd)
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, c.location)
			continue
		}

		currentHour := t.Hour()
		hour, carry := nextValidValue(currentHour, c.hour)
		if carry {
			t = t.AddDate(0, 0, 1)
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, c.location)
			continue
		} else if hour != currentHour {
			t = time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, c.location)
		}

		currentMinute := t.Minute()
		minute, carry := nextValidValue(currentMinute, c.minute)
		if carry {
			t = t.Add(time.Hour)
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, c.location)
			continue
		}

		// Set the minute
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, 0, 0, c.location)

		// Ensure the time is after the input timestamp
		if t.Unix() <= ts {
			t = t.Add(time.Minute)
			continue
		}

		// Verify the date is valid (e.g., not Feb 30)
		if t.Day() != day || int(t.Month()) != month {
			t = t.AddDate(0, 0, 1)
			t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, c.location)
			continue
		}

		// Return the result in UTC
		return t.UTC().Unix()
	}
}

// private functions

// Helper function to parse cron field (e.g., "*", "5", "*/2", "1,2,3")
func parseCronField(field string, min, max int) ([]int, error) {
	if field == "*" {
		values := make([]int, max-min+1)
		for i := min; i <= max; i++ {
			values[i-min] = i
		}
		return values, nil
	}
	if strings.HasPrefix(field, "*/") {
		step, err := strconv.Atoi(field[2:])
		if err != nil {
			return nil, fmt.Errorf("invalid step value: %s", err)
		}
		var values []int
		for i := min; i <= max; i += step {
			values = append(values, i)
		}
		return values, nil
	}
	if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		var values []int
		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil || val < min || val > max {
				return nil, fmt.Errorf("invalid value in list: %s", err)
			}
			values = append(values, val)
		}
		return values, nil
	}
	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format")
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil || start < min || start > max {
			return nil, fmt.Errorf("invalid range start: %s", err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil || end < min || end > max || end < start {
			return nil, fmt.Errorf("invalid range end: %s", err)
		}
		var values []int
		for i := start; i <= end; i++ {
			values = append(values, i)
		}
		return values, nil
	}
	val, err := strconv.Atoi(field)
	if err != nil || val < min || val > max {
		return nil, fmt.Errorf("invalid value: %s", err)
	}
	return []int{val}, nil
}

// Helper function to find the next valid value for a field
func nextValidValue(current int, allowed []int) (int, bool) {
	if len(allowed) == 0 {
		return 0, false
	}
	i := sort.SearchInts(allowed, current)
	if i < len(allowed) && allowed[i] == current {
		return current, false
	}
	if i == len(allowed) {
		return allowed[0], true
	}
	return allowed[i], false
}
