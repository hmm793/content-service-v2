package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	cron "github.com/robfig/cron/v3"
)

func main() {
    // Environment Variables
	err := godotenv.Load("app/environments/app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}

    // set scheduler berdasarkan zona waktu sesuai kebutuhan
    jakartaTime, _ := time.LoadLocation("Asia/Jakarta") 
    scheduler := cron.New(cron.WithLocation(jakartaTime))

    // stop scheduler tepat sebelum fungsi berakhir
    defer scheduler.Stop()

    // set task yang akan dijalankan scheduler
    // gunakan crontab string untuk mengatur jadwal
    scheduler.AddFunc("0 0 * * *", CleanTemporary)

    // start scheduler
    go scheduler.Start()

    // trap SIGINT untuk trigger shutdown.
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig
}


func CleanTemporary(){
    if err := os.RemoveAll(os.Getenv("CleanTemporaryPath"));err != nil {
        log.Fatal(err)
    }
    
    if err := os.Mkdir(os.Getenv("CleanTemporaryPath"), os.ModePerm); err != nil {
        log.Fatal(err)
    }
}