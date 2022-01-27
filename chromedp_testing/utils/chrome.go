package chrome

import (
	"context"
	"os"
	"time"

	"github.com/chromedp/cdproto/dom"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	ch "github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

const unblockScript = `(function(w, n, wn) {
  // Pass the Webdriver Test.
  Object.defineProperty(n, 'webdriver', {
    get: () => false,
  });

  // Pass the Languages Test.
  // Overwrite the plugins property to use a custom getter.
  Object.defineProperty(n, 'languages', {
    get: () => ['en-US', 'en'],
  });

  // Pass the Chrome Test.
  // We can mock this in as much depth as we need for the test.
  w.chrome = {
    runtime: {},
  };

  // Pass the Permissions Test.
  const originalQuery = wn.permissions.query;
  return wn.permissions.query = (parameters) => (
    parameters.name === 'notifications' ?
      Promise.resolve({ state: Notification.permission }) :
      originalQuery(parameters)
  );

})(window, navigator, window.navigator);`

var unblockTask = ch.ActionFunc(func(ctx context.Context) error {
	_, err := page.AddScriptToEvaluateOnNewDocument(unblockScript).Do(ctx)
	if err != nil {
		return err
	}
	return nil
})

func StartChrome(headless bool, extraOptions ...ch.ExecAllocatorOption) (context.Context, context.CancelFunc, error) {
	opts := append(ch.DefaultExecAllocatorOptions[:],
		ch.DisableGPU,
		// Set the headless flag to false to display the browser window
		ch.Flag("headless", headless),
		ch.Flag("start-fullscreen", true),
		//ch.Flag("disable-web-security", true),
		ch.UserDataDir(os.Getenv("HOME")), // dir to store cookies,local-storage, etc
	)

	if len(extraOptions) > 0 {
		opts = append(opts, extraOptions...)
	}

	if os.Getenv("CHROME") != "" {
		opts = append(opts, ch.ExecPath(os.Getenv("CHROME")))
	}

	ctx, cancel1 := ch.NewExecAllocator(context.Background(), opts...)

	ctx, cancel2 := ch.NewContext(ctx, ch.WithLogf(log.Printf))

	var cancel context.CancelFunc = func() {
		cancel2()
		cancel1()
	}

	err := ch.Run(ctx, unblockTask)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	return ctx, cancel, nil
}

func GetDescendant(node *cdp.Node, childPaths []int) *cdp.Node {
	if node == nil {
		return nil
	}

	for _, index := range childPaths {
		if index < 0 || index >= len(node.Children) {
			return nil
		}
		node = node.Children[index]
	}

	return node
}

func GetChildren(nodeID cdp.NodeID) []ch.Action {
	return []ch.Action{
		ch.ActionFunc(func(c context.Context) error { // Get the children of the table
			return dom.RequestChildNodes(nodeID).WithDepth(-1).Do(c)
		}),
		ch.Sleep(1 * time.Second)}
}
