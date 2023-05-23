package _map

import (
	"fmt"
	"time"
)

type Dict map[string]any

func MakeDict() Dict {
	return make(Dict, 15)
}

func (d Dict) Set(k string, v any) {
	d[k] = v
}

func (d Dict) Get(k string) any {
	v, ok := d[k]
	if !ok {
		return nil
	}
	return v
}

func (d Dict) Delete(k string) {
	delete(d, k)
}

func (d Dict) IterPrint() {
	for k, v := range d {
		fmt.Printf("[%s, %v]\n", k, v)
	}
}

func PrintMap() {
	d := MakeDict()
	d.Set("key 1", 0)
	d.Set("key 2", []int{1, 2, 3})
	d.Set("key 3", "xxx")
	d.Set("key 4", "should be deleted")
	fmt.Printf("get d's key 1: %v\n", d.Get("key 1"))
	fmt.Printf("get d's key 2: %v\n", d.Get("key 2"))
	fmt.Printf("get d's key 3: %v\n", d.Get("key 3"))
	fmt.Printf("get d's key 4: %v\n", d.Get("key 4"))
	fmt.Printf("get a not exist key in d: %s\n", d.Get("not exist"))

	d.Delete("key 4")
	d.IterPrint()

	// concurrent
	go func() {
		for i := 1; i < 1000; i++ {
			d.Set(string(i), i)
		}
	}()

	go func() {
		for i := 1; i < 1000; i++ {
			v := d.Get(string(1))
			fmt.Println(v)
		}
	}()
	time.Sleep(2 * time.Second)

	// pointer of value
	// p := &d["key 1"]
	// fmt.Println(p)
}
