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