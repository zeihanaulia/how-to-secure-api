# Bagaimana cara mengamankan API

REST API sudah menjadi sesuatu yang umum digunakan beberapa tahun belakangan ini. Disini kita coba untuk membuktikan konsep konsep bagaimana cara mengamankan API. Kalau kita googling dengan keyword "how to secure api" mata gw langsung tertuju ke link ini [OWASP - REST Security cheat sheet](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/REST_Security_Cheat_Sheet.md). Di link tersebut terdapat "contekan" bagaimana seharusnya REST API dijaga. Berikut Listnya.

- [ ] HTTPS
- [ ] Access Control
- [ ] [JWT](#JWT)
- [ ] API Keys
- [ ] Restrict HTTP methods
- [ ] Input validation
- [ ] Validate content types
- [ ] Validate request content types
  - [ ] Send safe response content types
  - [ ] Management endpoints
- [ ] Error handling
- [ ] Security headers
- [ ] CORS
- [ ] Sensitive information in HTTP requests
- [ ] HTTP Return Code

## JWT

JSON Web Token [JWT](https://tools.ietf.org/html/rfc7519). JWT adalah struktur data dalam bentuk JSON, yang datanya dapat di klaim / diakui oleh pengguna, yang dapat digunakan untuk mengontrol akses. Dengan proses verifikasi tanda tangan (signature verification) yang menggunakan [cryptographic signature](https://www.tutorialspoint.com/cryptography/cryptography_digital_signatures.htm) atau message authentication code ([MAC](https://en.wikipedia.org/wiki/Message_authentication_code)).

### Kapan kita menggunakan JWT

Kalo kita googling lagi dengan keyword "json web token", gw langsung melihat website [https://jwt.io/](https://jwt.io/). Dibagian [introduction](https://jwt.io/introduction/) dijelaskan. Ada 2 kasus yang dapat menggunakan JWT, yaitu.

- Authorization
- Information Exchange

Dalam kasus penjagaan REST API, Kita ada dikasus pertama yaitu Authorization, dimana kita ingin memberikan otorisasi / wewenang kepada pengguna API agar mereka bisa mengaksesnya. Contoh:

Kita punya 2 endpoint:

- /api/login
- /api/profile

API public adalah endpoint yang semua pengguna bisa akses. Misal. landing page, login page, dll. 

API profile adalah endpoint yang hanya pengguna yang punya wewenang yang dapat mengakses.

Bagaimana cara endpoint profile tau kalau dia admin. Dalam setiap pengakseannya akan disertakan Authorization `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.PZLMJBT9OIVG2qgp9hQr685oVYFgRgWpcSPmNcw6y7M` dan disetiap endpoint yang membutuhkan otorisasi akan mengecek validitas dari token itu. Jika token itu valid maka dia dapat mengaksesnya.

### Struktur JSON Web Token

Bentuk token seperti itu ada artinya tidak? Ada. Informasinya masih ada di website yang sama yaitu di jwt.io bagian [introduction](https://jwt.io/introduction/). Jadi token terdiri dari 3 bagian yang dibatasi oleh (.) dot `xxxx.YYYYY.zzzzz`. Yang artinya:

- Header
- Payload
- Signature

#### Header

Header terdiri dari 2 bagian, yaitu algoritma dan type.

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

header ini akan diencoded ke Base64Url.

Ada list algoritmanya ga? mau tau ada apa aja. Bisa di cek di sini [rfc7518#section-3](https://tools.ietf.org/html/rfc7518#section-3)

Kalo type ada apa aja? liat disini [rfc7519#page-11](https://tools.ietf.org/html/rfc7519#page-11)

#### Payload

Payload ini adalah data yang bisa di klaim / diakui. Payload terdiri dari 2 jenis, registered, public dan private.

- Registered claims

  Ini adalah data yang sudah ada sebelumnya. [Listnya](https://tools.ietf.org/html/rfc7519#section-4.1):
  - iss (issuer) = siapa yang nerbitin jwt, misal `https://api.mycompany.com/api/login`
  - exp (expiration time) = masa kadaluarsa token, misal token hanya bisa digunakan selama 1 jam
  - sub (subject) = untuk mengidentifikasi pemegang token dan harus unik, misal username, email, userid
  - aud (audience) = untuk mengetahui token ini untuk siapa. misal website, mobile app 1, mobile app 2 atau id audience yang kita bikin
  - "nbf" (Not Before) Claim = token tidak dapat digunakan sebelum jam sekian, misal token dapat digunakan 15 lagi
  - "iat" (Issued At) Claim = untuk mengetahui kapan token diterbitin, misal `1555000690` = Thursday, April 11, 2019 4:38:10 PM
  - "jti" (JWT ID) Claim = jwt id, fungsinya misal kita mau menyian jwt id di datbase dan sewaktu waktu kita bisa melakukan pencabutan hak akses bisa dengan memvalidasi apakah jwt id ini di cabut?

- Public claims

Public claims adalah data data yang seharusnya terdaftar disini [IANA JSON Web Token Registry](https://www.iana.org/assignments/jwt/jwt.xhtml). Jadi ini adalah registry atau data-data yang dibutuhkan untuk public, sehingga tidak ada salah pengertian mengenai data yang disajikan.

- Private claims

Private / Custom claims ini adalah field yang diberikan dan disetujui oleh kedua belah pihak penyedia dan pengguna

Sama seperti header, data payload ini akan di encode ke  Base64Url

#### Signature

Signature adalah gabungan dari Header + . + Payload + secret yang dibungkus didalam algoritma. Misal:

```bash
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret
)
```

Secret adalah kata rahasia yang digunakan untuk memverifikasi. kata rahasia ini biasanya hanya diketahui oleh penyedia. Pengguna hanya menerima token yang sudah berupa jwt.

#### Bagaimana cara kerja JWT

![image info](./images/how-jwt-works.png)

1. Client Meminta token, misal dengan api login
2. Disisi server akan memproses pembuatan JWT
3. Mengembalikan JWT ke client
4. Token diletakan di header: Authorization: bearer {token}, untuk mengakses api /profile
5. Server akan memproses validasi verifikasi
6. Pengembalian response ke Client
   1. (6.a) Jika berhasil akan mengambalikan response berhasil dan http code 200
   2. (6.b) Jika gagal akan mengembalikan response Unauthorize dan http code 401

### Kenapa kita harus menggunakan JWT

Bisa lihat divideo ini

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/67mezK3NzpU/0.jpg)](https://www.youtube.com/watch?v=67mezK3NzpU)

TODO: list kenapa menggunakan jwt dari video diatas dengan tulisan