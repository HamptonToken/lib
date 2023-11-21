package store

import (
    // "context"
    "bytes"
    "fmt"
    "os"
    "strings"

    shell "github.com/ipfs/go-ipfs-api"
)

/* high level APIs be here */

/*
>>> import ipfsApi
>>> api = ipfsApi.Client('127.0.0.1', 5001)
>>> res = api.add('test.txt')
>>> res
{'Hash': 'QmWxS5aNTFEc9XbMX1ASvLET1zrqEaTssqt33rVZQCQb22', 'Name': 'test.txt'}
>>> api.cat(res['Hash'])
'fdsafkljdskafjaksdjf\n'
*/

/*
const node = await IPFS.create()

const data = 'Hello, <YOUR NAME HERE>'

// add your data to to IPFS - this can be a string, a Buffer,
// a stream of Buffers, etc
const results = node.add(data)

// we loop over the results because 'add' supports multiple 
// additions, but we only added one entry here so we only see
// one log line in the output
for await (const { cid } of results) {
  // CID (Content IDentifier) uniquely addresses the data
  // and can be used to get it again.
  console.log(cid.toString())
}
*/
