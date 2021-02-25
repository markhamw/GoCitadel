package main

import (
	"testing"
)

func TestNothing(t *testing.T) {
	expected := 1
	if expected == 1 {
		t.Logf("success on this, expected %v and actual was %v", "1", expected)
	} else if expected == 2 {
		t.Errorf("failed as expected since %v is not %v ", "2", expected)
	}
}

func TestPlayerName(t *testing.T) {
	testchar := new(playerchar)
	testchar.Name = "testName"
	actual := testchar.GetName()
	expected := "testName"
	if expected != actual {
		t.Errorf("failed expected %v and got %v", "testName", expected)
	}
	if expected == actual {
		t.Logf("passed since %v is %v", "testName", expected)
	}

}
