package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {

	err := testRedisCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache but it shouldn´t be there")
	}

	err = testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache but it should be there")
	}

}

func TestRedisCache_Get(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testRedisCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("did not get correct value from cache")
	}

}

func TestRedisCache_Forget(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("values is still in cache but should be forgotten")
	}

}

func TestRedisCache_Empty(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("values is still in cache but should be empty")
	}

}

func TestRedisCache_EmptyByMatch(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.EmptyByMatch("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("values is still in cache but should be empty")
	}

	inCache, err = testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("values is not in cache but should be there")
	}

}
