package client

// UserClient - клиент для прокидывания запросов user-service
type UserClient struct {
}

// NewUserClient - конструктор UserClient
func NewUserClient() *UserClient {
	return &UserClient{}
}
