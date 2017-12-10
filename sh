sed -i "buffer" "s/package midi/package event/" *.go
sed -i "buffer" "s/DeltaTime/deltatime.DeltaTime/g" *.go
sed -i "buffer" "s/) deltatime.DeltaTime() /) DeltaTime() /g" *.go
sed -i "buffer" "s/_deltatime.DeltaTime(t /DeltaTime(t /g" *.go
sed -i "buffer" "s/event.deltatime.DeltaTime()/event.DeltaTime()/g" *test.go
sed -i "buffer" "s/0x00, //g" *test.go
