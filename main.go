/*
 * Go automate for BrowserStack
 * taken from sourcegraph go-selenium examples
 */

package main

import (
	"fmt"
	"github.com/rghose/go-selenium"
)

func ExampleFindElement() {
	var webDriver selenium.WebDriver
	var err error
	remoteUrl := "http://<user>:<key>@hub-cloud.browserstack.com/wd/hub"
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	if webDriver, err = selenium.NewRemote(caps, remoteUrl); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("http://www.google.com")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	var elem selenium.WebElement
	elem, err = webDriver.FindElement(selenium.ByName, "q")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}

	if err = elem.SendKeys("BrowserStack"); err != nil {
		fmt.Printf("Failed to send keys of element: %s\n", err)
		return
	}

	if err = elem.Submit(); err != nil {
		fmt.Printf("Failed to submit, because of: %+v", err)
		return
	}

	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	webDriver.Quit()
}

func main() {
	ExampleFindElement()
}
