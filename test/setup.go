package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/soracom/soracom-sdk-go"
)

const (
	defaultSandboxEndpoint = "https://api-sandbox.soracom.io"
	nSIM                   = 5
)

var (
	apiClient          *soracom.APIClient
	createdSubscribers []soracom.CreatedSubscriber
)

func main() {
	email := flag.String("email", "", "email address to create a sandbox account")
	password := flag.String("password", "", "password to create a sandbox account")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	if email == nil || *email == "" {
		fmt.Println("An argument 'email' is required.")
		os.Exit(1)
	}

	if password == nil || *password == "" {
		fmt.Println("An argument 'password' is required.")
		os.Exit(1)
	}

	apiClient = soracom.NewAPIClient(&soracom.APIClientOptions{
		Endpoint: defaultSandboxEndpoint,
	})
	if verbose != nil && *verbose {
		apiClient.SetVerbose(true)
	}

	err := setup(*email, *password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setup(email, password string) error {
	rand.Seed(time.Now().Unix())

	err := signup(email, password)
	if err != nil {
		return err
	}

	err = auth(email, password)
	if err != nil {
		return err
	}

	err = registerPaymentMethod()
	if err != nil {
		return err
	}

	// auth again to update token
	err = auth(email, password)
	if err != nil {
		return err
	}

	createdSubscribers = make([]soracom.CreatedSubscriber, 0, nSIM)
	for i := 0; i < nSIM; i++ {
		s, err := apiClient.CreateSubscriber()
		if err != nil {
			return err
		}
		createdSubscribers = append(createdSubscribers, *s)
	}

	err = registerSubscribers()
	if err != nil {
		return err
	}

	return nil
}

func signup(email, password string) error {
	err := apiClient.CreateOperator(email, password)
	if err != nil {
		return err
	}

	authKeyID := os.Getenv("SORACOM_AUTHKEY_ID_FOR_TEST")
	if authKeyID == "" {
		return errors.New("SORACOM_AUTHKEY_ID_FOR_TEST env var is required")
	}
	authKey := os.Getenv("SORACOM_AUTHKEY_FOR_TEST")
	if authKey == "" {
		return errors.New("SORACOM_AUTHKEY_FOR_TEST env var is required")
	}
	token, err := apiClient.GetSignupToken(email, authKeyID, authKey)
	if err != nil {
		return err
	}

	err = apiClient.VerifyOperator(token)
	if err != nil {
		return err
	}

	return nil
}

func auth(email, password string) error {
	err := apiClient.Auth(email, password)
	if err != nil {
		return err
	}

	return nil
}

func registerPaymentMethod() error {
	wp := &soracom.PaymentMethodInfoWebPay{
		Cvc:         "123",
		ExpireMonth: 12,
		ExpireYear:  getNextYear(),
		Name:        "SORAO TAMAGAWA",
		Number:      "4242424242424242", // https://webpay.jp/docs/mock_cards
	}
	err := apiClient.RegisterPaymentMethodWebPay(wp)
	if err != nil {
		return err
	}

	return nil
}

func getNextYear() int {
	return time.Now().AddDate(1, 0, 0).Year()
}

func registerSubscribers() error {
	for i, cs := range createdSubscribers {
		o := soracom.RegisterSubscriberOptions{
			RegistrationSecret: cs.RegistrationSecret,
			Tags:               soracom.Tags{},
		}
		if i%3 == 0 {
			o.Tags["soracom-sdk-go-test"] = "foo"
		}
		if i%3 == 1 {
			o.Tags["soracom-sdk-go-test"] = "hoge"
		}
		if i%3 == 2 {
			o.Tags["soracom-sdk-go-test"] = "beam-stats"
		}
		_, err := apiClient.RegisterSubscriber(cs.IMSI, o)
		if err != nil {
			return err
		}
		if i%4 == 0 {
			_, err := apiClient.ActivateSubscriber(cs.IMSI)
			if err != nil {
				return err
			}
		}
		if i%4 == 1 {
			_, err := apiClient.DeactivateSubscriber(cs.IMSI)
			if err != nil {
				return err
			}
		}
		if i%5 == 0 {
			_, err := apiClient.UpdateSubscriberSpeedClass(cs.IMSI, "s1.minimum")
			if err != nil {
				return err
			}
		}
		if i%5 == 1 {
			_, err := apiClient.UpdateSubscriberSpeedClass(cs.IMSI, "s1.slow")
			if err != nil {
				return err
			}
		}
		if i%5 == 2 {
			_, err := apiClient.UpdateSubscriberSpeedClass(cs.IMSI, "s1.standard")
			if err != nil {
				return err
			}
		}
		if i%5 == 3 {
			_, err := apiClient.UpdateSubscriberSpeedClass(cs.IMSI, "s1.fast")
			if err != nil {
				return err
			}
		}

		for j := 0; j < 10; j++ {
			t := time.Now().AddDate(0, 0, -10*j)
			ts := t.UnixNano() / 1000 / 1000
			err := apiClient.InsertAirStats(cs.IMSI, generateDummyAirStats(ts))
			if err != nil {
				return err
			}
		}

		for k := 0; k < 10; k++ {
			t := time.Now().AddDate(0, 0, -10*k)
			ts := t.UnixNano() / 1000 / 1000
			err := apiClient.InsertBeamStats(cs.IMSI, generateDummyBeamStats(ts))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getRandomSpeedClass() soracom.SpeedClass {
	switch rand.Intn(4) {
	case 0:
		return soracom.SpeedClassS1Minimum
	case 1:
		return soracom.SpeedClassS1Slow
	case 2:
		return soracom.SpeedClassS1Standard
	case 3:
		return soracom.SpeedClassS1Fast
	}
	return soracom.SpeedClassS1Standard
}

func generateDummyAirStats(ts int64) soracom.AirStats {
	ub := rand.Intn(1000000)
	up := ub / (rand.Intn(100) + 1)
	db := rand.Intn(1000000)
	dp := db / (rand.Intn(100) + 1)
	t := make(map[soracom.SpeedClass]soracom.AirStatsForSpeedClass)
	t[getRandomSpeedClass()] = soracom.AirStatsForSpeedClass{
		UploadBytes:     uint64(ub),
		UploadPackets:   uint64(up),
		DownloadBytes:   uint64(db),
		DownloadPackets: uint64(dp),
	}
	return soracom.AirStats{
		Unixtime: uint64(ts),
		Traffic:  t,
	}
}

func getRandomBeamType() soracom.BeamType {
	switch rand.Intn(11) {
	case 0:
		return soracom.BeamTypeInHTTP
	case 1:
		return soracom.BeamTypeInMQTT
	case 2:
		return soracom.BeamTypeInTCP
	case 3:
		return soracom.BeamTypeInUDP
	case 4:
		return soracom.BeamTypeOutHTTP
	case 5:
		return soracom.BeamTypeOutHTTPS
	case 6:
		return soracom.BeamTypeOutMQTT
	case 7:
		return soracom.BeamTypeOutMQTTS
	case 8:
		return soracom.BeamTypeOutTCP
	case 9:
		return soracom.BeamTypeOutTCPS
	case 10:
		return soracom.BeamTypeOutUDP
	}
	return soracom.BeamTypeInHTTP
}

func generateDummyBeamStats(ts int64) soracom.BeamStats {
	t := make(map[soracom.BeamType]soracom.BeamStatsForType)
	t[getRandomBeamType()] = soracom.BeamStatsForType{
		Count: uint64(rand.Intn(1000)),
	}
	return soracom.BeamStats{
		Unixtime: uint64(ts),
		Traffic:  t,
	}
}
