package cache

import "testing"

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
		t.Error("foo found in cache, and it shouldn't be there")
	}

	err = testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache")
	}

	err = testBadgerCache.Forget("foo")
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
		t.Error("did not get correct value from cache")
	}

}

func TestBadgerCache_Forget(t *testing.T) {
	err := testBadgerCache.Set("foo", "goo")
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
		t.Error("foo found but should have been removed")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := testBadgerCache.Set("foo", "goo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found but should have been removed")
	}
}

func TestBadgerCache_EmptyByMatch(t *testing.T) {
	err := testBadgerCache.Set("alpha", "goo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("alpha2", "goo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("beta", "goo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.EmptyByMatch("alpha")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha found but should have been removed")
	}

	inCache, err = testBadgerCache.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha2 found but should have been removed")
	}

	inCache, err = testBadgerCache.Has("beta")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("beta should not have been removed")
	}
}
