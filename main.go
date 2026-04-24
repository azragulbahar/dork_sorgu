package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Dork struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Query    string `json:"query"`
	Link     string `json:"link"`
}

func main() {
	http.HandleFunc("/generate", dorkHandler)
	fmt.Println("Dork Sorgu kullanıma hazır!")
	http.ListenAndServe(":8080", nil)
}

func dorkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Hedef eksik", http.StatusBadRequest)
		return
	}

	templates := []struct {
		Cat  string
		Name string
		Path string
	}{
		{"Giriş", "Yönetici Panelleri", "site:%s inurl:admin OR inurl:login"},
		{"Dosya", "PDF/Belge Sızıntıları", "site:%s filetype:pdf OR filetype:doc"},
		{"Dosya", "Excel/Veri Listeleri", "site:%s filetype:xls OR filetype:xlsx"},
		{"Sistem", "Konfigürasyon Dosyaları", "site:%s ext:env OR ext:conf OR ext:bak"},
		{"Sistem", "Log Kayıtları", "site:%s ext:log"},
		{"Hata", "SQL Hata Mesajları", "site:%s intext:\"sql syntax near\""},
		{"Hata", "PHP Hataları", "site:%s intext:\"Warning: mysql_connect()\""},
		{"Dizin", "Açık Dizinler (Index of)", "site:%s intitle:\"index of /\""},
		{"Gizli", "E-posta Adresleri", "site:%s intext:\"@%s\""},
		{"Dosya", "Metin Dosyaları", "site:%s filetype:txt OR filetype:sql"},
		{"Sistem", "Git Konfigürasyonu", "site:%s inurl:/.git/config -github.com"},
		{"Hata", "Django Debug Çıktısı", "site:%s \"DisallowedHost at /\" OR \"DEBUG = True\""},
		{"Dosya", "Kritik SQL Yedekleri", "site:%s filetype:sql \"MySQL dump\" OR \"pg_dump\""},
		{"Giriş", "Laravel Log Dosyası", "site:%s inurl:/storage/logs/laravel.log"},
		{"Gizli", "AWS Kimlik Bilgileri", "site:%s \"aws_access_key_id\" ext:php OR ext:js"},
		{"Dizin", "Hassas Backup Dizini", "site:%s intitle:\"index of\" \"backup\" | \"old\" | \"yedek\""},
		{"Web", "Potansiyel XSS Noktaları", "site:%s inurl:search.php?q= OR inurl:query="},
		{"Sistem", "Docker Yapılandırma", "site:%s \"docker-compose.yml\" intext:\"services:\""},
		{"Hata", "PHP PDO İstisnaları", "site:%s \"Uncaught PDOException: SQLSTATE\""},
		{"Giriş", "Yönetim Paneli (Brute Force Hedefi)", "site:%s inurl:/wp-login.php | inurl:/user/login"},
		{"Dosya", "Konfigürasyon (JSON)", "site:%s filename:config.json OR filename:settings.json"},
		{"Gizli", "RSA Private Key İfşası", "site:%s \"-----BEGIN RSA PRIVATE KEY-----\""},
		{"Hata", "Apache Sunucu Bilgisi", "site:%s \"Server at\" intitle:\"index of\""},
		{"Web", "Açık API Endpointleri", "site:%s inurl:/api/v1/ | inurl:/v2/api"},
		{"Sistem", "Environment (.env) Verisi", "site:%s \"DB_HOST=\" OR \"MAIL_PASSWORD=\""},
		{"Yedeklemeler", "Java Web Arşivi", "site:%s filetype:war"},
		{"Web", "Parametre Analizi (id=)", "site:%s inurl:id= OR inurl:page="},
		{"Sistem", "VNC Yapılandırma Dosyası", "site:%s filename:vnc.ini OR filename:vnc.conf"},
		{"Gizli", ".svn Kayıtları", "site:%s inurl:/.svn/entries"},
		{"Hata", "Yürütme Hatası", "site:%s intext:\"Runtime Error\" \"Server Error in\""},
		{"Dosya", "Shell Script Yedekleri", "site:%s filetype:sh OR filetype:bash"},
		{"Giriş", "login.php", "site:%s inurl:wp-login.php"},
		{"Dizin", "E-posta Arşiv Dizini", "site:%s intitle:\"index of\" \"mail\""},
		{"Sistem", "Nginx Konfigürasyonu", "site:%s filename:nginx.conf OR filename:proxy.conf"},
		{"Hata", "Java NullPointerException", "site:%s intext:\"java.lang.NullPointerException\""},
		{"Web", "Açık FTP Dizinleri", "site:%s inurl:ftp:// \"index of\""},
		{"Gizli", "SSH Key Sızıntısı", "site:%s \"id_rsa\" OR \"id_dsa\" ext:txt"},
		{"Dosya", "Kritik SQLite Veritabanı", "site:%s filetype:sqlite OR filetype:db"},
		{"Hata", "MySQL Bağlantı Uyarısı", "site:%s intext:\"Warning: mysql_connect()\""},
		{"Giriş", "Yönetim Paneli", "site:%s inurl:yonetim OR inurl:yonetici"},
	}

	var results []Dork
	for _, t := range templates {

		query := ""
		if t.Name == "E-posta Adresleri" {
			query = fmt.Sprintf(t.Path, domain, domain)
		} else {
			query = fmt.Sprintf(t.Path, domain)
		}

		searchURL := "https://www.google.com/search?q=" + url.QueryEscape(query)

		results = append(results, Dork{
			Category: t.Cat,
			Name:     t.Name,
			Query:    query,
			Link:     searchURL,
		})
	}

	json.NewEncoder(w).Encode(results)
}
