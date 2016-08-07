package comma

import (
	"testing"
)

func TestComma(t *testing.T) {
	cases := []struct {
	in , want  string
	}{
		{"123","123" },
		{"1234","1,234"},
		{"-123467","-123,467"},
		{"-1","-1"},
                {"-12345.45","-12,345.45"},
	}

	for _, c :=  range cases {
		got :=AddComma(c.in)
		if got  != c.want {
			t.Errorf("AddComma(%s) == %s , want %s", c.in , got , c.want)
		}
	}
}
