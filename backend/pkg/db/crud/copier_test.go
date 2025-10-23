package crud_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/flag"
	"majinyao.cn/my-app/backend/pkg/test"
)

type ABC struct {
	model.ModelReserved
	DeleteAt   gorm.DeletedAt
	Id         datatype.Id
	Oid        datatype.Oid
	Password   datatype.Password
	Secret     datatype.Secret
	Sensitive  datatype.Sensitive
	Id2        *datatype.Id
	Oid2       *datatype.Oid
	Password2  *datatype.Password
	Secret2    *datatype.Secret
	Sensitive2 *datatype.Sensitive
	Flag       sql.RawBytes
}

func (a *ABC) FlagBase64(b64 string) {
	f, _ := flag.FromBase64(b64)
	a.Flag = sql.RawBytes(f.ToBytes())
}

type Reserved struct {
	Reserved bool `json:"reserved" doc:"Is System Reserved"`
}

type DEF struct {
	Reserved
	DeleteAt   time.Time
	Id         string
	Oid        string
	Password   string
	Secret     string
	Sensitive  string
	Id2        string
	Oid2       string
	Password2  string
	Secret2    string
	Sensitive2 string
	FlagBase64 string
}

func (a *DEF) Flag(bs sql.RawBytes) {
	a.FlagBase64 = flag.FromBytes(bs).ToBase64()
}

func TestCopier(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	var id2 datatype.Id = 1234
	var oid2 datatype.Oid = []int64{4, 5, 6}
	var sensitive2 datatype.Sensitive = "654321"
	a := ABC{
		ModelReserved: model.ModelReserved{
			Reserved: sql.NullBool{Bool: true, Valid: true},
		},
		DeleteAt: gorm.DeletedAt{
			Time:  now,
			Valid: false,
		},
		Id:  123,
		Oid: []int64{1, 2, 3},
		Password: datatype.Password{
			Data: "123456",
		},
		Secret: datatype.Secret{
			Data: "123456",
		},
		Sensitive: "123456",
		Id2:       &id2,
		Oid2:      &oid2,
		Password2: &datatype.Password{
			Data: "654321",
		},
		Secret2: &datatype.Secret{
			Data: "654321",
		},
		Sensitive2: &sensitive2,
		Flag:       flag.FromBytes([]byte{1, 2, 3}).ToBytes(),
	}

	var b DEF
	err := copier.CopyWithOption(&b, &a, crud.GlobalCopierOption)
	if err != nil {
		t.Fatalf("CopyWithOption failed: %v", err)
	}
	b.FlagBase64 = "QAAAAAAAAAAAAAAAAAAA"
	err = copier.CopyWithOption(&a, &b, crud.GlobalCopierOption)
	if err != nil {
		t.Fatalf("CopyWithOption failed: %v", err)
	}
}
