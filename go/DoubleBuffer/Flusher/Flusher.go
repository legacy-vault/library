package Flusher

import "github.com/legacy-vault/library/go/DoubleBuffer/Item"

type Flusher func([]Item.Item) error
