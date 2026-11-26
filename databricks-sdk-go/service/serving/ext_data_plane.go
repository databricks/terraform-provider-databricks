package serving

import (
	"strings"
	"sync"

	goauth "golang.org/x/oauth2"
)

// DataPlaneService is an interface for services that access DataPlane.
type DataPlaneService interface {
	GetDataPlaneDetails(method string, params []string, refresh func(*DataPlaneInfo) (*goauth.Token, error), infoGetter func() (*DataPlaneInfo, error)) (string, *goauth.Token, error)
}

func NewDataPlaneService() DataPlaneService {
	return &dataPlaneServiceImpl{
		infos:  make(map[string]*DataPlaneInfo),
		tokens: make(map[string]*goauth.Token),
	}
}

type dataPlaneServiceImpl struct {
	infos  map[string]*DataPlaneInfo
	tokens map[string]*goauth.Token
	// This class can be shared across multiple threads.
	// This mutex is used to synchronize access to the infos and tokens maps.
	mu sync.Mutex
}

// GetDataPlaneDetails returns the endpoint URL and token. It returns a cached token if it is valid,
// otherwise it refreshes the token and returns the new token.
func (o *dataPlaneServiceImpl) GetDataPlaneDetails(method string, params []string, refresh func(*DataPlaneInfo) (*goauth.Token, error), infoGetter func() (*DataPlaneInfo, error)) (string, *goauth.Token, error) {
	key := o.generateKey(method, params)
	info, err := o.getInfo(key, infoGetter)
	if err != nil {
		return "", nil, err
	}
	token, err := o.refreshToken(key, info, refresh)
	if err != nil {
		return "", nil, err
	}
	return info.EndpointUrl, token, nil
}

func (o *dataPlaneServiceImpl) getInfo(key string, infoGetter func() (*DataPlaneInfo, error)) (*DataPlaneInfo, error) {
	info, infoOk := o.infos[key]
	if !infoOk {
		o.mu.Lock()
		defer o.mu.Unlock()
		info, infoOk = o.infos[key]
		if !infoOk {
			newInfo, err := infoGetter()
			if err != nil {
				return nil, err
			}
			o.infos[key] = newInfo
			info = newInfo
		}
	}
	return info, nil
}

func (o *dataPlaneServiceImpl) refreshToken(key string, info *DataPlaneInfo, refresh func(*DataPlaneInfo) (*goauth.Token, error)) (*goauth.Token, error) {
	token, tokenOk := o.tokens[key]
	if !tokenOk || !token.Valid() {
		o.mu.Lock()
		defer o.mu.Unlock()
		token, tokenOk = o.tokens[key]
		if !tokenOk || !token.Valid() {
			newToken, err := refresh(info)
			if err != nil {
				return nil, err
			}
			o.tokens[key] = newToken
			token = newToken
		}
	}
	return token, nil
}

func (o *dataPlaneServiceImpl) generateKey(method string, params []string) string {
	allElements := []string{method}
	allElements = append(allElements, params...)
	return strings.Join(allElements, "/")
}
