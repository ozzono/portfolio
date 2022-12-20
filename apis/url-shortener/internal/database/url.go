package database

import (
	"url-shortener/internal/models"
	"url-shortener/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	urlColl   = "urls"
	defaultDB = "url-shortener"
)

func (client *Client) AddURL(url *models.URL, debug bool) (*models.URL, error) {

	dbURL, found, err := client.FindURLBySource(url, false)
	if err != nil {
		return dbURL, errors.Wrap(err, "client.FindURLBySource")
	}
	if found {
		dbURL.Log("already existent", debug)
		return dbURL, nil
	}
	url.ID = primitive.NewObjectID()
	url.Shortened = utils.RString(5, 7)
	bsonURL, err := utils.ToDoc(url)
	if err != nil {
		return nil, errors.Wrap(err, "utils.ToDoc")
	}

	_, err = client.C.
		Database(defaultDB).
		Collection(urlColl).
		InsertOne(client.Ctx, bsonURL)
	if err != nil {
		return nil, errors.Wrap(err, "client.C.Database().Collection().InsertOne()")
	}

	url.Log("creating", debug)
	return url, nil
}

func (client *Client) findURL(filter bson.M, debug bool) (*models.URL, bool, error) {
	cursor, err := client.C.
		Database(defaultDB).
		Collection(urlColl).
		Find(client.Ctx, filter)
	if err != nil {
		return nil, false, errors.Wrap(err, "client.C.Database().Collection().Find()")
	}

	urls := []*models.URL{}
	for cursor.Next(client.Ctx) {
		u := &models.URL{}
		err = cursor.Decode(&u)
		if err != nil {
			return nil, false, errors.Wrap(err, "cursor.Decode")
		}
		urls = append(urls, u)
	}

	if len(urls) == 0 {
		return nil, false, nil
	}

	return urls[0], true, nil
}

func (client *Client) FindURLByID(url *models.URL, debug bool) (*models.URL, bool, error) {
	url, found, err := client.findURL(bson.M{"_id": url.ID}, debug)

	if err != nil {
		return nil, false, errors.Wrap(err, "client.findURL")
	}

	if !found {
		return nil, false, nil
	}

	url.Log("searching by id", debug)
	return url, found, nil
}

func (client *Client) FindURLByShortened(url *models.URL, debug bool) (*models.URL, bool, error) {
	url, found, err := client.findURL(bson.M{"shortened": url.Shortened}, debug)

	if err != nil {
		return nil, false, errors.Wrap(err, "client.findURL")
	}

	if !found {
		return nil, false, nil
	}

	url.Log("searching by shortened", debug)
	return url, found, nil
}

func (client *Client) FindURLBySource(url *models.URL, debug bool) (*models.URL, bool, error) {
	url, found, err := client.findURL(bson.M{"source": url.Source}, debug)
	if err != nil {
		return nil, false, errors.Wrap(err, "client.findURL")
	}

	if !found {
		return nil, false, nil
	}

	url.Log("searching by source", debug)
	return url, found, nil
}

func (client *Client) DelURL(url *models.URL) error {
	_, err := client.C.
		Database(defaultDB).
		Collection(urlColl).
		DeleteOne(client.Ctx, bson.M{"_id": url.ID})
	if err != nil {
		return errors.Wrap(err, "client.C.Database().Collection().DeleteOne()")
	}
	return nil
}

func (client *Client) IncrementURL(url *models.URL, debug bool) (*models.URL, error) {
	url, found, err := client.FindURLBySource(url, false)
	if err != nil {
		return nil, errors.Wrap(err, "client.FindURL")
	}
	if !found {
		return nil, errors.New(url.Source + " url not found")
	}
	url.Count++
	err = client.UpdateURL(url)
	if err != nil {
		return nil, errors.Wrap(err, "client.UpdateURL")
	}

	url.Log("incrementing counter", debug)
	return url, nil
}

func (client *Client) UpdateURL(url *models.URL) error {
	bsonURL, err := utils.ToDoc(url)
	if err != nil {
		return errors.Wrap(err, "utils.ToDoc")

	}
	update := bson.M{"$set": bsonURL}

	_, err = client.C.
		Database(defaultDB).
		Collection(urlColl).
		UpdateByID(client.Ctx, url.ID, update)
		// UpdateOne(
		// 	client.Ctx,
		// 	bson.M{"$set": url.Source},
		// 	bsonURL,
		// )
	if err != nil {
		return errors.Wrap(err, "client.C.Database().Collection().UpdateOne()")
	}
	return nil
}
