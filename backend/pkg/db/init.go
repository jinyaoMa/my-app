package db

import (
	"my-app/backend/pkg/code"
	"my-app/backend/pkg/enc"
	"my-app/backend/pkg/id"
)

var (
	idGenerator   id.IID
	codeGenerator code.ICode
	dataCipher    enc.ICipher
)

func init() {
	var err error
	idGenerator, err = id.Default()
	if err != nil {
		panic(err)
	}
	codeGenerator = code.New()
	dataCipher = enc.NewAesWithSalt("jinyaoMa")
}

func IdGenerator() id.IID {
	return idGenerator
}

func CodeGenerator() code.ICode {
	return codeGenerator
}

func DataCipher() enc.ICipher {
	return dataCipher
}

func SetIdGenerator(x id.IID) {
	idGenerator = x
}

func SetCodeGenerator(x code.ICode) {
	codeGenerator = x
}

func SetDataCipher(x enc.ICipher) {
	dataCipher = x
}
