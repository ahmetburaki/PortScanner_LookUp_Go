package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Alan adını girin: ")
	domain, _ := reader.ReadString('\n')
	domain = cleanInput(domain)

	subdomains := []string{"www", "mail", "ftp", "admin"} // Alt alan adı listesi

	for _, subdomain := range subdomains {
		subdomain := fmt.Sprintf("%s.%s", subdomain, domain) // Alt alan adını oluştur

		ipAddresses, err := net.LookupIP(subdomain) // IP adreslerini çözümle

		if err != nil {
			fmt.Printf("Subdomain: %s, Hata: %v\n", subdomain, err) // Hata durumunda hata mesajını yazdır ve devam et
			continue
		}

		fmt.Printf("Subdomain: %s\n", subdomain) // Alt alan adını yazdır
		fmt.Println("IP Adresleri:")             // IP adreslerini yazdır
		for _, ip := range ipAddresses {
			fmt.Println(ip)
		}
		fmt.Println()
	}
}

func cleanInput(input string) string {
	input = input[:len(input)-1] // Son karakter olan '\n' karakterini kaldırın
	return input
}
