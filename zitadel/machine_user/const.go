package machine_user

import (
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/user"
)

const (
	UserIDVar             = "user_id"
	userIDsVar            = "user_ids"
	userStateVar          = "state"
	UserNameVar           = "user_name"
	userNameMethodVar     = "user_name_method"
	loginNamesVar         = "login_names"
	preferredLoginNameVar = "preferred_login_name"
	nameVar               = "name"
	DescriptionVar        = "description"
	accessTokenTypeVar    = "access_token_type"
)

var (
	defaultAccessTokenType = user.AccessTokenType_name[0]
)
