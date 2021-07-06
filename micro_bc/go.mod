module microbc

go 1.13

require (
	git.aizinger.com/aizinger/go-base v0.0.0-20210621075611-7c171f66a252
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.5.0
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	google.golang.org/genproto v0.0.0-20200527145253-8367513e4ece // indirect
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.24.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
