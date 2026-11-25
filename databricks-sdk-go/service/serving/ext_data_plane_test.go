package serving

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	goauth "golang.org/x/oauth2"
)

type infoMock struct {
	called bool
	info   *DataPlaneInfo
	err    error
}

func (i *infoMock) DataPlaneInfoGetter() (*DataPlaneInfo, error) {
	i.called = true
	return i.info, i.err
}

type tokenRefreshSpy struct {
	called       bool
	expectedInfo *DataPlaneInfo
	token        *goauth.Token
	err          error
}

func (t *tokenRefreshSpy) TokenRefresh(info *DataPlaneInfo) (*goauth.Token, error) {
	t.expectedInfo = info
	t.called = true
	return t.token, t.err
}

func TestTokenNotCached(t *testing.T) {
	info := infoMock{
		info: &DataPlaneInfo{
			EndpointUrl:          "url",
			AuthorizationDetails: "authDetails",
		},
		err: nil,
	}
	s := tokenRefreshSpy{
		token: &goauth.Token{
			AccessToken: "token",
			TokenType:   "type",
			Expiry:      time.Now().Add(1 * time.Hour),
		},
		err: nil,
	}
	c := dataPlaneServiceImpl{
		infos:  make(map[string]*DataPlaneInfo),
		tokens: make(map[string]*goauth.Token),
	}
	url, token, err := c.GetDataPlaneDetails("method", []string{"params"}, s.TokenRefresh, info.DataPlaneInfoGetter)
	assert.NoError(t, err)
	assert.Equal(t, "url", url)
	assert.Equal(t, "token", token.AccessToken)
	assert.Equal(t, "type", token.TokenType)
	assert.True(t, token.Valid())
	assert.True(t, info.called)
	assert.True(t, s.called)
}

func TestTokenCached(t *testing.T) {
	info := infoMock{
		info: &DataPlaneInfo{
			EndpointUrl:          "url",
			AuthorizationDetails: "authDetails",
		},
		err: nil,
	}
	s := tokenRefreshSpy{
		token: &goauth.Token{
			AccessToken: "token",
			TokenType:   "type",
			Expiry:      time.Now().Add(1 * time.Hour),
		},
		err: nil,
	}
	c := dataPlaneServiceImpl{}
	c.infos = make(map[string]*DataPlaneInfo)
	c.tokens = make(map[string]*goauth.Token)
	c.infos["method/params"] = info.info
	c.tokens["method/params"] = s.token
	url, token, err := c.GetDataPlaneDetails("method", []string{"params"}, s.TokenRefresh, info.DataPlaneInfoGetter)
	assert.NoError(t, err)
	assert.Equal(t, "url", url)
	assert.Equal(t, "token", token.AccessToken)
	assert.Equal(t, "type", token.TokenType)
	assert.True(t, token.Valid())
	assert.False(t, info.called)
	assert.False(t, s.called)
}

func TestTokenExpired(t *testing.T) {
	info := infoMock{
		info: &DataPlaneInfo{
			EndpointUrl:          "url",
			AuthorizationDetails: "authDetails",
		},
		err: nil,
	}

	expired := &goauth.Token{
		AccessToken: "oldToken",
		TokenType:   "type",
		Expiry:      time.Now().Add(-1 * time.Hour),
	}
	s := tokenRefreshSpy{
		token: &goauth.Token{
			AccessToken: "token",
			TokenType:   "type",
			Expiry:      time.Now().Add(1 * time.Hour),
		},
		err: nil,
	}
	c := dataPlaneServiceImpl{}
	c.infos = make(map[string]*DataPlaneInfo)
	c.tokens = make(map[string]*goauth.Token)
	c.infos["method/params"] = info.info
	c.tokens["method/params"] = expired
	url, token, err := c.GetDataPlaneDetails("method", []string{"params"}, s.TokenRefresh, info.DataPlaneInfoGetter)
	assert.NoError(t, err)
	assert.Equal(t, "url", url)
	assert.Equal(t, "token", token.AccessToken)
	assert.Equal(t, "type", token.TokenType)
	assert.True(t, token.Valid())
	assert.False(t, info.called)
	assert.True(t, s.called)
}

func TestTokenInfoError(t *testing.T) {
	info := infoMock{
		info: nil,
		err:  assert.AnError,
	}
	s := tokenRefreshSpy{}
	c := dataPlaneServiceImpl{
		infos:  make(map[string]*DataPlaneInfo),
		tokens: make(map[string]*goauth.Token),
	}
	url, token, err := c.GetDataPlaneDetails("method", []string{"params"}, s.TokenRefresh, info.DataPlaneInfoGetter)
	assert.ErrorIs(t, err, assert.AnError)
	assert.Empty(t, url)
	assert.Nil(t, token)
	assert.True(t, info.called)
	assert.False(t, s.called)
}

func TestTokenRefreshError(t *testing.T) {
	info := infoMock{
		info: &DataPlaneInfo{
			EndpointUrl:          "url",
			AuthorizationDetails: "authDetails",
		},
		err: nil,
	}
	s := tokenRefreshSpy{
		token: nil,
		err:   assert.AnError,
	}
	c := dataPlaneServiceImpl{
		infos:  make(map[string]*DataPlaneInfo),
		tokens: make(map[string]*goauth.Token),
	}
	url, token, err := c.GetDataPlaneDetails("method", []string{"params"}, s.TokenRefresh, info.DataPlaneInfoGetter)
	assert.ErrorIs(t, err, assert.AnError)
	assert.Empty(t, url)
	assert.Nil(t, token)
	assert.True(t, info.called)
	assert.True(t, s.called)
}
