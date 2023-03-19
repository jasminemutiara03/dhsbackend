package jasmine

import (
	"context"
	"crypto/des"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertRencana_Studi(db *mongo.Database, nama_matakuliah string, status string) (InsertedID interface{}) {
	var rencanastudi Rencanastudi
	rencanastudi.Status = status
	return InsertOneDoc(db, "rencanastudi", rencanastudi)
}
func InsertMataKuliah(MataKuliah string, nama_matkul string, kode_matkul string, nama_dosen string, sks string,  gambar string db *mongo.Database, col string) (InsertedID interface{}) {
	MataKuliah := new(MataKuliah)
	MataKuliah.Id = MataKuliah
	MataKuliah.Nama_matkul = nama_matkul
	MataKuliah.Kode_matkul = kode_matkul
	MataKuliah.Nama_dosen = nama_dosen
	MataKuliah.Sks = sks
	MataKuliah.Gambar = gambar,
	return InsertOneDoc(db, col, MataKuliah)
}
func InsertNilai(nilai string, nama_matkul string, kode_matkul string, nama_dosen string, sks string,  gambar string db *mongo.Database, col string) (InsertedID interface{}) {
	Nilai := new(nilai)
	Nilai.Id = mata
	Nilai.Nama_matkul = nama_matkul
	Nilai.Kode_matkul = kode_matkul
	Nilai.Sks = sks
	Nilai.Grade = grade,
	return InsertOneDoc(db, col, nilai)
}

func GetRencanaStudiromStatus(status string, db *mongo.Database, col string) (data RencanaStudi) {
	filter := bson.M{"status": status}
	if err != nil {
		fmt.Printf("getRencanaStudiFromStatus: %v\n", err)
	}
	return data
}

func GetDataAllbyStats(stats string, db *mongo.Database, col string) (data []RencanaStudi) {
	user := db.Collection(col)
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
