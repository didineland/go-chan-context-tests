protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=require_unimplemented_servers=false:. 
--go-grpc_opt=paths=source_relative \
protofiles/data_streaming/streamingData.proto


protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=true:.  --go-grpc_opt=paths=source_relative protofiles\meteo-streaming\meteo_streaming.proto