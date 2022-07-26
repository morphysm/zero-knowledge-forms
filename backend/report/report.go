package report

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/golang/protobuf/proto"
	"github.com/status-im/go-waku/waku/v2/node"
	"github.com/status-im/go-waku/waku/v2/protocol/pb"
	"github.com/storyicon/sigverify"
	"golang.org/x/crypto/nacl/box"

	pblib "github.com/famed-submission-form/backend/pb"
)

func Decode(msg *pb.WakuMessage, privateKey string) (string, error) {
	payload, err := node.DecodePayload(msg, &node.KeyInfo{Kind: node.Asymmetric})
	if err != nil {
		log.Error("Error decoding waku message ", err)
		return "", err
	}

	encodedReport, err := UnmarshalEncodedReport(payload.Data)
	if err != nil {
		log.Error("Error unmarshalling encoded report ", err)
		return "", err
	}

	hexBytesReport, err := DecryptReport(encodedReport, privateKey)
	if err != nil {
		log.Error("Error decrypting report ", err)
		return "", err
	}

	decodedReport, err := UnmarshalDecryptedReportHex(hexBytesReport)
	if err != nil {
		log.Error("Error unmarshalling decrypted report ", err)
		return "", err
	}

	// TODO fix signature transmission
	cleanSig := string(decodedReport.Signature)[1 : len(decodedReport.Signature)-1]

	ok, err := VerifySig(string(decodedReport.EthAddress), cleanSig, string(decodedReport.Payload))
	if err != nil {
		log.Error("Error verifying signature ", err)
		return "", err
	}

	if !ok {
		log.Error("Signature invalid")
		return "", errors.New("signature invalid")
	}

	return string(decodedReport.Payload), nil
}

type EncodedReport struct {
	Version        string `json:"version"`
	Nonce          string `json:"nonce"`
	EphemPublicKey string `json:"ephemPublicKey"`
	Ciphertext     string `json:"ciphertext"`
}

func UnmarshalEncodedReport(data []byte) (EncodedReport, error) {
	encodedReport := EncodedReport{}
	err := json.Unmarshal(data, &encodedReport)
	if err != nil {
		return encodedReport, err
	}

	return encodedReport, nil
}

func UnmarshalDecryptedReportHex(data []byte) (*pblib.Report, error) {
	dst := make([]byte, hex.DecodedLen(len(data)))
	_, err := hex.Decode(dst, data)
	if err != nil {
		return &pblib.Report{}, err
	}

	decryptedReport := &pblib.Report{}
	err = proto.Unmarshal(dst, decryptedReport)
	if err != nil {
		return &pblib.Report{}, err
	}

	return decryptedReport, nil
}

func DecryptReport(encodedReport EncodedReport, privateKey string) ([]byte, error) {
	secretKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	nonceBytes, err := base64.StdEncoding.DecodeString(encodedReport.Nonce)
	if err != nil {
		return nil, err
	}

	var nonce [24]byte
	copy(nonce[:], nonceBytes)

	cypherBytes, err := base64.StdEncoding.DecodeString(encodedReport.Ciphertext)
	if err != nil {
		return nil, err
	}

	ephemPublicKeyBytes, err := base64.StdEncoding.DecodeString(encodedReport.EphemPublicKey)
	if err != nil {
		return nil, err
	}
	var ephemPublicKey [32]byte
	copy(ephemPublicKey[:], ephemPublicKeyBytes)

	decrypted, ok := box.Open(nil, cypherBytes, &nonce, &ephemPublicKey, &secretKey)
	if !ok {
		return nil, errors.New("decryption error")
	}

	return decrypted, nil
}

func VerifySig(from, sigHex, report string) (bool, error) {
	typedData := apitypes.TypedData{
		Types: map[string][]apitypes.Type{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
			"Report": {
				{Name: "report", Type: "string"},
				{Name: "fromAddress", Type: "string"},
			},
		},
		PrimaryType: "Report",
		Domain: apitypes.TypedDataDomain{
			Name:    "Ethereum Private Message over Waku",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(5),
		},
		Message: map[string]interface{}{
			"fromAddress": from,
			"report":      report,
		},
	}

	return sigverify.VerifyTypedDataHexSignatureEx(
		common.HexToAddress(from),
		typedData,
		sigHex,
	)
}
