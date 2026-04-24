# Google Dork Sorgu Otomasyonu

Bu proje, bir hedef domain için 40'tan fazla Google dork sorgusunu otomatik olarak üreten web tabanlı bir araçtır.

## Proje İçeriği
- **main.go**: Go diliyle yazılmış, dork üretimini yapan backend sunucusu.
- **index.html**: Kullanıcının domain girip sonuçları gördüğü arayüz.

## Kurulum ve Çalıştırma

1. Terminali açın ve proje klasörüne gidin.
2. Aşağıdaki komutu çalıştırarak backend sunucusunu başlatın:
   go run main.go
3. Sunucu çalıştıktan sonra `index.html` dosyasını tarayıcınızda açın.
4. Arama kutusuna hedef domaini (örneğin: ohu.edu.tr) yazın ve sorgulayın.