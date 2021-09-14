@echo off
mshta "javascript:var sh = new ActiveXObject( 'WScript.Shell' ); sh.Popup( 'App launched!', 10, 'Title!', 64 ); close()"