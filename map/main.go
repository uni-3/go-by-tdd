package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

// mapは参照型のため、データ構造への参照を渡すため、ポインタを使わなくても変更、取得時のアドレスが一緒になる

// Search keyの値を取得する
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Add keyとvalueの追加
func (d Dictionary) Add(key string, value string) {
	d[key] = value
}
