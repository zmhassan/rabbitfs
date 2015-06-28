package RaftServer

import (
	"fmt"
	"testing"
	"time"

	"github.com/lilwulin/rabbitfs/helper"
)

//TODO: refactor these test

var testPeers = []string{
	"127.0.0.1:8787",
	"127.0.0.1:8788",
	"127.0.0.1:8789",
}

func TestMultiRaftkv(t *testing.T) {
	defer helper.RemoveDirs(
		"./raft1", "./raft2", "./raft3", "./raft4",
	)
	fmt.Println("testing multi raft")

	rs1, err := NewRaftServer(
		nil,
		testPeers,
		"./raft1",
		testPeers[0],
		"/raft",
		500*time.Millisecond,
		0,
	)
	if err != nil {
		t.Error(err)
	}
	go rs1.ListenAndServe()
	time.Sleep(1000 * time.Millisecond)

	rs2, err := NewRaftServer(
		nil,
		testPeers,
		"./raft2",
		testPeers[1],
		"/raft",
		500*time.Millisecond,
		0,
	)
	if err != nil {
		t.Error(err)
	}
	go rs2.ListenAndServe()
	time.Sleep(300 * time.Millisecond)

	rs3, err := NewRaftServer(
		nil,
		testPeers,
		"./raft3",
		testPeers[2],
		"/raft",
		500*time.Millisecond,
		0,
	)
	if err != nil {
		t.Error(err)
	}
	go rs3.ListenAndServe()
	time.Sleep(1 * time.Second)
}

// func BenchmarkLevelDBKV(b *testing.B) {
// 	defer helper.RemoveDirs(
// 		"./raft1", "./leveldb1",
// 		"./raft2", "./leveldb2",
// 		"./raft3", "./leveldb3",
// 		"./raft4", "./leveldb4",
// 	)
// 	os.Mkdir("./raft1", 0700)
// 	os.Mkdir("./raft2", 0700)
// 	os.Mkdir("./raft3", 0700)
// 	fmt.Println("testing multi raftkv")
// 	// creating new leveldb kvstore
// 	kv1, _ := NewLevelDB("./leveldb1")
// 	kv2, _ := NewLevelDB("./leveldb2")
// 	kv3, _ := NewLevelDB("./leveldb3")

// 	rkv1, _ := NewRaftkv(
// 		testPeers,
// 		kv1,
// 		"./raft1",
// 		"http://127.0.0.1",
// 		8787,
// 		"/raft",
// 		500*time.Millisecond,
// 		0,
// 	)
// 	go rkv1.ListenAndServe()
// 	time.Sleep(300 * time.Millisecond)

// 	rkv2, _ := NewRaftkv(
// 		testPeers,
// 		kv2,
// 		"./raft2",
// 		"http://127.0.0.1",
// 		8788,
// 		"/raft",
// 		500*time.Millisecond,
// 		0,
// 	)
// 	go rkv2.ListenAndServe()
// 	time.Sleep(150 * time.Millisecond)

// 	rkv3, _ := NewRaftkv(
// 		testPeers,
// 		kv3,
// 		"./raft3",
// 		"http://127.0.0.1",
// 		8789,
// 		"/raft",
// 		500*time.Millisecond,
// 		0,
// 	)
// 	go rkv3.ListenAndServe()
// 	time.Sleep(200 * time.Millisecond)

// 	ops := 10000
// 	ben := bench.Start("RAFTKV-PUT")
// 	for i := 0; i < ops; i++ {
// 		_ = rkv2.Put([]byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", rand.Uint32())))
// 	}
// 	ben.End(ops)

// 	ben = bench.Start("RAFTKV-GET")
// 	for i := 0; i < ops; i++ {
// 		_, _ = rkv2.Get([]byte(fmt.Sprintf("%d", i)))
// 	}
// 	ben.End(ops)
// }

// func BenchmarkMemKV(b *testing.B) {
// 	defer helper.RemoveDirs(
// 		"./raft1",
// 		"./raft2",
// 		"./raft3",
// 	)
// 	os.Mkdir("./raft1", 0700)
// 	os.Mkdir("./raft2", 0700)
// 	os.Mkdir("./raft3", 0700)
// 	fmt.Println("testing multi raftkv")
// 	// creating new leveldb kvstore
// 	kv1 := NewMem()
// 	kv2 := NewMem()
// 	kv3 := NewMem()

// 	rkv1, _ := NewRaftkv(
// 		testPeers,
// 		kv1,
// 		"./raft1",
// 		"http://127.0.0.1",
// 		8787,
// 		"/raft",
// 		100*time.Millisecond,
// 		0,
// 	)
// 	go rkv1.ListenAndServe()
// 	time.Sleep(300 * time.Millisecond)

// 	rkv2, _ := NewRaftkv(
// 		testPeers,
// 		kv2,
// 		"./raft2",
// 		"http://127.0.0.1",
// 		8788,
// 		"/raft",
// 		100*time.Millisecond,
// 		0,
// 	)
// 	go rkv2.ListenAndServe()
// 	time.Sleep(150 * time.Millisecond)

// 	rkv3, _ := NewRaftkv(
// 		testPeers,
// 		kv3,
// 		"./raft3",
// 		"http://127.0.0.1",
// 		8789,
// 		"/raft",
// 		100*time.Millisecond,
// 		0,
// 	)
// 	go rkv3.ListenAndServe()
// 	time.Sleep(200 * time.Millisecond)

// 	ops := 300
// 	ben := bench.Start("RAFTKV-PUT")
// 	for i := 0; i < ops; i++ {
// 		err := rkv1.Put([]byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", rand.Uint32())))
// 		if err != nil {
// 			b.Error(err)
// 			return
// 		}
// 	}
// 	ben.End(ops)

// 	ben = bench.Start("RAFTKV-GET")
// 	for i := 0; i < ops; i++ {
// 		_, _ = rkv2.Get([]byte(fmt.Sprintf("%d", i)))
// 	}
// 	ben.End(ops)
// }