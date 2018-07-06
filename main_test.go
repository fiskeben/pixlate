package main

import "testing"

func Test_makeOutputFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "creates simple filename",
			args: args{
				filename: "foo.jpg",
			},
			want: "foo.pixlated.jpg",
		},
		{
			name: "creates name from filename with many dots",
			args: args{
				filename: "silly.name.with.many.dots.png",
			},
			want: "silly.name.with.many.dots.pixlated.png",
		},
		{
			name: "creates name from filename without an extension",
			args: args{
				filename: "foobar",
			},
			want: "foobar.pixlated",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeOutputFilename(tt.args.filename); got != tt.want {
				t.Errorf("makeOutputFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
