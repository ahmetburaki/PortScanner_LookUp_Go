This repository is prepared for the SyberBurger community and includes basic Go language concepts and simple applications. Everything shown in the training is for educational purposes.

# lookup.go

This program is a Go program that resolves subdomains for a user-inputted domain name and prints the IP addresses of each subdomain.

## Usage

1. Run the program.
2. The program will prompt you to enter a domain name.
3. Enter the domain name and press Enter.
4. The program will use the list of subdomains to resolve the IP addresses for each subdomain.
5. The resolved subdomain, IP addresses, and any error messages will be printed to the console.

## Example

```
Enter the domain name: example.com

Subdomain: www.example.com
IP Addresses:
192.0.2.1
192.0.2.2

Subdomain: mail.example.com
Error: lookup mail.example.com: no such host

Subdomain: ftp.example.com
IP Addresses:
192.0.2.3

Subdomain: admin.example.com
Error: lookup admin.example.com: no such host
```

# portScanner.go

This program is a Go program that performs a port scan on a user-inputted site address and port range to find open ports.

## Usage

1. Run the program.
2. The program will prompt you to enter a site address (e.g., example.com) and a port range.
3. Enter the site address and port range, and press Enter.
4. The program will scan the specified site address for ports within the specified range and identify open ports.
5. The discovered open ports will be printed to the console.

## Example

```
Enter a site address for scanning: example.com
Enter a port range for scanning: 1 1000
Scanning in progress...

Total open ports: 3
---------
1) Port 80 is open!
2) Port 443 is open!
3) Port 8080 is open!

Do you want to create a log file named "example.com_07-04-2023.txt"? (y/n)
```
