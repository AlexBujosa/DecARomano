package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T){
	want := "I"
	input := "1"
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
	}else {
		rom := ConvertDecToRom(inputINT)
		if rom != want {
			t.Error("Error")
		}
	}
}
func TestCase2(t *testing.T){
	want := "Error"
	input := "2.3"
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
		
	}else {
		rom := ConvertDecToRom(inputINT)
		if rom != want {
			t.Error("Error")
		}
	}
}
func TestCase3(t *testing.T){
	want := "Error"
	input := "ha"
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
	}else {
		rom := ConvertDecToRom(inputINT)
		if rom != want {
			t.Error("Error")
		}
	}
}
func TestCase4(t *testing.T){
	want := "Error"
	input := "3z"
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
	}else {
		rom := ConvertDecToRom(inputINT)
		if rom != want {
			t.Error("Error")
		}
	}
}
func TestCase5(t *testing.T){
	want := "XLIX"
	input := "3z"
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
	}else {
		rom := ConvertDecToRom(inputINT)
		if rom != want {
			t.Error("Error")
		}
	}
}
func TestCase6(t *testing.T){
	want := "III"
	input := "3"
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
	}else {
		rom := ConvertDecToRom(inputINT)
		if rom != want {
			t.Error("Error")
		}
	}
}