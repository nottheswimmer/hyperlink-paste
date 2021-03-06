# Hyperlink Paste

**Currently only supports macOS**. Run this and it tries to replace your clipboard with a sensible HTML hyperlink for a URL it contains.
Then it pastes.

## Prerequisites

- [Git](https://git-scm.com/) v2.13+
- [Go](https://golang.org/) v1.16+
- xcode

## Quick Start
```bash
git clone git@github.com:nottheswimmer/hyperlink-paste.git
cd hyperlink-paste
make build
```

To run tests:
```bash
make test
make test with=cover
```

## Usage

After you've built the application, go do the following to add a shortcut:

1. Open Automator.
2. Make a new Quick Action.
3. Make sure it receives 'no input' at all programs.
4. Select Run Apple Script and type in...
```applescript
on run {input, parameters}
    do shell script "{BUILD_DIRECTORY}/hyperlink_paste"

return input
end run
```
5. Save!

Now go to System Preferences > Keyboard > Shortcuts. 
Select Services from the sidebar and find your service. 
Add a shortcut by double clicking (none).

Finally go to System Preferences > Security > Privacy > Accessibility and 
add Automator and the preferred app to run the shortcut.

