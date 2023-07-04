# Go Dilinde Kodun Anlaşılması için Temel Konular

Bu doküman, Go dilinde yazılmış `portScanner.go` kodunun anlaşılması için gereken temel konuları kapsamaktadır. Aşağıda, bu kodu anlamak için bilinmesi gereken konuları bulabilirsiniz:

1. **Değişkenler ve Veri Türleri**

   - `totalPorts`: Tarama sırasında açık portların toplam sayısını tutmak için kullanılan bir tamsayı değişkeni.
   - `address`: Taranacak site adresini temsil eden bir dize (string) değişkeni.
   - `lb`, `ub`: Tarama yapılacak port aralığının alt sınırı (lower bound) ve üst sınırı (upper bound) olarak kullanılan tamsayı değişkenleri.
   - `openports`: Bulunan açık portların depolandığı bir tamsayı dizisi (slice).

   - Değişken Tanımlama:
     ```go
     var totalPorts = 0
     ```
   - Değişken Türleri:
     ```go
     var address string
     var lb, ub int
     var openports []int
     ```


2. **Fonksiyonlar ve Geri Dönüş Değerleri**

   - `WriteLog(pts []int, fname, address string)`: Açık portları bir log dosyasına yazan bir fonksiyon. `pts` parametresi açık portları içeren bir tamsayı dizisidir. `fname` parametresi log dosyasının adı, `address` parametresi ise taranan adresin kendisidir.
   - `GenName(ad string) string`: Taranan adresi ve tarih bilgisini kullanarak log dosyasının adını oluşturan bir fonksiyon. `ad` parametresi taranan adresi temsil eden bir dizedir.
   - `worker(ports chan int, address string, results chan int)`: Port taramasını gerçekleştiren işçi (worker) fonksiyonu. `ports` parametresi tarama yapılacak portları içeren bir kanal, `address` parametresi taranan adresi temsil eden bir dizedir. `results` parametresi ise açık portların raporlandığı bir kanaldır.


   - `WriteLog` fonksiyonu örneği:
     ```go
     func WriteLog(pts []int, fname, address string) {
         file, err := os.Create(fname)
         defer file.Close()
         if err != nil {
             panic(err)
         }
         fmt.Fprintln(file, address)
         for i := range pts {
             fmt.Fprintf(file, "Port %d açık!\n", pts[i])
         }
     }
     ```

   - `GenName` fonksiyonu örneği:
     ```go
     func GenName(ad string) string {
         dt := time.Now()
         a := dt.Format("07-04-2023") // ay-gün-yıl
         return ad + "_" + a + ".txt"
     }
     ```

   - `worker` fonksiyonu örneği:
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


3. **Paketlerin İçe Aktarılması**

   - `fmt`, `net`, `os`, `sort`, `time` gibi standart kütüphane paketlerinin kodda kullanıldığı belirtilmiştir.

   - Standart Kütüphane Paketleri:
     ```go
     import (
         "fmt"
         "net"
         "os"
         "sort"
         "time"
     )
     ```


4. **Döngüler ve Koşullu İfadeler**

   - `for` döngüleri kullanılarak işçi (worker) sayısına bağlı olarak goroutine'ler oluşturulur ve tarama işlemi gerçekleştirilir.
   - `if` koşulu, hata kontrolü ve işlemlerinin yanı sıra kullanıcının log dosyası oluşturup oluşturmak istemediğini belirlemek için kullanılır.

   - For Döngüsü:
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

   - If Koşulu:
     ```go
     if err != nil {
         results <- 0
         continue
     }

     if exit == "e" || exit == "E" {
         WriteLog(openports, GenName(address), address)
     }
     ```


5. **Dosya İşlemleri ve Zaman İşlemleri**

   - `WriteLog` fonksiyonunda, log dosyası oluşturulur ve tarama sonuçları dosyaya yazılır.
   - `GenName` fonksiyonunda, tarih bilgisi kullanılarak log dosyasının adı oluşturulur.
   - `time.Now()` ile mevcut zaman bilgisi alınır ve `Format()` yöntemi ile belirli bir formata dönüştürülür.

   - Dosya Oluşturma ve Kapatma:
     ```go
     file, err := os.Create(fname)
     defer file.Close()
     if err != nil {
         panic(err)
     }
     ```

   - Zaman Bilgisi Alma ve Formatlama:
     ```go
     dt := time.Now()
     a := dt.Format("01-02-2006")
     ```


6. **Ağ İşlemleri**

   - `net.Dial()` fonksiyonu ile TCP bağlantısı oluşturulur. `net` paketi, TCP/IP ağı üzerindeki bağlantıları yönetmek için kullanılır.
   Aşağıdaki örnekte TCP bağlantısı açılıp/kapatılır.

     ```go
     conn, err := net.Dial("tcp", address)
     conn.Close()
     ```


7. **Kanal (Channel) ve Goroutine'ler**

   - `ports` ve `results` kanalları, işçi (worker) goroutine'leri arasında port

 numaraları ve tarama sonuçlarının iletilmesi için kullanılır.
   - `go` anahtar kelimesi ile işçi (worker) goroutine'leri eşzamanlı olarak başlatılır.

   - Kanal Tanımlama:
     ```go
     ports := make(chan int, 100)
     results := make(chan int)
     ```

   - Goroutine Oluşturma:
     ```go
     go worker(ports, address, results)
     ```


8. **Kullanıcı Girişi**

   - `fmt.Fscan()` fonksiyonu kullanıcıdan giriş alır ve değişkenlere atama yapar. `os.Stdin` standart giriş akışını temsil eder.
   - `fmt.Scan()` fonksiyonu kullanıcının "e" veya "E" harfini girmesini bekler ve çıktıya göre log dosyası oluşturulup oluşturulmayacağına karar verir.

   - Kullanıcıdan Giriş Alma:
     ```go
     fmt.Fscan(os.Stdin, &address)
     fmt.Fscan(os.Stdin, &lb)
     fmt.Fscan(os.Stdin, &ub)
     fmt.Scan(&exit)
     ```