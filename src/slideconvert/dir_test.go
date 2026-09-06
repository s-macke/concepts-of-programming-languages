package main

import (
	"testing"

	"slideconvert/tools/present"
)

func TestPlayable(t *testing.T) {
	previous := present.PlayEnabled
	t.Cleanup(func() { present.PlayEnabled = previous })

	tests := []struct {
		name    string
		enabled bool
		code    present.Code
		want    bool
	}{
		{
			name:    "Go play directive",
			enabled: true,
			code:    present.Code{Play: true, Ext: ".go"},
			want:    true,
		},
		{
			name:    "Go code directive",
			enabled: true,
			code:    present.Code{Play: false, Ext: ".go"},
			want:    false,
		},
		{
			name:    "non-Go play directive",
			enabled: true,
			code:    present.Code{Play: true, Ext: ".java"},
			want:    false,
		},
		{
			name:    "playground disabled",
			enabled: false,
			code:    present.Code{Play: true, Ext: ".go"},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			present.PlayEnabled = tt.enabled
			if got := playable(tt.code); got != tt.want {
				t.Fatalf("playable(%+v) = %t; want %t", tt.code, got, tt.want)
			}
		})
	}
}
