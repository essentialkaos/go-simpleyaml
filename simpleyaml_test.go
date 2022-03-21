package simpleyaml

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const TEST_DATA = `
name: John Doe
age: 35
balance: 45.89
admin: true
categories:
   - category1
   - category2
meta:
  uid: 120
  gid: 350
test1:
    test2:
        test3:
            - test31
            - test32
            - test33
        test4: test99
array1:
   - test1
   -
   - test3
array2:
  - file: test1
    size: 100
  - file: test2
    size: 200
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type YamlSuite struct {
	yaml *Yaml
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&YamlSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ys *YamlSuite) SetUpSuite(c *C) {
	yaml, err := NewYaml([]byte(TEST_DATA))

	if err != nil {
		c.Fatal(err.Error())
	}

	ys.yaml = yaml
}

func (ys *YamlSuite) TestMarshal(c *C) {
	data, err := ys.yaml.Encode()

	c.Assert(err, IsNil)
	c.Assert(data, NotNil)
	c.Assert(data, Not(HasLen), 0)

	var yaml *Yaml

	data, err = yaml.MarshalYAML()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
	c.Assert(data, IsNil)
}

func (ys *YamlSuite) TestUnmarshal(c *C) {
	yaml := New()
	err := yaml.UnmarshalYAML([]byte(TEST_DATA))

	c.Assert(err, IsNil)
	c.Assert(yaml.IsMap(), Equals, true)

	yaml = nil

	err = yaml.UnmarshalYAML([]byte(TEST_DATA))

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
}

func (ys *YamlSuite) TestBasic(c *C) {
	yaml := New()

	c.Assert(yaml, NotNil)
	c.Assert(yaml.data, NotNil)

	yaml, err := NewYaml([]byte("v: [A,"))

	c.Assert(yaml, IsNil)
	c.Assert(err, NotNil)

	c.Assert(ys.yaml.Interface(), NotNil)
	c.Assert(yaml.Interface(), IsNil)
}

func (ys *YamlSuite) TestNil(c *C) {
	var err error

	var yaml *Yaml

	_, err = yaml.Get("unknown").Bool()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)

	_, err = yaml.Get("unknown").String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)

	_, err = yaml.Get("unknown").Int()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)

	_, err = yaml.Get("unknown").Float()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)

	_, err = yaml.Get("unknown").Array()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)

	_, err = yaml.Get("unknown").Map()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
}

func (ys *YamlSuite) TestDump(c *C) {
	c.Assert(ys.yaml.Get("unknown").Dump(), Equals, "")
	c.Assert(ys.yaml.Get("name").Dump(), Equals, "John Doe")
	c.Assert(ys.yaml.Get("age").Dump(), Equals, "35")
	c.Assert(ys.yaml.Get("balance").Dump(), Equals, "45.89")
	c.Assert(ys.yaml.Get("admin").Dump(), Equals, "true")
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
	c.Assert(err, Equals, ErrYAMLIsNil)
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
	c.Assert(err, Equals, ErrYAMLIsNil)
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
	c.Assert(err, Equals, ErrYAMLIsNil)
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
	c.Assert(err, Equals, ErrYAMLIsNil)
}

func (ys *YamlSuite) TestBytes(c *C) {
	var err error
	var val []byte

	val, err = ys.yaml.Get("name").Bytes()

	c.Assert(val, DeepEquals, []byte("John Doe"))
	c.Assert(err, IsNil)

	_, err = ys.yaml.Get("admin").Bytes()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrByteTypeAssertion)

	_, err = ys.yaml.Get("age").Bytes()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrByteTypeAssertion)

	_, err = ys.yaml.Get("balance").Bytes()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrByteTypeAssertion)

	_, err = ys.yaml.Get("unknown").Bytes()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
}

func (ys *YamlSuite) TestArray(c *C) {
	var err error
	var val []interface{}

	val, err = ys.yaml.Get("categories").Array()

	c.Assert(val, DeepEquals, []interface{}{"category1", "category2"})
	c.Assert(err, IsNil)
}

func (ys *YamlSuite) TestStringArray(c *C) {
	var err error
	var val []string

	val, err = ys.yaml.Get("categories").StringArray()

	c.Assert(val, DeepEquals, []string{"category1", "category2"})
	c.Assert(err, IsNil)

	val, err = ys.yaml.Get("array1").StringArray()

	c.Assert(val, DeepEquals, []string{"test1", "", "test3"})
	c.Assert(err, IsNil)

	var yaml *Yaml

	val, err = yaml.Get("categories").StringArray()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
	c.Assert(val, IsNil)

	val, err = ys.yaml.Get("admin").StringArray()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrArrayTypeAssertion)
	c.Assert(val, IsNil)

	val, err = ys.yaml.Get("array2").StringArray()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrStringArrayTypeAssertion)
	c.Assert(val, IsNil)
}

func (ys *YamlSuite) TestIsArray(c *C) {
	c.Assert(ys.yaml.Get("categories").IsArray(), Equals, true)
	c.Assert(ys.yaml.Get("name").IsArray(), Equals, false)
}

func (ys *YamlSuite) TestMap(c *C) {
	val, err := ys.yaml.Get("meta").Map()

	c.Assert(err, IsNil)
	c.Assert(val, NotNil)

	val, err = ys.yaml.Get("admin").Map()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrMapTypeAssertion)
	c.Assert(val, IsNil)

	c.Assert(ys.yaml.Get("meta").IsMap(), Equals, true)

	keys, err := ys.yaml.GetMapKeys()

	c.Assert(err, IsNil)
	c.Assert(keys, HasLen, 9)

	keys, err = ys.yaml.Get("admin").GetMapKeys()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrMapTypeAssertion)
	c.Assert(keys, IsNil)
}

func (ys *YamlSuite) TestMustArray(c *C) {
	var def = []interface{}{"test1", "test2"}

	val := ys.yaml.Get("categories").MustArray(def)

	c.Assert(val, NotNil)
	c.Assert(val, DeepEquals, []interface{}{"category1", "category2"})

	val = ys.yaml.Get("admin").MustArray(def)

	c.Assert(val, NotNil)
	c.Assert(val, DeepEquals, def)
}

func (ys *YamlSuite) TestMustMap(c *C) {
	var def = map[interface{}]interface{}{"test": "test1"}

	val := ys.yaml.Get("meta").MustMap(def)

	c.Assert(val, NotNil)
	c.Assert(val, DeepEquals, map[interface{}]interface{}{"uid": 120, "gid": 350})

	val = ys.yaml.Get("admin").MustMap(def)

	c.Assert(val, NotNil)
	c.Assert(val, DeepEquals, def)
}

func (ys *YamlSuite) TestMustString(c *C) {
	var def = "Bob Smith"

	val := ys.yaml.Get("name").MustString(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, "John Doe")

	val = ys.yaml.Get("admin").MustString(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, def)
}

func (ys *YamlSuite) TestMustStringArray(c *C) {
	var def = []string{"category3", "category4"}

	val := ys.yaml.Get("categories").MustStringArray(def)

	c.Assert(val, NotNil)
	c.Assert(val, DeepEquals, []string{"category1", "category2"})

	val = ys.yaml.Get("admin").MustStringArray(def)

	c.Assert(val, NotNil)
	c.Assert(val, DeepEquals, def)
}

func (ys *YamlSuite) TestMustInt(c *C) {
	var def = 20

	val := ys.yaml.Get("age").MustInt(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, 35)

	val = ys.yaml.Get("admin").MustInt(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, def)
}

func (ys *YamlSuite) TestMustFloat(c *C) {
	var def = 71.24

	val := ys.yaml.Get("balance").MustFloat(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, 45.89)

	val = ys.yaml.Get("admin").MustFloat(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, def)
}

func (ys *YamlSuite) TestMustBool(c *C) {
	var def = false

	val := ys.yaml.Get("admin").MustBool(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, true)

	val = ys.yaml.Get("name").MustBool(def)

	c.Assert(val, NotNil)
	c.Assert(val, Equals, def)
}

func (ys *YamlSuite) TestGetPath(c *C) {
	var err error
	var val string

	val, err = ys.yaml.GetPath("test1", "test2", "test4").String()

	c.Assert(err, IsNil)
	c.Assert(val, Not(Equals), "")
	c.Assert(val, Equals, "test99")

	val, err = ys.yaml.Get("test1").Get("test2").Get("test4").String()

	c.Assert(err, IsNil)
	c.Assert(val, Not(Equals), "")
	c.Assert(val, Equals, "test99")
}

func (ys *YamlSuite) TestCheckGet(c *C) {
	_, ok := ys.yaml.CheckGet("name")

	c.Assert(ok, Equals, true)

	_, ok = ys.yaml.CheckGet("unknown")

	c.Assert(ok, Equals, false)

	_, ok = ys.yaml.Get("admin").CheckGet("unknown")

	c.Assert(ok, Equals, false)
}

func (ys *YamlSuite) TestGetByIndex(c *C) {
	var err error
	var val string

	val, err = ys.yaml.Get("categories").GetByIndex(0).String()

	c.Assert(err, IsNil)
	c.Assert(val, Equals, "category1")

	val, err = ys.yaml.Get("categories").GetByIndex(999).String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
	c.Assert(val, Equals, "")

	val, err = ys.yaml.Get("admin").GetByIndex(0).String()

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrYAMLIsNil)
	c.Assert(val, Equals, "")
}

func (ys *YamlSuite) TestExistChecks(c *C) {
	c.Assert(ys.yaml.IsExist("admin"), Equals, true)
	c.Assert(ys.yaml.IsExist("unknown"), Equals, false)

	c.Assert(ys.yaml.IsPathExist("test1", "test2", "test4"), Equals, true)
	c.Assert(ys.yaml.IsPathExist("test1", "test20", "test4"), Equals, false)

	c.Assert(ys.yaml.Get("categories").IsIndexExist(1), Equals, true)
	c.Assert(ys.yaml.Get("categories").IsIndexExist(99), Equals, false)
}
