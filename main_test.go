package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPlan(t *testing.T) {
	plan, err := buildPlan()

	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.NotEmpty(t, plan)
}

func TestGenerateCalendar(t *testing.T) {
	defer removeFile()

	plan, _ := buildPlan()
	err := generateCalendar(plan)

	assert.Nil(t, err)
}

func removeFile() {
	if err := os.Remove(fileName()); err != nil {
		fmt.Println(err)
	}
}
