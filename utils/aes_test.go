package utils

import "testing"

const (
	testAESUtilAPIV3Key       = ""
	testAESUtilCiphertext     = ""
	testAESUtilNonce          = ""
	testAESUtilAssociatedData = ""
)

func TestDecryptAes256Gcm(t *testing.T) {
	type args struct {
		apiv3Key       string
		associatedData string
		nonce          string
		ciphertext     string
	}
	tests := []struct {
		name            string
		args            args
		wantCertificate string
		wantErr         bool
	}{
		{
			name: "decrypt certificate",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilCiphertext,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilAssociatedData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCertificate, err := DecryptAES256GCM(tt.args.apiv3Key, tt.args.associatedData, tt.args.nonce, tt.args.ciphertext)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptAES256GCM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCertificate != tt.wantCertificate {
				t.Errorf("DecryptAES256GCM() gotCertificate = %v, want %v", gotCertificate, tt.wantCertificate)
			}
		})
	}
}
