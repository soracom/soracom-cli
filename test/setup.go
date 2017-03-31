package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/soracom/soracom-sdk-go"
)

const (
	defaultSandboxEndpoint = "https://api-sandbox.soracom.io"
	nSIM                   = 5
	payJPPublishableKey    = "pk_test_5b4816eeedff49691d906902"
	stripePublishableKeyJP = "pk_test_5xUDApYcCbvIEZ2mcemGgd26"
	stripePublishableKeyG  = "pk_test_oDGnDF4fHODe1ARkDhPCvOlU"
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

type paymentMethodInfo struct {
	Cvc         string
	ExpireMonth int
	ExpireYear  int
	Name        string
	Number      string
}

func (pmi *paymentMethodInfo) getReader() io.Reader {
	v := url.Values{}
	v.Set("card[number]", pmi.Number)
	v.Set("card[cvc]", pmi.Cvc)
	v.Set("card[exp_month]", strconv.Itoa(pmi.ExpireMonth))
	v.Set("card[exp_year]", strconv.Itoa(pmi.ExpireYear))
	v.Set("card[name]", pmi.Name)
	b := ([]byte)(v.Encode())
	return bytes.NewBuffer(b)
}

func registerPaymentMethod() error {
	pmi := &paymentMethodInfo{
		Cvc:         "123",
		ExpireMonth: 12,
		ExpireYear:  getNextYear(),
		Name:        "SORAO TAMAGAWA",
		Number:      "4242424242424242", // https://pay.jp/docs/testcard
	}

	pt, err := getPayJPToken(pmi)
	if err != nil {
		return err
	}

	st, err := getStripeTokenJP(pmi)
	if err != nil {
		return err
	}

	p := &soracom.PaymentMethodInfoPayJP{
		PayJPToken:  pt,
		StripeToken: st,
	}

	err = apiClient.RegisterPaymentMethodPayJP(p)
	if err != nil {
		return err
	}

	return nil
}

type paymentGatewayResponse struct {
	ID string `json:"id"`
}

func getPayJPToken(pmi *paymentMethodInfo) (string, error) {
	return getPaymentGatewayToken(pmi, "https://api.pay.jp/v1/tokens", payJPPublishableKey)
}

func getStripeTokenJP(pmi *paymentMethodInfo) (string, error) {
	return getPaymentGatewayToken(pmi, "https://api.stripe.com/v1/tokens", stripePublishableKeyJP)
}

func getStripeTokenG(pmi *paymentMethodInfo) (string, error) {
	return getPaymentGatewayToken(pmi, "https://api.stripe.com/v1/tokens", stripePublishableKeyG)
}

func getPaymentGatewayToken(pmi *paymentMethodInfo, url, publishableKey string) (string, error) {
	req, err := http.NewRequest("POST", url, pmi.getReader())
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(publishableKey, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var r paymentGatewayResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return "", err
	}

	return r.ID, nil
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
