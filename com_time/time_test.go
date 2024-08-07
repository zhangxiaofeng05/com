package com_time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	t.Run("Specific date and time", func(t *testing.T) {
		testTime := time.Date(2024, 8, 8, 10, 30, 0, 0, time.UTC)
		expected := "2024-08-08 10:30:00"
		formattedTime := Format(testTime)
		require.Equal(t, expected, formattedTime)
	})

	t.Run("Zero time", func(t *testing.T) {
		zeroTime := time.Time{}
		expectedZero := "0001-01-01 00:00:00"
		formattedZeroTime := Format(zeroTime)
		require.Equal(t, expectedZero, formattedZeroTime)
	})

	t.Run("Current time", func(t *testing.T) {
		now := time.Now()
		formattedNow := Format(now)
		expectedNow := now.Format(defaultLayout)
		require.Equal(t, expectedNow, formattedNow)
	})
}
