package map_

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type WordFreq struct {
	Name string
	Freq int
}

type WordFreqs []*WordFreq

// 实现sort.Interface接口取元素数量方法
func (wfs WordFreqs) Len() int {
	return len(wfs)
}

// 实现sort.Interface接口比较元素方法
func (wfs WordFreqs) Less(i, j int) bool {
	return wfs[i].Freq < wfs[j].Freq
}

// 实现sort.Interface接口交换元素方法
func (wfs WordFreqs) Swap(i, j int) {
	wfs[i], wfs[j] = wfs[j], wfs[i]
}

// 实现wordFreqs翻转
func (wfs WordFreqs) Reverse() {
	if len(wfs) == 0 {
		return
	}
	for i, j := 0, len(wfs)-1; i < j; i, j = i+1, j-1 {
		wfs.Swap(i, j)
	}
}

// 实现打印
func (wfs WordFreqs) printInfo() {
	for _, wf := range wfs {
		fmt.Printf("%v: %d\n", wf.Name, wf.Freq)
	}
}

// read lyrics.txt in the same folder, return the word count with desec
func getLyricsWordCount(fileName string) WordFreqs {
	count := make(map[string]int)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("failed to open file %v \n", file)
		return WordFreqs{}
	}
	// defer close to the end
	defer file.Close()

	// create bufio
	scanner := bufio.NewScanner(file)

	// read the lyrics line by line
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			word = strings.ToLower(word)
			// can do this, but
			/*
				if _, ok :=  count[word]; ok {
					count[word]++
				} else {
					count[word] = 1
				}
			*/
			// When accessing a map key that doesn’t exist yet, one receives a zero value.
			// Often, the zero value is a suitable value, for example when using append or doing integer math.
			count[word] += 1
		}
	}

	// make wordFreqs
	wordFreqs := make(WordFreqs, 0, len(count))
	for k, v := range count {
		wordFreqs = append(wordFreqs, &WordFreq{k, v})
	}
	// fmt.Println(wordFreqs)

	// order
	sort.Sort(wordFreqs)

	return wordFreqs
}

func PrintMap() {
	// create
	testM := make(map[string]int)
	testM["alice"] = 1
	testM["jack"] = 2
	fmt.Println(testM)
	testM["adam"] = 3

	// retrive
	value, ok := testM["adam"]
	if ok {
		fmt.Println(value)
	}
	fmt.Println(testM["adam"])

	// delete
	delete(testM, "alice")
	fmt.Println(testM)

	// update
	testM["adam"] = 100
	fmt.Println(testM["adam"])

	// loop
	for k, v := range testM {
		fmt.Println(k, v)
	}

	// exercise
	wordFreqs := getLyricsWordCount("/Users/fangxiang/study/study_go/map/lyrics.txt")
	wordFreqs.printInfo()
	wordFreqs.Reverse()
	wordFreqs.printInfo()
}
