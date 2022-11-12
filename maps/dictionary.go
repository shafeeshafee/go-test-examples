package dictionary

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("Could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("Cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("Cannot update word because it doesn't exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	/*
		A map can return 2 values. The second value is a bool that shows if a key was found successfully.
	*/
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
