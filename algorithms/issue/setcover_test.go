package issue_test

import (
	"gnotes/algorithms/issue"
	"reflect"
	"sort"
	"testing"
)

func TestSetCover(t *testing.T) {
	statesNeed := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := map[string][]string{
		"kone":   {"id", "nv", "ut"},
		"ktwo":   {"wa", "id", "mt"},
		"kthree": {"or", "nv", "ca"},
		"kfour":  {"nv", "ut"},
		"kfive":  {"ca", "az"},
	}
	want := []string{"kone", "ktwo", "kthree", "kfive"}
	sort.Strings(want)

	result := issue.SetCover(statesNeed, stations)
	sort.Strings(result)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Wrong result %v, result must be %v", result, want)
	}
}
