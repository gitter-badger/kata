package format

import (
	"testing"
)

func TestFormat_Parse(t *testing.T) {
	type fields struct {
		Template string
		Json     bool
	}
	type args struct {
		template string
		data     interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"", fields{"", true}, args{"", &Format{}}, "{\"Template\":\"\",\"Json\":false}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Format{
				Template: tt.fields.Template,
				Json:     tt.fields.Json,
			}
			if got := f.Parse(tt.args.template, tt.args.data); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
