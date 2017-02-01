package simpleyaml

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type YamlSuite struct {
	yaml *Yaml
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&YamlSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ys *YamlSuite) SetUpSuite(c *C) {
	yaml, err := NewYaml([]byte(`
name: John Doe
age: 35
balance: 45.89
admin: true
categories:
   - category1
   - category2
test1:
    test2:
        test3:
            - test31
            - test32
            - test33
        test4: test99
`))

	if err != nil {
		c.Fatal(err.Error())
	}

	ys.yaml = yaml
}

func (ys *YamlSuite) TestBool(c *C) {
	var err error
	var val bool

	val, err = ys.yaml.Get("admin").Bool()

	c.Assert(val, Equals, true)
	c.Assert(err, IsNil)

	_, err = ys.yaml.Get("name").Bool()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrBoolTypeAssertion)

	_, err = ys.yaml.Get("age").Bool()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrBoolTypeAssertion)

	_, err = ys.yaml.Get("balance").Bool()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrBoolTypeAssertion)

	_, err = ys.yaml.Get("unknown").Bool()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrBoolTypeAssertion)
}

func (ys *YamlSuite) TestString(c *C) {
	var err error
	var val string

	val, err = ys.yaml.Get("name").String()

	c.Assert(val, Equals, "John Doe")
	c.Assert(err, IsNil)

	_, err = ys.yaml.Get("admin").String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrStringTypeAssertion)

	_, err = ys.yaml.Get("age").String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrStringTypeAssertion)

	_, err = ys.yaml.Get("balance").String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrStringTypeAssertion)

	_, err = ys.yaml.Get("unknown").String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrStringTypeAssertion)
}

func (ys *YamlSuite) TestInt(c *C) {
	var err error
	var val int

	val, err = ys.yaml.Get("age").Int()

	c.Assert(val, Equals, 35)
	c.Assert(err, IsNil)

	_, err = ys.yaml.Get("admin").Int()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrIntTypeAssertion)

	_, err = ys.yaml.Get("name").Int()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrIntTypeAssertion)

	_, err = ys.yaml.Get("balance").Int()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrIntTypeAssertion)

	_, err = ys.yaml.Get("unknown").Int()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrIntTypeAssertion)
}

func (ys *YamlSuite) TestFloat(c *C) {
	var err error
	var val float64

	val, err = ys.yaml.Get("balance").Float()

	c.Assert(val, Equals, 45.89)
	c.Assert(err, IsNil)

	_, err = ys.yaml.Get("name").Float()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrFloatTypeAssertion)

	_, err = ys.yaml.Get("age").Float()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrFloatTypeAssertion)

	_, err = ys.yaml.Get("admin").Float()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrFloatTypeAssertion)

	_, err = ys.yaml.Get("unknown").Float()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrFloatTypeAssertion)
}
