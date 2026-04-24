## DorkSorgu

**DorkSorgu**, siber güvenlikte sızma testlerinin ve güvenlik analizlerinin en başında yer alan bilgi toplama sürecini hızlandırmak amacıyla geliştirilmiş, **Go** tabanlı bir otomasyon aracıdır.Hedef kurum veya kişi hakkında internete açık kaynaklar üzerinden veri sızdırma işlemini otomatize ederek manuel iş yükünü ortadan kaldırır ve hata payını en aza indirmeyi hedefler.

## Özellikler
**40+ Kritik Dork Şablonu:** Yönetici panellerinden veritabanı hatalarına, hassas dosyalardan API anahtarlarına kadar geniş bir alanda profesyonel tarama yapar.
* **Modern Web GUI:** Kullanıcı dostu, karanlık tema (dark mode) prensibiyle tasarlanmış dinamik arayüz.
* **Tek Tıkla Aksiyon:** Hazırlanan sorguları doğrudan Google üzerinde çalıştırma veya panoya kopyalama özelliği.
* **İstatistik Paneli:** Taranan domain için üretilen toplam ve kategorize edilmiş dork sayılarını anlık takip eden sayaçlar.

## Teknik Altyapı
Uygulama, performans odaklı bir **Client-Server** mimarisiyle çalışır:
* **Backend:** Go (Golang).
* **Frontend:** HTML5, CSS3, JavaScript.
* **Veri İletişimi:** Üretilen dorklar, JSON formatında paketlenerek frontend tarafına iletilir ve dinamik olarak görselleştirilir.

## Proje Dosya Yapısı
Proje, modülerlik ve taşınabilirlik ilkelerine uygun olarak tasarlanmıştır:
* `main.go`: Projenin motoru HTTP sunucusunu yönetir ve dork şablonlarını barındırır.
* `index.html`: Kullanıcının etkileşime girdiği dinamik görsel arayüzdür.
* `go.mod`: Go paket yönetim ve bağımlılık dosyası.

##  Kurulum ve Çalıştırma
Uygulamayı yerel makinenizde (macOS, Linux veya Windows) çalıştırmak için şu adımları izleyin:

1. **Terminali Açın:** Proje klasörünün bulunduğu dizine gidin.
2. **Sunucuyu Başlatın:** ```bash
   go run main.go
3. **Arayüze Erişin:** Tarayıcınızın adres çubuğuna http://localhost:8080 yazarak panele giriş yapın.
4. **Analiz Süreci:** Hedef alan adını (domain) girin ve "Analiz Et" butonuna tıklayarak Go backend'inin dorkları üretmesini bekleyin
