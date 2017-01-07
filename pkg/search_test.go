package pkg
import "testing"
import "reflect"

func TestMatchFiles(t *testing.T) {
	testSet  := []string{"/something/something/else/hola","/something/something/else/hola/something else", "nomatchhere", "nomatchhere", "nomatchhere", "nomatchhere"}
	expected := []string{"/something/something/else/hola", "/something/something/else/hola/something else"}
	expectedBest := new(PriorityQueue);
	m1 := Match{value: "/something/something/else/hola", priority: 4, index: 0}
	m2 := Match{value: "/something/something/else/hola/something else", priority: 19, index: 1}
	expectedBest.Push(&m1)
	expectedBest.Push(&m2)


	results, best := MatchFiles("hola", testSet)


	if !reflect.DeepEqual(results, expected){
		t.Errorf("search failed to return the proper matches")
	}
	if !reflect.DeepEqual(best, expectedBest){
		t.Errorf("search failed to return the best matches")
	}

}
