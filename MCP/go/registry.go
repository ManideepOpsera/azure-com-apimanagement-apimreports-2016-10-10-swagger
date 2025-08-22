package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_reports "github.com/apimanagementclient/mcp-server/tools/reports"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_reports.CreateReports_listbyserviceTool(cfg),
	}
}
