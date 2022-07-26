package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/famed-submission-form/backend/report"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pblib "github.com/famed-submission-form/backend/pb"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	logging "github.com/ipfs/go-log"
	"github.com/status-im/go-waku/waku/v2/node"
	"github.com/status-im/go-waku/waku/v2/protocol"
	"github.com/status-im/go-waku/waku/v2/protocol/filter"
	"github.com/status-im/go-waku/waku/v2/protocol/pb"
	"github.com/status-im/go-waku/waku/v2/protocol/relay"
	"github.com/status-im/go-waku/waku/v2/utils"
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

	go writeLoop(ctx, wakuNode)
	go readLoop(ctx, wakuNode)

	// Wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("\n\n\nReceived signal, shutting down...")

	// shut the node down
	wakuNode.Stop()
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func write(ctx context.Context, wakuNode *node.WakuNode, msgContent string) {
	fmt.Println(wakuNode.PeerCount())

	msg, err := encode(msgContent)
	if err != nil {
		log.Error("Error encoding message: ", err)
	}

	_, err = wakuNode.Relay().Publish(ctx, msg)
	if err != nil {
		log.Error("Error sending a message: ", err)
	}
}

func writeLoop(ctx context.Context, wakuNode *node.WakuNode) {
	for {
		time.Sleep(2 * time.Second)
		write(ctx, wakuNode, "Hello world!")
		time.Sleep(28 * time.Second)
	}
}

func readLoop(ctx context.Context, wakuNode *node.WakuNode) {
	cf := filter.ContentFilter{
		Topic:         relay.DefaultWakuTopic,
		ContentTopics: []string{contentTopic.String()},
	}
	_, sub, err := wakuNode.Filter().Subscribe(ctx, cf)
	if err != nil {
		log.Error("Could not subscribe: ", err)
		return
	}

	for value := range sub.Chan {
		msg := value.Message()
		decode(msg)
	}
}

func decode(msg *pb.WakuMessage) {
	fmt.Println(msg)
	payload, err := node.DecodePayload(msg, &node.KeyInfo{Kind: node.Asymmetric})
	if err != nil {
		log.Error("Error decoding message: ", err)
		return
	}

	encodedReport, _ := report.Unmarshal(payload.Data)
	fmt.Println(encodedReport)

	// report, err := decodeReport(payload.Data)
	// if err != nil {
	// 	log.Error("Error unmarshalling", err)
	// 	return
	// }

	// log.Info("Received msg in msg, ", string(report))
}

func encode(msg string) (*pb.WakuMessage, error) {
	var version uint32 = 0
	var timestamp int64 = utils.GetUnixEpoch()

	reportBytes, err := encodeReport(msg)
	if err != nil {
		log.Error("Error encoding the report: ", err)
	}

	p := new(node.Payload)
	p.Data = reportBytes
	p.Key = &node.KeyInfo{Kind: node.None}

	payload, err := p.Encode(version)
	if err != nil {
		log.Error("Error encoding the payload: ", err)
		return nil, err
	}

	return &pb.WakuMessage{
		Payload:      payload,
		Version:      version,
		ContentTopic: contentTopic.String(),
		Timestamp:    timestamp,
	}, nil
}

func encodeReport(msg string) ([]byte, error) {
	report := &pblib.Report{
		Timestamp: uint64(time.Now().Unix()),
		Payload:   []byte(msg),
	}

	msgBytes, err := proto.Marshal(report)
	if err != nil {
		log.Error("Error encoding the report: ", err)
		return nil, err
	}

	return msgBytes, nil
}

func decodeReport(data []byte) (string, error) {
	report := &pblib.Report{}
	err := proto.Unmarshal(data, report)
	if err != nil {
		log.Error("Error unmarshalling", err)
		return "", err
	}

	return string(report.Payload), nil
}
