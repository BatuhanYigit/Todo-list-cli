**Proje:** Todo List

**Versiyon:** 0.1

**Tarih:** 2023-07-20

**Açıklama:**

Bu proje, kullanıcıların görevlerini yönetmelerine yardımcı olan bir todo list uygulamasıdır. Uygulama, sqlite3 veya txt dosyası olmak üzere iki farklı veritabanı seçeneği sunar.

**Kurulum:**

Uygulamayı çalıştırmak için öncelikle aşağıdaki gereksinimleri karşıladığınızdan emin olun:

* Go 1.18 veya üzeri
* sqlite3 veya txt dosyası

Uygulamayı çalıştırmak için aşağıdaki komutu kullanın:

```
go run main.go
```

**Kullanım:**

Uygulamayı kullanmaya başlamak için aşağıdaki adımları izleyin:

1. **Veritabanı seçin:**

```
**Veritabanı seçin:**
1 - sqlite3
2 - txt dosyası
```

2. **Bir görev ekleyin:**

```
**Bir görev ekleyin:**
Görev adı: Alışveriş
Durum: todo
```

3. **Tüm görevleri listeleyin:**

```
**Tüm görevleri listeleyin:**

ID | Görev | Durum
---|---|---|
1 | Alışveriş | todo
```

4. **Bir görevi güncelleyin:**

```
**Bir görevi güncelleyin:**
Görev ID: 1
Yeni durum: in progress
```

5. **Bir görevi silin:**

```
**Bir görevi silin:**
Görev ID: 1
```

**Örnek:**

```
**Todo List V0.1**

**1 - Database seçin**
**2 - Bir görev ekleyin**
**3 - Tüm görevleri listeleyin**
**4 - Bir görevi güncelleyin**
**5 - Bir görevi silin**
```

**Ek özellikler:**

* İleride uygulamaya aşağıdaki özellikleri eklemeyi planlıyorum:
    * Kullanıcı arayüzünü geliştirme
    * Görevlere tarih ve zaman ekleme
    * Görevlere hatırlatıcı ekleme
    * Birden çok kullanıcı desteği
