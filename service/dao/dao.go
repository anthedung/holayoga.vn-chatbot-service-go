package dao

import (
	"context"
	"google.golang.org/appengine/datastore"
	"vn.holayoga.dialogflow.service/model"
	"sync"
	"time"
	"errors"
	"google.golang.org/appengine/log"
	"github.com/sirupsen/logrus"
)

const (
	CategoryDataStoreEntity = "Category"
)

// only update Yoga Category every [x] minute or upon restart
type YogaCacheDao struct {
	*sync.RWMutex
	updateFreqInHours time.Duration
	ProjectID         string // determine where to fetch data

	// caches
	categories []model.YogaCategory

	// caches Entity name
	CategoryDataStoreEntity string
}

func (cache *YogaCacheDao) RefreshCacheIfUninitialized(ctx context.Context) error {
	if cache.categories == nil {
		return cache.RefreshCache(ctx)
	}

	return nil
}

func NewYogaCategoryCache(freqHours int, projectId string, categoryDataStoreEntity string) (*YogaCacheDao, error) {
	// if ctx is instantiated,
	// due to GAE Standard environment appengine context must be available from inflight request
	var cacheDao = &YogaCacheDao{
		updateFreqInHours:       time.Hour * time.Duration(freqHours),
		RWMutex:                 &sync.RWMutex{},
		ProjectID:               projectId,
		CategoryDataStoreEntity: categoryDataStoreEntity,
	}

	return cacheDao, nil
}

func (cache *YogaCacheDao) GetCategories() []model.YogaCategory {
	cache.RLock()
	defer cache.RUnlock()

	// lock in case to prevent concurrent modification of cache.categories in refresh()
	return cache.categories
}

func (cache *YogaCacheDao) GetCategoryByName(catName string) (*model.YogaCategory, error) {
	cache.RLock()
	defer cache.RUnlock()

	for _, cat := range cache.categories {
		if cat.Name == catName {
			return &cat, nil
		}
	}

	return nil, errors.New("category not found " + catName)
}

// combines all poses from different categories
func (cache *YogaCacheDao) GetAllPoses() []model.YogaPose {
	cache.RLock()
	defer cache.RUnlock()

	var poses []model.YogaPose
	for _, cat := range cache.categories {
		poses = append(poses, cat.Poses...)
	}

	return poses
}

// due to the low frequency of updates to DataStore, only refresh Cache per request
// gccloud_api_cost-- && performance++
func (cache *YogaCacheDao) RefreshCache(ctx context.Context) error {
	log.Infof(ctx, "refreshing cache...")

	// Updating categories
	log.Infof(ctx, "retrieveCategories...")
	categories, err := cache.retrieveCategories(ctx)

	if err != nil {
		log.Errorf(ctx, "error updating cache at ", err.Error())
		return err
	}

	cache.Lock()
	cache.categories = categories
	cache.Unlock()

	log.Infof(ctx, "cache refreshed...")
	logrus.Info("cache refreshed...")
	return nil
}

func (y *YogaCacheDao) retrieveCategories(ctx context.Context) ([]model.YogaCategory, error) {
	var categories []model.YogaCategory

	// Create a query to fetch all Task entities, ordered by "created".
	query := datastore.NewQuery(y.CategoryDataStoreEntity)

	keys, err := query.GetAll(ctx, &categories)

	if err != nil {
		log.Errorf(ctx, "can't get "+y.CategoryDataStoreEntity)
		return nil, err
	}

	// Set the id field on each Task from the corresponding key.
	for i, key := range keys {
		categories[i].ID = key.IntID()
	}

	return categories, nil
}
