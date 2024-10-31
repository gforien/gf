package semver

import (
	"testing"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected Version
		hasError bool
	}{
		{"1.2.3", Version{major: 1, minor: 2, patch: 3}, false},
		{"0.0.1", Version{major: 0, minor: 0, patch: 1}, false},
		{"10.0.0", Version{major: 10, minor: 0, patch: 0}, false},
		{"1.2", Version{}, true},
		{"1", Version{}, true},
		{"", Version{}, true},
		{"a.b.c", Version{}, true},
	}

	for _, test := range tests {
		result, err := FromString(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("FromString(%s): expected error=%v, got %v", test.input, test.hasError, err != nil)
		}
		if err == nil && result != test.expected {
			t.Errorf("FromString(%s): expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestString(t *testing.T) {
	v := Version{major: 1, minor: 2, patch: 3}
	expected := "1.2.3"
	if v.String() != expected {
		t.Errorf("Version.String(): expected %s, got %s", expected, v.String())
	}
}

func TestReleaseMajor(t *testing.T) {
	v := Version{major: 1, minor: 2, patch: 3}
	expected := Version{major: 2, minor: 0, patch: 0}
	result := v.ReleaseMajor()
	if result != expected {
		t.Errorf("Version.ReleaseMajor(): expected %v, got %v", expected, result)
	}
}

func TestReleaseMinor(t *testing.T) {
	v := Version{major: 1, minor: 2, patch: 3}
	expected := Version{major: 1, minor: 3, patch: 0}
	result := v.ReleaseMinor()
	if result != expected {
		t.Errorf("Version.ReleaseMinor(): expected %v, got %v", expected, result)
	}
}

func TestReleasePatch(t *testing.T) {
	v := Version{major: 1, minor: 2, patch: 3}
	expected := Version{major: 1, minor: 2, patch: 4}
	result := v.ReleasePatch()
	if result != expected {
		t.Errorf("Version.ReleasePatch(): expected %v, got %v", expected, result)
	}
}

func TestErrCannotParse(t *testing.T) {
	input := "1.2"
	_, err := FromString(input)
	if err == nil {
		t.Fatalf("FromString(%s): expected error, got nil", input)
	}

	if _, ok := err.(*ErrCannotParse); !ok {
		t.Fatalf("FromString(%s): expected *ErrCannotParse, got %T", input, err)
	}
}
