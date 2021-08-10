all: one
.PHONY: all

one:
	git add .
	git commit -m "sdsd"
	git push

update:
	go get github.com/kirigaikabuto/common16@main
	go get github.com/kirigaikabuto/users16@main

run:
	go get github.com/kirigaikabuto/common16@main
	go get github.com/kirigaikabuto/users16@main
	go build
	./16