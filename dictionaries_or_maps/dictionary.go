package dictionariesormaps

type Dictionary map[string]string

const (
	ErrorNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrorWordDoesNotExists = DictionaryErr("cannot update word because it not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		d[key] = value
	case ErrorWordExists:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExists
	case nil:
		d[key] = value
	default:
		return nil
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
