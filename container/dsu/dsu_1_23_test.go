//go:build go1.23
// +build go1.23

package dsu

import (
	"testing"
)

func TestNewGo123(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		wantNil  bool
		wantSize int
		wantComp int
	}{
		{
			name:     "valid size go1.23",
			n:        5,
			wantNil:  false,
			wantSize: 5,
			wantComp: 5,
		},
		{
			name:    "zero size go1.23",
			n:       0,
			wantNil: true,
		},
		{
			name:    "negative size go1.23",
			n:       -1,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsu := NewDSU(tt.n)
			if tt.wantNil {
				if dsu != nil {
					t.Errorf("NewDSU(%d) = %v, want nil", tt.n, dsu)
				}
				return
			}

			if dsu == nil {
				t.Fatalf("NewDSU(%d) = nil, want non-nil", tt.n)
			}

			if got := dsu.Size(); got != tt.wantSize {
				t.Errorf("Size() = %d, want %d", got, tt.wantSize)
			}

			if got := dsu.ComponentCount(); got != tt.wantComp {
				t.Errorf("ComponentCount() = %d, want %d", got, tt.wantComp)
			}
		})
	}
}

func TestComplexOperationsGo123(t *testing.T) {
	dsu := NewDSU(6)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	// Create some unions
	dsu.Union(0, 1)
	dsu.Union(2, 3)
	dsu.Union(4, 5)

	// Test connectivity
	if !dsu.Connected(0, 1) {
		t.Error("Connected(0, 1) = false, want true")
	}
	if !dsu.Connected(2, 3) {
		t.Error("Connected(2, 3) = false, want true")
	}
	if !dsu.Connected(4, 5) {
		t.Error("Connected(4, 5) = false, want true")
	}

	// Test non-connectivity
	if dsu.Connected(0, 2) {
		t.Error("Connected(0, 2) = true, want false")
	}

	// Verify component count
	if got := dsu.ComponentCount(); got != 3 {
		t.Errorf("ComponentCount() = %d, want 3", got)
	}
}
