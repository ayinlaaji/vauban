//Candidate Name: Ayinla Abdulsalam
//Email: ayinlaabdulsalam@gmail.com

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var url string = "https://qcvault.herokuapp.com/unlock_safe"

type xType struct {
	msg  string
	code string
}

func main() {
	maxNo := 999

	ch := make(chan xType)

	for i := 000; i <= maxNo; i++ {
		go work(i, ch)
	}

	v := <-ch
	fmt.Println(v.msg, v.code)
	close(ch)

}

func work(i int, ch chan xType) {

	no := fmt.Sprintf("%03d", i)
	sno := strings.Split(no, "")
	jsonStrB := fmt.Sprintf(`{"first":%s,"second":%s,"third":%s}`, sno[0], sno[1], sno[2])

	jsonStr := []byte(jsonStrB)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	msg := string(body)

	if msg != "Wrong code" {
		ch <- xType{msg, no}
	}
}
