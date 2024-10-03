package main

import (
	"fmt"
	"testing"
)

type data struct {
	Foo string
	Bar int
}

func TestDumperFromFile(t *testing.T) {
	mydata := &data{Foo: "abc", Bar: 123}
	Dumper := NewDumper(".", 100)
	defer Dumper.Stop()
	Dumper.AddLog(mydata)
	Dumper.Flush()
	Dumper.Stop()
	Dumper = NewDumperFromFile(Dumper.filepath)
	decoded := &data{}
	Dumper.GetFromFile(&decoded)
	a := *decoded == *mydata
	if !a {
		t.Errorf("Expected %v, got %v", mydata, decoded)
	}
}
func TestDumperBasic(t *testing.T) {

	mydata := data{Foo: "abc", Bar: 123}
	Dumper := NewDumper(".", 100)
	defer Dumper.Stop()
	Dumper.AddLog(&mydata)
	decoded := &data{}
	Dumper.GetLog(&decoded)
	a := *decoded == mydata
	if !a {
		t.Errorf("Expected %v, got %v", mydata, decoded)
	}
}

func TestDumperBasic2(t *testing.T) {

	mydata1 := data{Foo: "abc", Bar: 123}
	mydata2 := data{Foo: "dfg", Bar: 345}
	Dumper := NewDumper(".", 10000)
	defer Dumper.Stop()
	Dumper.AddLog(&mydata1)
	Dumper.AddLog(&mydata2)
	decoded1 := &data{}
	Dumper.GetLog(&decoded1)
	decoded2 := &data{}
	Dumper.GetLog(&decoded2)
	a := *decoded1 == mydata1
	b := *decoded2 == mydata2
	if !a {
		t.Errorf("Expected %v, got %v", mydata1, decoded1)
	}
	if !b {
		t.Errorf("Expected %v, got %v", mydata2, decoded2)
	}

}
func TestParquetFromMemdump(t *testing.T) {
	Dumper := NewDumperFromFile("./testdata/pdata_simple.memdump")
	defer Dumper.Close()
	decoded := &ParquetMessage{}
	Dumper.GetFromFile(&decoded)
	fmt.Println(*decoded)
	decoded2 := &ParquetMessage{}
	Dumper.GetFromFile(&decoded2)
	fmt.Println(*decoded2)
	l := NewLogger[ParquetMessage](".", 5000, 10)
	fmt.Println("STARTING")
	l.Start()

	for l1 := 0; l1 < 10000; l1++ {
		decoded3 := &ParquetMessage{}
		Dumper.GetFromFile(&decoded3)
		l.AddLog(*decoded3)
		fmt.Println(*decoded3)
	}
	defer l.Stop()

}
