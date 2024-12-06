package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type tmp struct {
	name     string
	typ      string
	children map[string]tmp
}

func TestGenerate(t *testing.T) {
	A := message{Time: 123, Type: 1}
	v := reflect.ValueOf(A)
	all := make(map[string]tmp)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct {

			t := tmp{name: v.Field(i).Type().Name(),
				typ:      v.Field(i).Type().Name(),
				children: make(map[string]tmp)}
			// t1 := tmp{name: "Type", typ: "[]int"}
			// t2 := tmp{name: "Time", typ: "[]int64"}
			// t3 := tmp{name: "RefTime", typ: "[]int64"}
			// t.children["Type"] = t1
			// t.children["Time"] = t2
			// t.children["RefTime"] = t3
			v2 := v.Field(i)

			for j := 0; j < v2.NumField(); j++ {

				typ := ""
				if v2.Field(j).Kind() == reflect.Int {
					typ = "[]int"
				} else if v2.Field(j).Kind() == reflect.String {
					typ = "[]string"
				} else if v2.Field(j).Kind() == reflect.Uint64 {
					typ = "[]uint64"
				} else if v2.Field(j).Kind() == reflect.Int64 {
					typ = "[]int64"
				} else if v2.Field(j).Kind() == reflect.Uint {
					typ = "[]uint"
				} else if v2.Field(j).Kind() == reflect.Int32 {
					typ = "[]int32"
				} else {

					// fmt.Println(v2.Field(j).Kind(), v2.Type().Field(j).Name)
				}
				if typ != "" {
					t2 := tmp{name: v2.Type().Field(j).Name,
						typ: typ,
					}
					t.children[v2.Type().Field(j).Name] = t2
				}
			}
			all[v.Type().Field(i).Name] = t
		} else if v.Field(i).Kind() == reflect.Int {
			typ := "[]int"
			t2 := tmp{name: v.Type().Field(i).Name,
				typ: typ,
			}
			all[v.Type().Field(i).Name] = t2
		} else if v.Field(i).Kind() == reflect.Int64 {
			typ := "[]int64"
			t2 := tmp{name: v.Type().Field(i).Name,
				typ: typ,
			}
			all[v.Type().Field(i).Name] = t2
		}

	}
	s1 := "package main\n"
	s1 += writeTypes(all)
	s1 += writeFunctions(all)
	fmt.Println(s1)
	s1 += writeAddFunc(all)

	f, _ := os.Create("dumpmodel.go")
	defer f.Close()
	f.WriteString(s1)
	// fmt.Println(s1)

}

func writeAddFunc(all map[string]tmp) string {
	s := ""
	s += "func AddData(data message, ref *PL){\n"
	for l1, l2 := range all {
		if len(l2.children) == 0 {
			s += "(*ref)." + l1 + " = append((*ref)." + l1 + ",data." + l1 + ")\n"
		} else {
			for l3, l4 := range l2.children {
				name := "(*ref)." + l1 + "." + l3
				s += name + " = append(" + name + ",data." + l1 + "." + l3 + ")\n"
				fmt.Println(l1, l3, l4)
			}
		}

	}
	s += "}\n"
	return s
}

func writeFunctions(all map[string]tmp) string {
	s := "func NewPL(length int) PL {\n"
	s += "return PL{\n"
	for i, _ := range all {
		if len(all[i].children) == 0 {
			s += i + ": make(" + all[i].typ + ",0,length),\n"
		} else {
			s += i + ": " + all[i].typ + "_gen{\n"
			for j, _ := range all[i].children {
				s += j + ": make(" + all[i].children[j].typ + ",0,length),\n"
			}
			s += "},\n"
		}

	}
	s += "}}\n"

	return s
}

func writeTypes(in map[string]tmp) string {
	s := "type PL struct {\n"
	for i, j := range in {
		if len(j.children) == 0 {
			s += i + " " + in[i].typ + "\n"
		} else {
			s += i + " " + in[i].typ + "_gen\n"
		}
	}
	s += "}\n"

	for i, k := range in {
		if len(k.children) > 0 {
			s += "type " + in[i].typ + "_gen struct {\n"
			for j, _ := range in[i].children {
				s += j + " " + in[i].children[j].typ + "\n"
			}
			s += "}\n"
		}

	}
	return s
}
