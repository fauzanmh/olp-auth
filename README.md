# Online Learning Platform Auth
## Deskripsi
Untuk repo ini digunakan sebagai microservice authentication, terdapat fitur login, create user, dan delete user

## Cara Install
### Pertama 
Jalankan perintah dibawah ini pada cmd:

	go get -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
	go mod vendor && swag init
atau ketikkan (menggunakan make file)

    make install

### Selanjutnya
Copy main.example.json lalu ubah namanya menjadi
main.json dan terakhir setting database mysql

    "database": {
        "mysql": {
            "host": "localhost",
            "port": "3306",
            "dbname": "olp_auth",
            "user": "root",
            "password": ""
        }
    }

setelah itu buka repo berikut
https://github.com/fauzanmh/olp-migration-auth



### Cara Menjalankan
Jalankan perintah dibawah ini pada cmd:

    cd .. && go get -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
    go get github.com/vektra/mockery/v2/.../ && \
    cd ${PROJECT_NAME} && swag init
atau ketikkan (menggunakan make file)

    make run


### Dokumentasi API (Swagger)

    http://localhost:8101/api/swagger/index.html

### Login Admin

    {
        "username": "system",
        "password": "12345678",
        "provider": "admin"
    }