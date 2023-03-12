package main

import (
	"fmt"
	"net"
	"net/http"
)

var Live = true
var Ready = true

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/-/health/liveness", liveness)
	http.HandleFunc("/-/health/readiness", readiness)
	http.HandleFunc("/changeState", changeState)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func root(w http.ResponseWriter, _ *http.Request) {
	ifaces, err := net.Interfaces()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not get network interfaces")
		return
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()

		if iface.Name == "lo" {
			continue
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Could not get IP addresses")
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			fmt.Fprintln(w, "Server's IP: ", ip)
		}
	}
}

func readiness(w http.ResponseWriter, _ *http.Request) {
	if Ready {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	fmt.Fprintln(w, Ready)
}

func liveness(w http.ResponseWriter, _ *http.Request) {
	if Live {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	fmt.Fprintln(w, Live)
}

func changeState(w http.ResponseWriter, r *http.Request) {
	if readiness := r.URL.Query().Get("readiness"); readiness != "" {
		Ready = !Ready
	}
	if liveness := r.URL.Query().Get("liveness"); liveness != "" {
		Live = !Live
	}

	fmt.Fprintln(w, "Ready: ", Ready, "; Live: ", Live)
}
