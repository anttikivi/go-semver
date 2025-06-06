// Copyright (c) 2025 Antti Kivi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package semver_test

import (
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/anttikivi/semver"
)

func TestVersionsSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input []string
		want  []string
	}{
		{
			[]string{
				"1.2.3",
				"1.0",
				"1.3",
				"2",
				"0.4.2",
			},
			[]string{
				"0.4.2",
				"1.0.0",
				"1.2.3",
				"1.3.0",
				"2.0.0",
			},
		},
		{
			[]string{
				"1.2.3",
				"1.0",
				"10",
				"1.3",
				"2",
				"0.4.2",
			},
			[]string{
				"0.4.2",
				"1.0.0",
				"1.2.3",
				"1.3.0",
				"2.0.0",
				"10.0.0",
			},
		},
		{
			[]string{
				"10",
				"2",
				"12",
				"1.2",
				"1.0",
				"1",
			},
			[]string{
				"1.0.0",
				"1.0.0",
				"1.2.0",
				"2.0.0",
				"10.0.0",
				"12.0.0",
			},
		},
		{
			[]string{
				"1-beta",
				"1",
				"2-beta",
				"2.0.1",
				"1-0.beta",
				"1-alpha",
				"1-alpha.1",
				"1.3",
				"1-alpha.beta",
			},
			[]string{
				"1.0.0-0.beta",
				"1.0.0-alpha",
				"1.0.0-alpha.1",
				"1.0.0-alpha.beta",
				"1.0.0-beta",
				"1.0.0",
				"1.3.0",
				"2.0.0-beta",
				"2.0.1",
			},
		},
		{
			[]string{
				"1.0.0-beta.2",
				"1.0.0-alpha.beta",
				"1.0.0-beta.11",
				"1.0.0",
				"1.0.0-alpha",
				"1.0.0-beta",
				"1.0.0-rc.1",
				"1.0.0-alpha.1",
			},
			[]string{
				"1.0.0-alpha",
				"1.0.0-alpha.1",
				"1.0.0-alpha.beta",
				"1.0.0-beta",
				"1.0.0-beta.2",
				"1.0.0-beta.11",
				"1.0.0-rc.1",
				"1.0.0",
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			versions := make(semver.Versions, len(tt.input))

			for j, s := range tt.input {
				v, _ := semver.ParseLax(s)
				if v == nil {
					t.Fatalf("Setup error: Version is nil for input %q", s)
				}

				versions[j] = v
			}

			sort.Sort(versions)

			x := make([]string, len(versions))

			for j, v := range versions {
				x[j] = v.String()
			}

			if !reflect.DeepEqual(x, tt.want) {
				t.Errorf("sort.Sort(%#v) = %#v, want %#v", tt.input, x, tt.want)
			}
		})
	}
}
