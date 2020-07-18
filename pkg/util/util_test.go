package util

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTimeJson(t *testing.T) {
	/*	now := Time(time.Now())
		t.Log(now)
		src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
		p := new(Person)
		err := json.Unmarshal([]byte(src), p)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(p)
		t.Log(time.Time(p.Birthday))
		js, _ := json.Marshal(p)
		t.Log(string(js))*/

	dt, err := NewDatetime(time.Now())
	if err != nil {
		fmt.Println("dt", err)
		return
	}
	per := Person{
		Id:       1,
		Name:     "22",
		Birthday: *dt,
	}

	by, _ := json.Marshal(per)
	per2 := Person{}
	json.Unmarshal(by, &per2)

	fmt.Println("-----mar", string(by))
	fmt.Println("-----unmar", per2.Birthday)

}

func TestSha256(t *testing.T) {
	//fmt.Println(Sha256("123456"))

	var a = []int{1, 2, 3, 45}
	for i := range a {
		fmt.Println(i, ":", a[i])
	}

}
