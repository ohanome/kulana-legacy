package fetcher

import (
	"crypto/tls"
	"io/ioutil"
	"kulana/hostinfo"
	"kulana/misc"
	"kulana/output"
	"log"
	"net/http"
)

func CreateHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}
}

func FetchHTTPHost(url string, foreignId string, checkSSLCert bool) output.Output {
	client := CreateHTTPClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "ohano *kulana/"+misc.Version)
	start := misc.MicroTime()
	resp, err := client.Do(req)
	end := misc.MicroTime()
	defer client.CloseIdleConnections()
	if err != nil {
		return output.Output{
			Url:           url,
			Status:        0,
			Time:          0,
			Destination:   "",
			ContentLength: 0,
			Content:       "",
			ForeignID:     foreignId,
			IpAddress:     "",
			Certificate: output.CertificateOutput{
				ValidUntil: "",
				Issuer:     "",
			},
		}
	}

	statusCode := resp.StatusCode
	responseTime := (end - start) * 1000
	responseTimeRounded := misc.Round(responseTime, 0.000001)

	var destination string
	location, err := resp.Location()
	if err != nil {
		destination = url
	} else {
		destination = location.String()
	}

	body, err := ioutil.ReadAll(resp.Body)
	contentLength := int64(len(body))

	response := output.Output{
		Url:           url,
		Status:        statusCode,
		Time:          responseTimeRounded,
		Destination:   destination,
		ContentLength: contentLength,
		Content:       string(body),
		ForeignID:     foreignId,
		IpAddress:     hostinfo.ResolveIPAddress(url),
		Certificate: output.CertificateOutput{
			ValidUntil: "",
			Issuer:     "",
		},
	}

	if checkSSLCert {
		certIsValid, validUntil, issuer := hostinfo.CheckCertificate(url)
		if certIsValid {
			response.Certificate.Valid = true
			response.Certificate.ValidUntil = validUntil.Format("2006-01-02 15:04:05")
			response.Certificate.Issuer = issuer
		} else {
			response.Certificate.Valid = false
		}
	}

	return response
}
