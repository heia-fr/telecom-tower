// Copyright 2015 Jacques Supcik, HEIA-FR
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Test programm for Telecom Tower


package main

import (
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/lyrics"
	"log"
        "github.com/ant0ine/go-json-rest/rest"
        "net/http"
        "net/url"
        "os"
	"os/signal"
)

func main() {
	log.Println("Booting tower...")
        matrix := ledmatrix.NewMatrix()
        api := rest.NewApi()
        api.Use(rest.DefaultDevStack...)
        c := make(chan os.Signal, 1)
	signalChannel := make(chan os.Signal, 1)
	lyrics.Init(128)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		s := <-c
		log.Println("Shutting down tower... please wait for message finish")
                log.Println("Shutting down tower now")
                lyrics.Shutdown()
                os.Exit(0)
		signalChannel <- s
	}()
        router, err := rest.MakeRouter(
            rest.Get("/lyrics/#text", func(w rest.ResponseWriter, req *rest.Request) {
                text, err := url.QueryUnescape(req.PathParam("text"))
                log.Print(text)
                if err != nil {
                    log.Fatal(err)
                }
 
	        
  	        writer := ledmatrix.NewWriter(matrix)


		bg := ledmatrix.RGB(0, 0, 0)
	        writer.Spacer(64,bg)
		writer.WriteText(
			text,
			bitmapfont.F68,
			ledmatrix.RGB(150, 200, 255),
			bg,
		)
                
    		writer.Spacer(64, bg)
                writer.SetPos(writer.Pos()-127)
                log.Println(writer.Pos())
		log.Println("Roll!")

		lyrics.Roll(writer)
                log.Println("Roll ended")
	
	}),
        )
        if err != nil {
            log.Fatal(err)
        }
        api.SetApp(router)
        log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

}
