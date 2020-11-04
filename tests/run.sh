################################
####   THE TESTING SCRIPT   ####
################################

# Setting up the number of tries and hit URL
url="http://localhost:3000/"
tries=10000

#####################
####   NODE.JS   ####
#####################

# run the app (forked)
cd nodejs
node app.js >/dev/null 2>&1 &
pid=$!

# wait for the app
sleep 5

# run the curl command $tries times
sumnode=0
for i in `seq 1 $tries`
do
    rtime=$(curl -s -w '%{time_total}' -o /dev/null $url)
    sumnode=$(echo $sumnode + $rtime | bc)
done

# kill the forked process and return to the root
kill $pid
wait $pid 2>/dev/null
cd ..

####################
####     GO     ####
####################

# run the app (forked) (build if not builded)
cd go
if [ ! -f main.go ]; then
    go build main.go
fi
./main >/dev/null 2>&1 &
pid=$!

# wait for the start
sleep 5

# run the curl command $tries timres
sumgo=0
for i in `seq 1 $tries`
do
    rtime=$(curl -s -w '%{time_total}' -o /dev/null $url)
    sumgo=$(echo $sumgo + $rtime | bc)
done

# kill the forked process
kill $pid
wait $pid 2>/dev/null
cd ..

#########################
####   FINAL STATS   ####
#########################

printf 'Node.js\t\t~%.7f ms\n' $(echo $sumnode/$tries\*1000 | bc -l)
printf 'Go\t\t~%.7f ms\n' $(echo $sumgo/$tries\*1000 | bc -l)
