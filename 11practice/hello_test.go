package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Ashutosh"},
			result: "Hello Ashutosh!",
		},
		{
			items:  []string{"Ashutosh", "Mishra"},
			result: "Hello, Ashutosh, Mishra!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v) , got %s", st.result, st.items, s)
		}
	}
}
