// Copyright (C) 2017 ScyllaDB

package mermaidclient

import (
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/go-cmp/cmp"
	"github.com/scylladb/mermaid/internal/timeutc"
	"github.com/scylladb/mermaid/uuid"
)

func TestTaskSplit(t *testing.T) {
	t.Parallel()

	table := []struct {
		S  string
		T  string
		ID uuid.UUID
	}{
		{
			S:  "repair/d7d4b241-f7fe-434e-bc8e-6185b30b078a",
			T:  "repair",
			ID: uuid.MustParse("d7d4b241-f7fe-434e-bc8e-6185b30b078a"),
		},
		{
			S:  "d7d4b241-f7fe-434e-bc8e-6185b30b078a",
			ID: uuid.MustParse("d7d4b241-f7fe-434e-bc8e-6185b30b078a"),
		},
	}

	for i, test := range table {
		tp, id, err := TaskSplit(test.S)
		if err != nil {
			t.Error(i, err)
		}
		if tp != test.T {
			t.Error(i, tp)
		}
		if id != test.ID {
			t.Error(i, id)
		}
	}
}

func TestUUIDFromLocation(t *testing.T) {
	t.Parallel()

	u0 := uuid.MustRandom()
	u1, err := uuidFromLocation("http://bla/bla/" + u0.String() + "?param=true")
	if err != nil {
		t.Fatal(err)
	}
	if u1 != u0 {
		t.Fatal(u1, u0)
	}
}

func TestParseStartDate(t *testing.T) {
	t.Parallel()

	const epsilon = 50 * time.Millisecond

	table := []struct {
		S string
		D time.Duration
		E string
	}{
		{
			S: "now",
			D: nowSafety,
		},
		{
			S: "now-5s",
			E: "start date cannot be in the past"},
		{
			S: "now+5s",
			E: "start date must be at least in",
		},
		{
			S: "now+1h",
			D: time.Hour,
		},
		{
			S: timeutc.Now().Add(-5 * time.Second).Format(time.RFC3339),
			E: "start date cannot be in the past"},
		{
			S: timeutc.Now().Add(5 * time.Second).Format(time.RFC3339),
			E: "start date must be at least in",
		},
		{
			S: timeutc.Now().Add(time.Hour).Format(time.RFC3339),
			D: time.Hour,
		},
	}

	for i, test := range table {
		startDate, err := ParseStartDate(test.S)

		msg := ""
		if err != nil {
			msg = err.Error()
		}
		if test.E != "" || msg != "" {
			if !strings.Contains(msg, test.E) {
				t.Error(i, msg)
			}
			continue
		}

		s := truncateToSecond(time.Time(startDate))
		now := truncateToSecond(timeutc.Now())
		diff := now.Add(test.D).Sub(s)
		if diff < 0 {
			diff *= -1
		}
		if diff > epsilon {
			t.Fatal(i, startDate, test.D, diff, test.S)
		}
	}
}

func truncateToSecond(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func TestFormatTimeZero(t *testing.T) {
	t.Parallel()

	if s := FormatTime(strfmt.DateTime(time.Time{})); s != "" {
		t.Error(s)
	}
}

func TestFormatTimeNonZero(t *testing.T) {
	t.Parallel()

	tz, _ := timeutc.Now().Local().Zone()

	if s := FormatTime(strfmt.DateTime(timeutc.Now())); !strings.Contains(s, tz) {
		t.Error(s)
	}
}

func TestFormatMultiHostError(t *testing.T) {
	t.Parallel()

	msg := `Error: create cluster: connectivity check failed: 192.168.100.13: unknown network bla; 192.168.100.22: unknown network bla; 192.168.100.12: unknown network bla; 192.168.100.23: unknown network bla; 192.168.100.11: unknown network bla; 192.168.100.21: unknown network bla`

	golden := `Error: create cluster: connectivity check failed
 192.168.100.13: unknown network bla
 192.168.100.22: unknown network bla
 192.168.100.12: unknown network bla
 192.168.100.23: unknown network bla
 192.168.100.11: unknown network bla
 192.168.100.21: unknown network bla`

	if diff := cmp.Diff(FormatMultiHostError(msg, " "), golden); diff != "" {
		t.Fatal(diff)
	}
}

func TestFormatTables(t *testing.T) {
	t.Parallel()

	table := []struct {
		Name      string
		Threshold int
		Tables    []string
		AllTables bool
		Golden    string
	}{
		{
			Name:      "empty with threshold",
			Threshold: 2,
			Golden:    "(0 tables)",
		},
		{
			Name:      "empty unlimited",
			Threshold: -1,
			Golden:    "(0 tables)",
		},
		{
			Name:      "unlimited",
			Threshold: -1,
			Tables:    []string{"a", "b"},
			Golden:    "(a, b)",
		},
		{
			Name:   "one table",
			Tables: []string{"a"},
			Golden: "(1 table)",
		},
		{
			Name:   "no threshold",
			Tables: []string{"a", "b"},
			Golden: "(2 tables)",
		},
		{
			Name:      "above threshold",
			Threshold: 1,
			Tables:    []string{"a", "b"},
			Golden:    "(2 tables)",
		},
		{
			Name:      "below threshold",
			Threshold: 1,
			Tables:    []string{"a"},
			Golden:    "(a)",
		},
		{
			Name:      "all tables above threshold",
			Threshold: 1,
			Tables:    []string{"a", "b"},
			AllTables: true,
			Golden:    "all (2 tables)",
		},
		{
			Name:      "all tables below threshold",
			Threshold: 1,
			Tables:    []string{"a"},
			AllTables: true,
			Golden:    "all (a)",
		},
	}

	for i := range table {
		test := table[i]

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			if s := FormatTables(test.Threshold, test.Tables, test.AllTables); s != test.Golden {
				t.Errorf("FormatTables() expected %s got %s", test.Golden, s)
			}
		})
	}
}
