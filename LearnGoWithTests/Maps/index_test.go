package HelloWorld

import (
	"testing"
)

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// ##################################################################################################################

/*
 * Declaring a Map
 *  |_ is somewhat similar to an array.
 *  |_ Except, it starts with the map keyword and requires two types.
 *  |_ The first is the key type, which is written inside the [].
 *  |_ The second is the value type, which goes right after the [].
 *
 *
 */
func TestSearch(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	t.Run("known word", func(t *testing.T) {
		//dictionary := map[string]string{"test": "this is just a test"}
		//got := Search(dictionary, "test")
		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)
	})
	// The basic search was very easy to implement, but what will happen if we supply a word that's not in our dictionary?
}

// An interesting property of maps is that you can modify them without passing them as a pointer.
// This is because map is a reference type. Meaning it holds a reference to the underlying data structure, much like a pointer. The underlying data structure is a hash table, or hash map, and you can read more about hash tables
// Maps being a reference is really good, because no matter how big a map gets there will only be one copy.
func (d Dictionary) Search(word string) (string, error) {
	//return d[word], nil

	// In order to make this pass, we are using an interesting property of the map lookup.
	// It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// ##################################################################################################################

func (d Dictionary) Add(word, definition string) error {
	//d[word] = definition
	_, err := d.Search(word)
	//return nil
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// I decided to create an assertStrings helper to make the implementation more general.
func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Map will not throw an error if the value already exists.
// Instead, they will go ahead and overwrite the value with the newly provided value.
// This can be convenient in practice, but makes our function name less than accurate.
// Add should not modify existing values. It should only add new words to our dictionary.
// Refactor -> Add is not overwrite old val.
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertError(t, nil, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)                // 看额外错误信息，e.g：新增已有的key值
		assertDefinition(t, dictionary, word, definition) //
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func (e DictionaryErr) Error() string {
	return string(e)
}

// ##################################################################################################################

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func (d Dictionary) Update(word, definition string) error {
	d[word] = definition // write enough code to make it pass
	return nil
}

// ##################################################################################################################

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Excepted %q to be deleted", word)
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// ##################################################################################################################
//func Search(dictionary map[string]string, word string) string {
//	//return ""
//	return dictionary[word]
//}
