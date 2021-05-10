package reader

import (
	"bufio"
	"os"
)

func ReadStdin() []byte {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 10000)

	for {
		n, err := r.Read(buf)
		if n == 0 {
			break
		} // TODO: 確保したスライスを超過した場合の条件分岐を追加
		if err != nil {
			panic(err)
		}
	}

	return buf
}
