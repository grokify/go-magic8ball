package magic8ball

import (
	"errors"
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

// ErrResponseNotFound is the error when a supplied 8 ball response isn't found.
var ErrResponseNotFound = errors.New("response not found")

func responseIndex(r string) (int, error) {
	for i, res := range responses {
		if r == res {
			return i, nil
		}
	}
	return -2, ErrResponseNotFound
}

// ResponseType returns whether a response is positive, neutral, or negative.
// Values returned are `1` for positive, `0` for neutral, and `-1` for negative.
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

// Shake returns a Magic 8 Ball response. It returns an error if `crypto/rand` fails.
func Shake() string {
	return responses[Intn(uint(len(responses)))]
}
