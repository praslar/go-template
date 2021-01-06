package models

type ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            LoginRequest           `json:"input"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLData struct {
	LoginResponse *LoginResponse `json:"loginAction,omitempty"`
}

type GraphQLResponse struct {
	Data   GraphQLData    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

type LoginResponse struct {
	Token         string      `json:"token"`
	ListPlatform  interface{} `json:"list_platform"`
	UserInfo      interface{} `json:"user_info"`
	PermissionIDs []string    `json:"permission_ids"`
	Stations      interface{} `json:"stations"`
}

type Mutation struct {
	CreateUser *LoginResponse
}
