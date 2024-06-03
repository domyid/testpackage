package idname

import "github.com/whatsauth/itmodel"

func IDName(Pesan itmodel.IteungMessage) (reply string) {

	return "Hai.. hai..  kakak atas nama :\n" + Pesan.Alias_name + "\n berhasil absensi\nmakasih"
}
