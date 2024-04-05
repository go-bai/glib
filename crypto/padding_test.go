package crypto

import (
	"reflect"
	"testing"
)

func TestPKCS7Padding(t *testing.T) {

	blockSize := 3

	type args struct {
		data      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "pad three 3",
			args: args{
				data:      []byte{},
				blockSize: blockSize,
			},
			want: []byte{3, 3, 3},
		},
		{
			name: "pad two 2",
			args: args{
				data:      []byte{103},
				blockSize: blockSize,
			},
			want: []byte{103, 2, 2},
		},
		{
			name: "pad one 1",
			args: args{
				data:      []byte{103, 111},
				blockSize: blockSize,
			},
			want: []byte{103, 111, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PKCS7Padding(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS7Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}
