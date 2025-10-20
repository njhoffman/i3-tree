Task: Fix branch character usage

Reduce to single characters config  branches.connect_h to `├` and branches.connect_v to `└`, use
branches.horizontal to fill in the missing characters.  This will help with the focus highlighting.  

Focused highlighting, given the following sample output: 
`
    [root][splith] root
    ├──[output][output] __i3
    │  └──[con][splith] content
    │     └──[workspace] 󰊓 __i3_scratch
    └──[output][output] eDP-1
       ├──[dockarea][dockarea] topdock
       │  └──[con] (Polybar) polybar-top-powerline_eDP-1
       ├──[con][splith] content
       │  ├──[workspace][splith] 󰊓 1
       │  │  └──[con] (Alacritty) main
`
If the last line `[con] (Alacritty) main` is the focused window, the highlighting of the branches 
from root should follow this pattern (highlighted characters replaced with +)
`
    [root][splith] root
    +──[output][output] __i3
    +  └──[con][splith] content
    +     └──[workspace] 󰊓 __i3_scratch
    +++[output][output] eDP-1
       +──[dockarea][dockarea] topdock
       +  └──[con] (Polybar) polybar-top-powerline_eDP-1
       +++[con][splith] content
       │  +++[workspace][splith] 󰊓 1
       │  │  +++[con] (Alacritty) main
`
Set the default config formatting.focus_brackets.foreground to 81
