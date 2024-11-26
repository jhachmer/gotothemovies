package util

import (
	"os"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		values []T
		fn     func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{name: "Filter Modulo 2",
			args: args[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{2, 4, 6, 8},
		},
		{name: "Filter Empty result",
			args: args[int]{
				values: []int{1, 3, 5, 7, 9},
				fn: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.values, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindValidFiles(t *testing.T) {
	tempDir := t.TempDir()
	// log.Println(tempDir)
	temp1, err := os.CreateTemp(tempDir, "test*.mp3")
	if err != nil {
		t.Fatal("Failed to create temp file")
	}
	temp2, err := os.CreateTemp(tempDir, "test*.mp4")
	if err != nil {
		t.Fatal("Failed to create temp file")
	}
	defer temp1.Close()
	defer temp2.Close()
	defer os.Remove(temp1.Name())
	defer os.Remove(temp2.Name())
	type args struct {
		root string
		ext  string
	}
	tests := []struct {
		name    string
		args    args
		want    []DirFiles
		wantErr bool
	}{
		{name: "", args: args{
			root: tempDir,
			ext:  ".mp4",
		}, want: []DirFiles{{Name: temp2.Name()}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindValidFiles(tt.args.root, tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindValidFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindValidFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[TValue any, TResult any] struct {
		values []TValue
		fn     func(TValue) TResult
	}
	type testCase[TValue any, TResult any] struct {
		name string
		args args[TValue, TResult]
		want []TResult
	}
	tests := []testCase[int, int]{
		{name: "Map Sqaured Numbers", args: args[int, int]{values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: func(i int) int {
			return i * i
		}}, want: []int{1, 4, 9, 16, 25, 36, 49, 64, 81}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.values, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args[TValue any, TResult any] struct {
		values       []TValue
		initialValue TResult
		fn           func(TResult, TValue) TResult
	}
	type testCase[TValue any, TResult any] struct {
		name string
		args args[TValue, TResult]
		want TResult
	}
	tests := []testCase[int, int]{
		{name: "[1,2,3] = 6", args: args[int, int]{
			values:       []int{1, 2, 3},
			initialValue: 0,
			fn: func(i int, j int) int {
				return i + j
			},
		}, want: 6},
		{name: "[1,2,3] + initial {2} = 8", args: args[int, int]{
			values:       []int{1, 2, 3},
			initialValue: 2,
			fn: func(i int, j int) int {
				return i + j
			},
		}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.values, tt.args.initialValue, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
