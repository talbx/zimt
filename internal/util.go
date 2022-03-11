package internal

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

func Split(time string) (hours string, minutes string) {
	var split = strings.Split(time, ":")
	hours = split[0]
	minutes = split[1]
	return
}

func ToMinutes(hours string, minutes string) int {
	realHours, _ := strconv.Atoi(hours)
	realMinutes, _ := strconv.Atoi(minutes)
	return realHours*60 + realMinutes
}

func AddLeadingZero(num int) string {
	if num < 10 {
		return "0" + strconv.Itoa(num)
	}
	return strconv.Itoa(num)
}


func ReadFromYaml(list int) string {
	filename := "zimt.yaml"
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var entries []ZimtEntry
	err = yaml.Unmarshal(yamlFile, &entries)

	length := len(entries)

	reversed := reverse(entries)

	if length < list {
		list = length
	}

	var resultList []ZimtEntry
	for i := 0; i < list; i++ {
		resultList = append(resultList, reversed[i])
	}

	marshal, _ := yaml.Marshal(&resultList)
	sprintf := fmt.Sprintf("---\n%s\n\n", string(marshal))

	return sprintf
}

func reverse(entries []ZimtEntry)[]ZimtEntry {
	var copy = entries
	for i, j := 0, len(copy)-1; i < j; i, j = i+1, j-1 {
		copy[i], copy[j] = copy[j], copy[i]
	}
	return copy
}

func WriteToYaml(entry ZimtEntry) {

	filename := "zimt.yaml"
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var entries []ZimtEntry
	err = yaml.Unmarshal(yamlFile, &entries)

	fmt.Println(entries)

	if err != nil {
		panic(err)
	}

	newEntries := append(entries, entry)
	entries = newEntries
	fmt.Println(entries)

	yamlData, err := yaml.Marshal(entries)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(&yamlData)
	fileName := filename
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}
}
