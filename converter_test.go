package iso8583

import (
	"strings"
	"testing"
)

func TestBitmapToBinary(t *testing.T) {
	testList := []struct {
		hex, want   string
		errcontains string
	}{
		{"42 10 00 11 02 C0 48 04", "0100001000010000000000000001000100000010110000000100100000000100", ""},
		{"4210001102C04804", "0100001000010000000000000001000100000010110000000100100000000100", ""},
		{"4210001102C0480n", "", "invalid hex"},
	}
	for _, l := range testList {
		b, err := BitmapToBinary(l.hex)
		if err != nil {
			if l.errcontains == "" {
				t.Fatalf("BitmapToBinary returned wrong error value, got %v, expected nil", err)
			}
			if !strings.Contains(err.Error(), l.errcontains) {
				t.Errorf("BitmapToBinary returned invalid error, expected '%s', got %v", l.errcontains, err)
			}
			continue
		}
		if l.errcontains != "" {
			t.Errorf("BitmapToBinary returned invalid error value, expected '%s', got nil", l.errcontains)
		}
		if b != l.want {
			t.Errorf("BitmapToBinary returned wrong value, got %s, expected %s", b, l.want)
		}
	}

}

func TestHexToBinary(t *testing.T) {
	testList := []struct {
		hex, want   string
		errcontains string
	}{
		{"42", "01000010", ""},
		{"4x", "", "invalid hex"},
		{"42 10 00 11 02 C0 48 04", "", "invalid hex"},
		{"4210001102C04804", "0100001000010000000000000001000100000010110000000100100000000100", ""},
	}
	for _, l := range testList {
		b, err := HexToBinary(l.hex)
		if err != nil {
			if l.errcontains == "" {
				t.Fatalf("HexToBinary returned wrong error value, got %v, expected nil", err)
			}
			if !strings.Contains(err.Error(), l.errcontains) {
				t.Errorf("HexToBinary returned invalid error, expected '%s', got %v", l.errcontains, err)
			}
			continue
		}
		if l.errcontains != "" {
			t.Errorf("HexToBinary returned invalid error value, expected '%s', got nil", l.errcontains)
		}
		if b != l.want {
			t.Errorf("HexToBinary returned wrong value, got %s, expected %s", b, l.want)
		}
	}

}
