package main

/*
 * MIT License
 *
 * Copyright (c) 2018 yasukun
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

import (
	"context"
	"errors"

	"github.com/go-redis/redis"
)

type PriceHandler struct {
	Client *redis.Client
}

// NewPriceHandler ...
func NewPriceHandler(client *redis.Client) *PriceHandler {
	return &PriceHandler{
		Client: client,
	}
}

// (p PriceHandler) Price ...
func (p PriceHandler) Price(ctx context.Context, key string) (string, error) {
	var result string
	resp, err := p.Client.LRange(key, -1, -1).Result()
	if err != nil {
		return result, err
	}
	if len(resp) == 0 {
		return result, errors.New("response length 0")
	}
	return resp[0], nil

}

// (p PriceHandler) Prices ...
func (p PriceHandler) Prices(ctx context.Context, key string, start, stop int16) ([]string, error) {
	var resp []string
	resp, err := p.Client.LRange(key, int64(start), int64(stop)).Result()
	if err != nil {
		return resp, err
	}
	return resp, nil
}
