package output

import (
	"reflect"
	"testing"

	output "output/readWrite"
)

func TestReadAscii(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := output.ReadAscii(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAscii() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}
