package util

import "sync"

// RoundRobin 是一个实现轮询算法的结构体
type RoundRobin struct {
	keys []string
	idx  int
	mu   sync.Mutex
}

var APP_KEYS = []string{
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI5MWViNTdkMC1kNmY1LTAxM2ItNGU3Yy0wMjJlZGE2MTJmNmMiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjg0MzM3MzgyLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii1mOTAzNGM2NC0wY2MyLTQxNmUtYWUyMS1hZWQ2YjkzOTYxYTYifQ.v29Ut6ma2q6On6HCQSiKn94QQYtdNn7TBHI0Nj38GCM",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI3ZDIwNTVjMC1kNmY1LTAxM2ItNGU3YS0wMjJlZGE2MTJmNmMiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjg0MzM3MzQ4LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii00ZmU3ZDczYi0yN2FjLTQyN2UtODQwOS01YzBmODBkMTgyMGYifQ.S8CgAGI91RMJRggjlmaiUzWmE9k39NpNbzA4yLbH68E",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJjMjdiZDAzMC1kN2I4LTAxM2ItYWFjNC00YWQ1NTRmZjEzNWIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjg0NDIxMjE2LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii05Y2U5N2UyOS1lNzM2LTRkMjgtOGM3ZC02ODE3OWUxODlmMDYifQ.gfZhUwIdE0uFwI_1nBLYI9lBLlmxrtEWEfAJzO6-pJU",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI5YTUyMmUxMC1kN2I4LTAxM2ItYzdlOC00Mjg1NWU5MDJkNTYiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjg0NDIxMTQ4LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii02MTYzOWU0MC1iN2I0LTRhYTItOGNjYS1mMmYxZWE2MTBmNmIifQ.O4psO_Pcn9wv8wEnCh3w9GzQC0jabi4gE124RR-v3Os",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIxZThkM2Y4MC0yYjQ3LTAxMzgtZjZhZC0wYmZlODVlOGVlNjQiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNTgxMDE4MzE1LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjI1NDcyNzY5Ny1xcS1jIn0.7gpNhq3LtRT_tyDoMzdeQgtueKYE3L-T7iTTkPgD5MA",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJkYTNhOGM1MC1lMTE4LTAxM2EtMWU0NS0xOTMxZmRiMDY1M2IiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjU3MzA0NTUwLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6ImZtIn0.QLU9I6oBR_MJXpApWm0bUisMiqELUe8AioSwyj-BZ5U",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI5YzFkNDE2MC1hZjc3LTAxM2ItNTdhMy0wYWJiMzhhMGNmM2MiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjc5OTk1MTg4LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii0xZWIxZjJhMi02MzVmLTQwOTAtODVjZC1mNTEzNjIxMjRjNTYifQ.Svhl8lRfHO2ff2HRva9uEbExDHJIibESi6owyhUAiRo",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMDk3MzJhMC1hZjczLTAxM2ItYTAwNy03MzRjN2QzMWQyMzYiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjc5OTkzMjg5LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii04MWVkNTdhYi1mYWRiLTRhMjctOWI4NS05MzkyOTZjZWQ3MzIifQ.HcToEZIkJ4x1ShTtQtNWYNSQ1KM33tQpH203UbDG8wM",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIyNjc0Y2MyMC1hZjczLTAxM2ItYTAwNS03MzRjN2QzMWQyMzYiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjc5OTkzMjcyLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii0wZGI0MDYyZS05OThmLTQyN2EtYjRjNy0zODcwNTFiMTA0YjEifQ.XBRrvG79QCYpHQf44U_CNiapBLi3sC_D8jM5F4WBWeY",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJiZjVkYTA3MC1hZjcxLTAxM2ItYTAwMy03MzRjN2QzMWQyMzYiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjc5OTkyNjcwLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6ImZtLWxlYWd1ZSJ9.tqF9YUwags7-TVspYp-tHiIJBiK0wDE-5aQea9JJ0_I",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJjYTAyZjliMC1hZmM3LTAxM2ItYTAwYy03MzRjN2QzMWQyMzYiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNjgwMDI5NjI0LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Im5ld2dhbWUifQ.ceIpNQlvUFBHbcmdOPQQ_4Poy5EHQTt4io6hbzA_9Zw",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJmMjZjZWNhMC04YmI3LTAxM2MtNTVlYS00MmI5M2U5NTM4ZjAiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MjEyMDc2LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6ImZtLWFwcGx5In0.xNkQ_GsMgoDPjYWy2XNngleze-94AAnmmNGx6qAQzHc",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJlZmYyYjRkMC02YmZlLTAxM2MtYzg5My01MjFkYWMzN2YxYTEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzAwNzI0MTI5LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii03ZDQyNmQ5Zi05YmEzLTQyOTctODJkZS0xMjQ2NTEyMjhjMjQifQ.vAZSK0qzCw6nQzroK58YUdeUvkfBWsDw3DE6udEwnZs",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJiMDIzYWRlMC04ZDQzLTAxM2MtYTdjYS0zZTlkOTIyZjEwNGEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MzgyMDQ2LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii0xMDg0ZjJkYy1lZmY5LTQxYTYtYTllMC1hODczNjZlMzI3ZDMifQ.2RnB1hdvGLTUE2XrKNIhfr1gWTaDOjpdVVl4amOmt0I",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJhOTRjNDg3MC04ZDQzLTAxM2MtN2E3ZS02ZTVhNmFhZGIzMTMiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MzgyMDM0LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii1hZDdlZmFiYy1lYmI4LTQzZTUtOTU0ZS1hZTc4ZDhmODlhZDkifQ.CGxf67FO5wuY2vX-IW_bdjuCxTAaM5BxnVJpHR27Bq0",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIyM2Q3MjNkMC04ZDQ0LTAxM2MtNTZkMy03YTg0MDBmODVjMmUiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MzgyMjQwLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii04OTNmNzFkNC01MjFiLTQwMTktOGYzYS1hYjlhMTYyNTBmYjUifQ.Zul9ujL6piBfj-dspN6WU8dbhfVKNyzalYW8YaId5dA",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIxZTI1ZjcyMC04ZDQ0LTAxM2MtMDg0NS0yMjIxMDUzM2M3NWEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MzgyMjMwLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii1lZDY0ZGVlYi1iNWE0LTRlZDQtOTllZS1iZGVlYjRmYjJmMzMifQ.racm-zwmf8oMmLarovguvFq9LRjmkX7kYgzbm20yPnA",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJlN2Y1NTNjMC04ZDQ0LTAxM2MtMDg0OS0yMjIxMDUzM2M3NWEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MzgyNTY5LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii1hN2IxZDdlNS1mMDM3LTRjNTAtOGE0Mi03YTAxN2UzMTQwNjUifQ.nnBc_g8pktIGcysKTcShOaB6TMTspvOr9PzMf8f_0hs",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJlMzc1MDI5MC04ZDQ0LTAxM2MtMDg0Ny0yMjIxMDUzM2M3NWEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0MzgyNTYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Ii04YzdhMmE1Yy03MGRjLTRlYjUtYjc0My1hOTc2ZTM0NjU1ODcifQ.YcM-je40wLjew-4ckiRi58GyFZIsMBvfwK06vXLS9Eg",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIxODg1NGRhMC1kOTM5LTAxM2MtM2U4Yy00YWU0ZjgxNmFiMmMiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzEyNzMzNzg1LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6ImZtX2xlYWd1ZSJ9.nW-304SsXAC9__GVTyN3CuzB97TBwlFn_M95ecrqxgE",
	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJhYmZkNmQ2MC1kYjhmLTAxM2MtNzU5OC02NmY5NjQzNDM5NzQiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzEyOTkwODcxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6ImZtX2xlYWd1ZV8yIn0.0sE_EhS6Owoz4VzZA4qz7CTaC6M8lTHLMzlMqc1GAoc",
}
var APP_KEY *RoundRobin

func init() {
	APP_KEY = NewRoundRobin()
}

// NewRoundRobin 初始化并返回一个 RoundRobin 结构体
func NewRoundRobin() *RoundRobin {
	return &RoundRobin{
		keys: APP_KEYS,
		idx:  0,
	}
}

// GetNext 返回下一个要使用的 key
func (r *RoundRobin) GetNext() string {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := r.keys[r.idx]
	r.idx = (r.idx + 1) % len(r.keys)
	return key
}
