@echo oFF
if "%1" neq "1" (
>"%temp%\tmp.vbs" echo set WshShell = WScript.CreateObject^(^"WScript.Shell^"^)
>>"%temp%\tmp.vbs" echo WshShell.Run chr^(34^) ^& %0 ^& chr^(34^) ^& ^" 1^",0
start /d "%temp%" tmp.vbs
exit

)
"D:\Program Files\ConEmu\ConEmu64.exe"  "-runlist" "cmd -new_console:d:E:\project\社区\src\jachunPM\jachunPM_commom ||| cmd -new_console:d:E:\project\社区\src\jachunPM\jachunPM_http ||| cmd -new_console:d:E:\project\社区\src\jachunPM\jachunPM_project ||| cmd -new_console:d:E:\project\社区\src\jachunPM\jachunPM_user||| cmd -new_console:d:E:\project\社区\src\jachunPM\jachunPM_test