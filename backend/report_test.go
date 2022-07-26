package main_test

import (
	"encoding/hex"
	pblib "github.com/famed-submission-form/backend/pb"
	"github.com/golang/protobuf/proto"
	"testing"

	"github.com/famed-submission-form/backend/report"

	"github.com/stretchr/testify/assert"
)

func Test_DecryptReport(t *testing.T) {
	t.Run("DecryptReport", func(t *testing.T) {
		encodedReport := report.EncodedReport{
			Version:        "x25519-xsalsa20-poly1305",
			EphemPublicKey: "Jx7Pj5Osxqbpa90Ws+zHwRIaICJ2P+tN+eD5VUytEDQ=",
			Nonce:          "G6ko7Z0TVhb6gNO9ACi7J2u4Bynmwpuq",
			Ciphertext:     "7+MoHRsxXKCvnPUqUX3s8KKrmsqks02FKLaVOPt5kL6ne5cQCs0wVFuiMpfSCXpd7fzK+RzhCcgP/s7G9hMGKmi4Fgt2tTDSTphWxUuuplGkp79U0bcG1a9SCVyzAaCz+hC+crVOE/SbGy24+fWkXiFjAew86y4nV/Kn5ZfqnePoHIxESobFx+VqRsRydY3ziFYGy+M/BPC6vaYyTzs5c1b81E14TvZ2MpzOkpKXGyr+vlElFi4SfDx1rqaPM8oYxAYparC0W3E82qLFOCi4zm30teoQW8fbZC0Clh5PSrz/lttRtU8L5GgpZlmr8B9xrE+1/AJeWk5MP8iNGELmDC8w/90MdAjCB/pI42kCFMXhpOSLc7DAT3G6RcM6T4qpmmc36cIx8ZFJ+lsCqhCZGpF6KDOoUsAQM9Yjlo0Ex21dWTPjcb6OIS7xdjko7YRAmB68ciM4ms/71FwposkP/5ElakMtb+mvqgrNfQ2Dgsa01LqbhWalQiJwzgCnY5CNH/ccAM9I75N4UOjtYy2bldZvOwJoac2OEG7uBt5rb8OVfrwaUUezqZbaTWfnItdQJq8=",
		}
		hexBytesReport, err := report.DecryptReport(encodedReport)
		assert.NoError(t, err)
		assert.Equal(t, string(hexBytesReport), "08b8f7a0b0a330121348657265206973206d657373616765202331321a2a3078386262393245624161323431383537454435613534453934343161463662306134323833314434612286012230783335633465376664666436643536353230303031653764613061323362376134343337616664393639313536356531323662323534303033356462363734306531373735633464636130623330383831343636613536336638336633316134663363656466323761653532623337386464616663306638343139633537626262316222")
		
		dst := make([]byte, hex.DecodedLen(len(hexBytesReport)))
		hex.Decode(dst, hexBytesReport)
		assert.NoError(t, err)

		decodedreport := &pblib.Report{}
		err = proto.Unmarshal(dst, decodedreport)
		if err != nil {
			assert.NoError(t, err)
		}

		// TODO fix signature transmission
		cleanSig := string(decodedreport.Signature)[1 : len(decodedreport.Signature)-1]

		ok, err := report.VerifySig(string(decodedreport.EthAddress), cleanSig, string(decodedreport.Payload))
		assert.NoError(t, err)
		assert.True(t, ok)
	})
}
