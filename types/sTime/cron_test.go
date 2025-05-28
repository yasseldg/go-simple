package sTime

import (
	"testing"
	"time"
)

func TestNewCron(t *testing.T) {
	tests := []struct {
		name     string
		location string
		wantErr  bool
	}{
		{"UTC", "UTC", false},
		{"America/New_York", "America/New_York", false},
		{"Invalid", "Invalid/Location", true},
		{"Empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var loc *time.Location
			var err error

			if tt.location == "" {
				loc = time.UTC
			} else {
				loc, err = time.LoadLocation(tt.location)
				if (err != nil) != tt.wantErr {
					t.Errorf("time.LoadLocation() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantErr {
					return
				}
			}

			cron := NewCron(loc)
			if cron == nil {
				t.Error("NewCron() returned nil when no error expected")
			}
		})
	}
}

func TestCronSetFields(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	cron := NewCron(loc)

	tests := []struct {
		name    string
		setter  func(string) error
		value   string
		wantErr bool
	}{
		{"Valid Minute", cron.SetMinute, "*/5", false},
		{"Invalid Minute", cron.SetMinute, "70", true},
		{"Valid Hour", cron.SetHour, "0-12", false},
		{"Invalid Hour", cron.SetHour, "25", true},
		{"Valid Day", cron.SetDay, "1,15", false},
		{"Invalid Day", cron.SetDay, "32", true},
		{"Valid Month", cron.SetMonth, "1-6", false},
		{"Invalid Month", cron.SetMonth, "13", true},
		{"Valid DayOfWeek", cron.SetDayOfWeek, "0-6", false},
		{"Invalid DayOfWeek", cron.SetDayOfWeek, "7", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestCronActive(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	cron := NewCron(loc)

	// Intentar activar sin configurar campos
	err := cron.Active()
	if err == nil {
		t.Error("Active() should return error when fields are not set")
	}

	// Configurar todos los campos
	cron.SetMinute("*")
	cron.SetHour("*")
	cron.SetDay("*")
	cron.SetMonth("*")
	cron.SetDayOfWeek("*")

	// Intentar activar con todos los campos configurados
	err = cron.Active()
	if err != nil {
		t.Errorf("Active() error = %v, wantErr false", err)
	}

	if !cron.IsActive() {
		t.Error("IsActive() should return true after successful activation")
	}
}

func TestCronNext(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	cron := NewCron(loc)

	// Configurar para ejecutar cada minuto
	cron.SetMinute("*")
	cron.SetHour("*")
	cron.SetDay("*")
	cron.SetMonth("*")
	cron.SetDayOfWeek("*")
	cron.Active()

	now := time.Now().Unix()
	next := cron.Next(now)

	if next <= now {
		t.Errorf("Next() = %v, want > %v", next, now)
	}

	// Probar con timestamp 0
	if cron.Next(0) != 0 {
		t.Error("Next(0) should return 0")
	}

	// Probar con cron inactivo
	inactiveCron := NewCron(loc)
	if inactiveCron.Next(now) != 0 {
		t.Error("Next() should return 0 for inactive cron")
	}
}

func TestParseCronField(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		min     int
		max     int
		wantErr bool
	}{
		{"Wildcard", "*", 0, 59, false},
		{"Step", "*/5", 0, 59, false},
		{"List", "1,2,3", 0, 59, false},
		{"Single", "30", 0, 59, false},
		{"Invalid Step", "*/abc", 0, 59, true},
		{"Invalid List", "1,abc,3", 0, 59, true},
		{"Out of Range", "70", 0, 59, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseCronField(tt.field, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCronField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNextValidValue(t *testing.T) {
	tests := []struct {
		name      string
		current   int
		allowed   []int
		want      int
		wantCarry bool
	}{
		{"Next in Range", 3, []int{1, 3, 5}, 3, false},
		{"Next with Carry", 4, []int{1, 3, 5}, 5, false},
		{"Wrap Around", 6, []int{1, 3, 5}, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, carry := nextValidValue(tt.current, tt.allowed)
			if got != tt.want || carry != tt.wantCarry {
				t.Errorf("nextValidValue() = (%v, %v), want (%v, %v)", got, carry, tt.want, tt.wantCarry)
			}
		})
	}
}

func TestCronNextWithNewYorkLocation(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("Error loading New York location: %v", err)
	}

	cron := NewCron(loc)

	// Configurar para ejecutar a las 22:05 todos los días
	cron.SetMinute("5")
	cron.SetHour("22")
	cron.SetDay("*")
	cron.SetMonth("*")
	cron.SetDayOfWeek("*")
	if err := cron.Active(); err != nil {
		t.Fatalf("Error activating cron: %v", err)
	}

	// Función auxiliar para imprimir fechas en ambos formatos
	printNextTime := func(ts int64) {
		nyTime := time.Unix(ts, 0).In(cron.location)
		utcTime := time.Unix(ts, 0).UTC()
		t.Logf("Next execution - NY: %s, UTC: %s", nyTime.Format("2006-01-02 15:04:05"), utcTime.Format("2006-01-02 15:04:05"))
	}

	// Test 1: 25 de febrero al 2 de marzo 2024
	t.Run("February-March Transition 2024", func(t *testing.T) {
		startTime := time.Date(2024, 2, 25, 22, 0, 0, 0, cron.location)
		current := startTime.Unix()

		t.Log("\nTest 1: 25 de febrero al 2 de marzo 2024")
		for i := 0; i < 6; i++ {
			next := cron.Next(current)
			if next == 0 {
				t.Fatalf("Next() returned 0 for timestamp %d", current)
			}
			printNextTime(next)
			current = next
		}
	})

	// Test 2: 28 de abril al 3 de mayo 2024
	t.Run("April-May Transition 2024", func(t *testing.T) {
		startTime := time.Date(2024, 4, 28, 22, 0, 0, 0, cron.location)
		current := startTime.Unix()

		t.Log("\nTest 2: 28 de abril al 3 de mayo 2024")
		for i := 0; i < 6; i++ {
			next := cron.Next(current)
			if next == 0 {
				t.Fatalf("Next() returned 0 for timestamp %d", current)
			}
			printNextTime(next)
			current = next
		}
	})

	// Test 3: 25 de febrero al 2 de marzo 2025
	t.Run("February-March Transition 2025", func(t *testing.T) {
		startTime := time.Date(2025, 2, 25, 22, 0, 0, 0, cron.location)
		current := startTime.Unix()

		t.Log("\nTest 3: 25 de febrero al 2 de marzo 2025")
		for i := 0; i < 6; i++ {
			next := cron.Next(current)
			if next == 0 {
				t.Fatalf("Next() returned 0 for timestamp %d", current)
			}
			printNextTime(next)
			current = next
		}
	})

	// Test 4: 28 de abril al 3 de mayo 2025
	t.Run("April-May Transition 2025", func(t *testing.T) {
		startTime := time.Date(2025, 4, 28, 22, 0, 0, 0, cron.location)
		current := startTime.Unix()

		t.Log("\nTest 4: 28 de abril al 3 de mayo 2025")
		for i := 0; i < 6; i++ {
			next := cron.Next(current)
			if next == 0 {
				t.Fatalf("Next() returned 0 for timestamp %d", current)
			}
			printNextTime(next)
			current = next
		}
	})
}

func TestCronSetSchedule(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	cron := NewCron(loc)

	tests := []struct {
		name    string
		spec    string
		wantErr bool
	}{
		{
			name:    "Valid Schedule",
			spec:    "*/5 0-12 1,15 1-6 0-6",
			wantErr: false,
		},
		{
			name:    "Invalid Fields Count",
			spec:    "*/5 0-12 1,15 1-6",
			wantErr: true,
		},
		{
			name:    "Invalid Minute",
			spec:    "70 0-12 1,15 1-6 0-6",
			wantErr: true,
		},
		{
			name:    "Invalid Hour",
			spec:    "*/5 25 1,15 1-6 0-6",
			wantErr: true,
		},
		{
			name:    "Invalid Day",
			spec:    "*/5 0-12 32 1-6 0-6",
			wantErr: true,
		},
		{
			name:    "Invalid Month",
			spec:    "*/5 0-12 1,15 13 0-6",
			wantErr: true,
		},
		{
			name:    "Invalid DayOfWeek",
			spec:    "*/5 0-12 1,15 1-6 7",
			wantErr: true,
		},
		{
			name:    "Empty Spec",
			spec:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cron.SetSchedule(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
