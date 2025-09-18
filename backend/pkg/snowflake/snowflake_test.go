package snowflake_test

import (
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestSnowflake(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	epoch := time.Now()

	tries := 32
	ids := make(chan int64, tries)
	var g errgroup.Group
	for i := 1; i <= tries; i++ {
		nodeId := i
		g.Go(func() error {
			gen, err := snowflake.New(snowflake.Options{
				Epoch:    epoch,
				NodeBits: 7,
				StepBits: 14,
				NodeId:   nodeId,
			})
			if err != nil {
				return err
			}

			ids <- gen.Generate()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}

	idMap := make(map[int64]struct{})
	for {
		select {
		case id := <-ids:
			if _, exist := idMap[id]; exist {
				t.Fatal("there should not be duplicate ids")
			}
			idMap[id] = struct{}{}
		default:
			return
		}
	}
}
