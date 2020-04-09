package client

import (
	"bytes"
	"github.com/cbeuw/Cloak/internal/common"
	"github.com/cbeuw/Cloak/internal/multiplex"
	"testing"
	"time"
)

func TestMakeAuthenticationPayload(t *testing.T) {
	tests := []struct {
		authInfo   authInfo
		expPayload authenticationPayload
		expSecret  [32]byte
	}{
		{
			authInfo{
				Unordered: false,
				SessionId: 3421516597,
				UID: []byte{
					0x4c, 0xd8, 0xcc, 0x15, 0x60, 0x0d, 0x7e,
					0xb6, 0x81, 0x31, 0xfd, 0x80, 0x97, 0x67, 0x37, 0x46},
				ServerPubKey: &[32]byte{
					0x21, 0x8a, 0x14, 0xce, 0x49, 0x5e, 0xfd, 0x3f,
					0xe4, 0xae, 0x21, 0x3e, 0x51, 0xf7, 0x66, 0xec,
					0x01, 0xd0, 0xb4, 0x87, 0x86, 0x9c, 0x15, 0x9b,
					0x86, 0x19, 0x53, 0x6e, 0x60, 0xe9, 0x51, 0x42},
				ProxyMethod:      "shadowsocks",
				EncryptionMethod: multiplex.E_METHOD_PLAIN,
				MockDomain:       "d2jkinvisak5y9.cloudfront.net",
				WorldState: common.WorldState{
					Rand: bytes.NewBuffer([]byte{
						0xf1, 0x1e, 0x42, 0xe1, 0x84, 0x22, 0x07, 0xc5,
						0xc3, 0x5c, 0x0f, 0x7b, 0x01, 0xf3, 0x65, 0x2d,
						0xd7, 0x9b, 0xad, 0xb0, 0xb2, 0x77, 0xa2, 0x06,
						0x6b, 0x78, 0x1b, 0x74, 0x1f, 0x43, 0xc9, 0x80}),
					Now: func() time.Time { return time.Unix(1579908372, 0) },
				},
			},
			authenticationPayload{
				randPubKey: [32]byte{
					0xee, 0x9e, 0x41, 0x4e, 0xb3, 0x3b, 0x85, 0x03,
					0x6d, 0x85, 0xba, 0x30, 0x11, 0x31, 0x10, 0x24,
					0x4f, 0x7b, 0xd5, 0x38, 0x50, 0x0f, 0xf2, 0x4d,
					0xa3, 0xdf, 0xba, 0x76, 0x0a, 0xe9, 0x19, 0x19},
				ciphertextWithTag: [64]byte{
					0x71, 0xb1, 0x6c, 0x5a, 0x60, 0x46, 0x90, 0x12,
					0x36, 0x3b, 0x1b, 0xc4, 0x79, 0x3c, 0xab, 0xdd,
					0x5a, 0x53, 0xc5, 0xed, 0xaf, 0xdb, 0x10, 0x98,
					0x83, 0x96, 0x81, 0xa6, 0xfc, 0xa2, 0x1e, 0xb0,
					0x89, 0xb2, 0x29, 0x71, 0x7e, 0x45, 0x97, 0x54,
					0x11, 0x7d, 0x9b, 0x92, 0xbb, 0xd6, 0xce, 0x37,
					0x3b, 0xb8, 0x8b, 0xfb, 0xb6, 0x40, 0xf0, 0x2c,
					0x6c, 0x55, 0xb9, 0xfc, 0x5d, 0x34, 0x89, 0x41},
			},
			[32]byte{
				0xc7, 0xc6, 0x9b, 0xbe, 0xec, 0xf8, 0x35, 0x55,
				0x67, 0x20, 0xcd, 0xeb, 0x74, 0x16, 0xc5, 0x60,
				0xee, 0x9d, 0x63, 0x1a, 0x44, 0xc5, 0x09, 0xf6,
				0xe0, 0x24, 0xad, 0xd2, 0x10, 0xe3, 0x4a, 0x11},
		},
	}
	for _, tc := range tests {
		func() {
			payload, sharedSecret := makeAuthenticationPayload(tc.authInfo)
			if payload != tc.expPayload {
				t.Errorf("payload doesn't match:\nexp %v\ngot %v", tc.expPayload, payload)
			}
			if sharedSecret != tc.expSecret {
				t.Errorf("secret doesn't match:\nexp %x\ngot %x", tc.expPayload, payload)
			}
		}()
	}
}
