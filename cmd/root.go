package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/eh-am/i3-tree/cmd/internal"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var flagHelp = `i3-tree generates a user friendly view of the i3 tree

EXAMPLES
# display focused workspace
i3-tree

# display all non empty workspaces
i3-tree all

# show a specific workspace (for example, workspace 6)
i3-tree 6

# show focused workspace, with no colors
i3-tree --render=no-color

# use mock data (useful if you don't have i3 running)
i3-tree --from=mock

# watch mode: refresh every 5 seconds (using default interval)
i3-tree --watch=0

# watch mode: refresh every 2 seconds
i3-tree -w 2

# watch mode: refresh every 10 seconds
i3-tree --watch 10
`

var fetchStratName *string
var renderStratName *string
var watchInterval *int

var rootFs *flag.FlagSet
var root *ffcli.Command

func init() {
	rootFs = flag.NewFlagSet("root", flag.ExitOnError)

	fetchStratName = rootFs.String(
		"from",
		string(internal.FromI3),
		"where to fetch the tree from. available: "+fmt.Sprintf("%s", internal.AvailableFetchStrats),
	)

	renderStratName = rootFs.String(
		"render",
		string(internal.ConsoleStrat), // Default
		"where/how to render the output to. available: "+fmt.Sprintf("%s", internal.AvailableRendererStrats),
	)

	watchInterval = rootFs.Int(
		"watch",
		-1,
		"watch mode: refresh every N seconds (default: 5 if flag is used without value)",
	)
	rootFs.IntVar(watchInterval, "w", -1, "shorthand for --watch")

	root = &ffcli.Command{
		Name:       "i3-tree",
		ShortUsage: "i3-tree",
		LongHelp:   flagHelp,
		ShortHelp:  "Print the i3 tree in a user friendly format",
		FlagSet:    rootFs,
		Exec:       rootExec,
	}
}

func rootExec(ctx context.Context, args []string) error {
	fetcher, err := internal.NewFetcher(*fetchStratName)
	if err != nil {
		return err
	}

	pruneArg := ""
	if len(args) > 0 {
		pruneArg = args[0]
	}
	pruner, err := internal.NewPruner(pruneArg)
	if err != nil {
		return err
	}

	renderer, err := internal.NewRenderer(*renderStratName)
	if err != nil {
		return err
	}

	i3tv := i3treeviewer.NewI3TreeViewer(
		fetcher,
		pruner,
		renderer,
	)

	// Determine watch interval
	interval := *watchInterval

	// -1 means flag was not set (no watch mode)
	// 0 means flag was set with value 0, which we treat as default (5 seconds)
	// Any positive value is used as-is
	if interval == -1 {
		// No watch mode
		return i3tv.View()
	}

	if interval == 0 {
		interval = 5 // default interval when --watch is used with 0
	}

	// Watch mode: loop forever
	for {
		// Clear screen
		clearScreen()

		// Render tree
		if err := i3tv.View(); err != nil {
			return err
		}

		// Wait for interval
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// clearScreen clears the terminal screen
func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
