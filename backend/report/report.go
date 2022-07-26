package report

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"golang.org/x/crypto/nacl/box"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

type EncodedReport struct {
	Version        string `json:"version"`
	Nonce          string `json:"nonce"`
	EphemPublicKey string `json:"ephemPublicKey"`
	Ciphertext     string `json:"ciphertext"`
}

func Unmarshal(data []byte) (EncodedReport, error) {
	encodedReport := EncodedReport{}
	err := json.Unmarshal(data, &encodedReport)
	if err != nil {
		return encodedReport, err
	}

	return encodedReport, nil
}

func DecryptReport(encodedReport EncodedReport) ([]byte, error) {
	secretKeyBytes, err := hex.DecodeString("955088f92ec369e81f362ba9dfad86bcd49ec5c70fdd68612b666455f8ee6fd3")
	if err != nil {
		return nil, err
	}

	fmt.Println(len(secretKeyBytes))

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	nonceBytes, err := base64.StdEncoding.DecodeString(encodedReport.Nonce)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(nonceBytes))
	var nonce [24]byte
	copy(nonce[:], nonceBytes)

	cypherBytes, err := base64.StdEncoding.DecodeString(encodedReport.Ciphertext)

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
		ethcommon.HexToAddress(from),
		typedData,
		sigHex,
	)
}
