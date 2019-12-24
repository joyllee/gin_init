package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"local/gin_init/common/errors"
	"local/gin_init/util/httpclient"
	"net/http"
)

type CallClient struct {
	httpclient.HttpClient
	Host string
	Uri  string
}

func (c *CallClient) PostCall(req, resp interface{}) (err error) {
	traceId := uuid.New().String()
	jsonstr, err := json.Marshal(req)
	if err != nil {
		return errors.NewRespErr(errors.ECodeUnImplement, err)
	}
	httpRequest, err := c.NewPostRequestWithJson(c.Host, c.Uri, string(jsonstr))
	if err != nil {
		logrus.Error("GetUserList create post request fail", err)
		return err
	}

	var rs *http.Response
FOR:
	for i := 0; i < 3; i++ {
		rs, err = c.DoRequest(traceId, httpRequest)
		if err != nil {
			continue FOR
		}
		defer func() {
			if err := rs.Body.Close(); err != nil {
				logrus.Error(err)
			}
		}()
		break FOR
	}
	if err != nil {
		return errors.NewRespErr(errors.ECodeInternalError, err)
	}

	if rs.StatusCode != http.StatusOK {
		return errors.NewRespErr(errors.ECodeInternalError, err)
	}

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return errors.NewRespErr(errors.ECodeUnkown, err)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return errors.NewRespErr(errors.ECodeUnImplement, err)
	}

	return nil
}
