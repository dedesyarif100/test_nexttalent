Cara Running Rest API :
1. Masuk Folder Document pada project
2. Running SQL Dengan Laragon, link download https://laragon.org/download/index.html
3. Input Configurasi
    1. Input Hostname / IP : 127.0.0.1
    2. User : root
    3. Password : root
    4. Port : 43306

4. Import database nexttalent_test.sql
5. Buka Project Rest API, running go run main.go
6. Import Collection Postman
7. Jalankan Rest API
8. Port pada project, berjalan pada port 8080
9. Contoh hit Rest API :
    http://localhost:8080/GetCountry/dede
    http://localhost:8080/GetCurrentTime/?timezone=Asia/Jakarta
