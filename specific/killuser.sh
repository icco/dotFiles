pgrep -u pi
ps -fp $(pgrep -u pi)
killall -KILL -u pi 
userdel -r pi
