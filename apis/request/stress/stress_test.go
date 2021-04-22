package stress

import (
	"fmt"
	"log"
	"os"
	"request/models"
	"request/requests"
	"request/utils"
	"sync"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testToken string
)

func init() {
	os.Setenv("MONGOHOSTNAME", "localhost")
	_, err := models.DefaultDB()
	if err != nil {
		log.Fatal(err)
	}
}

func TestStress(t *testing.T) {
	userCount := 30
	reqCount := 1000
	rps := 100 // requests per second

	t.Log("Ping test")
	if err := requests.Ping(); err != nil {
		t.Fatal(err)
	}
	user := models.User{
		Name:     "request Challenge",
		Email:    "challenge@me.more",
		Password: "winner lottery ticket",
	}
	t.Log("gimme the token")
	t.Logf("email   : %s", user.Email)
	t.Logf("password: %s", user.Password)
	outToken, err := requests.Login(user)
	if err != nil {
		t.Log("no token for u")
		t.Fatal(err)
	}
	testToken = outToken

	users := []models.User{}
	errs := map[string][]error{}
	wg := sync.WaitGroup{}
	t.Logf("gimme %d users", userCount)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			user := models.User{
				Email:    fmt.Sprintf("%s@%s.test", utils.RString(5, 8), utils.RString(3, 5)),
				Name:     utils.RString(5, 8),
				Password: utils.RString(6, 10),
			}
			err := requests.AddUser(user, testToken)
			if err != nil {
				t.Logf("%d - add user err: %v", i, err)
				errs["user"] = append(errs["user"], err)
				return
			}
			users = append(users, user)
		}()
	}
	wg.Wait()

	if len(errs) > 0 {
		c := 0
		for key := range errs {
			for i := range errs[key] {
				t.Log(errs[key][i])
				c++
			}
		}
		t.Fatalf("%d total user err", c)
	}

	if len(users) != userCount {
		t.Logf("%d users were left behind", userCount-len(users))
	}
	t.Logf("%d available users", len(users))

	reqs := []models.ActivationRequest{}
	t.Log("Ready for the first load?")
	for i := 0; i < reqCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			reqID, err := requests.AddReq(users[utils.RInt(len(users)-1)].ID.Hex(), testToken)
			if err != nil {
				t.Logf("%d - add req err: %v", i, err)
				errs["req"] = append(errs["req"], err)
				return
			}
			reqOID, err := primitive.ObjectIDFromHex(reqID)
			if err != nil {
				t.Logf("primitive.ObjectIDFromHex err: %v", err)
				errs["req"] = append(errs["req"], err)
				return
			}
			reqs = append(reqs, models.ActivationRequest{ID: reqOID})
		}()
		if (i % rps) == 0 {
			time.Sleep(time.Second)
			if utils.RInt(2) == 1 {
				rps++
				t.Logf("new rps %d", rps)
				// I added a 50% chance to increase the rps each time it is reached
				// the application endured all the tests, even with higher rps
			}
		}
		if (i % (reqCount / 10)) == 0 {
			t.Logf("[%04d of %04d]", i, reqCount)
		}
	}
	wg.Wait()
	t.Log("ok, that was easy")

	if len(reqs) != reqCount {
		t.Logf("%d reqs were left behind", reqCount-len(reqs))
		// I've added many debug msgs and comparisons trying to figure why I lost these requests
		// yet still, I did not find why
		// the best I'm able to guess is that it happens due to the way I programmed the loop
	}
	t.Logf("%d available reqs", len(reqs))

	t.Log("Ready for the second load?")
	status := []string{}
	for key := range models.StatusEnum {
		status = append(status, key)
	}
	for i := range reqs {
		go func() {
			err := requests.UpdateReq(reqs[i].ID.Hex(), users[utils.RInt(len(users)-1)].ID.Hex(), status[utils.RInt(len(status)-1)], testToken)
			if err != nil {
				t.Logf("%d - update req err: %v", i, err)
				errs["upd"] = append(errs["upd"], err)
				return
			}
		}()
		if (i)%rps == 0 {
			time.Sleep(time.Second)
			if utils.RInt(2) == 1 {
				rps++
				t.Logf("new rps %d", rps)
			}
		}
		if (i % (reqCount / 10)) == 0 {
			t.Logf("[%04d of %04d]", i, len(reqs))
		}
	}
	c := 0
	for key := range errs {
		for i := range errs[key] {
			t.Log(errs[key][i])
			c++
		}
	}
	if c > 0 {
		t.Fatalf("err count %d", c)
	}
	t.Log("Easy, right?")
}
