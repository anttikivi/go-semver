package semver_test

import (
	"testing"

	"github.com/anttikivi/go-semver"
)

var tests = []struct { //nolint:gochecknoglobals // the test cases need to be global
	in  string
	out string
}{
	{"", ""},

	{"0.1.0-alpha.24+sha.19031c2.darwin.amd64", "0.1.0-alpha.24"},
	{"0.1.0-alpha.24+sha.19031c2-darwin-amd64", "0.1.0-alpha.24"},

	{"bad", ""},
	{"1-alpha.beta.gamma", ""},
	{"1-pre", ""},
	{"1+meta", ""},
	{"1-pre+meta", ""},
	{"1.2-pre", ""},
	{"1.2+meta", ""},
	{"1.2-pre+meta", ""},
	{"1.0.0-alpha", "1.0.0-alpha"},
	{"1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"1.0.0-beta", "1.0.0-beta"},
	{"1.0.0-beta.2", "1.0.0-beta.2"},
	{"1.0.0-beta.11", "1.0.0-beta.11"},
	{"1.0.0-rc.1", "1.0.0-rc.1"},
	{"1", ""},
	{"1.0", ""},
	{"1.0.0", "1.0.0"},
	{"1.2", ""},
	{"1.2.0", "1.2.0"},
	{"1.2.3-456", "1.2.3-456"},
	{"1.2.3-456.789", "1.2.3-456.789"},
	{"1.2.3-456-789", "1.2.3-456-789"},
	{"1.2.3-456a", "1.2.3-456a"},
	{"1.2.3-pre", "1.2.3-pre"},
	{"1.2.3-pre+meta", "1.2.3-pre"},
	{"1.2.3-pre.1", "1.2.3-pre.1"},
	{"1.2.3-zzz", "1.2.3-zzz"},
	{"1.2.3", "1.2.3"},
	{"1.2.3+meta", "1.2.3"},
	{"1.2.3+meta-pre", "1.2.3"},
	{"1.2.3+meta-pre.sha.256a", "1.2.3"},

	{"vbad", ""},
	{"v1-alpha.beta.gamma", ""},
	{"v1-pre", ""},
	{"v1+meta", ""},
	{"v1-pre+meta", ""},
	{"v1.2-pre", ""},
	{"v1.2+meta", ""},
	{"v1.2-pre+meta", ""},
	{"v1.0.0-alpha", "1.0.0-alpha"},
	{"v1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"v1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"v1.0.0-beta", "1.0.0-beta"},
	{"v1.0.0-beta.2", "1.0.0-beta.2"},
	{"v1.0.0-beta.11", "1.0.0-beta.11"},
	{"v1.0.0-rc.1", "1.0.0-rc.1"},
	{"v1", ""},
	{"v1.0", ""},
	{"v1.0.0", "1.0.0"},
	{"v1.2", ""},
	{"v1.2.0", "1.2.0"},
	{"v1.2.3-456", "1.2.3-456"},
	{"v1.2.3-456.789", "1.2.3-456.789"},
	{"v1.2.3-456-789", "1.2.3-456-789"},
	{"v1.2.3-456a", "1.2.3-456a"},
	{"v1.2.3-pre", "1.2.3-pre"},
	{"v1.2.3-pre+meta", "1.2.3-pre"},
	{"v1.2.3-pre.1", "1.2.3-pre.1"},
	{"v1.2.3-zzz", "1.2.3-zzz"},
	{"v1.2.3", "1.2.3"},
	{"v1.2.3+meta", "1.2.3"},
	{"v1.2.3+meta-pre", "1.2.3"},
	{"v1.2.3+meta-pre.sha.256a", "1.2.3"},

	{"reginaldbad", ""},
	{"reginald1-alpha.beta.gamma", ""},
	{"reginald1-pre", ""},
	{"reginald1+meta", ""},
	{"reginald1-pre+meta", ""},
	{"reginald1.2-pre", ""},
	{"reginald1.2+meta", ""},
	{"reginald1.2-pre+meta", ""},
	{"reginald1.0.0-alpha", "1.0.0-alpha"},
	{"reginald1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"reginald1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"reginald1.0.0-beta", "1.0.0-beta"},
	{"reginald1.0.0-beta.2", "1.0.0-beta.2"},
	{"reginald1.0.0-beta.11", "1.0.0-beta.11"},
	{"reginald1.0.0-rc.1", "1.0.0-rc.1"},
	{"reginald1", ""},
	{"reginald1.0", ""},
	{"reginald1.0.0", "1.0.0"},
	{"reginald1.2", ""},
	{"reginald1.2.0", "1.2.0"},
	{"reginald1.2.3-456", "1.2.3-456"},
	{"reginald1.2.3-456.789", "1.2.3-456.789"},
	{"reginald1.2.3-456-789", "1.2.3-456-789"},
	{"reginald1.2.3-456a", "1.2.3-456a"},
	{"reginald1.2.3-pre", "1.2.3-pre"},
	{"reginald1.2.3-pre+meta", "1.2.3-pre"},
	{"reginald1.2.3-pre.1", "1.2.3-pre.1"},
	{"reginald1.2.3-zzz", "1.2.3-zzz"},
	{"reginald1.2.3", "1.2.3"},
	{"reginald1.2.3+meta", "1.2.3"},
	{"reginald1.2.3+meta-pre", "1.2.3"},
	{"reginald1.2.3+meta-pre.sha.256a", "1.2.3"},

	{"reggiebad", ""},
	{"reggie1-alpha.beta.gamma", ""},
	{"reggie1-pre", ""},
	{"reggie1+meta", ""},
	{"reggie1-pre+meta", ""},
	{"reggie1.2-pre", ""},
	{"reggie1.2+meta", ""},
	{"reggie1.2-pre+meta", ""},
	{"reggie1.0.0-alpha", "1.0.0-alpha"},
	{"reggie1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"reggie1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"reggie1.0.0-beta", "1.0.0-beta"},
	{"reggie1.0.0-beta.2", "1.0.0-beta.2"},
	{"reggie1.0.0-beta.11", "1.0.0-beta.11"},
	{"reggie1.0.0-rc.1", "1.0.0-rc.1"},
	{"reggie1", ""},
	{"reggie1.0", ""},
	{"reggie1.0.0", "1.0.0"},
	{"reggie1.2", ""},
	{"reggie1.2.0", "1.2.0"},
	{"reggie1.2.3-456", "1.2.3-456"},
	{"reggie1.2.3-456.789", "1.2.3-456.789"},
	{"reggie1.2.3-456-789", "1.2.3-456-789"},
	{"reggie1.2.3-456a", "1.2.3-456a"},
	{"reggie1.2.3-pre", "1.2.3-pre"},
	{"reggie1.2.3-pre+meta", "1.2.3-pre"},
	{"reggie1.2.3-pre.1", "1.2.3-pre.1"},
	{"reggie1.2.3-zzz", "1.2.3-zzz"},
	{"reggie1.2.3", "1.2.3"},
	{"reggie1.2.3+meta", "1.2.3"},
	{"reggie1.2.3+meta-pre", "1.2.3"},
	{"reggie1.2.3+meta-pre.sha.256a", "1.2.3"},
}

func TestIsValid(t *testing.T) { //nolint:funlen // lot's of test cases
	t.Parallel()

	tests := []struct {
		v    string
		want bool
	}{
		{"", false},

		{"0.1.0-alpha.24+sha.19031c2.darwin.amd64", true},
		{"0.1.0-alpha.24+sha.19031c2-darwin-amd64", true},

		{"bad", false},
		{"1-alpha.beta.gamma", false},
		{"1-pre", false},
		{"1+meta", false},
		{"1-pre+meta", false},
		{"1.2-pre", false},
		{"1.2+meta", false},
		{"1.2-pre+meta", false},
		{"1.0.0-alpha", true},
		{"1.0.0-alpha.1", true},
		{"1.0.0-alpha.beta", true},
		{"1.0.0-beta", true},
		{"1.0.0-beta.2", true},
		{"1.0.0-beta.11", true},
		{"1.0.0-rc.1", true},
		{"1", false},
		{"1.0", false},
		{"1.0.0", true},
		{"1.2", false},
		{"1.2.0", true},
		{"1.2.3-456", true},
		{"1.2.3-456.789", true},
		{"1.2.3-456-789", true},
		{"1.2.3-456a", true},
		{"1.2.3-pre", true},
		{"1.2.3-pre+meta", true},
		{"1.2.3-pre.1", true},
		{"1.2.3-zzz", true},
		{"1.2.3", true},
		{"1.2.3+meta", true},
		{"1.2.3+meta-pre", true},
		{"1.2.3+meta-pre.sha.256a", true},

		{"vbad", false},
		{"v1-alpha.beta.gamma", false},
		{"v1-pre", false},
		{"v1+meta", false},
		{"v1-pre+meta", false},
		{"v1.2-pre", false},
		{"v1.2+meta", false},
		{"v1.2-pre+meta", false},
		{"v1.0.0-alpha", true},
		{"v1.0.0-alpha.1", true},
		{"v1.0.0-alpha.beta", true},
		{"v1.0.0-beta", true},
		{"v1.0.0-beta.2", true},
		{"v1.0.0-beta.11", true},
		{"v1.0.0-rc.1", true},
		{"v1", false},
		{"v1.0", false},
		{"v1.0.0", true},
		{"v1.2", false},
		{"v1.2.0", true},
		{"v1.2.3-456", true},
		{"v1.2.3-456.789", true},
		{"v1.2.3-456-789", true},
		{"v1.2.3-456a", true},
		{"v1.2.3-pre", true},
		{"v1.2.3-pre+meta", true},
		{"v1.2.3-pre.1", true},
		{"v1.2.3-zzz", true},
		{"v1.2.3", true},
		{"v1.2.3+meta", true},
		{"v1.2.3+meta-pre", true},
		{"v1.2.3+meta-pre.sha.256a", true},

		{"semverbad", false},
		{"semver1-alpha.beta.gamma", false},
		{"semver1-pre", false},
		{"semver1+meta", false},
		{"semver1-pre+meta", false},
		{"semver1.2-pre", false},
		{"semver1.2+meta", false},
		{"semver1.2-pre+meta", false},
		{"semver1.0.0-alpha", false},
		{"semver1.0.0-alpha.1", false},
		{"semver1.0.0-alpha.beta", false},
		{"semver1.0.0-beta", false},
		{"semver1.0.0-beta.2", false},
		{"semver1.0.0-beta.11", false},
		{"semver1.0.0-rc.1", false},
		{"semver1", false},
		{"semver1.0", false},
		{"semver1.0.0", false},
		{"semver1.2", false},
		{"semver1.2.0", false},
		{"semver1.2.3-456", false},
		{"semver1.2.3-456.789", false},
		{"semver1.2.3-456-789", false},
		{"semver1.2.3-456a", false},
		{"semver1.2.3-pre", false},
		{"semver1.2.3-pre+meta", false},
		{"semver1.2.3-pre.1", false},
		{"semver1.2.3-zzz", false},
		{"semver1.2.3", false},
		{"semver1.2.3+meta", false},
		{"semver1.2.3+meta-pre", false},
		{"semver1.2.3+meta-pre.sha.256a", false},
	}
	for _, tt := range tests {
		name := tt.v
		if name == "" {
			name = "empty"
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if ok := semver.IsValid(tt.v); ok != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.v, ok, !ok)
			}
		})
	}
}

