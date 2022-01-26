package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTime_CreateNewTimeObj(t *testing.T) {
	testTable := []struct {
		name string

		h, m, s int
	}{
		{
			name: "should create time obj",

			h: 1, m: 1, s: 1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			time := CreateNewTimeObj(tt.h, tt.m, tt.s)

			assert.Equal(t, tt.h, time.hours)
			assert.Equal(t, tt.h, time.minutes)
			assert.Equal(t, tt.h, time.seconds)
		})
	}
}
