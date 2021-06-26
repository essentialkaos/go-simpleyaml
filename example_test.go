package simpleyaml

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"sort"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewYaml() {
	data, err := ioutil.ReadFile("file.yml")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	yml, err := NewYaml(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(yml.Get("abcd").MustInt(10))
}

func ExampleYaml_Interface() {
	yml, _ := NewYaml([]byte("abcd: 123"))

	fmt.Println(yml.Get("abcd").Interface() != nil)

	// Output: true
}

func ExampleYaml_Encode() {
	yml, _ := NewYaml([]byte("abcd: 123"))

	data, err := yml.Encode()

	fmt.Println(err)
	fmt.Println(string(data))

	// Output:
	// <nil>
	// abcd: 123
}

func ExampleYaml_MarshalYAML() {
	yml, _ := NewYaml([]byte("abcd: 123"))

	data, err := yml.MarshalYAML()

	fmt.Println(err)
	fmt.Println(string(data))

	// Output:
	// <nil>
	// abcd: 123
}

func ExampleYaml_UnmarshalYAML() {
	yml := New()
	err := yml.UnmarshalYAML([]byte("abcd: 123"))

	fmt.Println(err)
	fmt.Println(yml.Get("abcd").MustInt(0))

	// Output:
	// <nil>
	// 123
}

func ExampleYaml_Int() {
	yml, _ := NewYaml([]byte("abcd: 123"))
	v, err := yml.Get("abcd").Int()

	fmt.Println(err)
	fmt.Println(v)

	// Output:
	// <nil>
	// 123
}

func ExampleYaml_Float() {
	yml, _ := NewYaml([]byte("abcd: 1.23"))
	v, err := yml.Get("abcd").Float()

	fmt.Println(err)
	fmt.Println(v)

	// Output:
	// <nil>
	// 1.23
}

func ExampleYaml_Bool() {
	yml, _ := NewYaml([]byte("abcd: true"))
	v, err := yml.Get("abcd").Bool()

	fmt.Println(err)
	fmt.Println(v)

	// Output:
	// <nil>
	// true
}

func ExampleYaml_String() {
	yml, _ := NewYaml([]byte("abcd: John Doe"))
	v, err := yml.Get("abcd").String()

	fmt.Println(err)
	fmt.Println(v)

	// Output:
	// <nil>
	// John Doe
}

func ExampleYaml_Bytes() {
	yml, _ := NewYaml([]byte("abcd: John Doe"))
	v, err := yml.Get("abcd").Bytes()

	fmt.Println(err)
	fmt.Println(string(v))

	// Output:
	// <nil>
	// John Doe
}

func ExampleYaml_Map() {
	data := `abcd:
  user1: John
  user2: Bob`

	yml, _ := NewYaml([]byte(data))
	v, err := yml.Get("abcd").Map()

	fmt.Println(err)
	fmt.Println(v["user1"])
	fmt.Println(v["user2"])

	// Output:
	// <nil>
	// John
	// Bob
}

func ExampleYaml_Array() {
	data := `abcd:
  - 123
  - Bob`

	yml, _ := NewYaml([]byte(data))
	v, err := yml.Get("abcd").Array()

	fmt.Println(err)
	fmt.Println(v[0])
	fmt.Println(v[1])

	// Output:
	// <nil>
	// 123
	// Bob
}

func ExampleYaml_StringArray() {
	data := `abcd:
  - John
  - Bob`

	yml, _ := NewYaml([]byte(data))
	v, err := yml.Get("abcd").StringArray()

	fmt.Println(err)
	fmt.Println(v[0])
	fmt.Println(v[1])

	// Output:
	// <nil>
	// John
	// Bob
}

func ExampleYaml_MustArray() {
	data := `abcd:
  - John
  - Bob`

	yml, _ := NewYaml([]byte(data))

	v := yml.Get("abcd").MustArray([]interface{}{"Mary", "Elen"})

	fmt.Println(v[1])

	v = yml.Get("unknown").MustArray([]interface{}{"Mary", "Elen"})

	fmt.Println(v[1])

	// Output:
	// Bob
	// Elen
}

func ExampleYaml_MustMap() {
	data := `abcd:
  user1: John
  user2: Bob`

	yml, _ := NewYaml([]byte(data))

	v := yml.Get("abcd").MustMap(map[interface{}]interface{}{"user1": "Mary"})

	fmt.Println(v["user1"])

	v = yml.Get("unknown").MustMap(map[interface{}]interface{}{"user1": "Mary"})

	fmt.Println(v["user1"])

	// Output:
	// John
	// Mary
}

func ExampleYaml_MustString() {
	yml, _ := NewYaml([]byte("abcd: John Doe"))

	v := yml.Get("abcd").MustString("Jane Doe")

	fmt.Println(v)

	v = yml.Get("unknown").MustString("Jane Doe")

	fmt.Println(v)

	// Output:
	// John Doe
	// Jane Doe
}

func ExampleYaml_MustStringArray() {
	data := `abcd:
  - John
  - Bob`

	yml, _ := NewYaml([]byte(data))

	v := yml.Get("abcd").MustStringArray([]string{"Mary", "Elen"})

	fmt.Println(v[0])

	v = yml.Get("unknown").MustStringArray([]string{"Mary", "Elen"})

	fmt.Println(v[0])

	// Output:
	// John
	// Mary
}

func ExampleYaml_MustInt() {
	yml, _ := NewYaml([]byte("abcd: 123"))

	v := yml.Get("abcd").MustInt(456)

	fmt.Println(v)

	v = yml.Get("unknown").MustInt(456)

	fmt.Println(v)

	// Output:
	// 123
	// 456
}

func ExampleYaml_MustFloat() {
	yml, _ := NewYaml([]byte("abcd: 1.23"))

	v := yml.Get("abcd").MustFloat(5.67)

	fmt.Println(v)

	v = yml.Get("unknown").MustFloat(5.67)

	fmt.Println(v)

	// Output:
	// 1.23
	// 5.67
}

func ExampleYaml_MustBool() {
	yml, _ := NewYaml([]byte("abcd: true"))

	v := yml.Get("abcd").MustBool(true)

	fmt.Println(v)

	v = yml.Get("unknown").MustBool(true)

	fmt.Println(v)

	// Output:
	// true
	// true
}

func ExampleYaml_Get() {
	data := `
test1:
    test2:
        test3: John Doe`

	yml, _ := NewYaml([]byte(data))

	fmt.Println(yml.Get("test1").Get("test2").Get("test3").MustString(""))

	// Output:
	// John Doe
}

func ExampleYaml_GetPath() {
	data := `
test1:
    test2:
        test3: John Doe`

	yml, _ := NewYaml([]byte(data))

	fmt.Println(yml.GetPath("test1", "test2", "test3").MustString(""))

	// Output:
	// John Doe
}

func ExampleYaml_GetByIndex() {
	data := `abcd:
  - John
  - Bob`

	yml, _ := NewYaml([]byte(data))

	fmt.Println(yml.Get("abcd").GetByIndex(1).MustString(""))

	// Output:
	// Bob
}

func ExampleYaml_IsExist() {
	yml, _ := NewYaml([]byte("abcd: 123"))

	fmt.Println(yml.IsExist("abcd"))
	fmt.Println(yml.IsExist("unknown"))

	// Output:
	// true
	// false
}
func ExampleYaml_IsPathExist() {
	data := `
test1:
    test2:
        test3: John Doe`

	yml, _ := NewYaml([]byte(data))

	fmt.Println(yml.IsPathExist("test1", "test2", "test3"))
	fmt.Println(yml.IsPathExist("test1", "test2", "test4"))

	// Output:
	// true
	// false
}

func ExampleYaml_IsIndexExist() {
	data := `abcd:
  - John
  - Bob`

	yml, _ := NewYaml([]byte(data))

	fmt.Println(yml.Get("abcd").IsIndexExist(1))
	fmt.Println(yml.Get("abcd").IsIndexExist(2))

	// Output:
	// true
	// false
}

func ExampleYaml_GetMapKeys() {
	data := `abcd:
  user1: John
  user2: Bob`

	yml, _ := NewYaml([]byte(data))
	keys, _ := yml.Get("abcd").GetMapKeys()

	sort.Strings(keys)

	fmt.Println(keys)

	// Output:
	// [user1 user2]
}
