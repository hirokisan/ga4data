CREDENTIAL_PATH := "./credential.json"
OUTPUT_PATH := "../../const_columns_gen.go"

gen:
	CREDENTIAL_PATH=${CREDENTIAL_PATH} OUTPUT_PATH=${OUTPUT_PATH} go run main.go
	go fmt ${OUTPUT_PATH}
