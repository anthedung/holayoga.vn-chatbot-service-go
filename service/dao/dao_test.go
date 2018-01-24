package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testCache *YogaCacheDao
var categoryEntityTest = "CategoryUnitTest"

func init() {
	/* load test data */
	testCache, _ = NewYogaCategoryCache(1, "newagent-4790c", categoryEntityTest)
}

func TestNewYogaCategoryCache(t *testing.T) {
	dao, err := NewYogaCategoryCache(1, "newagent-4790c", categoryEntityTest)

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
