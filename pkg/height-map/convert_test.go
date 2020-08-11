package heightmap

import (
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	testFiles := []string{"N44E033", "N48E009", "N48E010"}
	for _, file := range testFiles {
		heightData, err := os.Open("testdata/" + file + ".hgt")
		assert.NoError(t, err)
		defer heightData.Close()
		_, _ = Convert(heightData)
		assert.NoError(t, err)
	}
}

func TestBigEndianSignedInt16(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			args: args{b: []byte{0xFF, 0xFF}},
			want: -1,
		},
		{
			args: args{b: []byte{0, 0xFF}},
			want: 255,
		},
		{
			args: args{b: []byte{0, 0}},
			want: 0,
		},
		{
			args: args{b: []byte{0x80, 0x00}},
			want: -32768,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigEndianSignedInt16(tt.args.b); got != tt.want {
				t.Errorf("BigEndianSignedInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizedHeight(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			args: args{b: []byte{0xFF, 0xFF}},
			want: -1 + 32768,
		},
		{
			args: args{b: []byte{0, 0xFF}},
			want: 255 + 32768,
		},
		{
			args: args{b: []byte{0, 0}},
			want: 0 + 32768,
		},
		{
			args: args{b: []byte{0x80, 0x00}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizedHeight(tt.args.b); got != tt.want {
				t.Errorf("NormalizedHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
