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

			qrc, err := qrcode.New(fmt.Sprintf("http://207.162.100.117:5173/verify/%s", uid))
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
