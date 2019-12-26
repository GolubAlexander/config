package config

import "testing"

func Test_isPointer(t *testing.T) {
	type args struct {
		cfg interface{}
	}
	type testStruct struct {
		field string
	}
	ts := testStruct{}
	type typeMap map[string]string
	m := typeMap{}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is a pointer",
			args: args{cfg: &ts},
			want: true,
		},
		{
			name: "is not a pointer",
			args: args{cfg: ts},
			want: false,
		},
		{
			name: "is a map and is not a pointer",
			args: args{cfg: m},
			want: false,
		},
		{
			name: "is a map and is a pointer",
			args: args{cfg: &m},
			want: true,
		},
		{
			name: "is a nil",
			args: args{cfg: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPointer(tt.args.cfg); got != tt.want {
				t.Errorf("isPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFile(t *testing.T) {
	type config struct {
		Test string `json:"test" yaml:"test"`
	}
	type args struct {
		cfg        interface{}
		pathToFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "struct is not a pointer",
			args: args{
				cfg:        config{},
				pathToFile: "./config-examples/valid.json",
			},
			wantErr: true,
		},
		{
			name: "file is not exists",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/not-existed-file.json",
			},
			wantErr: true,
		},
		{
			name: "file is not exists",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/not-existed-file.json",
			},
			wantErr: true,
		},
		{
			name: "type is not implemented",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/not.implemented-type",
			},
			wantErr: true,
		},
		{
			name: "json valid",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/valid.json",
			},
			wantErr: false,
		},
		{
			name: "json invalid",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/invalid.json",
			},
			wantErr: true,
		},
		{
			name: "yml valid",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/valid.yml",
			},
			wantErr: false,
		},
		{
			name: "yml invalid",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/invalid.yml",
			},
			wantErr: true,
		},
		{
			name: "yaml valid",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/valid.yaml",
			},
			wantErr: false,
		},
		{
			name: "yaml invalid",
			args: args{
				cfg:        &config{},
				pathToFile: "./config-examples/invalid.yaml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromFile(tt.args.cfg, tt.args.pathToFile); (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromBytes(t *testing.T) {
	type config struct {
		Test string `json:"test" yaml:"test"`
	}
	type args struct {
		cfg  interface{}
		data []byte
		t    typeFile
	}
	validJSON := []byte(`{ "test": "test"}`)
	invalidJSON := []byte(`{ "test": "test"`)
	validYAML := []byte(`test: 'test'`)
	invalidYAML := []byte(`- test: 'test'`)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "json valid",
			args: args{
				cfg:  &config{},
				data: validJSON,
				t:    TypeJson,
			},
			wantErr: false,
		},
		{
			name: "json invalid",
			args: args{
				cfg:  &config{},
				data: invalidJSON,
				t:    TypeJson,
			},
			wantErr: true,
		},
		{
			name: "yaml valid",
			args: args{
				cfg:  &config{},
				data: validYAML,
				t:    TypeYaml,
			},
			wantErr: false,
		},
		{
			name: "yaml invalid",
			args: args{
				cfg:  &config{},
				data: invalidYAML,
				t:    TypeYaml,
			},
			wantErr: true,
		},
		{
			name: "structure is not pointer",
			args: args{
				cfg:  config{},
				data: invalidYAML,
				t:    TypeYaml,
			},
			wantErr: true,
		},
		{
			name: "data is nil",
			args: args{
				cfg:  &config{},
				data: nil,
				t:    TypeYaml,
			},
			wantErr: true,
		},
		{
			name: "data is empty",
			args: args{
				cfg:  &config{},
				data: []byte(""),
				t:    TypeYaml,
			},
			wantErr: true,
		},
		{
			name: "not implemented",
			args: args{
				cfg:  &config{},
				data: []byte("1234"),
				t:    TypeUnknown,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromBytes(tt.args.cfg, tt.args.data, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("FromBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
