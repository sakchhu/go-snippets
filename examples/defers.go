// Usage of defer snippets

package examples

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// 1. destroy/close/free owned resources (resources can be connections, files etc.) on function return
// defer when used like this closely matches RAII philosophy if you've used Rust/C++
func properFetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close() // this guarantees that this resource is freed after we return from this function.

	savePath := path.Base(resp.Request.URL.Path)
	if savePath == "/" {
		savePath = "index.html"
	}

	saveFile, err := os.Create(savePath)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(saveFile, resp.Body)
	if closeErr := saveFile.Close(); closeErr == nil {
		err = closeErr
	}

	return savePath, n, err
}

// 2. inspect/modify function results if necessary
func verify(id string) (isAdmin bool) {
	// simple example of some json data being processed

	type Role uint8
	const (
		adminRole     = Role(0)
		moderatorRole = Role(1)
	)
	users := map[string]Role{"John": adminRole, "Tom": moderatorRole, "Farah": adminRole}

	fmt.Printf("userRoles (0 = admin, 1 = moderator): %+v\n", users)

	defer func() {
		isAdmin = false                              // just don't grant access to anyone.
		fmt.Printf("Verify(%s) = %v\n", id, isAdmin) // inspect
	}()

	isAdmin = users[id] == adminRole
	return
}

func DeferVerify() {
	log.SetPrefix("verify: ")

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("usage: verify <username>")
	}

	verify(args[0])
}

func DeferFetch() {
	log.SetPrefix("defer-fetch: ")

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("usage: defer-fetch <url>")
	}

	url := args[0]
	f, n, err := properFetch(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("%s saved to %q, size: %d", url, f, n)
}
