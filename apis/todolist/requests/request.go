package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"todo/models"
	"todo/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testURL  = "http://localhost:8888"
	maxRetry = 5
	retryNap = 100
)

//Although very good, the idea of making such a kind of test never occurred to me before
func Ping() error {
	res, err := utils.DefaultRequest("GET", testURL+"/ping", "")
	if err != nil {
		return fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("handler not available: %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if string(body) != "pong" {
		return fmt.Errorf("unexpected body: expecting 'pong' - found %s", string(body))
	}
	return nil
}

func Login(user models.User) (string, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("json.Marshal err: %v", err)
	}
	res, err := utils.DefaultRequest("POST", testURL+"/user/login", string(payload))
	if err != nil {
		return "", fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d\nBody: %+v", res.StatusCode, string(body))
	}
	output := map[string]string{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal err: %v", err)
	}
	_, ok := output["token"]
	if !ok {
		_, ok = output["msg"]
		if ok {
			return "", fmt.Errorf("failed to get token: %s", output["msg"])
		}
		return "", fmt.Errorf("failed to get token: %s", string(body))
	}
	return output["token"], nil
}

func AddUser(user models.User, token string) error {
	payload, err := json.Marshal(map[string]string{
		"email":    user.Email,
		"password": user.Password,
		"token":    token,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal err: %v", err)
	}
	c := 1
	res, err := utils.DefaultRequest("POST", testURL+"/user/add", string(payload))
	if err != nil {
		for res, err = utils.DefaultRequest("POST", testURL+"/user/add", string(payload)); err != nil && c < maxRetry; c++ {
			time.Sleep(time.Duration(retryNap) * time.Millisecond)
		}
		if err != nil {
			return fmt.Errorf("defaultRequest err: %v", err)
		}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	output := map[string]string{}
	err = json.Unmarshal(body, &output)
	_, ok := output["success"]
	if !ok {
		_, ok = output["msg"]
		if ok {
			return fmt.Errorf("failed to add user: %+v", output["msg"])
		}
		return fmt.Errorf("failed to add user: %+v", string(body))
	}
	return nil
}

func GetUser(user models.User, token string) (models.User, error) {
	payload, err := json.Marshal(map[string]string{
		"token": token,
		"email": user.Email,
	})
	if err != nil {
		return models.User{}, fmt.Errorf("json.Marshal err: %v", err)
	}
	res, err := utils.DefaultRequest("POST", testURL+"/user/get", string(payload))
	if err != nil {
		return models.User{}, fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.User{}, fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	if res.StatusCode != 200 {
		return models.User{}, fmt.Errorf("StatusCode: %d\nBody: %+v", res.StatusCode, string(body))
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		output := map[string]string{}
		json.Unmarshal(body, &output)
		_, ok := output["msg"]
		if ok {
			return models.User{}, fmt.Errorf("failed to get user: %v", output["msg"])
		}
		return models.User{}, fmt.Errorf("failed to get user: %v", body)
	}
	return user, nil
}

func UpdateUser(user models.User, token string) error {
	payload, err := json.Marshal(map[string]string{
		"email":    user.Email,
		"token":    token,
		"password": user.Password,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal err: %v", err)
	}
	res, err := utils.DefaultRequest("POST", testURL+"/user/update", string(payload))
	if err != nil {
		return fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	output := map[string]string{}
	err = json.Unmarshal(body, &output)
	_, ok := output["success"]
	if !ok {
		_, ok = output["msg"]
		if ok {
			return fmt.Errorf("failed to add user: %+v", string(body))
		}
		return fmt.Errorf("failed to add user: %+v", string(body))
	}
	return nil
}
func AddReq(userID, token string) (string, error) {
	payload, err := json.Marshal(map[string]string{
		"token":   token,
		"user_id": userID,
	})
	res, err := utils.DefaultRequest("POST", testURL+"/request/add", string(payload))
	if err != nil {
		return "", fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	output := map[string]string{}
	err = json.Unmarshal(body, &output)
	_, ok := output["msg"]
	if ok {
		return "", fmt.Errorf("failed to add request: %s", output["msg"])
	}
	reqID, err := primitive.ObjectIDFromHex(output["request_id"])
	if err != nil {
		return "", fmt.Errorf("primitive.ObjectIDFromHex err: %v", err)
	}
	return reqID.Hex(), nil
}

func GetReq(reqID, token string) (models.ActivationRequest, error) {
	payload, err := json.Marshal(map[string]string{
		"token":      token,
		"request_id": reqID,
	})
	res, err := utils.DefaultRequest("POST", testURL+"/request/get", string(payload))
	if err != nil {
		return models.ActivationRequest{}, fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.ActivationRequest{}, fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	req := models.ActivationRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		output := map[string]string{}
		json.Unmarshal(body, &output)
		_, ok := output["msg"]
		if ok {
			return models.ActivationRequest{}, fmt.Errorf("failed to get request: %s", output["msg"])
		}
		return models.ActivationRequest{}, fmt.Errorf("failed to get request: %s", string(body))
	}
	return req, nil
}

func UpdateReq(reqID, userID, status, token string) error {
	payload, err := json.Marshal(map[string]string{
		"token":      token,
		"request_id": reqID,
		"user_id":    userID,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal err: %v", err)
	}
	res, err := utils.DefaultRequest("POST", fmt.Sprintf("%s/request/%s?test=test", testURL, status), string(payload))
	if err != nil {
		return fmt.Errorf("defaultRequest err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	output := map[string]string{}
	err = json.Unmarshal(body, &output)
	_, ok := output["success"]
	if !ok {
		_, ok = output["msg"]
		if ok {
			return fmt.Errorf("failed to request data: %s", output["msg"])
		}
		return fmt.Errorf("failed to request data: %s", string(body))
	}
	return nil
}
