package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"sync"
)

type portArray []int

func (i *portArray) String() string {
	var sb strings.Builder
	for n, v := range *i {
		if n > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf(":%d", v))
	}
	return sb.String()
}

func (i *portArray) Set(value string) error {
	if s, err := strconv.Atoi(value); err == nil {
		*i = append(*i, s)
	} else {
		log.Printf("Error adding parameter: %s: %s\n", value, err)
		return err
	}
	return nil
}

func main() {
	var ports portArray
	var wg sync.WaitGroup
	log.Println("Program started:", os.Args)
	flag.Var(&ports, "port", "List of ports")
	flag.Parse()

	log.Printf("Starting %d servers...\n", len(ports))
	for _, p := range ports {
		listenStr := fmt.Sprintf(":%d", p)
		wg.Add(1)
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				handleRequest(w, r, listenStr)
			})
			srv := http.Server{Addr: listenStr, Handler: mux}
			log.Printf("Listening On %s", listenStr)
			srv.ListenAndServe()
		}()
	}
	log.Println()
	wg.Wait()
	log.Printf("Terminating.")
}

func handleRequest(w http.ResponseWriter, r *http.Request, listenstring string) {
	// w.WriteHeader(status)
	reqdump, _ := httputil.DumpRequest(r, true)
	w.Write([]byte(fmt.Sprintf("Received on: %s\n", listenstring)))
	w.Write(reqdump)
	//#fmt.Fprint(w, r.String())
	log.Printf("%s: %s: %s", r.Method, r.URL, r.UserAgent())
}
