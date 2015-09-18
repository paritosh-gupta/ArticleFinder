#ArticleFinder

Installation
-
    
        git clone https://github.com/paritosh-gupta/ArticleFinder
    
  - Dependencies
    - Go,
    - Java 8(My sincere apologies for the jvm)
    - Python
    - Docker for easy deployment
- Setting Up the Python Goose Server
    -
    > The tool uses python [Goose] to extract text,images and the title from a web url.

    ```sh
    $ git clone https://github.com/grangier/python-goose.git
    $ cd python-goose
    $ pip install -r requirements.txt
    $ python setup.py install
    ```
  - Browse to `Backend>goose` (where you cloned this repository)
  ```sh
    $ python gooseServer.py
    ``` 
    The server will run at localhost 8081
    
- DBPedia Spotlight installation
    - 
    > [DBPedia Spotlight] is used to extract entities out of the Database. This thing eats ram for breakfast. I would recommend 10 gig or more in your server. But I got it run on a mac book with 8 gigs of ram.
    
    - Download the `DBPedia-spotlight-0.7.jar` from https://github.com/dbpedia-spotlight/dbpedia-spotlight/releases
    - Download the en_4+8.tar.gz from http://spotlight.sztaki.hu/downloads/. If you have more than 8gb of ram to spare download the en_2+2.tar.gz might which might give better performance(untested)
    ```sh
    $ tar -xvf en_4+8.tar.gz 
    $ #`Make sure JAVA_HOME is in your path`
    $ #`if not for mac jdk8 use`
    $ export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_60.jdk/Contents/Home` 
    $ java -Xmx6G -Xms3g -jar dbpedia-spotlight-0.7.jar ~/Downloads/en_4+8 http://localhost:2222/rest 
    ```
	- ###Argument Explanatation
		- `Xmx6G` : Max ram usage.Bump to 8 if you have more
		- `Xms3g` : Minx ram usage.Bump to 6 if you have more
		- `jar dbpedia-spotlight-0.7.jar` :The jar file
		- `~/Downloads/en_4+8` : The location of the model tar extraction
		- ` http://localhost:222/rest`:The address where server will run
 
 - The spotlight server should start in some time
    
	- DBPedia lookup installation
	-
	> DBPedia Lookup is a service used to obtain information about entities present in dbpedia.It also powers the autocomplete in search.
	
	Install [Maven] first:-
	- /run Server dbpedia-lookup-index-3.8
	- Comming Soon