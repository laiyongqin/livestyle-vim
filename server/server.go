package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	c := make(chan string, 256)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if r.Method == "POST" {
			b, err := ioutil.ReadAll(r.Body)
			r.Body.Close()
			if err == nil {
				c <- string(b)
			}
		}
		fmt.Fprint(w, "OK")
	})
	http.Handle("/browser", websocket.Handler(func(ws *websocket.Conn) {
		go func() {
			for {
				var r string
				err := websocket.Message.Receive(ws, &r)
				if err != nil {
					log.Println(err)
				} else {
					log.Println(r)
				}
			}
		}()
		for {
			s := <-c
			log.Println(s)
			err := websocket.Message.Send(ws, s)
			if err != nil {
				break
			}
		}
	}))
	err := http.ListenAndServe(":54000", nil)
	if err != nil {
		log.Fatal(err)
	}
}