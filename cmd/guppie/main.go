package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ttacon/gophish"
	"github.com/ttacon/pretty"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Usage: "The host to connect to",
			},
			cli.StringFlag{
				Name:  "token",
				Usage: "The API token to use to connect",
			},
		},
		Commands: []cli.Command{
			{
				Name:        "sending-profiles",
				Aliases:     []string{"sp"},
				Usage:       "Manipulate sending profiles",
				Subcommands: sendingProfileCommands(),
			},
			{
				Name:        "templates",
				Aliases:     []string{"t"},
				Usage:       "Manipulate templates",
				Subcommands: templatesCommands(),
			},
			{
				Name:        "landing-pages",
				Aliases:     []string{"lp"},
				Usage:       "Manipulate landing pages",
				Subcommands: landingPagesCommands(),
			},
			{
				Name:        "groups",
				Aliases:     []string{"g"},
				Usage:       "Manipulate groups",
				Subcommands: groupCommands(),
			},
			{
				Name:        "campaigns",
				Aliases:     []string{"c"},
				Usage:       "Manipulate campaigns",
				Subcommands: campaignCommands(),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func sendingProfileCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list",
			Usage: "List all sending profiles",
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.SendingProfiles.ListSendingProfiles()
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Retrieve a specific sending profile",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "profile-id",
					Usage: "The ID of the sending profile to retrieve",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.SendingProfiles.GetSendingProfile(
					c.Int("profile-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a specific sending profile",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "profile-id",
					Usage: "The ID of the sending profile to retrieve",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.SendingProfiles.DeleteSendingProfile(
					c.Int("profile-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
	}
}

func templatesCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list",
			Usage: "List all templates",
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Templates.ListTemplates()
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Retrieve a specific template",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "template-id",
					Usage: "The ID of the template to retrieve",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Templates.GetTemplate(
					c.Int("template-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a specific template",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "template-id",
					Usage: "The ID of the template to delete",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Templates.DeleteTemplate(
					c.Int("template-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
	}
}

func landingPagesCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list",
			Usage: "List all landing pages",
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.LandingPages.ListLandingPages()
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Retrieve a specific landing page",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "page-id",
					Usage: "The ID of the landing page to retrieve",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.LandingPages.GetLandingPage(
					c.Int("page-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a specific landing page",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "page-id",
					Usage: "The ID of the landing page to delete",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.LandingPages.DeleteLandingPage(
					c.Int("page-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:    "import",
			Aliases: []string{"imp"},
			Usage:   "Import a landing page",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url",
					Usage: "The URL of the page to import",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.LandingPages.ImportSite(gophish.ImportSiteRequest{
					URL: c.String("url"),
				},
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
	}
}

func groupCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list",
			Usage: "List all groups",
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Groups.ListGroups()
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Retrieve a specific group",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "group-id",
					Usage: "The ID of the group to retrieve",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Groups.GetGroup(
					c.Int("group-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a specific group",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "group-id",
					Usage: "The ID of the group to delete",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Groups.DeleteGroup(
					c.Int("group-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
	}
}

func campaignCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list",
			Usage: "List all campaigns",
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Campaigns.ListCampaigns()
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Retrieve a specific campaign",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "campaign-id",
					Usage: "The ID of the campaign to retrieve",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Campaigns.GetCampaign(
					c.Int("campaign-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a specific campaign",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "campaign-id",
					Usage: "The ID of the campaign to delete",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Campaigns.DeleteCampaign(
					c.Int("campaign-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "complete",
			Usage: "Complete a specific campaign",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "campaign-id",
					Usage: "The ID of the campaign to complete",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Campaigns.CompleteCampaign(
					c.Int("campaign-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
		{
			Name:  "stats",
			Usage: "Retrieve the stats for a specific campaign",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "campaign-id",
					Usage: "The ID of the campaign to retrieve stats for",
				},
			},
			Action: func(c *cli.Context) error {
				client := gophish.NewClient(
					c.GlobalString("host"),
					c.GlobalString("token"),
				)
				profiles, err := client.Campaigns.GetCampaignSummary(
					c.Int("campaign-id"),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				pretty.Println(profiles)
				return nil
			},
		},
	}
}
