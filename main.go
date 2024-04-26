package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type Store interface {
	Set(key string, val any)
	Get(key string) (val any, err error)
}

type MyStore struct {
	mp sync.Map
}

func (m *MyStore) Set(key string, val any) {
	m.mp.Store(key, val)
}
func (m *MyStore) Get(key string) (val any, err error) {
	val, ok := m.mp.Load(key)
	if ok != true {
		return nil, errors.New("нет такого значения")
	}
	return val, nil
}
func NewStore() Store {
	store := &MyStore{mp: sync.Map{}}
	//store.mp = sync.Map{}
	return store
}

type MyErr struct {
	er string
}

func (err MyErr) Error() string {
	return "my err"
}

type MyErr2 struct {
}

func (err MyErr2) Error() string {
	return "my err"
}
func main() {
	err1 := MyErr{er: "ss"}
	err2 := MyErr{er: "ss"}
	fmt.Println(err2 == err1)
	fmt.Println(errors.Is(err1, err2))
	arr := make([]int, 10)
	arr[1] = 5
	fmt.Println("cap", cap(arr))
	fmt.Println("len", len(arr))
	fmt.Println("arr", arr)
	//check(NewStore())
	//changeString()
	//getSlices2()
	//focusWithCap()
	//getSlices1()
	changeSliceCap()
	rangeStr()

	//mp := make(map[string]string)
	//editMap(mp)//map всегда передается по ссылке
	//fmt.Println("map")
	//fmt.Println(mp)
	//fmt.Println(3 << 4) //3*(2**4)
}

func check(s Store) {
	s.Set("1", "one")
	s.Set("2", 2)
	fmt.Println(s.Get("1"))
	fmt.Println(s.Get("2"))
	fmt.Println(s.Get("3"))
}
func getSlices1() {
	arr1 := []int{1, 2, 3, 4, 5}
	//если капасити поменять то происходит перевыделение и поведение другое
	arr2 := arr1[:4]
	fmt.Println(arr2)
	fmt.Println(cap(arr2))
	arr3 := append(arr2[:2], 9, 8, 7)

	fmt.Println(arr1) //[1, 2, 9, 8, 7]
	fmt.Println(arr2) //[1, 2, 9, 8]
	fmt.Println(arr3) //[1, 2, 9, 8, 7]
}
func changeString() {
	s0 := "some string"
	s1 := []byte(s0)
	fmt.Println(cap(s1))
	fmt.Println(len(s1))
	s2 := s1[5:]
	fmt.Println(cap(s2))
	fmt.Println(len(s2))
	s2[3] = 'o'
	fmt.Println(string(s2))
	fmt.Println(string(s1))
	fmt.Println(string(s0))
}

func getSlices2() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(1)
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	fmt.Println(2)
	printSlice(s)

	// Extend its length.
	s = s[:4]
	fmt.Println(3)
	printSlice(s)

	//s = s[:]
	//printSlice(s)
	// Drop its first two values.
	s = s[2:]
	fmt.Println(4)
	printSlice(s)
	s = s[:4]
	fmt.Println(5)
	printSlice(s)

	s2 := s[1:]
	s2[0] = 9
	fmt.Println(6)
	printSlice(s)
	fmt.Println(7)
	printSlice(s2)
}

func focusWithCap() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func changeSliceCap() {
	sl := []int{1, 2, 3, 4, 5}
	printSlice(sl)
	sl2 := sl[:3:3]
	sl2 = append(sl2, 1)
	printSlice(sl)
	printSlice(sl2)
}

func rangeStr() {
	str := "my_strапррёing"
	for _, v := range str {
		//fmt.Print(v) // bytes
		fmt.Printf("%c", v)
	}
	fmt.Println()

	str2 := "это стр¢ока"

	for _, v := range str2 {
		//fmt.Print(v) // bytes
		fmt.Printf("%c", v)
	}
	fmt.Println()
	fmt.Println(len(str2))
	fmt.Printf("%c", str2[4])

}

func editMap(mp map[string]string) {
	mp["test"] = "just test"
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
