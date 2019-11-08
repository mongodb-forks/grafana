package setting

type OAuthOrgInfo struct {
	Role  string
	OrgId int64
}

type OAuthInfo struct {
	ClientId, ClientSecret       string
	Scopes                       []string
	AuthUrl, TokenUrl            string
	Enabled                      bool
	EmailAttributeName           string
	EmailAttributePath           string
	RoleAttributePath            string
	AllowedDomains               []string
	HostedDomain                 string
	ApiUrl                       string
	AllowSignup                  bool
	Name                         string
	TlsClientCert                string
	TlsClientKey                 string
	TlsClientCa                  string
	TlsSkipVerify                bool
	SendClientCredentialsViaPost bool
	OrgMapping                   map[string]OAuthOrgInfo
}

type OAuther struct {
	OAuthInfos map[string]*OAuthInfo
}

var OAuthService *OAuther
