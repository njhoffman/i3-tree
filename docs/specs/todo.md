Task: Add logging functionality with watch-log command switch

Add `default_log_path` option to config with a default value of `/tmp/i3-tree.log`. 

`│  └──[con] (Alacritty) main`

`(-W|--watch-log)` Acts like watch except also outputs log statements below the tree summarizing the
captured changes between tree redraws.  An optional file path may be added after this switch to
indicate output file for logged statements, otherwise use `default_log_path`, this will contain all
of the logged messages with timestamps.  For the output below the tree, new log statements should 
be displayed top-down, and don't show log messages that extend beyond the window height 
(no scrolling). Changes should focus on modifications to windows and workspaces. Don't show changes 
to window focus.  Output messages in a colorful and easy to read format, emphasizing actions 
(i.e. 'moved window', 'created window', 'added workspace', 'resized window').


