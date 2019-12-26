package config

import "testing"

func Test_detectType(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want typeFile
	}{
		{
			name: "json extension",
			args: args{ext: ".json"},
			want: TypeJson,
		},
		{
			name: "yml extension",
			args: args{ext: ".yml"},
			want: TypeYaml,
		},
		{
			name: "yaml extension",
			args: args{ext: ".yaml"},
			want: TypeYaml,
		},
		{
			name: "unknown extension",
			args: args{ext: ".exe"},
			want: TypeUnknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectType(tt.args.ext); got != tt.want {
				t.Errorf("detectType() = %v, want %v", got, tt.want)
			}
		})
	}
}
