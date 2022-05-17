package store

import (
    // "context"
    "bytes"
    "fmt"
    "os"
    "strings"

    shell "github.com/ipfs/go-ipfs-api"
)

/* high level APIs here */

/*
>>> import ipfsApi
>>> api = ipfsApi.Client('127.0.0.1', 5001)
>>> res = api.add('test.txt')
>>> res
{'Hash': 'QmWxS5aNTFEc9XbMX1ASvLET1zrqEaTssqt33rVZQCQb22', 'Name': 'test.txt'}
>>> api.cat(res['Hash'])
'fdsafkljdskafjaksdjf\n'
*/
