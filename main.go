package main

import (
	"fmt"
	"github.com/JoshStrobl/trunk"
	"github.com/lrstanley/girc"
	"github.com/narwhalirc/tusk"
	"strings"
	"time"
)

func Parse(c *girc.Client, e girc.Event, m tusk.NarwhalMessage) {
	cleanCommand := strings.Replace(m.Message, ".", "", 1)

	if !strings.HasPrefix(m.Message, ".") { // Only accept messages which start with . (commands)
		return
	}

	switch cleanCommand { // Anyone can use these
	case "budgie": // Budgie (Generic Info)
		c.Cmd.Reply(e,
			`Budgie, our flagship desktop environment, is made available over at https://github.com/solus-project/budgie-desktop. 
If you have an issue with Budgie, please go to its dedicated issue tracker and file an issue, so we can work towards addressing it for all of our users!
			`)
	case "contribute", "getinvolved": // Get Involved
		c.Cmd.Reply(e, "We always love more contributions from even more people! If you want to contribute to Solus, we'd love for you to check out https://getsol.us/articles/contributing/getting-involved/en/")
	case "dev", "phab": // Link our Development Tracker
		c.Cmd.Reply(e, "We have a development tracker which enables you to report issues, file package requests, and more. Feel free to check it out at https://dev.getsol.us")
	case "docs", "help", "helpcenter": // Link our Help Center docs
		c.Cmd.Reply(e, "We have documentation available on our Help Center, available at https://getsol.us/help-center/home/")
	case "download", "get": // Getting the latest Solus ISO
		c.Cmd.Reply(e, "You can get the latest release of Solus and any of its editions over at https://getsol.us/download/")
	case "eopkg": // Link to our package manager documentation
		c.Cmd.Reply(e, "Solus uses its own, unique, package manager called eopkg. To learn about package management on Solus, go to https://getsol.us/articles/package-management/")
	case "eta": // We don't give them
		c.Cmd.Reply(e, "Solus does not provide ETAs. It's ready when it's ready.")
	case "facebook": // Link to our Facebook
		c.Cmd.Reply(e, "Solus has a Facebook account at https://facebook.com/get.solus")
	case "flarum", "forums": // Link to our Forums
		c.Cmd.Reply(e, "We have forums available to discuss a wide range of topics. Feel free to check it out at https://dev.getsol.us. You can sign in with your Dev Tracker or GitHub account!")
	case "guidelines", "rules": // Link to our Community Guidelines
		c.Cmd.Reply(e, "Solus always aims to provide a friendly, healthy environment for all users. Please ensure you read and follow our Community Guidelines, which can be found at https://getsol.us/articles/contributing/community-guidelines/en/")
		c.Cmd.Reply(e, "In the event you have a concern or issue with another member of our community, please reach out to a member of the Core Team immediately.")
	case "mastodon": // Link to our Mastodon
		c.Cmd.Reply(e, "Solus has a Mastodon account at https://mastodon.cloud/@SolusProject")
	case "packaging": //  Packaging Documentation
		c.Cmd.Reply(e, "Want to get started with packaging under Solus? Check out https://getsol.us/articles/packaging/")
	case "reddit": // Link to our sub-reddit
		c.Cmd.Reply(e, "Solus has its own subreddit at https://reddit.com/r/SolusProject")
	case "roadmap": // Link to our Roadmap
		c.Cmd.Reply(e, "Solus has a dedicated Roadmap page where you can find out what our short-term and long-term plans are, available at https://getsol.us/solus/roadmap/")
	case "social": // General Social Response
		c.Cmd.Reply(e, "Solus has a multitude of accounts across various social networks and mediums. You can see all the places we are at by going to https://getsol.us/articles/contributing/getting-involved/en/#social-media")
	case "twitter": // Link to our Twitter
		c.Cmd.Reply(e, "Solus has a Twitter account at https://twitter.com/SolusProject")
	}

	if m.Admin { // If the issuer is an admin
		switch m.Command {
		case "frozen": // Freeze for sync
			c.Cmd.Topic(m.Channel, "Solus Development | Stable: Syncing | Unstable: Frozen")
		case "synced": // Just synced
			SetToSynced(c, m)
		}
	}
}

// SetToSynced will set our topic to our synced message
func SetToSynced(c *girc.Client, m tusk.NarwhalMessage) {
	isoFormat := "2006-01-02T15:04:05-07:00"
	now := time.Now()
	nowISO := now.Format(isoFormat)

	trunk.LogInfo(fmt.Sprintf("%s performed a sync on %s and updated the topic", m.Issuer, nowISO))
	c.Cmd.Topic(m.Channel, fmt.Sprintf("Solus Development | Stable: Synced (Last on %s) | Unstable: Unfrozen", nowISO))
}
