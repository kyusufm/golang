go mod init (nama folder)

create new main 

syntax
go run (file)
- menjalankan file tersebut
go build 
- membuat/build file dengan nama sesuai folder
go build (nama program)
- membuild dengan nama app (nama program)


package,
- menjelaskan bahwa semua file yang ada "package (main)" saling terkait satu sama lain


1. executable
2. library

kadang bikin package yang tidak langsung di eksekusi.
nanti ada 

cara mengetahui exec/lib, yang exec, harus masuk kedalam package main.

library bisa selain main.

penamaan fungsi di library, harus diawali huruf besar agar bisa diakses dari main.
jika mengawali dengan huruf kecil tidak bisa diakses dari fungsi main, tapi bisa diakses oleh fungsi lain di dalam file tersebut.


menggunakan import.
2.memanggil library dari package berbeda yang kita buat.
3.memanggil package lain yang dibikin orang lain.
1.memanggil default library dari golang (misal fmt.)

func main
wajib, karena ini fungsi utama yang akan di execute.

Variable, tipe data, dan default value
kenapa harus mendefenisikan tipe data ?
karena golang termasuk static type (serperti java, c#)

tipe umum- string, int, float64, bool

kalau ingin membuat variable tanpa harus mengisi typenya atau mengetik var, kita bisa mengetik nama :=

tidak bisa di bikin 2 kali, misal age:=20, kemudian diketik lagi age:=30 tidak bisa. karena sudah di deklarasi diatas. kalau ingin mengganti nilainya langsung diketik age = 20.


if condition
percabangan / kondisi
untuk if, akan return value boolean (true/false)

switch case

    
