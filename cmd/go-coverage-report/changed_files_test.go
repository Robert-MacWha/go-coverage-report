package main

import (
	"reflect"
	"testing"
)

func Test_parseChangedFiles(t *testing.T) {
	type args struct {
		data    []byte
		prefix  string
		ignored []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"empty",
			args{
				data:    []byte("[]"),
				prefix:  "",
				ignored: []string{},
			},
			[]string{},
			false,
		},
		{
			"sample",
			args{
				data:    []byte(`["file1", "foo/file2", "bar/file3"]`),
				prefix:  "github.com",
				ignored: []string{"foo"},
			},
			[]string{"github.com/file1", "github.com/bar/file3"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseChangedFiles(tt.args.data, tt.args.prefix, tt.args.ignored)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseChangedFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseChangedFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
