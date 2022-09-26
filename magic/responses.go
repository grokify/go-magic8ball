package magic

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var responses = []string{
	// positive
	"It is certain.",
	"It is decidedly so.",
	"Without a doubt.",
	"Yes definitely.",
	"You may rely on it.",

	"As I see it, yes.",
	"Most likely.",
	"Outlook good.",
	"Yes.",
	"Signs point to yes.",

	// neutral
	"Reply hazy, try again.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",

	// negative
	"Don't count on it.",
	"My reply is no.",
	"My sources say no.",
	"Outlook not so good.",
	"Very doubtful.",
}

var ErrResponseNotFound = errors.New("response not found")

func responseIndex(r string) (int, error) {
	for i, res := range responses {
		if r == res {
			return i, nil
		}
	}
	return -2, ErrResponseNotFound
}

func ResponseType(r string) (int, error) {
	idx, err := responseIndex(r)
	if err != nil {
		return -2, err
	}
	if idx < 10 {
		return 1, nil
	} else if idx < 15 {
		return 0, nil
	}
	return -1, nil
}

func Shake() (string, error) {
	idx, err := cryptoRandInt(len(responses))
	if err != nil {
		return "", err
	}
	return responses[idx], nil
}

func cryptoRandInt(max int) (int, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(num.Int64()), nil
}
