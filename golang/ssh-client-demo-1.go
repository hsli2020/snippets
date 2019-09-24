
// Golang SSH Client: Multiple Commands, Crypto & Goexpect Examples
// Golang SSH client examples

// Single Command Example

package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {

	host := ""
	port := "22"
	user := ""
	pass := ""
	cmd  := "ps"

	// get host public key
	hostKey := getHostKey(host)

	// ssh client config
	config := &amp;ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		// allow any host key to be used (non-prod)
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		// verify host public key
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		// optional host key algo list
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
		// optional tcp connect timeout
		Timeout:         5 * time.Second,
	}

	// connect
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// start session
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// setup standard out and error
	// uses writer interface
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// run single command
	err = sess.Run(cmd)
	if err != nil {
		log.Fatal(err)
	}

}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}

package main

// Multiple Command Example

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	// Uncomment to store output in variable
	//"bytes"
)

func main() {

	username := ""
	password := ""
	hostname := ""
	port := ""

	// SSH client config
	config := &amp;ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Non-production only
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to host
	client, err := ssh.Dial("tcp", hostname+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer sess.Close()

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Uncomment to store output in variable
	//var b bytes.Buffer
	//sess.Stdout = &amp;b
	//sess.Stderr = &amp;b

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		log.Fatal(err)
	}

	// send the commands
	commands := []string{
		"pwd",
		"whoami",
		"echo 'bye'",
		"exit",
	}
	for _, cmd := range commands {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Wait for sess to finish
	err = sess.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// Uncomment to store in variable
	//fmt.Println(b.String())

}

/*
Golang SSH Client Supported Ciphers

The Golang SSH Client specifies the default preference for ciphers (see preferredCiphers list):

    aes128-gcm@openssh.com
    chacha20Poly1305ID
    aes128-ctr
    aes192-ctr
    aes256-ctr”

The Golang SSH Client lists supported ciphers that are not recommend(see supportedCiphers list):

    aes128-ctr
    aes192-ctr
    aes256-ctr
    aes128-gcm@openssh.com
    chacha20Poly1305ID
    arcfour256
    arcfour128
    arcfour
    aes128cbcID
    tripledescbcID

Take note in cipherModes :

    CBC mode is insecure and so is not included in the default config.(See http://www.isg.rhul.ac.uk/~kp/SandPfinal.pdf). If absolutely needed, it’s possible to specify a custom Config to enable it. You should expect that an active attacker can recover plaintext if you do.
    3des-cbc is insecure and is not included in the default config.

A custom Config example to additionally allow cbc may look like this below:
package main
...
    var sshconfig ssh.Config
    sshconfig.SetDefaults()
    sshconfig.Ciphers = append(sshconfig.Ciphers, "aes128-cbc", "aes192-cbc", "aes256-cbc", "3des-cbc", "des-cbc")
    config := &ssh.ClientConfig{...}
...


Likewise you could explicitly list your custom ciphers in ClientConfig:
package main
...
	config := &ssh.ClientConfig{
		Config: ssh.Config {
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-cbc", "aes192-cbc", "aes256-cbc", "3des-cbc", "des-cbc"},
		},
		//Auth:
		//Timeout:
	}
...
*/

//Golang Expect Example

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/google/goexpect"
	"github.com/google/goterm/term"
	"golang.org/x/crypto/ssh"
)

const (
	timeout = 10 * time.Minute
)

func main() {

	host := ""
	port := "22"
	user := ""
	pass := ""
	cmd1 := "pwd"
	cmd2 := "whoami"
	promptRE := regexp.MustCompile("\\$")

	// get host public key
	hostKey := getHostKey(host)

	sshClt, err := ssh.Dial("tcp", host+":"+port, &amp;ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
		// allow any host key to be used (non-prod)
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		// verify host public key
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	})
	if err != nil {
		glog.Exitf("ssh.Dial(%q) failed: %v", host, err)
	}
	defer sshClt.Close()

	e, _, err := expect.SpawnSSH(sshClt, timeout)
	if err != nil {
		glog.Exit(err)
	}
	defer e.Close()

	e.Expect(promptRE, timeout)
	e.Send(cmd1 + "\n")
	result1, _, _ := e.Expect(promptRE, timeout)
	e.Send(cmd2 + "\n")
	result2, _, _ := e.Expect(promptRE, timeout)
	e.Send("exit\n")

	fmt.Println(term.Greenf("Done!\n"))
	fmt.Printf("%s: result:\n %s\n\n", cmd1, result1)
	fmt.Printf("%s: result:\n %s\n\n", cmd2, result2)

}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to pull key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}
