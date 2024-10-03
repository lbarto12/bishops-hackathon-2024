package methods

import (
	"fmt"
	"formgen/util"
	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"os"
)

func GenerateVoters(n int64) error {

	frontUrl, ok := os.LookupEnv("FRONTEND_URL")
	if !ok {
		return fmt.Errorf("env variable FRONTEND_URL not found")
	}

	for i := range n {
		name := fmt.Sprintf("FNAME%d LNAME%d", i, i)
		card := util.RandomHealthCard(8)
		err := os.Mkdir(fmt.Sprintf("out/%s-%s", name, card), os.ModePerm)
		if err != nil {
			return err
		}

		var canUUIDs []string

		for j, candidate := range []string{uuid.New().String(), uuid.New().String(), uuid.New().String()} {

			uid := fmt.Sprintf("%d%s", j+1, candidate[1:])
			canUUIDs = append(canUUIDs, uid)

			qrc, err := qrcode.New(fmt.Sprintf("%s/verify/%s", frontUrl, uid))
			if err != nil {
				return err
			}

			w, err := standard.New(fmt.Sprintf("out/%s-%s/%s", name, card, uid))
			if err != nil {
				return err
			}

			if err = qrc.Save(w); err != nil {
				return err
			}
		}

		err = CreateVoter(name, card, canUUIDs)
		if err != nil {
			return err
		}
	}
	return nil
}
