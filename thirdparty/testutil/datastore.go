package testutil

import (
	ds2 "github.com/ipfs/go-ipfs/thirdparty/datastore2"
	"gx/ipfs/QmcLrdzNZ2PLEWY9MQaZyNJNSXWaqADJqzPTwWGJRrMqT1/go-datastore"
	syncds "gx/ipfs/QmcLrdzNZ2PLEWY9MQaZyNJNSXWaqADJqzPTwWGJRrMqT1/go-datastore/sync"
)

func ThreadSafeCloserMapDatastore() ds2.ThreadSafeDatastoreCloser {
	return ds2.CloserWrap(syncds.MutexWrap(datastore.NewMapDatastore()))
}
