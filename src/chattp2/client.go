package chattp2

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

// Client represents a chattp2 client
type Client struct {
	sender     string
	receiver   string
	httpClient http.Client
}

// NewClient returns new Client
func NewClient(sender, receiver string) (*Client, error) {
	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //@TODO: Configure with server keys
		},
	}

	return &Client{
		sender:     sender,
		receiver:   receiver,
		httpClient: client,
	}, nil
}

// Run chattp2 client.
func (c *Client) Run(ctx context.Context) {
	go c.connect()
	go c.send()

	select {}
}

type messageReader struct {
	r io.Reader
}

// Read input from p.
// Here you can apply transformations the input message.
func (mr messageReader) Read(p []byte) (n int, err error) {
	n, err = mr.r.Read(p)

	return
}

// Connect to chattp2 server. Wait to receive messages and print them to stdout
func (c *Client) connect() {
	req, err := http.NewRequest("GET", "https://localhost:4430/connect?user="+c.sender, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	for {
		_, _ = io.Copy(os.Stdout, messageReader{res.Body})
	}
}

// Send messages. Reads input from stdin and sent to chattp2.
func (c *Client) send() {
	fmt.Println("Write a message: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		c.sendMessage(s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// SendMessage actually sends the message to chattp2 through HTTP2.
func (c *Client) sendMessage(content string) {
	msg := Message{content, c.sender, c.receiver}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("POST", "https://localhost:4430/send", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
