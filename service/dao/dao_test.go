package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"vn.holayoga.dialogflow.service/model"
)

var testCache *YogaCacheDao
var categoryEntityTest = "CategoryUnitTest"

func init() {
	/* load test data */
	testCache, _ = NewYogaCategoryCache(1, categoryEntityTest)
	ctx, done, err := aetest.NewContext()
	model.InitDataStore(ctx, categoryEntityTest)

	if err != nil {
		panic("cannot init ctx")
	}
	defer done()
	testCache.RefreshCache(ctx)
}

func TestNewYogaCategoryCache(t *testing.T) {
	dao, err := NewYogaCategoryCache(1, categoryEntityTest)
	ctx, done, err := aetest.NewContext()
	model.InitDataStore(ctx, categoryEntityTest)

	if err != nil {
		t.Fatal(err)
	}
	defer done()
	dao.RefreshCache(ctx)

	assert.Nil(t, err)
	assert.NotZero(t, len(dao.GetCategories()))
}

func TestYogaCacheDao_GetAllPoses(t *testing.T) {
	poses := testCache.GetAllPoses()

	assert.Equal(t, 2, len(poses))
}

func TestYogaCacheDao_GetCategoryByName(t *testing.T) {
	cat, _ := testCache.GetCategoryByName("cơ bản")

	assert.NotNil(t, cat)
}

//TODO: Race test for refresh update
