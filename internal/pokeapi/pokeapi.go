package pokeapi

import (
    "net/http"
    "time"
    "github.com/zdravce988/pokedex/internal/pokecache"
    "io"
    "fmt"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
    httpClient http.Client
    cache pokecache.Cache
}

func NewClient() Client {
    return Client{
        httpClient: http.Client{
            Timeout: time.Minute,
        },
        cache: pokecache.NewCache(5 * time.Second), 
    }
}

func (c *Client) getData(url string) ([]byte, error) {
    data, exists := c.cache.Get(url)
    if exists {
        return data, nil
    }
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode > 399 {
        return nil, fmt.Errorf("bad status code: %v", resp.StatusCode)
    }
    defer resp.Body.Close()

    data, err = io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    c.cache.Add(url, data)
    return data, nil
}
