package main_test

import (
	"testing"

	"github.com/status-im/go-waku/waku/v2/protocol/pb"
	"github.com/stretchr/testify/assert"

	"github.com/famed-submission-form/backend/report"
)

func TestDecode(t *testing.T) {
	t.Run("DecodeReport", func(t *testing.T) {
		msg := &pb.WakuMessage{
			Payload:      []byte("{\"version\":\"x25519-xsalsa20-poly1305\",\"nonce\":\"pcT2bagPL+9OfVhR/NYaDNXBhLM8OogC\",\"ephemPublicKey\":\"acay4GausOL5fSV5UQdLp32bvWWpinADLSeuIJzSkBo=\",\"ciphertext\":\"WgUz/y07bxP2W5L1lnr4/tW3wckitLtLpNnu+4dw5nlaOTbK43kqAKZDpb75+WWExe6mPru5Ci4rV/wmJxqx+YDg7GXLzNdDHQGSJpCUBRXNl1KXThLqyN8NDYQn4RXriEW/KxuY/bVRvWyPWzKQlff3lsoqHcAeVtCKK0NrnNs0J/tgDJiWgPSvdaiWkChXPppnCRT8wM2YRJMsdACTmXlhcF1/7NXh3VXov6tdfWybVkIql2CHRcfHuM/HzjiPIbnAdTLacOTdsn8otWz1F5Z3EuS8Q4oKqttEwAgUXld5aqagBY1Wuh/KhSJlDn6WTIi4HLarhdpjoMJwsSMyctNW83Pw8qWSy/iZdVTDW5I4MDrsvqkhun38YO3fm0gny3AlcDGtPBocDOMh7hNWDnPSnB8RAb1rhN5idzfswhmkuZ9eUwEbl56fDAvrC9rnL1829KtN+avBJoLIkVIxpqDN6w8obllI2OvrYarSNVmWcamzU3jiItXGfXFU086wOWYMbqhO8uyBC+n/C69IH9IsiNtNLNDPF0eA7JTlBk1hbguICf+cZ9d4H7EKScrA\"}"),
			ContentTopic: "/relay-reactjs-chat/1/chat/proto",
			Timestamp:    1658846479522000000,
			Proof:        []byte("\n\200\002\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\022 \000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\032 \000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\" \000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000* \000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\0002 \000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000"),
		}

		result, err := report.Decode(msg, "955088f92ec369e81f362ba9dfad86bcd49ec5c70fdd68612b666455f8ee6fd3")

		assert.NoError(t, err)
		assert.Equal(t, "Here is message #1", result)
	})
}

func Test_DecryptReport(t *testing.T) {
	t.Run("DecryptReport", func(t *testing.T) {
		encodedReport := report.EncodedReport{
			Version:        "x25519-xsalsa20-poly1305",
			EphemPublicKey: "yF2JLwlbHNkj+ZrF5oc0XWeXQXNedUsRcBb4X9sJEhM=",
			Nonce:          "afLkOr7dpv3lhPZ+JdLmF0GeWsDm2suY",
			Ciphertext:     "RY/nei/dPcqE335pt1aEqQs87QVDCoJDNwdfEi12rxkN6JM//LCS4mhI+ERA4MXQa855TUeNfNuyCqdMvdn6Z+uXOtufdSK9X0vguFQm6bNWneAxndYgVjG5d7TQrjIYetm24kBtX/yYfkLzkHuk3Z+ukNwKgH86rbRaF0rtJsH1plSb9rrmG8ZEtX8DM8lu4jvhGh1iYMrHO9HALdP1DJmyxrk3faxAcVYAwx129qkmmkMS7Iw2pXnRrqcHn9/7uxavxPbDTjycksVr4Ao5WdwGeZj2NiFlXkkSYl1NzZZgx6mFENI3ThbIg0TbdlycctExPktwIo62tHWkTv+bx+pUaScVkn2yt8cbYr+xGSSiatFFxS0mHdrtSBO60JpD+jDB9McnrTLQhXAJZOgkLXLKmGa9MxySFkMifk2xgvefx2cCZajPvHgtfqwPkpIN0xa4V+Kru0sWTi9Cq4fO9qDaG/J6r7dsby+QikisXLeF8V+t+cRREnkGMD8iGA8U/rSWPIl6fD5I7t6H8RXsJ4mieyRxao8YXbbkHBMJtZF4ZF1J99iYVv2dZgSfqTm9",
		}
		hexBytesReport, err := report.DecryptReport(encodedReport, "955088f92ec369e81f362ba9dfad86bcd49ec5c70fdd68612b666455f8ee6fd3")
		assert.NoError(t, err)

		decodedReport, err := report.UnmarshalDecryptedReportHex(hexBytesReport)
		if err != nil {
			assert.NoError(t, err)
		}

		// TODO fix signature transmission
		cleanSig := string(decodedReport.Signature)[1 : len(decodedReport.Signature)-1]

		ok, err := report.VerifySig(string(decodedReport.EthAddress), cleanSig, string(decodedReport.Payload))
		assert.NoError(t, err)
		assert.True(t, ok)
	})
}
