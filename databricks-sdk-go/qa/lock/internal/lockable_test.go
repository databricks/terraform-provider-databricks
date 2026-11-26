package internal

import (
	"testing"
)

func TestGenerateBlobName(t *testing.T) {
	type testStruct struct {
		A string
		B int
	}

	cases := []struct {
		name string
		in   interface{}
		want string
	}{
		{
			name: "string",
			in:   "hello",
			want: "hello",
		},
		{
			name: "struct",
			in:   testStruct{A: "hello", B: 1},
			want: "testStruct:A=hello;B=1",
		},
		{
			name: "struct pointer",
			in:   &testStruct{A: "hello", B: 1},
			want: "testStruct:A=hello;B=1",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GenerateBlobName(c.in)
			if got != c.want {
				t.Errorf("generateBlobName(%v) == %s, want %s", c.in, got, c.want)
			}
		})
	}
}

func TestGenerateBlobName_Panic(t *testing.T) {
	type testStruct struct {
		C chan int
	}
	cases := []struct {
		name string
		in   interface{}
	}{
		{
			name: "unsupported type",
			in:   1,
		},
		{
			name: "unsupported type pointer",
			in:   &[]int{1},
		},
		{
			name: "unsupported type struct",
			in:   testStruct{C: make(chan int)},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("generateBlobName(%v) did not panic", c.in)
				}
			}()
			GenerateBlobName(c.in)
		})
	}
}
