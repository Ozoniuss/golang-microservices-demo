# Golang Microservices Demo

Work in progress...

Word before...
--------------

This project is inspired from a take-home assignment I received during an interview. Obviously, I cannot reveal the exact statement, nor the points that were evaluated, so I will instead provide a short summary. Basically, a large file with a bunch of json items representing ports was provided (which is likely the `ports.json` file taken from [this repository](https://github.com/marchah/sea-ports), and you had to implement a REST client that reads the file, and a ports microservice that can process the ports and persist them (at least that's how I understood it). The REST client had limited memory resources, meaning that it could not load the entire file into memory. There were some other points and many implementation details were left to choice, but those are not relevant for this demo.

That assignment was the first time I've been exposed to microservices, thus my past experience working on a REST API with [gorrila/mux](https://github.com/gorilla/mux) (archived now, but it's what I used in the assignment implementation at the time) and various monolith toy projects in university didn't make this project less intimidating. Looking back at the code I wrote at the time, there were many things I find questionable now 😅, and I definitely made some weird design choices. I'm still happy with how it turned out back then, but many parts could be improved.

With this project, I decided to actually improve it. Perhaps "improve" is not the right word, since I ended up pretty much rewriting everything 😁. Though, I also changed a bit the task. I still implemented the ports service for the sake of nostalgia (and that I had the file already, didn't really want to spend time modelling some other entity etc...😛), but the fact that the ports couldn't be processed at once made me think about [online algorithms](https://en.wikipedia.org/wiki/Online_algorithm), or algorithms in general where memory is limited. So I've decided to change the purpose of the client a bit, which is detailed in the section below.


Project overview
----------------

The app comes with a `files-streamer` REST API, which can communicate with multiple services. Each of those services can implement some sort of "online" algorithm. The idea behind it is that any client can connect to this API and stream a file of arbitrary size from their node, which the file streamer processes in real time by converting the incoming stream data to a protobuf format that can be then streamed further to a microservice. The microservice then performs the task, providing some real-time notifications about the data that is processed back to the API, which are forwarded back to the client through the API. These tasks may be anything, from computing the average of a long list of numbers to inserting some items in the database (which are both included in the demo).

For extended functionality, the files-streamer will also allow streaming files that are local to the API. It will also provide the required format for the file if streamed directly from the client, and offer a RESTful interface to perform these operations.

The entire process must happen in a way that makes a good trade-off between the available memory and the time it takes to process the entire file.

The client API and the microservices will be deployed on my home server and will be available on my domain, [www.ozoniuss.com](www.ozoniuss.com) (unreachable at the time of writing this). The file stream is also going to be encrypted with TLS. However, to avoid exploits, this will only be publicly available under some sort of authorization mechanism that I will think about once completed.