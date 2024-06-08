PROTOC_IMG_PATH = ""
PROTO_REPOSITORY_PATH = ""
PROTO_DIRECTORY_PATH = ""

.PHONY: toy
toy: $(SRCS)
	go build -o toy cmd/main.go

.PHONY: dev_server
dev_server: $(SRCS)
	gin -p 3005 -a 8080 -x frontend_ops/ -i -d cmd run dev_server

