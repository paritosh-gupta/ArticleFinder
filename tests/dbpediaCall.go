// http://spotlight.dbpedia.org/rest/annotate \
//   --data-urlencode "text=President Obama called Wednesday on Congress to extend a tax break
//   for students included in last year's economic stimulus package, arguing
//   that the policy provides more generous assistance." \
//   --data "confidence=0.2" \
//   --data "support=20"
package main

import (
// "encoding/xml"
"fmt"
"net/url"
"io/ioutil"
"net/http"
// "bytes"
)

func main(){


	v := url.Values{}
	v.Set("text", data)
	v.Add("powered_by", "no")
	v.Add("policy", "whitelist")
	v.Add("confidence", "0.5")
	// v.Add("support","20")
	// v.Add("content-type","application/x-www-form-urlencoded")
	// v.Add("Accept","application/json")
	// status,result:=PostRequest("http://spotlight.sztaki.hu:2222/rest/annotate",v)
	status,result:=PostRequest("http://n2-51-219.dhcp.drexel.edu:2222","/rest/annotate",v)
	
	fmt.Println(status)
	ioutil.WriteFile("result.xml", result, 0644)
	//fmt.Println(string(result))

}



func PostRequest(baseUrl string,path string,urlValues url.Values) (string, []byte){
	
	Url, err := url.Parse(baseUrl)
	if err != nil {
        panic("invalid Url for get request")
    }
	Url.Path += path
    Url.RawQuery = urlValues.Encode()

	client := &http.Client{}
	fmt.Println(Url.String())
	req, err := http.NewRequest("POST",Url.String(),nil)
	// req.Header.Add("Accept", "application/json")
		req.Header.Add("Accept", "text/xml")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return resp.Status, body
}

const data=`Outdoor brand Timberland is debuting an ad campaign called "Made for the Modern Trail" that highlights everyday adventures, exploring what the outdoors mean to urban shoppers.The brand's footwear and apparel is more often worn to outdoor concerts on the waterfront than to hike in the mountains, the company learned from three years of research. Now Timberland says it will promote styles that are stylish and versatile as part of a fall campaign breaking September 24.The effort caters to millennials, who Timberland reintroduced itself to in 2013 with the "Best Then. Better Now." campaign, but is also designed to appeal to a broader subset of consumers, which the brand calls the "outdoor life-styler," who are both trendy and outdoorsy."These are folks who have a real interest and appreciation of the outdoors but they also care about style," said Jim Davey, VP-global marketing at Timberland, who said that research showed this consumer mindset extends across 18- to 60-year-olds. "It was interesting that this 'outdoor life-styler' is a much larger market than we thought."The campaign features more than 500 pieces of digital content including paid and owned media, digital product and lifestyle videos, banners and lookbooks. "We've got a strong sense of how to find that consumer online and be very specific and targeted," said Mr. Davey, adding that Timberland is using data more aggressively than it ever has before. The company plans to invest more than half of its media budget for the 2015 fiscal year in digital experiences. Mr. Davey declined to reveal the full budget for the campaign.Parent company VF Corp. spent $9.7 million on U.S. measured-media for Timberland in 2014, according to Kantar Media.Timberland also tapped British artist Jamie N Commons, who will release his debut album this week, to create two exclusive tracks for the campaign. He will also be featuerd at events and in digital videos as part of the push. "As we thought about showing up in unexpected places and lifestyle segments ... music was a great chance for us to connect with consumers in a very unexpected way," said Mr. Davey.The brand will also highlight other influencers throughout the fall with a push called "Mark Makers," which documents a day in the life of various artists.As part of the new effort, Timberland is also boosting its efforts aimed at women, a market that has grown for the brand over the last 18 months. The company will invest more in women's media outlets including Teen Vogue, Nylon, People StyleWatch, InStyle and Marie Claire."We've seen so much momentum around the women's business, especially in boots," said Mr. Davey. "Women certainly respect the outdoor-proven quality in our boots and they're really surprised by the style. … The big piece is the ability to be not just a footwear brand."The VF Corp.-owned brand aims to bring in $3.1 billion in global revenue by 2019, up from $1.8 billion in 2014, according to annual reports. And tapping into the women's market is an area of potential growth for the brand, Mr. Davey said.Timberland worked with Yard on the campaign concept and video assets, an in-house creative team on print assets and imagery, Assembly on media strategy and buying, MediaBlaze on editorial content, and Coyne PR on public relations."`
