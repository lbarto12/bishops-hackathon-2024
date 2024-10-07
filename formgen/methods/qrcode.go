package methods

import (
	"crypto/sha256"
	"fmt"
	"formgen/util"
	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"os"
	"slices"
)

func GenerateVoters(n int64) error {

	frontUrl, ok := os.LookupEnv("FRONTEND_URL")
	if !ok {
		return fmt.Errorf("env variable FRONTEND_URL not found")
	}

	salt, ok := os.LookupEnv("CANDIDATE_SALT")
	if !ok {
		panic("SALT NOT SET")
	}

	for i := range n {
		name := fmt.Sprintf("FNAME%d LNAME%d", i, i)
		card := util.RandomHealthCard(8)
		err := os.Mkdir(fmt.Sprintf("out/%s-%s", name, card), os.ModePerm)
		if err != nil {
			return err
		}

		var canUUIDs []string

		candidates := make([]string, 0)
		var canVerify []string
	restart:
		candidates = []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
		canVerify = make([]string, 0)
		for _, candidate := range candidates {
			id := candidate + salt
			ha := sha256.Sum256([]byte(id))
			uid := fmt.Sprintf("%x", ha[:2])
			if slices.Contains(canVerify, uid) {
				goto restart
			}
			canVerify = append(canVerify, fmt.Sprintf("%x", ha[:2]))
		}

		for j, candidate := range candidates {

			uid := fmt.Sprintf("%d%s", j+1, candidate[1:])
			canUUIDs = append(canUUIDs, uid)

			qrc, err := qrcode.New(fmt.Sprintf("%s/verify/%s", frontUrl, uid))
			if err != nil {
				return err
			}

			w, err := standard.New(fmt.Sprintf("out/%s-%s/%s-%s", name, card, canVerify[j], uid))
			if err != nil {
				return err
			}

			if err = qrc.Save(w); err != nil {
				return err
			}
		}

		err = CreateVoter(name, card, canUUIDs, canVerify)
		if err != nil {
			return err
		}
	}
	return nil
}
