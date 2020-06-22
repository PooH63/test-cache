package controllers

import (
	"errors"
	"github.com/RGRU/go-memorycache"
	"github.com/test-cache/internal/api"
	"time"
)

// Cache функция для работы с кэшом.
func Cache(cache *memorycache.Cache) api.HandlerFunc {
	return func(c *api.Context) error {
		var result api.Result
		params := c.Request.URL.Query()

		if len(params) == 0 {
			return errors.New(api.EMPTY_QUERY_PARAMS_ERROR)
		}

		for key, value := range params {

			if params[key][0] == "" {
				return errors.New(api.EMPTY_PARAM_VALUE_ERROR)
			}

			stored, ok := cache.Get(key)
			if !ok {
				err := cache.Set(key, value[0], 5*time.Minute)
				if err != nil {
					return errors.New(err.Error())
				}

				result = api.STORED
			} else {
				if stored == value[0] {
					result = api.IDENTICAL
				} else {
					result = api.DIFFER
				}
			}

			c.NewResponseSuccess(result.String(), value[0], stored)

			break
		}

		return nil
	}
}
