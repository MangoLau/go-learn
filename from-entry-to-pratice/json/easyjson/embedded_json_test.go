package easyjson

import (
	"encoding/json"
	"fmt"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}

}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}
}

func TestJsoniter(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

func TestGoJson(t *testing.T) {
	e := new(Employee)
	err := gojson.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := gojson.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}

}

func BenchmarkEmbeddedJson(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkEasyJson(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		e.UnmarshalJSON([]byte(jsonStr))
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkJsoniter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoJson(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := gojson.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := gojson.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkEmbeddedJsonEncode(b *testing.B) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkEasyJsonEncode(b *testing.B) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkJsoniterEncode(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoJsonEncode(b *testing.B) {
	e := new(Employee)
	err := gojson.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkEmbeddedJsonDecode(b *testing.B) {
	e := new(Employee)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkEasyJsonDecode(b *testing.B) {
	e := Employee{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.UnmarshalJSON([]byte(jsonStr))
	}
	b.StopTimer()
}

func BenchmarkJsoniterDeocde(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoJsonDecode(b *testing.B) {
	e := new(Employee)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := gojson.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
