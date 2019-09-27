// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// publishEvent
func doSomeWork() {

	threads := os.Getenv("EVENTS_THREADS")
	if threads == "" {
		threads = "1"
	}
	timeout := os.Getenv("EVENTS_TIMEOUT")
	if timeout == "" {
		timeout = "2"
	}

	cmd := exec.Command("stress", "--cpu", threads, "-v", "--timeout", timeout)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("events: exec error ", err)
	}
}

func handleDefaultRequest(w http.ResponseWriter, r *http.Request) {

	remote := r.RemoteAddr
	if r.Header.Get("X-FORWARDED-FOR") != "" {
		remote = r.Header.Get("X-FORWARDED-FOR")
	}

	log.Printf("events: handling request %s from %s", r.URL.Path, remote)
	doSomeWork()
	log.Printf("events: done with request %s from %s", r.URL.Path, remote)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("events: Hostname error ", err)
	}
	fmt.Fprintf(w, "Hello %s from %s\n", remote, hostname)
}

func main() {

	eventsPort := os.Getenv("events_PORT")
	if eventsPort == "" {
		eventsPort = "80"
	}

	http.HandleFunc("/", handleDefaultRequest)
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("events: started service listening on port: %s", eventsPort)
	err := http.ListenAndServe(":"+eventsPort, nil)
	if err != nil {
		log.Fatal("events: ListenAndServer error ", err)
	}
}
