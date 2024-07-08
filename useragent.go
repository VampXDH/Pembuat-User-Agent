package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Data untuk user agent
var (
	browserList = []string{
		"Chrome", "Firefox", "Safari", "Edge", "Opera", "Internet Explorer", "Brave", "Vivaldi", "Yandex", "UC Browser",
		"Mozilla", "Netscape", "Konqueror", "Lynx", "Links", "Midori", "Epiphany", "SeaMonkey", "Pale Moon", "Maxthon",
		"Sleipnir", "Waterfox", "Avant Browser", "Camino", "Comodo IceDragon", "Dolphin", "Flock", "GNU IceCat", "K-Meleon",
		"NetFront", "OmniWeb", "Puffin", "QuteBrowser", "Slimjet", "Tor Browser", "Vivaldi", "W3M", "Yandex Browser",
	}
	osList = []string{
		"Windows NT 10.0; Win64; x64",
		"Windows NT 10.0; WOW64",
		"Windows NT 6.3; Win64; x64",
		"Windows NT 6.1; Win64; x64",
		"Windows NT 6.1; WOW64",
		"Windows NT 6.0",
		"Windows NT 5.1",
		"Windows NT 5.0",
		"Windows 98; Win 9x 4.90",
		"Windows 95",
		"Macintosh; Intel Mac OS X 10_15_7",
		"Macintosh; Intel Mac OS X 10_14_6",
		"Macintosh; Intel Mac OS X 10_13_6",
		"Macintosh; Intel Mac OS X 10_12_6",
		"Macintosh; Intel Mac OS X 10_11_6",
		"X11; Linux x86_64",
		"X11; Ubuntu; Linux x86_64",
		"X11; Fedora; Linux x86_64",
		"X11; Debian; Linux x86_64",
		"iPhone; CPU iPhone OS 14_6 like Mac OS X",
		"iPhone; CPU iPhone OS 13_5 like Mac OS X",
		"iPhone; CPU iPhone OS 12_4 like Mac OS X",
		"iPad; CPU OS 14_6 like Mac OS X",
		"iPad; CPU OS 13_5 like Mac OS X",
		"iPad; CPU OS 12_4 like Mac OS X",
		"Android 10; SM-G973F",
		"Android 9; Pixel 3",
		"Android 8; Pixel 2",
		"Android 7; Nexus 6P",
		"Linux i686",
		"Linux armv7l",
		"Linux aarch64",
	}
	androidVersionList = []string{
		"10.0", "9.0", "8.0", "7.0", "6.0",
	}
	iOSVersionList = []string{
		"14.6", "13.5", "12.4", "11.3", "10.2",
	}
	brandList = []string{
		"Samsung", "Google", "Apple", "Huawei", "Xiaomi", "OnePlus", "Sony", "LG", "HTC", "Nokia",
	}
)

