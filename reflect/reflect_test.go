package main

import (
	"log"
	"reflect"
	"testing"
)

func TestReflectType(tc *testing.T) {
	g := Gift{sender: "James",
		recipient: "Linda",
		number:    1,
		contents:  "Teddy",
	}
	t := reflect.TypeOf(g)
	if kind := t.Kind(); kind != reflect.Struct {
		log.Fatal("Invalid struct type")
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		log.Printf("Field %03d: %-10.10s %v", i, f.Name, f.Type.Kind())
	}
}

func TestReflectMethod(tt *testing.T) {
	r := reindeer("rudolph")
	r.takeoff()
	r.land()
	r.togglenose()
	t := reflect.TypeOf(r)
	log.Printf("nummethod=%d", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		log.Printf("Field %03d: %-10.10s", i, m.Name)
	}
}
