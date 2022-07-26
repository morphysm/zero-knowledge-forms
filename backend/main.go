package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/famed-submission-form/backend/report"
	logging "github.com/ipfs/go-log"
	"github.com/status-im/go-waku/waku/v2/node"
	"github.com/status-im/go-waku/waku/v2/protocol"
	"github.com/status-im/go-waku/waku/v2/protocol/filter"
	"github.com/status-im/go-waku/waku/v2/protocol/relay"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	log          = logging.Logger("basic2")
	contentTopic = protocol.NewContentTopic("relay-reactjs-chat", 1, "chat", "proto")
)

func main() {
	lvl, err := logging.LevelFromString("info")
	if err != nil {
		panic(err)
	}
	logging.SetAllLoggers(lvl)

	privateKey := os.Getenv("ETH_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("ETH_PRIVATE_KEY is not set")
	}

	hostAddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprint("0.0.0.0:0"))

	key, err := randomHex(32)
	if err != nil {
		log.Error("Could not generate random key")
		return
	}
	prvKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Error(err)
		return
	}

	ctx := context.Background()

	wakuNode, err := node.New(ctx,
		node.WithPrivateKey(prvKey),
		node.WithHostAddress(hostAddr),
		node.WithWakuRelayAndMinPeers(1),
		node.WithWakuFilter(false),
	)
	if err != nil {
		log.Error(err)
		return
	}

	if err := wakuNode.Start(); err != nil {
		log.Error(err)
		return
	}

	// TODO move to json
	err = wakuNode.DialPeer(ctx, "/ip4/34.121.100.108/tcp/30303/p2p/16Uiu2HAmVkKntsECaYfefR1V2yCR79CegLATuTPE6B9TxgxBiiiA")
	if err != nil {
		log.Error("Could not connect to peer: " + err.Error())
		return
	}

	go readLoop(ctx, wakuNode, privateKey)

	// Wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("\n\n\nReceived signal, shutting down...")

	// Shut the node down
	wakuNode.Stop()
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func readLoop(ctx context.Context, wakuNode *node.WakuNode, privateKey string) {
	cf := filter.ContentFilter{
		Topic:         relay.DefaultWakuTopic,
		ContentTopics: []string{contentTopic.String()},
	}
	_, sub, err := wakuNode.Filter().Subscribe(ctx, cf)
	if err != nil {
		log.Error("Could not subscribe:  ", err)
		return
	}

	for value := range sub.Chan {
		msg := value.Message()
		payload, err := report.Decode(msg, privateKey)
		if err != nil {
			log.Error("Could not subscribe:  ", err)
			return
		}

		fmt.Println(payload)
	}
}