// Fungsi untuk membuat user agent acak
func RandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())

	browser := browserList[rand.Intn(len(browserList))]
	os := getRandomOS()
	version := randomVersion()

	var userAgent string
	switch browser {
	case "Chrome":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
	case "Firefox":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 Firefox/%s", os, version, version)
	case "Safari":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%s Safari/605.1.15", os, version)
	case "Edge":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Edg/%s", os, version, version)
	case "Opera":
		userAgent = fmt.Sprintf("Opera/9.80 (%s) Presto/2.12.388 Version/%s", os, version)
	case "Internet Explorer":
		userAgent = fmt.Sprintf("Mozilla/5.0 (compatible; MSIE %s; %s; Trident/6.0)", version, os)
	case "Brave":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Brave/%s", os, version, version)
	case "Vivaldi":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Vivaldi/%s", os, version, version)
	case "Yandex":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 YaBrowser/%s", os, version, version)
	case "UC Browser":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 UC Browser/%s", os, version, version)
	case "Mozilla":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 Firefox/%s", os, version, version)
	case "Netscape":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 Firefox/%s", os, version, version)
	case "Konqueror":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) KHTML/5.0 (like Gecko) Konqueror/%s", os, version)
	case "Lynx":
		userAgent = fmt.Sprintf("Lynx/%s %s", version, os)
	case "Links":
		userAgent = fmt.Sprintf("Links (%s)", os)
	case "Midori":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Midori/%s", os, version, version)
	case "Epiphany":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Epiphany/%s", os, version, version)
	case "SeaMonkey":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 SeaMonkey/%s", os, version, version)
	case "Pale Moon":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 PaleMoon/%s", os, version, version)
	case "Maxthon":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Maxthon/%s", os, version, version)
	case "Sleipnir":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Sleipnir/%s", os, version, version)
	case "Waterfox":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 Waterfox/%s", os, version, version)
	case "Avant Browser":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Avant Browser/%s", os, version, version)
	case "Camino":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Camino/%s", os, version, version)
	case "Comodo IceDragon":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 IceDragon/%s", os, version, version)
	case "Dolphin":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Dolphin/%s", os, version, version)
	case "Flock":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Flock/%s", os, version, version)
	case "GNU IceCat":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 IceCat/%s", os, version, version)
	case "K-Meleon":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 K-Meleon/%s", os, version, version)
	case "NetFront":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 NetFront/%s", os, version, version)
	case "OmniWeb":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 OmniWeb/%s", os, version, version)
	case "Puffin":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Puffin/%s", os, version, version)
	case "QuteBrowser":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 QuteBrowser/%s", os, version, version)
	case "Slimjet":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Slimjet/%s", os, version, version)
	case "Tor Browser":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Tor Browser/%s", os, version, version)
	case "W3M":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 W3M/%s", os, version, version)
	case "Yandex Browser":
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 YaBrowser/%s", os, version, version)
	default:
		userAgent = fmt.Sprintf("Mozilla/5.0 (%s)", os)
	}

	return userAgent
}

// Fungsi untuk menghasilkan versi acak
func randomVersion() string {
	return fmt.Sprintf("%d.%d.%d", rand.Intn(100)+1, rand.Intn(100)+1, rand.Intn(100)+1)
}

// Fungsi untuk menghasilkan OS acak
func getRandomOS() string {
	rand.Seed(time.Now().UnixNano())

	os := osList[rand.Intn(len(osList))]
	switch {
	case strings.Contains(os, "Android"):
		os += fmt.Sprintf("; Mobile %s", getRandomAndroidVersion())
	case strings.Contains(os, "iPhone"):
		os += fmt.Sprintf("; CPU iPhone OS %s like Mac OS X", getRandomiOSVersion())
	case strings.Contains(os, "iPad"):
		os += fmt.Sprintf("; CPU OS %s like Mac OS X", getRandomiOSVersion())
	case strings.Contains(os, "Windows"):
		os += fmt.Sprintf("; Win%d; %s", rand.Intn(2)+1, getRandomBrand())
	case strings.Contains(os, "Macintosh"):
		os += fmt.Sprintf(" Intel Mac OS X 10_%d_%d", rand.Intn(10)+10, rand.Intn(9)+5)
	case strings.Contains(os, "Linux"):
		os += fmt.Sprintf("; %s", getRandomLinuxVariant())
	}
	return os
}

// Fungsi untuk menghasilkan versi Android acak
func getRandomAndroidVersion() string {
	return androidVersionList[rand.Intn(len(androidVersionList))]
}

// Fungsi untuk menghasilkan versi iOS acak
func getRandomiOSVersion() string {
	return iOSVersionList[rand.Intn(len(iOSVersionList))]
}

// Fungsi untuk menghasilkan merek HP acak
func getRandomBrand() string {
	return brandList[rand.Intn(len(brandList))]
}

// Fungsi utama untuk contoh penggunaan
func main() {
	// Contoh penggunaan fungsi RandomUserAgent
	for i := 0; i < 10; i++ {
		fmt.Println(RandomUserAgent())
	}
}
