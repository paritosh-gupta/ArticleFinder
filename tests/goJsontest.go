package main

import (
	"encoding/json"
	"fmt"
)



func main() {

	 var jsonBlob = []byte(`[ { "description": "Two months after its launch, Showtime's streaming-only subscriber service has been eye-opening for the network.", "image": "http://www.adweek.com/files/imagecache/node-detail/2015_Sep/showtime-streaming-hed-2015.png", "text": "Two months after its launch, Showtime's streaming-only subscriber service has been eye-opening for the network.\n\nSince the service, which is also called Showtime, went live on July 7\u2014exactly three months after HBO debuted its own standalone streaming platform, HBO Now\u2014Showtime Networks President David Nevins has been receiving detailed, data-fueled reports about its growth and usage each day. Having long been limited to getting monthly reports about subscriber trends for the premium cable network, he now browses detailed updates each morning, learning how many subscriptions were sold and what the service's usage looks like.\n\n\"You can see on a nightly basis exactly what people are watching, and it's fascinating,\" Nevins said.\n\nOne discovery: many subscribers are opting to stream the network's live feed, instead of seeking out on-demand programming. \"The livestreaming has done better than we expected,\" said Nevins of the feature, something which HBO Now doesn't offer. \"As people are discovering Showtime, a lot of them just want to see what's on the livestreaming, so they're watching it there.\"\n\nThe service is currently available on iOS devices, Roku and PlayStation Vue for $10.99 per month, $4 less than HBO Now. Hulu subscribers can purchase it as a premium add-on for $8.99.\u00a0Nevins said he expects to add more partners before the Oct. 4 season premieres of Homeland and The Affair.\n\nWhile Showtime isn't releasing specific streaming numbers yet\u2014the network also has more than 23 million households subscribe via cable and satellite\u2014Nevins said early results have been promising. It's also put Showtime in the new position of no longer being separated from its users by a middleman.\n\n\"Showtime has been in the subscription business for 40 years, but we never had any subscribers. Now we do,\" said chairman and CEO Matthew Blank. \"David came to the company five years ago, and in his first week on the job, he asked questions that you would know the answer to if you actually had a direct relationship with your subscribers. Now we're going to be able to answer those questions a lot better for ourselves.\"\n\nDespite the sudden influx of data, Nevins isn't tailoring Showtime to his streaming subscribers\u2014at least, not yet. \"I'm \u00a0very well aware that the vast majority of the people that watching Showtime are watching through the MVPDs (multichannel video programming distributors, such as cable or satellite services), although their behavior is changing, too. They're streaming more and watching more delayed,\" said Nevins, who estimates that the live tune-in for one of his original series represents\u00a0\"probably a fifth\" of the audience that tunes in during the first seven days.\n\nIn addition to the live feed's popularity, Nevins has been\u00a0also surprised by the large number of streaming subscribers who are watching the network's documentaries. (The company has been bulking up its documentaries, with what Nevins called \"seven meaningful films\" on the slate, including American Dream/American Knightmare, about Death Row Records cofounder Suge Knight, airing Sept. 26.)\n\nBlank attributed the documentary interest to the new ways that users are interacting with Showtime via the streaming service (as well as Showtime Anytime, its streaming option for cable and satellite subscribers), where they are actively selecting and browsing through content, rather than watching whatever happens to be airing on the network at the time.\n\n\"Because you're forced to menu, in a very slick, promotional way, I think you take away a lot more from what's on the service and what Showtime is,\" Blank said. \"You're not relying on some external promotional vehicle to bring you to that documentary. You now have that great customer interface at your disposal, where you're going to get to immerse yourself with things like docs, and maybe you're going to want to watch them more. I think we're going to see much more of that.\"", "title": "Here's What Surprised Showtime After It Began Offering a Streaming Service" } ]
	`)
	
	type Animal struct {
		Name  string
		Order string
		Bitch string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}