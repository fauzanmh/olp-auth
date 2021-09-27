# INSTALL PROGRAM
install:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
	go mod vendor && swag init

# RUN PROGRAM
run:
	air -c config/.air.toml