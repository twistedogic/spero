package poll

import (
	"context"
	"testing"
	"time"

	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/twistedogic/spero/pkg/message"
)

type testCase struct {
	pointer       int
	payload, want [][]byte
}

func (tc *testCase) next() bool {
	tc.pointer++
	if max := len(tc.payload) - 1; tc.pointer > max {
		return false
	}
	return true
}

func (tc *testCase) check(t *testing.T, msg *service.Message) {
	want := string(tc.want[tc.pointer])
	b, err := msg.AsBytes()
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want:%s, got:%s", want, got)
	}
}

func (tc *testCase) poll(ctx context.Context) (*service.Message, error) {
	content := tc.payload[tc.pointer]
	return message.NewContentHashMessage(content), nil
}

func TestPoller(t *testing.T) {
	cases := map[string]testCase{
		"base": testCase{
			payload: [][]byte{
				[]byte(`a`),
				[]byte(`b`),
				[]byte(`c`),
				[]byte(`d`),
			},
			want: [][]byte{
				[]byte(`a`),
				[]byte(`b`),
				[]byte(`c`),
				[]byte(`d`),
			},
		},
	}
	for name := range cases {
		tc := cases[name]
		t.Run(name, func(t *testing.T) {
			poller := New(tc.poll, time.Millisecond)
			ctx := context.TODO()
			if err := poller.Connect(ctx); err != nil {
				t.Fatal(err)
			}
			for {
				got, _, err := poller.Read(ctx)
				if err != nil {
					t.Fatal(err)
				}
				tc.check(t, got)
				if !tc.next() {
					break
				}
			}
			if err := poller.Close(ctx); err != nil {
				t.Fatal(err)
			}
		})
	}
}
