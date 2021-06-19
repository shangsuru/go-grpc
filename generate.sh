protoc calculator/calculator.proto --go_out=plugins=grpc:./calculator
protoc prime_number/prime_number.proto --go_out=plugins=grpc:./prime_number
protoc average/average.proto --go_out=plugins=grpc:./average
protoc maximum/maximum.proto --go_out=plugins=grpc:./maximum
