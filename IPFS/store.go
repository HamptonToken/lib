/* store api*/
package store

import (
    // "context"
    "bytes"
    "fmt"
    "os"
    "strings"

    shell "github.com/ipfs/go-ipfs-api"
)

type Directory struct {
	inode

	// Internal cache with added entries to the directory, its cotents
	// are synched with the underlying `unixfsDir` node in `sync()`.
	entriesCache map[string]FSNode

	lock sync.Mutex
	// TODO: What content is being protected here exactly? The entire directory?

	ctx context.Context

	// UnixFS directory implementation used for creating,
	// reading and editing directories.
	unixfsDir uio.Directory

	modTime time.Time
}

func store() {
    //
    // Connect to your local IPFS deamon running in the background.
    //
    // Where your local node is running on localhost:5001
    sh := shell.NewShell("localhost:5001")

    //
    // Add the file to IPFS
    //

    cid, err := sh.Add(strings.NewReader("hello world!"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
    }
    fmt.Printf("added %s\n", cid)

    //
    // Get the data from IPFS and save the contents to a file.
    //
    out := fmt.Sprintf("%s.txt", cid)
    err = sh.Get(cid, out)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
    }

    //
    // Get the data from IPFS and output the contents into `string` format
    // and output into the terminal console.
    //
    data, err := sh.Cat(cid)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
    }

    // ...so we convert it to a string by passing it through
    // a buffer first. A 'costly' but useful process.
    // https://golangcode.com/convert-io-readcloser-to-a-string/
    buf := new(bytes.Buffer)
    buf.ReadFrom(data)
    newStr := buf.String()
    fmt.Printf("data %s", newStr)
}
