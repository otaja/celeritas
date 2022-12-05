package cache

import (
	"testing"
)

func TestBadgerCache_Has(t *testing.T) {
	err := testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cahce but it shouldnt be there")
	}

	_ = testBadgerCache.Set("foo", "bar")

	inCache, err = testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cahce but it should be there")
	}

	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}
}

func TestBadgerCache_Get(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testBadgerCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("should get bar but did not get it")
	}
}

func TestBadgerCache_Forget(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("get foo but shoul not get anything")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := testBadgerCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("get alpha but shoul not get anything")
	}

}

func TestBadgerCache_EmptyByCache(t *testing.T) {
	err := testBadgerCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("alpha2", "beta2")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("beta", "foo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.EmptyByMatch("a")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("get alpha but shoul not get anything")
	}

	inCache, err = testBadgerCache.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("get alpha2 but shoul not get anything")
	}

	inCache, err = testBadgerCache.Has("beta")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("should get beta but get nothing")
	}
}
