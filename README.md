#save data in Redis#
===================

A blank for the future microservice. 
Saving and returning data using Redis.

Application for tracking visited links.

The app provides two resources:

resource for loading visits

	POST /links {"links": [ ] }

and a resource for getting statistics

	GET /domains?from=xxx&to=yyy
where xxx and yyy is keys in Redis

The first resource is used to send an array 
of links in a POST request to the service.
Addresses are cut from links
domains and stored in the database. 
The key for saving and subsequent search 
is the time of receipt of the request.

The second resource is used for getting 
a GET request for a list of unique domains
visited during the specified period.
The request specifies a time interval.

###Warnings:###
---------------
- the application was tested only in LinuxMint
- for testing and debugging, the Redis key is output to the terminal. 
to remove this action, comment out line 47 in the *workRedis.go* file
- for identification in Redis, a string is added to the key
(lines 23 and 46 in the *workRedis.go* file)
