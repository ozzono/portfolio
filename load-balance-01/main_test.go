package loadbalance

import (
	"math/rand"
	"testing"
	"time"
)

var (
	clientXPList = arrayToType([]int{10, 10, 10, 20, 20, 30, 30, 30, 20, 60})
)

func TestDefault(t *testing.T) {

	testCS := []int{}
	for i := 0; i <= 6; i++ {
		for j := i; j > 0; j-- {
			testCS = append(testCS, i)
		}
	}
	unnavailableCS := []int{}
	max, err := balance(arrayToType(shuffle(testCS)), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 0
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}

func TestCase1(t *testing.T) {
	css := []uniqueType{
		{id: 1, score: 60},
		{id: 2, score: 20},
		{id: 3, score: 95},
		{id: 4, score: 75},
	}
	customers := []uniqueType{
		{id: 1, score: 90},
		{id: 2, score: 20},
		{id: 3, score: 70},
		{id: 4, score: 40},
		{id: 5, score: 60},
		{id: 6, score: 10},
	}

	max, err := balance(css, customers, []int{2, 4})
	if err != nil {
		t.Fatal(err)
	}

	expected := 1
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}

func TestCase2(t *testing.T) {
	testCS := []int{11, 21, 31, 3, 4, 5}
	unnavailableCS := []int{10, 10, 10, 20, 20, 30, 30, 30, 20, 60}
	max, err := balance(arrayToType(testCS), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 0
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}

// This test was the most tricky one to understand in ruby
// Although it was easy to translate, once understood
func TestCase3(t *testing.T) {
	testCS := []int{}
	for i := 0; i < 1000; i++ {
		testCS = append(testCS, 0)
	}
	testCS[998] = 100
	unnavailableCS := []int{}
	for i := 0; i < 10000; i++ {
		unnavailableCS = append(unnavailableCS, 10)
	}
	max, err := balance(arrayToType(testCS), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 999
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}

func TestCase4(t *testing.T) {

	testCS := []int{1, 2, 3, 4, 5, 6}
	unnavailableCS := []int{}
	max, err := balance(arrayToType(testCS), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 0
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}
func TestCase5(t *testing.T) {

	testCS := []int{100, 2, 3, 3, 4, 5}
	unnavailableCS := []int{}
	max, err := balance(arrayToType(testCS), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 1
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}
func TestCase6(t *testing.T) {

	testCS := []int{100, 99, 88, 3, 4, 5}
	unnavailableCS := []int{1, 3, 2}
	max, err := balance(arrayToType(testCS), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 0
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}
func TestCase7(t *testing.T) {

	testCS := []int{100, 99, 88, 3, 4, 5}
	unnavailableCS := []int{4, 5, 6}
	max, err := balance(arrayToType(testCS), clientXPList, unnavailableCS)
	if err != nil {
		t.Fatal(err)
	}
	expected := 3
	if max != expected {
		t.Fatalf("expecting %d as output; got %d", expected, max)
	}
}

func shuffle(input []int) []int {
	output := []int{}
	for len(input) > 0 {
		index := rand.New(rand.NewSource(time.Now().UnixNano() + int64(len(input)))).Intn(len(input))
		output = append(output, input[index])
		input = append(input[:index], input[index+1:]...)
	}
	return output
}
