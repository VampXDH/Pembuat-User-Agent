package useragent

import (
	"fmt"
	"math/rand"
	"time"
)

// Daftar browser yang akan digunakan
var browsers = []string{
	"Mozilla/5.0",
	"Opera/9.80",
	"Safari/537.36",
	"Chrome/91.0",
}

// Daftar sistem operasi yang akan digunakan
var operatingSystems = []string{
	"Windows NT 10.0; Win64; x64",
	"Macintosh; Intel Mac OS X 10_15_7",
	"X11; Linux x86_64",
	"iPhone; CPU iPhone OS 14_6 like Mac OS X",
	"Android 11",
}

// Inisialisasi seed untuk fungsi random
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Fungsi untuk menghasilkan User-Agent secara random
func GenerateUserAgent() string {
	browser := browsers[rand.Intn(len(browsers))]
	operatingSystem := operatingSystems[rand.Intn(len(operatingSystems))]
	return fmt.Sprintf("%s (%s)", browser, operatingSystem)
}
