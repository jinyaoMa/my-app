package h3server_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	libp2phttp "github.com/libp2p/go-libp2p/p2p/http"
	"github.com/libp2p/go-libp2p/p2p/net/gostream"

	// "majinyao.cn/my-app/backend/pkg/h3server"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestH3Server(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	// s := h3server.New(http.DefaultServeMux)
	// err := s.Run(8080, 8443, "testdata/cert.pem", "testdata/key.pem")
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func TestH3Server2(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	return

	serverStreamHost, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/udp/0/quic-v1"))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := gostream.Listen(serverStreamHost, "/testiti-test")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	go func() {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			resp := fmt.Sprintf("Hi %s!", body)
			w.Write([]byte(resp))
		})
		server := &http.Server{}
		server.Serve(listener)
	}()

	// tr := &http.Transport{}
	// tr.RegisterProtocol("libp2p", NewTransport(clientHost, ProtocolOption("/testiti-test")))
	// client := &http.Client{Transport: tr}

	// buf := bytes.NewBufferString("Hector")
	// res, err := client.Post(fmt.Sprintf("libp2p://%s/hello", serverStreamHost.ID().String()), "text/plain", buf)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer res.Body.Close()
	// text, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if string(text) != "Hi Hector!" {
	// 	t.Errorf("expected Hi Hector! but got %s", text)
	// }

	// t.Log(string(text))
}

func TestH3Server3(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	return

	serverStreamHost, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/udp/0/quic-v1"))
	if err != nil {
		log.Fatal(err)
	}

	server := libp2phttp.Host{
		StreamHost: serverStreamHost,
	}

	// A server with a simple echo protocol
	server.SetHTTPHandler("/echo/1.0.0", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/octet-stream")
		io.Copy(w, r.Body)
	}))
	go server.Serve()
	defer server.Close()

	clientStreamHost, err := libp2p.New(libp2p.NoListenAddrs)
	if err != nil {
		log.Fatal(err)
	}

	client := libp2phttp.Host{StreamHost: clientStreamHost}

	// Make an HTTP request using the Go standard library, but over libp2p
	// streams. If the server were listening on an HTTP transport, this could
	// also make the request over the HTTP transport.
	httpClient, err := client.NamespacedClient("/echo/1.0.0", peer.AddrInfo{ID: server.PeerID(), Addrs: server.Addrs()})

	// Only need to Post to "/" because this client is namespaced to the "/echo/1.0.0" protocol.
	resp, err := httpClient.Post("/", "application/octet-stream", strings.NewReader("Hello HTTP"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	// Output: Hello HTTP
}
