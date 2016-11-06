Live API provides an easy to use client-side interface for interacting with [Facebook Live API](https://developers.facebook.com/docs/videos/live-video/getting-started) from your desktop.

## Setup
#### option 1
You'll need a working Go installation. 
Run 
```bash
go get github.com/kn9ts/frodo
go get github.com/andela-ooranagwa/live-api
```

Then run the app 

```bash
cd $GOPATH/src/github.com/andela-ooranagwa/live-api
go run main.go
```
That starts the app on `localhost:3102`

#### option 2
Just clone this repo and visit the `index.html` page on your browser 😉

### Note
You'll need to update [this section in index.html](https://github.com/andela-ooranagwa/live-api/blob/master/index.html#L69) with a [Facebook APP ID](https://developers.facebook.com/docs/apps/register)

### How to Use
Start the app and visit the index page.

To start using live api and get your RTMP stream url and keys, click the large play button and follow the dialog prompt. Put the RTMP stream URL into your encoder of choice (a good choice is Open Broadcasting Software, OBS, which can be downloaded for free [here](https://obsproject.com/)).

For OBS-specific broadcasts, open OBS, click 'Settings' then 'Stream' and paste in your RTMP stream URL and key, then click 'Ok'. And click `start streaming`.

Head back to the browser pop up; notice that the stream has begun but still in preview mode. Make any necessary updates you desire like adding a description. When you are ready to go live, click the go live button. 

That's all.

### Features
* Add/Update live stream description
* View details of live stream
* Start/Stop live stream 