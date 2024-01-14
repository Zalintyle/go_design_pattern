package demo

import "testing"

func Test_Suv_Builder(t *testing.T) {
	suvBuilder := GetBuilder("suv")
	director := NewDirector(suvBuilder)
	director.Construct()
	car := suvBuilder.GetCar()
	if car.SeatsType != "suv seats" {
		t.Fatal("car seats type error")
	}
	if car.EngineType != "suv engine" {
		t.Fatal("car engine type error")
	}
	if car.Number != 6 {
		t.Fatal("car number error")
	}
}

func Test_Mpv_Builder(t *testing.T) {
	mpvBuilder := GetBuilder("mpv")
	director := NewDirector(mpvBuilder)
	director.Construct()
	car := mpvBuilder.GetCar()
	if car.SeatsType != "mpv seats" {
		t.Fatal("car seats type error")
	}
	if car.EngineType != "mpv engine" {
		t.Fatal("car engine type error")
	}
	if car.Number != 8 {
		t.Fatal("car number error")
	}
}

func TestDirector_SetBuilder(t *testing.T) {
	suvBuilder := GetBuilder("suv")
	director := NewDirector(suvBuilder)
	car := director.Construct()
	if car.SeatsType != "suv seats" {
		t.Fatal("car seats type error")
	}
	if car.EngineType != "suv engine" {
		t.Fatal("car engine type error")
	}
	if car.Number != 6 {
		t.Fatal("car number error")
	}

	mpvBuilder := GetBuilder("mpv")
	director.SetBuilder(mpvBuilder)
	car = director.Construct()
	if car.SeatsType != "mpv seats" {
		t.Fatal("car seats type error")
	}
	if car.EngineType != "mpv engine" {
		t.Fatal("car engine type error")
	}
	if car.Number != 8 {
		t.Fatal("car number error")
	}
}
