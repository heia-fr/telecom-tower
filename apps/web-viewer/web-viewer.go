package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/nats-io/nats"
	"html/template"
	"log"
	"net/http"
)

const (
	subject = "tower.frames"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Bitmap [][]string

func handler(w http.ResponseWriter, r *http.Request) {

	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	frameChannel := make(chan ledmatrix.Stripe)
	c.BindRecvChan(subject, frameChannel)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// readloop
	go func(c *websocket.Conn) {
		for {
			if _, _, err := c.NextReader(); err != nil {
				c.Close()
				break
			}
		}
	}(conn)

	for {
		frame := <-frameChannel
		if err := conn.WriteJSON(frame); err != nil {
			return
		}
	}
}

func Tower(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/main.html"))
	tmpl.Execute(w, "T")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tower", Tower)
	r.HandleFunc("/stream", handler)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", r)
	http.ListenAndServe(":8100", nil)
}
