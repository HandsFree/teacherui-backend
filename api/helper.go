package api

import (
    "fmt"
    "log"
    "time"

    "github.com/patrickmn/go-cache"

    "github.com/gin-gonic/gin"

    "github.com/HandsFree/teacherui-backend/cfg"
)

// ApiLayer is a layer which handles manipulation of
// sending and retrieving data to the beaconing API

// API is the main instance to the api helper
// this performs any api requests necessary
var API *CoreAPIManager

// timeout for api requests (set to 120 seconds temporarily)
const timeout = 120 * time.Second

// GetRedirectBaseLink returns the link for
// redirecting the api tokens
func GetRedirectBaseLink() string {
    if cfg.Beaconing.Server.CallbackURL == "" {
        log.Fatal("Server Host not defined in config!")
    }

    return cfg.Beaconing.Server.CallbackURL + "/api/v1/token/"
}

// GetLink returns the callback link with a trailing slash
func GetLink() string {
    return cfg.Beaconing.Server.CallbackURL + "/"
}

// SetupAPIHelper sets up an instanceof the API manager
// should not be called more than once (in theory!)
func SetupAPIHelper() {
    API = newAPIHelper()
}

// CoreAPIManager manages all of the api middleman requests, etc.
// as well as caching any json/requests that are frequently requested
type CoreAPIManager struct {
    APIPath string
    cache   *cache.Cache
}

// getPath creates an API path, appending on the given beaconing URL
// "https://core.beaconing.eu/api/", this makes concatenation painless
// as well as it slaps the access token on the end
func (a *CoreAPIManager) getPath(s *gin.Context, args ...string) string {
    path := a.APIPath
    for _, arg := range args {
        path += arg
    }
    return fmt.Sprintf("%s", path)
}

func Cache() *cache.Cache {
    return API.cache
}

func newAPIHelper() *CoreAPIManager {
    return &CoreAPIManager{
        APIPath: cfg.Beaconing.Server.BeaconingAPIRoute,
        cache:   cache.New(30*time.Minute, 10*time.Minute),
    }
}
