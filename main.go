package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	
	produk := &pb.Product{
		Id: 1,
		NamaProduk: "Laptop Gaming",
		Kategori: "Elektronik",
		Harga: 1500000,
		Stok: 25,
		Deskripsi: "Laptop gaming dengan spesifikasi tinggi untuk pengalaman bermain game yang luar biasa.",
		Spesifikasi: &pb.Spesifikasi{
			Procesor: "Intel Core i7",
			Ram: "16GB",
			Penyimpanan: "1TB SSD",
			Layar: "15.6 inci",
			Grafis: "NVIDIA RTX 3060",
		},
		Merek: "BrandX",
		TanggalRilis: "2023-07-15",
		Rating: 4.5,
		Ulasan: []*pb.Ulasan{
			{
				IdUlasan: 1,
				NamaPelanggan: "Budi",
				TanggalUlasan:"2023-07-20",
				Komentar:"Laptopnya sangat cepat dan lancar untuk bermain game.",
				Rating: 5,
			},
			{
				IdUlasan: 2,
				NamaPelanggan: "Siti",
				TanggalUlasan:"2023-07-22",
				Komentar:"Desainnya keren tapi agak panas kalau dipakai lama.",
				Rating: 4,
			},
		},
	}

	data, err := proto.Marshal(produk)
	if err != nil {
		log.Fatal("marshal error", err)
	}

	//compact binary wire format
	fmt.Println(data)

	testProduk := &pb.Product{}
	if err = proto.Unmarshal(data, testProduk); err != nil{
		log.Fatal("Unmarshal error", err)
	}

	for _, product := range testProduk.GetNamaProduk() {
		fmt.Println(product)
	}

}