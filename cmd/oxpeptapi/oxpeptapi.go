package main

// The API connecting Oxide to PeptideO
// Versioned the API endpt as v2.0.0 to match the peer functions of HLF v2.0
// TODO: Add endpoints associated with chaincode versions, beginning at v1.0.0

import (
	"flag"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	// Add the fabric SDK, but don't reference yet
	_ "github.com/hyperledger/fabric-sdk-go"
)

var (
	port = flag.String("port", "8080", "port")
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hit top-level route entry point!\n")
}

func apiv2RootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hit apiv2 root!\n")
}

func peerv2RootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hit peerv2 root!\n")
}

func lifecyclev2RootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hit lifecyclev2 root!\n")
}

func chaincodev2RootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hit chaincodev2 root!\n")
}

func main() {
	flag.Parse()

	var router = mux.NewRouter()
	var apiv2 = router.PathPrefix("/v2").Subrouter()
	var peerv2 = apiv2.PathPrefix("/peer").Subrouter()
	// `peer lifecycle` subcommands
	//package|install|queryinstalled|getinstalledpackage|approveformyorg|checkcommitreadiness|commit|querycommitted
	var lifecyclev2 = peerv2.PathPrefix("/lifecycle").Subrouter()
	// `peer chaincode` subcommands
	//install|instantiate|invoke|package|query|signpackage|upgrade|list
	var chaincodev2 = peerv2.PathPrefix("/chaincode").Subrouter()

	router.HandleFunc("/", IndexHandler)
	apiv2.HandleFunc("/", apiv2RootHandler)
	peerv2.HandleFunc("/", peerv2RootHandler)
	lifecyclev2.HandleFunc("/", lifecyclev2RootHandler)
	chaincodev2.HandleFunc("/", chaincodev2RootHandler)

	lifecyclev2.HandleFunc("/package", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit package!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit install!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/queryinstalled", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit queryinstalled!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/getinstalledpackage", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit getinstalledpackage!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/approveformyorg", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit approveformyorg!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/checkcommitreadiness", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit checkcommitreadiness!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/commit", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit commit!\n")
		w.WriteHeader(http.StatusOK)
	})

	lifecyclev2.HandleFunc("/querycommitted", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit querycommitted!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit install!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/instantiate", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit instantiate!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/invoke", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit invoke!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/package", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit package!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit query!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/signpackage", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit signpackage!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/upgrade", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit upgrade!\n")
		w.WriteHeader(http.StatusOK)
	})

	chaincodev2.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hit list!\n")
		w.WriteHeader(http.StatusOK)
	})

	apiv2.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	peerv2.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	lifecyclev2.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	chaincodev2.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	http.ListenAndServe(":"+*port, router)
}
