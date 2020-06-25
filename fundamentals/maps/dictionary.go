package maps

type (
	// Dictionary : A map which contains words and their definitions
	Dictionary map[string]string
	// DictionaryErr : A custom error type made for the Dictionary type, implements the error interface as it has an Error() string method
	DictionaryErr string
)

const (
	// ErrNotFound : A custom error for when a word cannot be found in the Dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists : A custom error for when a word already exists in the Dictionary
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist : A custom error for when a word does not exist in the Dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word as it does not exist")
)

// Error : the Error() string method that makes DictionaryErr implement the error interface
func (e DictionaryErr) Error() string {
	// Any type with an Error() string method fulfils the error interface.
	return string(e)
}

// Search : Finds a given word in the dictionary, throws an error if it cannot be found
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add : Adds a word and definition to the dictionary, throws an error if it already exists
func (d Dictionary) Add(word, definition string) error {
	// Don't need to use a pointer as a map is a reference type
	// The underlying data structure is a hash table or hash map
	// Maps being a reference is really good, because no matter how big a map gets there will only be one copy
	// Maps can be nil, writing to a nil map will cause a panic so never initialise an empty map (don't do this "var m map[string]string")
	_, err := d.Search(word)
	// Using a switch like this provides the extra safety net in case an unexpected error is returned
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

// Update : Updates a definition in the dictionary, throws an error if the word doesn't exist
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	// Using a switch like this provides the extra safety net in case an unexpected error is returned
	switch err {
	case ErrNotFound:
		// We could reuse ErrNotFound and not add a new error. However, it is often better to have a precise error for when an update fails.
		// Having specific errors gives you more information about what went wrong.
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete : Deletes a word and definition in the dictionary, does nothing if the word doesn't exist
func (d Dictionary) Delete(word string) {
	// Delete is built into Go and doesn't return anything so we follow the same notion and don't check if it exists
	delete(d, word)
}
