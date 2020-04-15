# WorkoutRecorder
Web app for recording workout
![Screen Shot 2020-04-04 at 7 54 46 PM](https://user-images.githubusercontent.com/52692945/78995847-3ac05380-7b7e-11ea-9b58-6788ce05dc77.png)

#HTTP request
use "net/http" package for http request
muxHTTP.HandleFunc("/", handler)
HandleFunc registers the handler function for the given pattern.
This essentially means that if somebody access address / like "localhost:8080/" or like "localhost:8080/hello", program call "handler" func.