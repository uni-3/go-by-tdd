package main

type Dictionary map[string]string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add key because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update key because it does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// mapは参照型のため、データ構造への参照を渡すため、ポインタを使わなくても変更、取得時のアドレスが一緒になる

// Search keyの値を取得する
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Add keyとvalueの追加。なかったら追加。すでにあれば、なにもしない
func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update 存在するkeyの更新。無いときはエラー
func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

// Delete 指定したkeyを削除。dがnil or 存在しないkeyのときはなにもしない
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
