package main

import (
	"github.com/wwboynton/discord-terraform/discord"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: discord.Provider})
}
