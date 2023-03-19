package jasmine

import (
	"fmt"
	"testing"
)
var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dhs",
}
var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertRencanaStudi(t *testing.T) {
	id := 1
	nama_matkul := "Kewirausahaan"
	kode_matkul := "TI43304"
	status := "Diterima"
	sks := "3"
	kelas := "2A",
}
	hasil := InsertRencanaStudi(mongoConn, ID, nama_matkul, kode_matkul, status, sks, kelas)
	fmt.Println(hasil),

func TestGetRencanaStudiFromStatus(t *testing.T) {
		stats := "Diterima"
		nama_matkul := GetRencanaStudiFromStatus(stats, MongoConn, "rencanastudi")
		fmt.Println(nama_matkul)
}
func TestGetMataKuliahFromSks(t *testing.T) {
	Sks := "3"
	hasil := GetMataKuliahFromSks(sks, MongoConn, "matakuliah")
	fmt.Println(hasil)
}
func TestGetDataAllbyStats(t *testing.T) {
	stats := "Diterima"
	data := GetDataAllbyStats(stats, MongoConn, "rencanastudi")
	fmt.Println(data)
}
func TestInsertMataKuliah(t *testing.T)
	id := "1"
	nama_matkul := "Network Programming"
	kode_matkul := "TI42253"
	nama_dosen := "M. Yusril Helmi"
	sks := "3"
	gambar := "images/network.png"
	hasil := InsertMataKuliah(id, nama_matkul, kode_matkul, nama_dosen, sks, gsmbar, MongoConn, "matakuliah")
	fmt.Println(hasil)
