package response

import "github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/abstraction"

type Meta struct {
	Success bool                        `json:"success" default:"true"`
	Message string                      `json:"message" default:"true"`
	Info    *abstraction.PaginationInfo `json:"info"`
}
