package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "something",
	}

	t.Run("known key", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "something"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrUnknownDicKey

		assertErrors(t, err, ErrUnknownDicKey)

		assertStrings(t, err.Error(), want.Error())
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new key", func(t *testing.T) {
		dictionary := Dictionary{}
		addKey := "test"
		addValue := "something"
		err := dictionary.Add(addKey, addValue)
		assertNoError(t, err)
		assertDefinitions(t, dictionary, addKey, addValue)
	})

	t.Run("add existing key", func(t *testing.T) {
		existedKey := "test"
		existedValue := "something"
		dictionary := Dictionary{
			existedKey: existedValue,
		}

		err := dictionary.Add(existedKey, "anotherthing")
		assertErrors(t, err, ErrExistedDicKey)
		assertDefinitions(t, dictionary, existedKey, existedValue)
	})

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertErrors(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %s, want %s", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertDefinitions(t *testing.T, d Dictionary, key, want string) {
	t.Helper()

	got, err := d.Search(key)

	if err != nil {
		t.Fatal("should find the key: ", err)
	}

	if want != got {
		assertStrings(t, got, want)
	}
}
