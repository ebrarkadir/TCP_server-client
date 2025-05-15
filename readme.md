# 🔌 TCP Server-Client Communication System (Go)

Bu proje, Go dili ile geliştirilmiş bir **TCP sunucu-istemci** uygulamasıdır. Sistem, özel bir binary protokol kullanarak iki uç arasında veri iletimi sağlar. Performans ve sistem verimliliği için çeşitli optimizasyon teknikleri entegre edilmiştir.

---

## ⚙️ Özellikler

- 🧱 **TCP Server / TCP Client**
- ✉️ **Custom Binary Protocol**
- 📈 **Profiling ve Performans Ölçümü**
- 🚀 **Optimizasyon Teknikleri:**
  - `pprof` ile CPU profili
  - `sync.Pool` & Worker Pool
  - Epoll & Netpoll (non-blocking IO)
  - `SO_REUSEPORT` desteği
  - eBPF entegrasyon yapısına uygun mimari
  - ARM uyumluluğu

---

## 🔗 İletişim Yapısı

```
client <------------------> server
          TCP stream
         read / write
```

---

## 🧱 Protokol Formatı

```
0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string
type    | length  | data
```

Mesaj 3 parçadan oluşur:

- `type`: Mesaj türü (örneğin JSON, TEXT, XML)
- `length`: Veri uzunluğu
- `data`: İçerik (string)

```go
const (
	MessageTypeJSON = 1
	MessageTypeText = 2
	MessageTypeXML  = 3
)
```

---

## 🚀 Kullanım

### Sunucu Başlatma
```bash
cd server
go run main.go
```

### İstemci Başlatma
```bash
cd client
go run main.go
```

İstemci, sunucuya 150.000 özel formatta mesaj gönderir.

---

## 📊 Performans

| Profil Durumu   | Süre          |
|------------------|---------------|
| `with pprof`     | 1.1420958 s   |
| `without pprof`  | 1.0709970 s   |

> Profil açıkken CPU ölçümleri `cpu.out` dosyasına yazılır.

---
