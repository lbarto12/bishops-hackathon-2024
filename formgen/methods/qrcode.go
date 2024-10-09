package methods

import (
	"crypto/sha256"
	"fmt"
	"formgen/postgres"
	"formgen/util"
	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"os"
	"slices"
)

func ResetDB() error {
	db, err := postgres.Database()
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM voter`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM voter_reg`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE polls SET votes = 0`)
	if err != nil {
		return err
	}
	return nil
}

func GenerateStaticVoters() error {
	//static hard coded salt for demo
	salt := "90p34ujtra3w4htfgaip3w4948hytq3p3948ythfaewihnflae8ryghpa4hetga98er8hgzergh"
	for _, voter := range voterStatic {
		var canVerify = make([]string, 3)
		candidates := []string{voter.can1, voter.can2, voter.can3}
		for _, candidate := range candidates {
			id := candidate + salt
			ha := sha256.Sum256([]byte(id))
			uid := fmt.Sprintf("%x", ha[:2])
			canVerify = append(canVerify, uid)
		}
		err := GenerateQRCode(voter.name, voter.healthCard, candidates, canVerify)
		if err != nil {
			return err
		}
		err = CreateVoter(voter.name, voter.healthCard, candidates)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateQRCode(name, card string, canUUID, canVerify []string) error {
	frontUrl, ok := os.LookupEnv("FRONTEND_URL")
	if !ok {
		return fmt.Errorf("env variable FRONTEND_URL not found")
	}
	err := os.Mkdir(fmt.Sprintf("out/%s-%s", name, card), os.ModePerm)
	if err != nil {
		return err
	}
	for j, candidate := range canUUID {
		qrc, err := qrcode.New(fmt.Sprintf("%s/verify/%s", frontUrl, candidate))
		if err != nil {
			return err
		}

		w, err := standard.New(fmt.Sprintf("out/%s-%s/%s_%s", name, card, canVerify[j], candidate))
		if err != nil {
			return err
		}
		if err = qrc.Save(w); err != nil {
			return err
		}
	}
	return nil
}

// GenerateVoters accepts an int 'n', and generates entries in the database for 'n' voters, and generates
// '3n' QR codes for users to scan to place their votes
func GenerateVoters(n int64) error {
	salt, ok := os.LookupEnv("CANDIDATE_SALT")
	if !ok {
		panic("SALT NOT SET")
	}

	for i := range n {
		name := fmt.Sprintf("FNAME%d LNAME%d", i, i)
		card := util.RandomHealthCard(8)

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

		for _, candidate := range candidates {
			canUUIDs = append(canUUIDs, candidate)
		}
		err := GenerateQRCode(name, card, canUUIDs, canVerify)
		if err != nil {
			return err
		}
		err = createRegVoter(name, card, canUUIDs, canVerify)
		if err != nil {
			return err
		}
		err = CreateVoter(name, card, canUUIDs)
		if err != nil {
			return err
		}
	}
	return nil
}
