package base62

import (
	"math"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const b = 62

// Func encode the given db ID to a base62 string
func ToBase62(num int) string {
	r := num % b
	res := string(base[r])
	div := num / b
	q := int(math.Floor(float64(div)))

	for q != 0 {
		r = q % b
		temp := q / b
		q = int(math.Floor(float64(temp)))
		res = string(base[int(r)]) + res
	}

	return string(res)
}

// Func decodes a given base62 string to db ID
func ToBase10(str string) int {
	res := 0
	for _, r := range str {
		res = (b * res) + strings.Index(base, string(r))
	}
	return res
}

// curl -X POST \
//  http://localhost:9000/v1/short \
//  -H 'cache-control: no-cache' \
//  -H 'content-type: application/json' \
//  -d '{
//  "url":
// "https://www.dropbox.com/scl/fi/9fgfciz4wdowwqm979rv8/How-to-Build-Performant-Servers-with-GRPC-and-Protocol-Buffers.paper?rlkey=130ljrh24zgp9zcbjsawzyemh&dl=0"
// }'
