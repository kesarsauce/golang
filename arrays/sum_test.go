package arrays

import "testing"
import "reflect"

func TestSumArray(t *testing.T) {
	compareGotWant := func (got, want int, t testing.TB){
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("5 element array", func(t *testing.T){
		numbers := [] int{2,3,4,5,6}
		got := SumArray(numbers)
		want := 20
		compareGotWant(got, want, t)
	})
}

func TestSumAll(t *testing.T) {
	compareGotWant := func (got, want []int, t testing.TB){
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("2 arrays", func(t *testing.T){
		firstArray := [] int{1,3,6}
		secondArray:= [] int{0,2}
		got := SumAll(firstArray, secondArray)
		want := [] int{10,2}
		compareGotWant(got, want, t)
	})
}