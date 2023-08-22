package dictionary

var (
	ErrNotFound         = DictionaryErr("could not find the word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word that does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, defn string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = defn
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, defn string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = defn
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
