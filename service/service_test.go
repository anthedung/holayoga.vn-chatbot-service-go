package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"vn.holayoga.dialogflow.service/test"
	"vn.holayoga.dialogflow.service/service/dao"
)

var testCache *dao.YogaCacheDao
var yogaservice *YogaService
var categoryEntityTest = "CategoryUnitTest"

func init() {
	/* load test data */
	testCache, _ = dao.NewYogaCategoryCache(1, categoryEntityTest)
	ctx, done, err := aetest.NewContext()
	test.InitDataStore(ctx, categoryEntityTest)

	if err != nil {
		panic("cannot init ctx")
	}
	defer done()
	testCache.RefreshCache(ctx)

	yogaservice, _ = NewYogaService(testCache)
}

func TestYogaService_GetCoursesByCategory(t *testing.T) {
	c, _ := yogaservice.GetCoursesByCategory("cơ bản")

	assert.NotEmpty(t, c)
}

func TestNewYogaService_NameStringCompare(t *testing.T) {
	assert.Equal(t, "10 ngày cơ bản", "10 ngày cơ bản") // was false due to escaping in Messenger
}

//TODO: Race test for refresh update
