# Basic Concepts for Understanding Go Language Code

This document covers the fundamental concepts necessary for understanding the `portScanner.go` code written in the Go language. Below, you can find the topics you need to know to comprehend this code:

1. **Variables and Data Types**

   - `totalPorts`: An integer variable used to keep track of the total number of open ports during scanning.
   - `address`: A string variable representing the site address to be scanned.
   - `lb`, `ub`: Integer variables used as the lower bound and upper bound of the port range to be scanned.
   - `openports`: An integer slice used to store the discovered open ports.

   - Variable Declaration:
     ```go
     var totalPorts = 0
     ```
   - Variable Types:
     ```go
     var address string
     var lb, ub int
     var openports []int
     ```

2. **Functions and Return Values**

   - `WriteLog(pts []int, fname, address string)`: A function that writes the open ports to a log file. The `pts` parameter is an integer slice containing the open ports, `fname` parameter represents the name of the log file, and the `address` parameter represents the scanned address.
   - `GenName(ad string) string`: A function that generates the name of the log file using the scanned address and date information. The `ad` parameter is a string representing the scanned address.
   - `worker(ports chan int, address string, results chan int)`: A worker function that performs the port scanning. The `ports` parameter is a channel containing the ports to be scanned, the `address` parameter represents the scanned address, and the `results` parameter is a channel used to report the open ports.

   - Example of the `WriteLog` function:
     ```go
     func WriteLog(pts []int, fname, address string) {
         file, err := os.Create(fname)
         defer file.Close()
         if err != nil {
             panic(err)
         }
         fmt.Fprintln(file, address)
         for i := range pts {
             fmt.Fprintf(file, "Port %d is open!\n", pts[i])
         }
     }
     ```

   - Example of the `GenName` function:
     ```go
     func GenName(ad string) string {
         dt := time.Now()
         a := dt.Format("01-02-2006")
         return ad + "_" + a + ".txt"
     }
     ```

   - Example of the `worker` function:
     ```go
     func worker(ports chan int, address string, results chan int) {
         for p := range ports {
             address := fmt.Sprintf("%s:%d", address, p)
             conn, err := net.Dial("tcp", address)
             if err != nil {
                 results <- 0
                 continue
             }
             conn.Close()
             results <- p
         }
     }
     ```

3. **Importing Packages**

   - The code imports standard library packages such as `fmt`, `net`, `os`, `sort`, and `time`.

   - Standard Library Packages:
     ```go
     import (
         "fmt"
         "net"
         "os"
         "sort"
         "time"
     )
     ```

4. **Loops and Conditional Statements**

   - The code uses `for` loops to create goroutines based on the number of workers and perform the scanning process.
   - `if` conditions are used for error handling, operations, and to determine if

 the user wants to create a log file.

   - For Loop:
     ```go
     for i := 0; i < cap(ports); i++ {
         go worker(ports, address, results)
     }

     for i := lb; i <= ub; i++ {
         ports <- i
     }

     for i := lb; i < ub; i++ {
         port := <-results
         if port != 0 {
             openports = append(openports, port)
         }
     }
     ```

   - If Condition:
     ```go
     if err != nil {
         results <- 0
         continue
     }

     if exit == "e" || exit == "E" {
         WriteLog(openports, GenName(address), address)
     }
     ```

5. **File Operations and Time Operations**

   - In the `WriteLog` function, a log file is created and the scan results are written to the file.
   - The `GenName` function generates the name of the log file using the date information.
   - `time.Now()` retrieves the current time, and the `Format()` method formats it according to a specific layout.

   - File Creation and Closing:
     ```go
     file, err := os.Create(fname)
     defer file.Close()
     if err != nil {
         panic(err)
     }
     ```

   - Retrieving Time and Formatting:
     ```go
     dt := time.Now()
     a := dt.Format("07-04-2023") // mm-dd-yy
     ```

6. **Network Operations**

   - The `net.Dial()` function establishes a TCP connection. The `net` package is used to manage connections on the TCP/IP network.
   In the following example, a TCP connection is opened and closed:

     ```go
     conn, err := net.Dial("tcp", address)
     conn.Close()
     ```

7. **Channels and Goroutines**

   - The `ports` and `results` channels are used to communicate the port numbers and scan results between worker goroutines.
   - The `go` keyword is used to start worker goroutines concurrently.

   - Channel Declaration:
     ```go
     ports := make(chan int, 100)
     results := make(chan int)
     ```

   - Creating Goroutines:
     ```go
     go worker(ports, address, results)
     ```

8. **User Input**

   - The `fmt.Fscan()` function takes user input and assigns it to variables. `os.Stdin` represents the standard input stream.
   - The `fmt.Scan()` function waits for the user to input "e" or "E" and determines whether to create a log file based on the output.

   - Reading User Input:
     ```go
     fmt.Fscan(os.Stdin, &address)
     fmt.Fscan(os.Stdin, &lb)
     fmt.Fscan(os.Stdin, &ub)
     fmt.Scan(&exit)
     ```
