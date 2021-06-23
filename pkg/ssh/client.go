package ssh

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	Client *ssh.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewSsh(user, host, authType, p string, port int) (s *Ssh) {

	config := &ssh.ClientConfig{
		User:            user,
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if authType == "password" {
		config.Auth = []ssh.AuthMethod{
			ssh.Password(p),
		}
	} else {

	}

	hostport := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", hostport, config)
	if err != nil {
		log.Fatal("connect ssh err: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	s = &Ssh{Client: client, Ctx: ctx, Cancel: cancel}
	return
}

func (s *Ssh) Close() {
	s.Client.Close()
}
