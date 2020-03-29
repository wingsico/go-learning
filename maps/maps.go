package maps

// ErrUnknownDicKey : 字典中没有相关的key
const (
	ErrUnknownDicKey = DictionaryErr("字典中找不到该key")
	ErrExistedDicKey = DictionaryErr("已存在的key, 不可重复添加")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary : 键值都是字符串的map
type Dictionary map[string]string

// Search : find the value by the key in the dictionary
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrUnknownDicKey
	}
	return definition, nil
}

// Add : 给字典增加键值对
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrUnknownDicKey:
		d[key] = value
	case nil:
		return ErrExistedDicKey
	default:
		return err
	}
	return nil
}
