// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package modules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/uberfx/core"
)

func TestWithName_OK(t *testing.T) {
	h := mci()
	option := WithName("foo")
	err := option(h)
	require.NoError(t, err, "Should be able to set name")

	assert.Equal(t, "foo", h.Name)
}

func TestWithRoles_OK(t *testing.T) {
	h := mci()
	option := WithRoles("foo", "bar")
	err := option(h)
	require.NoError(t, err, "Should be able to set name")

	assert.Equal(t, []string{"foo", "bar"}, h.Roles)
}

func mci() *core.ModuleCreateInfo {
	return &core.ModuleCreateInfo{
		Host: core.NullServiceHost(),
	}
}