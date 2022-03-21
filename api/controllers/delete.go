package controllers

import (
	"api/config"
)

func DeleteAll() error {

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return err
	}

	for _, j := range length {
		err := config.Rdb.Del(config.RedisCtx, j).Err()
		if err != nil {
			return err
		}

	}

	return nil
}

