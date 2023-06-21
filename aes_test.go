package rice

import "testing"

func TestAESEncrypt(t *testing.T) {
	key := NextUUID()
	type args struct {
		keyString   string
		plainString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{keyString: key, plainString: "😊"}, "😊", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESEncrypt(tt.args.keyString, tt.args.plainString)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got2, err := AESDecrypt(tt.args.keyString, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got2 != tt.want {
				t.Errorf("AESDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
