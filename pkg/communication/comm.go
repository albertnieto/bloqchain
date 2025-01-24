package communication

import (
    "crypto/tls"
    "log"
    "net"
)

// CommChannel is the interface for secure communication
type CommChannel interface {
    StartServer() error
    DialPeer(address string) (net.Conn, error)
}

// defaultTLSChannel is a basic implementation using TLS
type defaultTLSChannel struct {
    certFile string
    keyFile  string
    listener net.Listener
}

// NewTLSChannel initializes a TLS-based communication channel
func NewTLSChannel(certFile, keyFile string) (CommChannel, error) {
    return &defaultTLSChannel{
        certFile: certFile,
        keyFile:  keyFile,
    }, nil
}

// StartServer starts a TLS server that accepts incoming connections
func (c *defaultTLSChannel) StartServer() error {
    cert, err := tls.LoadX509KeyPair(c.certFile, c.keyFile)
    if err != nil {
        return err
    }

    config := &tls.Config{Certificates: []tls.Certificate{cert}}
    ln, err := tls.Listen("tcp", ":4433", config) // Example: listens on port 4433
    if err != nil {
        return err
    }
    c.listener = ln

    go func() {
        for {
            conn, err := ln.Accept()
            if err != nil {
                log.Println("Error accepting connection:", err)
                continue
            }
            go handleConnection(conn)
        }
    }()
    log.Println("TLS server started on port 4433")
    return nil
}

// DialPeer creates an outgoing TLS connection to a peer
func (c *defaultTLSChannel) DialPeer(address string) (net.Conn, error) {
    cert, err := tls.LoadX509KeyPair(c.certFile, c.keyFile)
    if err != nil {
        return nil, err
    }

    config := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
    return tls.Dial("tcp", address, config)
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    // In a real blockchain node, you'd handle P2P messages, block data, etc.
    log.Println("New TLS connection established:", conn.RemoteAddr())
}