func TestIsValidPrefix(t *testing.T) { //nolint:funlen // lot's of test cases
	t.Parallel()

	tests := []struct {
		v    string
		want bool
	}{
		{"", false},

		{"0.1.0-alpha.24+sha.19031c2.darwin.amd64", true},
		{"0.1.0-alpha.24+sha.19031c2-darwin-amd64", true},

		{"bad", false},
		{"1-alpha.beta.gamma", false},
		{"1-pre", false},
		{"1+meta", false},
		{"1-pre+meta", false},
		{"1.2-pre", false},
		{"1.2+meta", false},
		{"1.2-pre+meta", false},
		{"1.0.0-alpha", true},
		{"1.0.0-alpha.1", true},
		{"1.0.0-alpha.beta", true},
		{"1.0.0-beta", true},
		{"1.0.0-beta.2", true},
		{"1.0.0-beta.11", true},
		{"1.0.0-rc.1", true},
		{"1", false},
		{"1.0", false},
		{"1.0.0", true},
		{"1.2", false},
		{"1.2.0", true},
		{"1.2.3-456", true},
		{"1.2.3-456.789", true},
		{"1.2.3-456-789", true},
		{"1.2.3-456a", true},
		{"1.2.3-pre", true},
		{"1.2.3-pre+meta", true},
		{"1.2.3-pre.1", true},
		{"1.2.3-zzz", true},
		{"1.2.3", true},
		{"1.2.3+meta", true},
		{"1.2.3+meta-pre", true},
		{"1.2.3+meta-pre.sha.256a", true},

		{"vbad", false},
		{"v1-alpha.beta.gamma", false},
		{"v1-pre", false},
		{"v1+meta", false},
		{"v1-pre+meta", false},
		{"v1.2-pre", false},
		{"v1.2+meta", false},
		{"v1.2-pre+meta", false},
		{"v1.0.0-alpha", true},
		{"v1.0.0-alpha.1", true},
		{"v1.0.0-alpha.beta", true},
		{"v1.0.0-beta", true},
		{"v1.0.0-beta.2", true},
		{"v1.0.0-beta.11", true},
		{"v1.0.0-rc.1", true},
		{"v1", false},
		{"v1.0", false},
		{"v1.0.0", true},
		{"v1.2", false},
		{"v1.2.0", true},
		{"v1.2.3-456", true},
		{"v1.2.3-456.789", true},
		{"v1.2.3-456-789", true},
		{"v1.2.3-456a", true},
		{"v1.2.3-pre", true},
		{"v1.2.3-pre+meta", true},
		{"v1.2.3-pre.1", true},
		{"v1.2.3-zzz", true},
		{"v1.2.3", true},
		{"v1.2.3+meta", true},
		{"v1.2.3+meta-pre", true},
		{"v1.2.3+meta-pre.sha.256a", true},

		{"semverbad", false},
		{"semver1-alpha.beta.gamma", false},
		{"semver1-pre", false},
		{"semver1+meta", false},
		{"semver1-pre+meta", false},
		{"semver1.2-pre", false},
		{"semver1.2+meta", false},
		{"semver1.2-pre+meta", false},
		{"semver1.0.0-alpha", true},
		{"semver1.0.0-alpha.1", true},
		{"semver1.0.0-alpha.beta", true},
		{"semver1.0.0-beta", true},
		{"semver1.0.0-beta.2", true},
		{"semver1.0.0-beta.11", true},
		{"semver1.0.0-rc.1", true},
		{"semver1", false},
		{"semver1.0", false},
		{"semver1.0.0", true},
		{"semver1.2", false},
		{"semver1.2.0", true},
		{"semver1.2.3-456", true},
		{"semver1.2.3-456.789", true},
		{"semver1.2.3-456-789", true},
		{"semver1.2.3-456a", true},
		{"semver1.2.3-pre", true},
		{"semver1.2.3-pre+meta", true},
		{"semver1.2.3-pre.1", true},
		{"semver1.2.3-zzz", true},
		{"semver1.2.3", true},
		{"semver1.2.3+meta", true},
		{"semver1.2.3+meta-pre", true},
		{"semver1.2.3+meta-pre.sha.256a", true},

		{"sembad", false},
		{"sem1-alpha.beta.gamma", false},
		{"sem1-pre", false},
		{"sem1+meta", false},
		{"sem1-pre+meta", false},
		{"sem1.2-pre", false},
		{"sem1.2+meta", false},
		{"sem1.2-pre+meta", false},
		{"sem1.0.0-alpha", false},
		{"sem1.0.0-alpha.1", false},
		{"sem1.0.0-alpha.beta", false},
		{"sem1.0.0-beta", false},
		{"sem1.0.0-beta.2", false},
		{"sem1.0.0-beta.11", false},
		{"sem1.0.0-rc.1", false},
		{"sem1", false},
		{"sem1.0", false},
		{"sem1.0.0", false},
		{"sem1.2", false},
		{"sem1.2.0", false},
		{"sem1.2.3-456", false},
		{"sem1.2.3-456.789", false},
		{"sem1.2.3-456-789", false},
		{"sem1.2.3-456a", false},
		{"sem1.2.3-pre", false},
		{"sem1.2.3-pre+meta", false},
		{"sem1.2.3-pre.1", false},
		{"sem1.2.3-zzz", false},
		{"sem1.2.3", false},
		{"sem1.2.3+meta", false},
		{"sem1.2.3+meta-pre", false},
		{"sem1.2.3+meta-pre.sha.256a", false},

		{"sebad", false},
		{"se1-alpha.beta.gamma", false},
		{"se1-pre", false},
		{"se1+meta", false},
		{"se1-pre+meta", false},
		{"se1.2-pre", false},
		{"se1.2+meta", false},
		{"se1.2-pre+meta", false},
		{"se1.0.0-alpha", true},
		{"se1.0.0-alpha.1", true},
		{"se1.0.0-alpha.beta", true},
		{"se1.0.0-beta", true},
		{"se1.0.0-beta.2", true},
		{"se1.0.0-beta.11", true},
		{"se1.0.0-rc.1", true},
		{"se1", false},
		{"se1.0", false},
		{"se1.0.0", true},
		{"se1.2", false},
		{"se1.2.0", true},
		{"se1.2.3-456", true},
		{"se1.2.3-456.789", true},
		{"se1.2.3-456-789", true},
		{"se1.2.3-456a", true},
		{"se1.2.3-pre", true},
		{"se1.2.3-pre+meta", true},
		{"se1.2.3-pre.1", true},
		{"se1.2.3-zzz", true},
		{"se1.2.3", true},
		{"se1.2.3+meta", true},
		{"se1.2.3+meta-pre", true},
		{"se1.2.3+meta-pre.sha.256a", true},
	}
	for _, tt := range tests {
		name := tt.v
		if name == "" {
			name = "empty"
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if ok := semver.IsValidPrefix(tt.v, "semver", "se"); ok != tt.want {
				t.Errorf("IsValidPrefix(%q, %q, %q) = %v, want %v", tt.v, "semver", "se", ok, !ok)
			}
		})
	}
}