package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Msg struct {
	Int    int64       `json:"int"`
	String string      `json:"string"`
	Bool   bool        `json:"bool"`
	Json   interface{} `json:"json"`
}
type Data map[string]bool

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		var data Msg

		var msgTypes Data

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println("body", body)
		if err = json.Unmarshal(body, &data); err != nil {
			w.Write([]byte("unmarshal" + err.Error()))
			return
		}

		fmt.Println("Int", data.Int)
		fmt.Println("String", data.String)

		fmt.Println("Bool", data.Bool)
		fmt.Println("Interface", data.Json)

		val := reflect.ValueOf(data.Json)
		fmt.Println(val.MapKeys())
		for _, e := range val.MapKeys() {
			fmt.Println("key", e, "value", val.MapIndex(e))
			fmt.Println("key", e.String(), val.MapIndex(e).Interface().(bool))
			msgTypes[e.String()] = val.MapIndex(e).Interface().(bool)
		}

		fmt.Println("msgTypeInterface", msgTypes)
		fmt.Println("marshal", msgTypes["marshal"])
		fmt.Println("unmarshal", msgTypes["unmarshal"])
		fmt.Println("unmhal", msgTypes["unmhal"])

		// for key, value := range msgs {
		// 	fmt.Println("key", key, "value", value)
		// }
		// va
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
