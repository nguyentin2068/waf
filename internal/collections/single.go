// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package collections

import (
	"fmt"
	"strings"

	"github.com/nguyentin2068/waf/collection"
	"github.com/nguyentin2068/waf/internal/corazarules"
	"github.com/nguyentin2068/waf/types"
	"github.com/nguyentin2068/waf/types/variables"
)

type Single struct {
	data     string
	variable variables.RuleVariable
}

var _ collection.Single = &Single{}

// NewSingle creates a new Single.
func NewSingle(variable variables.RuleVariable) *Single {
	return &Single{
		variable: variable,
	}
}

func (c *Single) FindAll() []types.MatchData {
	return []types.MatchData{
		&corazarules.MatchData{
			Variable_: c.variable,
			Value_:    c.data,
		},
	}
}

func (c *Single) Get() string {
	return c.data
}

func (c *Single) Set(value string) {
	c.data = value
}

func (c *Single) Name() string {
	return c.variable.Name()
}

func (c *Single) Reset() {
	c.data = ""
}

func (c *Single) Format(res *strings.Builder) {
	res.WriteString(c.variable.Name())
	res.WriteString(": ")
	res.WriteString(c.data)
}

func (c *Single) String() string {
	return fmt.Sprintf("%s: %s", c.variable.Name(), c.data)
}
