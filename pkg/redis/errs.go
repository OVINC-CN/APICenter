package redis

import "fmt"

type ConcurrencyLocked struct {
	Key string
}

func (c *ConcurrencyLocked) Error() string {
	return fmt.Sprintf("concurrency locked for %s", c.Key)
}
