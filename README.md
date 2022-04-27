Langkah - Langkah :


clone


buat database dengan nama 'content-service' dan sesuai kan username dan password pada file database.go di dalam 'content-service-v2\app\internal\config\configuration.go pada bagian development


run dengan perintah 'go run app/cmd/main.go' di terminal


Untuk seeder, pertama uncomment seeder comment pada file 'content-service-v2\app\cmd\main.go kemudian jalankan pada terminal perintah 'go run app/cmd/main.go'. Jangan lupa untuk comment kembali


Untuk Cronjob sesuaikan CleanTemporaryPath pada environment variables pada folder environments

jalankan dengan perintah : go run app/cmd/cronjob/cronjob.go


Terima kasih