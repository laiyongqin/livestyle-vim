package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	c := make(chan string, 256)
	name, path := "", ""
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if r.Method == "POST" {
			b, err := ioutil.ReadAll(r.Body)
			r.Body.Close()
			if err != nil {
				fmt.Fprint(w, "NG")
				return
			}
			c <- string(b)
		}
		fmt.Fprint(w, "OK")
	})
	http.HandleFunc("/shutdown", func (w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	})
	http.HandleFunc("/vim", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		name = r.FormValue("name")
		path = r.FormValue("path")
		fmt.Fprint(w, "OK")
	})
	http.Handle("/browser", websocket.Handler(func(ws *websocket.Conn) {
		go func() {
			for {
				var r string
				err := websocket.Message.Receive(ws, &r)
				if err != nil {
					break
				} else if name != "" && path != "" {
					cmd := exec.Command(path, "--servername", name, "--remote-expr", fmt.Sprintf("livestyle#reply(%q)", r))
					cmd.Run()
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
