package internal

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReverse(t *testing.T){

	// given
	entry1 := ZimtEntry{Type: "yo"}
	entry2 := ZimtEntry{Type: "ai"}
	var e = []ZimtEntry{entry1, entry2}

	// when
	reversed := reverse(e)

	// then
	assert.Equal(t,reversed[0], entry2)
	assert.Equal(t,reversed[1], entry1)
}

func TestSplit(t *testing.T) {
	// given
	time := "12:31"

	// when
	hours, minutes := Split(time)

	// then
	assert.Equal(t, hours, 12)
	assert.Equal(t, minutes, 31)
}

func TestToMinutes(t *testing.T) {
	// given
	hours, minutes := "1", "30"

	// when
	totalMinutes := ToMinutes(hours, minutes)

	// then
	assert.Equal(t, totalMinutes, 90)
}

func TestAddLeadingZero(t *testing.T) {

	// given
	var num1 = 2

	// when
	str1 := AddLeadingZero(num1)

	// then
	assert.Equal(t, str1, "02")
}

func TestAddLeadingZero2(t *testing.T) {

	// given
	var num2 = 11

	// when
	str2 := AddLeadingZero(num2)

	// then
	assert.Equal(t, str2, "11")
}
func TestAddLeadingZero3(t *testing.T) {

	// given
	var num3 = 10

	// when
	str3 := AddLeadingZero(num3)

	// then
	assert.Equal(t, str3, "10")
}
