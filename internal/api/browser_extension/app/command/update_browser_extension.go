package command

import (
	"github.com/google/uuid"
	"github.com/twofas/2fas-server/internal/api/browser_extension/domain"
)

type UpdateBrowserExtension struct {
	BrowserExtensionId string `uri:"extension_id" validate:"required,uuid4"`
	Name               string `json:"name" validate:"not_blank,lte=64"`
	BrowserName        string `json:"browser_name" validate:"lte=255"`
	BrowserVersion     string `json:"browser_version" validate:"lte=32"`
}

type UpdateBrowserExtensionHandler struct {
	Repository domain.BrowserExtensionRepository
}

func (h *UpdateBrowserExtensionHandler) Handle(cmd *UpdateBrowserExtension) error {
	id, _ := uuid.Parse(cmd.BrowserExtensionId)

	browserExtension, err := h.Repository.FindById(id)

	if err != nil {
		return err
	}

	if cmd.Name != "" {
		browserExtension.Name = cmd.Name
	}

	if cmd.BrowserVersion != "" {
		browserExtension.BrowserName = cmd.BrowserName
	}

	if cmd.BrowserVersion != "" {
		browserExtension.BrowserVersion = cmd.BrowserVersion
	}

	return h.Repository.Update(browserExtension)
}
