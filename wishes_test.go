package moduleb

import "testing"

func TestGoodMorning(t *testing.T) {
	want := "Good Morning!!!"

	if got := GoodMorning(); got != want {
		t.Errorf("GoodMorning() = %q, want %q", got, want)
	}
}
