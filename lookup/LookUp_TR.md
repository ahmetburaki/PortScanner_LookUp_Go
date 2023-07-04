# Go Dilinde Kodun Anlaşılması için Temel Konular

Bu doküman, Go dilinde yazılmış `lookup.go` koduun anlaşılması için gereken temel konuları kapsamaktadır. Aşağıda, bu kodu anlamak için bilinmesi gereken konuları bulabilirsiniz:

1. **Temel Go Dil Özellikleri**

   - Değişkenler ve Veri Türleri: Go dilinde değişkenlerin nasıl tanımlanacağı ve farklı veri türlerinin nasıl kullanılacağı öğrenilmelidir.
   - İfadeler ve Operatörler: Aritmetik operatörler, ilişkisel operatörler ve mantıksal operatörler gibi ifadeler ve operatörlerin nasıl kullanıldığına dair temel bilgiye sahip olunmalıdır.
   - Döngüler: `for` döngüsü kullanılarak tekrarlı işlemler yapılabilir.
   - Koşullar: `if`, `else if` ve `else` ifadeleri ile koşullu ifadeler oluşturulabilir.
   - Fonksiyonlar: Fonksiyonların nasıl tanımlandığı, çağrıldığı ve geri dönüş değerleriyle nasıl çalıştığı öğrenilmelidir.


   - Değişkenler ve Veri Türleri:
     ```go
     var age int = 30
     var name string = "John Doe"
     ```
   - İfadeler ve Operatörler:
     ```go
     x := 5
     y := 10
     sum := x + y
     ```
   - Döngüler:
     ```go
     for i := 0; i < 5; i++ {
         fmt.Println(i)
     }
     ```
   - Koşullar:
     ```go
     if x > 10 {
         fmt.Println("x is greater than 10")
     } else if x < 10 {
         fmt.Println("x is less than 10")
     } else {
         fmt.Println("x is equal to 10")
     }
     ```
   - Fonksiyonlar:

     ```go
     func add(x, y int) int {
         return x + y
     }
     ```


2. **Paketler ve Importlar**

   - `import` ifadesi ile paketlerin içe aktarılması ve kullanılması sağlanır. Bu kodda `bufio`, `fmt`, `net` ve `os` paketleri kullanılmaktadır.

    ```go
   import (
       "fmt"
       "net"
       "os"
   )
   ```

3. **Konsol Girişi/Çıkışı**

   - `fmt` paketi kullanılarak konsola çıktı yazdırma (`Println`, `Printf`) ve kullanıcıdan giriş alma (`Scan`, `ReadString`) işlemleri yapılır. Bu kodda kullanıcıdan bir alan adı girdisi alınmaktadır.


   - Çıktı Yazdırma:
     ```go
     fmt.Println("Hello, World!")
     ```
   - Giriş Alma:
     ```go
     var name string
     fmt.Print("Enter your name: ")
     fmt.Scan(&name)
     ```

4. **Dize İşlemleri**

   - Dize manipülasyonları: Dize birleştirme (`+` operatörü veya `Sprintf` fonksiyonu), dize bölme ve dize uzunluğunu alma gibi işlemler yapılabilir.

   - Dize Birleştirme:
     ```go
     str1 := "Hello"
     str2 := "World"
     result := str1 + " " + str2
     ```
   - Dize Uzunluğunu Alma:
     ```go
     str := "Hello"
     length := len(str)
     ```

5. **Hata Yönetimi**

   - `error` veri türü ve `if` ifadeleri kullanılarak hata yönetimi yapılır. Hata durumlarının nasıl kontrol edileceği ve işleneceği öğrenilmelidir. Bu kodda `LookupIP` fonksiyonu bir hata döndürebilir.


   ```go
   result, err := SomeFunction()
   if err != nil {
       fmt.Println("Error:", err)
   } else {
       fmt.Println("Result:", result)
   }
   ```


6. **Ağ İşlemleri**

   - `net` paketi, ağ işlemleri için kullanılır. IP adreslerini çözümlemek için `LookupIP` fonksiyonu kullanılır.

   ```go
   ipAddresses, err := net.LookupIP("www.example.com")
   if err != nil {
       fmt.Println("Error:", err)
   } else {
       fmt.Println("IP Addresses:")
       for _, ip := range ipAddresses {
           fmt.Println(ip)
       }
   }
   ```
   
   Bu kod, "www.example.com" alan adının IP adreslerini çözümleyerek ekrana yazdırır. 

    1. `net.LookupIP("www.example.com")` ifadesi kullanılarak "www.example.com" alan adının IP adresleri çözümlenir. Bu işlem, DNS (Alan Adı Sistemi) çözümlemesi yaparak ilgili alan adının bağlı olduğu IP adreslerini bulmayı sağlar.
    2. Sonuçlar, `ipAddresses` değişkenine ve olası bir hata da `err` değişkenine atanır.
    3. `if err != nil` ifadesi kullanılarak bir hata olup olmadığı kontrol edilir. Eğer bir hata varsa, `fmt.Println("Error:", err)` ifadesi çalışır ve hatayı ekrana yazdırır.
    4. Eğer hata yoksa, yani çözümleme başarılıysa, `fmt.Println("IP Addresses:")` ifadesiyle IP adreslerini yazdırmaya başlanır.
    5. `for _, ip := range ipAddresses` döngüsü kullanılarak `ipAddresses` dizisi üzerindeki her bir IP adresi için döngü işlemi gerçekleştirilir.
    6. Her bir IP adresi, `fmt.Println(ip)` ifadesiyle ekrana yazdırılır.

        Sonuç olarak, bu kod "www.example.com" alan adının çözümlenmiş IP adreslerini ekrana yazdırır. Eğer çözümleme işlemi başarılı olursa, IP adresleri sırasıyla görüntülenir. Eğer bir hata oluşursa, hata mesajı ekrana yazdırılır.
