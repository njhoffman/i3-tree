
I want to extend the functionality of this go project "i3-tree (https://github.com/eh-am/i3-tree)" that outputs a hierarchal tree of the current
i3 containers.  Here are useful references:

general i3 functionality: https://i3wm.org/docs/userguide.html
i3 ipc: https://i3wm.org/docs/ipc.html (especially https://build.i3wm.org/docs/ipc.html#_tree_reply)
go-i3 ipc interface: https://pkg.go.dev/go.i3wm.org/i3?utm_source=godoc

Follow existing test patterns, add new tests for key functionality.  Always run tests after each
bath of changes, and if tests are successful run the command 'i3-tree raw' so I can visually inspect
the current tree. Always reference this file, but watch docs/specs/scratch.md for current
instructions.  Keep track of approved changes in CHANGELOG.md, after each approved change bump the
version (use semver versioning) by a minor version, and commit with a succint message.   

Everytime scratch.md changes follow the instructions, run tests and 'i3-tree raw' output if tests are successful.  If the first line of scratch.md changes consider it a
new task with previous changes approved, and update CHANGELOG, bump version, and commit previous changes.

View scratch.md for additional instructions.
